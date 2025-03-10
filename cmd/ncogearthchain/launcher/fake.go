package launcher

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ncog-Earth-Chain/forest-base/inter/idx"
	cli "gopkg.in/urfave/cli.v1"

	"github.com/Ncog-Earth-Chain/ncogearthchain/integration/makegenesis"
	"github.com/ethereum/go-ethereum/cryptod"
)

// FakeNetFlag enables special testnet, where validators are automatically created
var FakeNetFlag = cli.StringFlag{
	Name:  "fakenet",
	Usage: "'n/N' - sets coinbase as fake n-th key from genesis of N validators.",
}

/* func getFakeValidatorKey(ctx *cli.Context) *ecdsa.PrivateKey {
	num, _, err := parseFakeGen(ctx.GlobalString(FakeNetFlag.Name))
	if err != nil {
		return nil
	}

	if num == 0 {
		return nil
	}

	return makegenesis.FakeKey(int(num))
} */

func getFakeValidatorKey(ctx *cli.Context) *cryptod.PrivateKey {
	num, _, err := parseFakeGen(ctx.GlobalString(FakeNetFlag.Name))
	if err != nil {
		return nil
	}

	if num == 0 {
		return nil
	}

	return makegenesis.FakeKey(int(num))
}

func parseFakeGen(s string) (id idx.ValidatorID, num int, err error) {
	parts := strings.SplitN(s, "/", 2)
	if len(parts) != 2 {
		err = fmt.Errorf("use %%d/%%d format")
		return
	}

	var u32 uint64
	u32, err = strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return
	}
	id = idx.ValidatorID(u32)

	u32, err = strconv.ParseUint(parts[1], 10, 32)
	num = int(u32)
	if num < 0 || int(id) > num {
		err = fmt.Errorf("key-num should be in range from 1 to validators (<key-num>/<validators>), or should be zero for non-validator node")
		return
	}

	return
}
