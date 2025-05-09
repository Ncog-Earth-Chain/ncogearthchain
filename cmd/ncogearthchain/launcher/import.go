package launcher

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	"github.com/Ncog-Earth-Chain/forest-base/hash"
	"github.com/Ncog-Earth-Chain/forest-base/inter/idx"
	"github.com/Ncog-Earth-Chain/forest-base/kvdb"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/status-im/keycard-go/hexutils"
	"gopkg.in/urfave/cli.v1"

	"github.com/Ncog-Earth-Chain/ncogearthchain/gossip"
	"github.com/Ncog-Earth-Chain/ncogearthchain/gossip/emitter"
	"github.com/Ncog-Earth-Chain/ncogearthchain/integration"
	"github.com/Ncog-Earth-Chain/ncogearthchain/inter"
	"github.com/Ncog-Earth-Chain/ncogearthchain/utils/iodb"
	"github.com/Ncog-Earth-Chain/ncogearthchain/utils/ioread"
)

type restrictedEvmBatch struct {
	kvdb.Batch
}

func (v *restrictedEvmBatch) Put(key []byte, value []byte) error {
	if len(key) != 32 {
		if !bytes.HasPrefix(key, []byte("secure-key-")) && !bytes.HasPrefix(key, []byte("c")) {
			return errors.New("not expected prefix for EVM history dump")
		}
	}
	return v.Batch.Put(key, value)
}

func importEvm(ctx *cli.Context) error {
	if len(ctx.Args()) < 1 {
		utils.Fatalf("This command requires an argument.")
	}

	cfg := makeAllConfigs(ctx)

	rawProducer := integration.DBProducer(path.Join(cfg.Node.DataDir, "chaindata"), cacheScaler(ctx))
	gdb, err := makeRawGossipStore(rawProducer, cfg)
	if err != nil {
		log.Crit("DB opening error", "datadir", cfg.Node.DataDir, "err", err)
	}
	defer gdb.Close()

	for _, fn := range ctx.Args() {
		log.Info("Importing EVM storage from file", "file", fn)
		if err := importEvmFile(fn, gdb); err != nil {
			log.Error("Import error", "file", fn, "err", err)
			return err
		}
		log.Info("Imported EVM storage from file", "file", fn)
	}

	return nil
}

func importEvmFile(fn string, gdb *gossip.Store) error {
	// Open the file handle and potentially unwrap the gzip stream
	fh, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer fh.Close()

	var reader io.Reader = fh
	if strings.HasSuffix(fn, ".gz") {
		if reader, err = gzip.NewReader(reader); err != nil {
			return err
		}
		defer reader.(*gzip.Reader).Close()
	}

	return iodb.Read(reader, &restrictedEvmBatch{gdb.EvmStore().EvmKvdbTable().NewBatch()})
}

func importEvents(ctx *cli.Context) error {
	if len(ctx.Args()) < 1 {
		utils.Fatalf("This command requires an argument.")
	}

	// avoid P2P interaction, API calls and events emitting
	genesis := getNcogearthchainGenesis(ctx)
	cfg := makeAllConfigs(ctx)
	cfg.Ncogearthchain.Protocol.EventsSemaphoreLimit.Size = math.MaxUint32
	cfg.Ncogearthchain.Protocol.EventsSemaphoreLimit.Num = math.MaxUint32
	cfg.Ncogearthchain.Emitter.Validator = emitter.ValidatorConfig{}
	cfg.Ncogearthchain.TxPool.Journal = ""
	cfg.Node.IPCPath = ""
	cfg.Node.HTTPHost = ""
	cfg.Node.WSHost = ""
	cfg.Node.NoUSB = true
	cfg.Node.P2P.ListenAddr = ""
	cfg.Node.P2P.NoDiscovery = true
	cfg.Node.P2P.BootstrapNodes = nil
	cfg.Node.P2P.DiscoveryV5 = false
	cfg.Node.P2P.BootstrapNodesV5 = nil
	cfg.Node.P2P.StaticNodes = nil
	cfg.Node.P2P.TrustedNodes = nil

	err := importEventsToNode(ctx, cfg, genesis, ctx.Args()...)
	if err != nil {
		return err
	}

	return nil
}

func importEventsToNode(ctx *cli.Context, cfg *config, genesis integration.InputGenesis, args ...string) error {
	node, svc, nodeClose := makeNode(ctx, cfg, genesis)
	defer nodeClose()
	startNode(ctx, node)

	for _, fn := range args {
		log.Info("Importing events from file", "file", fn)
		if err := importEventsFile(svc, fn); err != nil {
			log.Error("Import error", "file", fn, "err", err)
			return err
		}
	}
	return nil
}

func checkEventsFileHeader(reader io.Reader) error {
	headerAndVersion := make([]byte, len(eventsFileHeader)+len(eventsFileVersion))
	err := ioread.ReadAll(reader, headerAndVersion)
	if err != nil {
		return err
	}
	if bytes.Compare(headerAndVersion[:len(eventsFileHeader)], eventsFileHeader) != 0 {
		return errors.New("expected an events file, mismatched file header")
	}
	if bytes.Compare(headerAndVersion[len(eventsFileHeader):], eventsFileVersion) != 0 {
		got := hexutils.BytesToHex(headerAndVersion[len(eventsFileHeader):])
		expected := hexutils.BytesToHex(eventsFileVersion)
		return errors.New(fmt.Sprintf("wrong version of events file, got=%s, expected=%s", got, expected))
	}
	return nil
}

func importEventsFile(srv *gossip.Service, fn string) error {
	// Watch for Ctrl-C while the import is running.
	// If a signal is received, the import will stop.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	// Open the file handle and potentially unwrap the gzip stream
	fh, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer fh.Close()

	var reader io.Reader = fh
	if strings.HasSuffix(fn, ".gz") {
		if reader, err = gzip.NewReader(reader); err != nil {
			return err
		}
		defer reader.(*gzip.Reader).Close()
	}

	// Check file version and header
	if err := checkEventsFileHeader(reader); err != nil {
		return err
	}

	stream := rlp.NewStream(reader, 0)

	start := time.Now()
	last := hash.Event{}

	batch := make(inter.EventPayloads, 0, 8*1024)
	batchSize := 0
	maxBatchSize := 8 * 1024 * 1024
	epoch := idx.Epoch(0)
	txs := 0
	events := 0

	processBatch := func() error {
		if batch.Len() == 0 {
			return nil
		}
		done := make(chan struct{})
		err := srv.DagProcessor().Enqueue("", batch.Bases(), true, nil, func() {
			done <- struct{}{}
		})
		if err != nil {
			return err
		}
		<-done
		last = batch[batch.Len()-1].ID()
		batch = batch[:0]
		batchSize = 0
		return nil
	}

	for {
		select {
		case <-interrupt:
			return fmt.Errorf("interrupted")
		default:
		}
		e := new(inter.EventPayload)
		err = stream.Decode(e)
		if err == io.EOF {
			err = processBatch()
			if err != nil {
				return err
			}
			break
		}
		if err != nil {
			return err
		}
		if e.Epoch() != epoch || batchSize >= maxBatchSize {
			err = processBatch()
			if err != nil {
				return err
			}
		}
		epoch = e.Epoch()
		batch = append(batch, e)
		batchSize += 1024 + e.Size()
		txs += e.Txs().Len()
		events++
	}
	srv.WaitBlockEnd()
	log.Info("Events import is finished", "file", fn, "last", last.String(), "imported", events, "txs", txs, "elapsed", common.PrettyDuration(time.Since(start)))

	return nil
}
