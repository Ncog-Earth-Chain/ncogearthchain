# Ncogearthchain 

EVM-compatible chain secured by the Forest consensus algorithm.
With NCOG Earth Chain, you can experience the [fastest blockchain](https://www.ncogchain.earth/fastest-blockchain) transactions.

## Building the source

Building `ncogearthchain` requires both a Go (version 1.14 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make ncogearthchain
```
The build output is ```build/ncogearthchain``` executable.

## Running `ncogearthchain`

Going through all the possible command line flags is out of scope here,
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `ncogearthchain` instance.

### Launching a network

Launching `ncogearthchain` for a network:

```shell
$ ncogearthchain --genesis /path/to/genesis.g
```

### Configuration

As an alternative to passing the numerous flags to the `ncogearthchain` binary, you can also pass a
configuration file via:

```shell
$ ncogearthchain --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ ncogearthchain --your-favourite-flags dumpconfig
```

#### Validator

New validator private key may be created with `ncogearthchain validator new` command.

To launch a validator, you have to use `--validator.id` and `--validator.pubkey` flags to enable events emitter.

```shell
$ ncogearthchain --nousb --validator.id YOUR_ID --validator.pubkey 0xYOUR_PUBKEY
```

`ncogearthchain` will prompt you for a password to decrypt your validator private key. Optionally, you can
specify password with a file using `--validator.password` flag.

#### Participation in discovery

Optionally you can specify your public IP to straighten connectivity of the network.
Ensure your TCP/UDP p2p port (5050 by default) isn't blocked by your firewall.

```shell
$ ncogearthchain --nat extip:1.2.3.4
```

## Dev

### Running testnet

The network is specified only by its genesis file, so running a testnet node is equivalent to
using a testnet genesis file instead of a mainnet genesis file:
```shell
$ ncogearthchain --genesis /path/to/testnet.g # launch node
```

It may be convenient to use a separate datadir for your testnet node to avoid collisions with other networks:
```shell
$ ncogearthchain --genesis /path/to/testnet.g --datadir /path/to/datadir # launch node
$ ncogearthchain --datadir /path/to/datadir account new # create new account
$ ncogearthchain --datadir /path/to/datadir attach # attach to IPC
```

### Testing

Forest has extensive unit-testing. Use the Go tool to run tests:
```shell
go test ./...
```

If everything goes well, it should output something along these lines:
```
ok  	github.com/Ncog-Earth-Chain/go-ncogearthchain/app	0.033s
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/cmd/cmdtest	[no test files]
ok  	github.com/Ncog-Earth-Chain/go-ncogearthchain/cmd/ncogearthchain	13.890s
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/cmd/ncogearthchain/metrics	[no test files]
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/cmd/ncogearthchain/tracing	[no test files]
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/crypto	[no test files]
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/debug	[no test files]
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/ethapi	[no test files]
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/eventcheck	[no test files]
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/eventcheck/basiccheck	[no test files]
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/eventcheck/gaspowercheck	[no test files]
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/eventcheck/heavycheck	[no test files]
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/eventcheck/parentscheck	[no test files]
ok  	github.com/Ncog-Earth-Chain/go-ncogearthchain/evmcore	6.322s
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/gossip	[no test files]
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/gossip/emitter	[no test files]
ok  	github.com/Ncog-Earth-Chain/go-ncogearthchain/gossip/filters	1.250s
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/gossip/gasprice	[no test files]
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/gossip/occuredtxs	[no test files]
?   	github.com/Ncog-Earth-Chain/go-ncogearthchain/gossip/piecefunc	[no test files]
ok  	github.com/Ncog-Earth-Chain/go-ncogearthchain/integration	21.640s
```

Also it is tested with [fuzzing](./FUZZING.md).


### Operating a private network (fakenet)

Fakenet is a private network optimized for your private testing.
It'll generate a genesis containing N validators with equal stakes.
To launch a validator in this network, all you need to do is specify a validator ID you're willing to launch.

Pay attention that validator's private keys are deterministically generated in this network, so you must use it only for private testing.

Maintaining your own private network is more involved as a lot of configurations taken for
granted in the official networks need to be manually set up.

To run the fakenet with just one validator (which will work practically as a PoA blockchain), use:
```shell
$ ncogearthchain --fakenet 1/1
```

To run the fakenet with 5 validators, run the command for each validator:
```shell
$ ncogearthchain --fakenet 1/5 # first node, use 2/5 for second node
```

If you have to launch a non-validator node in fakenet, use 0 as ID:
```shell
$ ncogearthchain --fakenet 0/5
```

After that, you have to connect your nodes. Either connect them statically or specify a bootnode:
```shell
$ ncogearthchain --fakenet 1/5 --bootnodes "enode://verylonghex@1.2.3.4:5050"
```

### Running the demo

For the testing purposes, the full demo may be launched using:
```shell
cd demo/
./start.sh # start the Ncogearthchain processes
./stop.sh # stop the demo
./clean.sh # erase the chain data
```
