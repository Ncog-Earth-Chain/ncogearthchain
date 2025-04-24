package launcher

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/cloudflare/circl/sign/mldsa/mldsa87"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/cryptod"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	cli "gopkg.in/urfave/cli.v1"

	"github.com/Ncog-Earth-Chain/ncogearthchain/inter/validatorpk"
	"github.com/Ncog-Earth-Chain/ncogearthchain/valkeystore"
)

// addFakeValidatorKey adds a fake validator key to the keystore.
func addFakeValidatorKey(ctx *cli.Context, key interface{}, pubkey validatorpk.PubKey, valKeystore valkeystore.RawKeystoreI) {
	if key == nil || valKeystore.Has(pubkey) {
		return
	}

	var err error
	switch k := key.(type) {
	case *mldsa87.PrivateKey:
		// Add MLDSA87 private key
		keyBytes := cryptod.FromMLDsa87(k)
		err = valKeystore.Add(pubkey, keyBytes, validatorpk.FakePassword)
	case *ecdsa.PrivateKey:
		// Add ECDSA private key for backward compatibility
		keyBytes := crypto.FromECDSA(k)
		err = valKeystore.Add(pubkey, keyBytes, validatorpk.FakePassword)
	default:
		utils.Fatalf("Unsupported key type for validator key: %T", key)
	}

	if err != nil {
		utils.Fatalf("Failed to add fake validator key: %v", err)
	}
}

// getValKeystoreDir returns the directory path for validator keystore.
func getValKeystoreDir(cfg node.Config) string {
	_, _, keydir, err := cfg.AccountConfig()
	if err != nil {
		utils.Fatalf("Failed to setup account config: %v", err)
	}
	return keydir
}

// makeValidatorPasswordList reads password lines from the file specified by the global --validator.password flag.
func makeValidatorPasswordList(ctx *cli.Context) []string {
	if path := ctx.GlobalString(validatorPasswordFlag.Name); path != "" {
		text, err := ioutil.ReadFile(path)
		if err != nil {
			utils.Fatalf("Failed to read password file: %v", err)
		}
		lines := strings.Split(string(text), "\n")
		// Sanitise DOS line endings.
		for i := range lines {
			lines[i] = strings.TrimRight(lines[i], "\r")
		}
		return lines
	}
	if ctx.GlobalIsSet(FakeNetFlag.Name) {
		return []string{validatorpk.FakePassword}
	}
	return nil
}

// unlockValidatorKey unlocks the validator key in the keystore.
// old 18-apr
func unlockValidatorKey(ctx *cli.Context, pubKey validatorpk.PubKey, valKeystore valkeystore.KeystoreI) error {
	var err error
	for trials := 0; trials < 3; trials++ {
		prompt := fmt.Sprintf("Unlocking validator key %s | Attempt %d/%d", pubKey.String(), trials+1, 3)
		password := getPassPhrase(prompt, false, 0, makeValidatorPasswordList(ctx))
		err = valKeystore.Unlock(pubKey, password)
		if err == nil {
			log.Info("Unlocked validator key", "pubkey", pubKey.String())
			return nil
		}
	}
	// All trials expended to unlock account, bail out
	return err
}

// unlockValidatorKey unlocks the validator key in the keystore.
/* func unlockValidatorKey(ctx *cli.Context, pubKey validatorpk.PubKey, valKeystore valkeystore.KeystoreI) error {
	var err error
	for trials := 0; trials < 3; trials++ {
		prompt := fmt.Sprintf("Unlocking validator key %s | Attempt %d/%d", pubKey.String(), trials+1, 3)
		password := getPassPhrase(prompt, false, 0, makeValidatorPasswordList(ctx))
		publicKeyFile := validatorpk.PubKey{
			Type: pubKey.Type,
			Raw:  pubKey.Bytes()[1:33],
		}

		err = valKeystore.Unlock(publicKeyFile, password)
		if err == nil {
			log.Info("Unlocked validator key", "pubkey", publicKeyFile)
			return nil
		}
	}
	// All trials expended to unlock account, bail out
	return err
} */
