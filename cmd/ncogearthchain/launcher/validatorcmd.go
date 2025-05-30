package launcher

import (
	"fmt"
	"path"
	"strings"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/cryptod"
	"gopkg.in/urfave/cli.v1"

	"github.com/Ncog-Earth-Chain/ncogearthchain/inter/validatorpk"
	"github.com/Ncog-Earth-Chain/ncogearthchain/valkeystore"
	"github.com/Ncog-Earth-Chain/ncogearthchain/valkeystore/encryption"
)

var (
	validatorCommand = cli.Command{
		Name:     "validator",
		Usage:    "Manage validators",
		Category: "VALIDATOR COMMANDS",
		Description: `

Create a new validator private key.

It supports interactive mode, when you are prompted for password as well as
non-interactive mode where passwords are supplied via a given password file.
Non-interactive mode is only meant for scripted use on test networks or known
safe environments.

Make sure you remember the password you gave when creating a new validator key.
Without it you are not able to unlock your validator key.

Note that exporting your key in unencrypted format is NOT supported.

Keys are stored under <DATADIR>/keystore/validator.
It is safe to transfer the entire directory or the individual keys therein
between Ncogearthchain nodes by simply copying.

Make sure you backup your keys regularly.`,
		Subcommands: []cli.Command{
			{
				Name:   "new",
				Usage:  "Create a new validator key",
				Action: utils.MigrateFlags(validatorKeyCreate),
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.KeyStoreDirFlag,
					utils.PasswordFileFlag,
				},
				Description: `
    ncogearthchain validator new

Creates a new validator private key and prints the public key.

The key is saved in encrypted format, you are prompted for a passphrase.

You must remember this passphrase to unlock your key in the future.

For non-interactive use the passphrase can be specified with the --validator.password flag:

Note, this is meant to be used for testing only, it is a bad idea to save your
password to file or expose in any other way.
`,
			},
			{
				Name:   "convert",
				Usage:  "Convert an account key to a validator key",
				Action: utils.MigrateFlags(validatorKeyConvert),
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.KeyStoreDirFlag,
				},
				ArgsUsage: "<account address> <validator pubkey>",
				Description: `
    ncogearthchain validator convert

Converts an account private key to a validator private key and saves in the validator keystore.
`,
			},
		},
	}
)

// validatorKeyCreate creates a new validator key into the keystore defined by the CLI flags.
/* func validatorKeyCreate(ctx *cli.Context) error {
	cfg := makeAllConfigs(ctx)
	utils.SetNodeConfig(ctx, &cfg.Node)

	password := getPassPhrase("Your new validator key is locked with a password. Please give a password. Do not forget this password.", true, 0, utils.MakePasswordList(ctx))

	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		utils.Fatalf("Failed to create account: %v", err)
	}
	privateKey := crypto.FromECDSA(privateKeyECDSA)
	publicKey := validatorpk.PubKey{
		Raw:  crypto.FromECDSAPub(&privateKeyECDSA.PublicKey),
		Type: validatorpk.Types.Secp256k1,
	}

	valKeystore := valkeystore.NewDefaultFileRawKeystore(path.Join(getValKeystoreDir(cfg.Node), "validator"))
	err = valKeystore.Add(publicKey, privateKey, password)
	if err != nil {
		utils.Fatalf("Failed to create account: %v", err)
	}

	// Sanity check
	_, err = valKeystore.Get(publicKey, password)
	if err != nil {
		utils.Fatalf("Failed to decrypt the account: %v", err)
	}

	fmt.Printf("\nYour new key was generated\n\n")
	fmt.Printf("Public key:                  %s\n", publicKey.String())
	fmt.Printf("Path of the secret key file: %s\n\n", valKeystore.PathOf(publicKey))
	fmt.Printf("- You can share your public key with anyone. Others need it to validate messages from you.\n")
	fmt.Printf("- You must NEVER share the secret key with anyone! The key controls access to your validator!\n")
	fmt.Printf("- You must BACKUP your key file! Without the key, it's impossible to operate the validator!\n")
	fmt.Printf("- You must REMEMBER your password! Without the password, it's impossible to decrypt the key!\n\n")
	return nil
} */

// validatorKeyCreate creates a new validator key into the keystore defined by the CLI flags.
func validatorKeyCreate(ctx *cli.Context) error {
	cfg := makeAllConfigs(ctx)
	utils.SetNodeConfig(ctx, &cfg.Node)

	password := getPassPhrase("Your new validator key is locked with a password. Please give a password. Do not forget this password.", true, 0, utils.MakePasswordList(ctx))

	// Generate a new MLDSA87 private key
	privateKeyMLDSA87, err := cryptod.GenerateMLDsa87Key()
	if err != nil {
		utils.Fatalf("Failed to create account: %v", err)
	}
	privateKey := cryptod.FromMLDsa87(privateKeyMLDSA87)
	publicKeyRaw := cryptod.FromMLDsa87Pub(privateKeyMLDSA87.Public().(*cryptod.PublicKey))

	// Convert the MLDsa87 private key to a hex string
	/* hexPrivateKey, err := cryptod.MLDsa87ToHex(privateKeyMLDSA87) // Pass the original key object here
	if err != nil {
		fmt.Printf("Error converting MLDsa87 key to hex: %v\n", err)
	} else {
		fmt.Printf("MLDsa87 private key in hex: %s\n", hexPrivateKey)
	} */

	// fmt.Println("privateKey", privateKey)
	// fmt.Println("Length of privateKey:", len(privateKey))

	// fmt.Println("publicKeyRaw", publicKeyRaw)
	// fmt.Println("Length of publicKeyRaw:", len(publicKeyRaw))

	publicKey := validatorpk.PubKey{
		Raw:  publicKeyRaw,
		Type: validatorpk.Types.MLDsa87, // Ensure this type exists in your implementation
	}

	//fmt.Println("publicKey", publicKey.String())

	//fmt.Println("test1", "test 1")

	valKeystore := valkeystore.NewDefaultFileRawKeystore(path.Join(getValKeystoreDir(cfg.Node), "validator"))

	//fmt.Println("test2", "test 2")

	// if valKeystore != nil {
	// 	fmt.Println("valKeystore", valKeystore)
	// }

	err = valKeystore.Add(publicKey, privateKey, password)

	//fmt.Println("test3", "test 3")

	if err != nil {
		//fmt.Println("test5", "test 5")
		utils.Fatalf("Failed to create account: %v", err)
	}

	//fmt.Println("test4", "test 4")

	// fmt.Printf("\nYour new key was generated\n\n")

	// Sanity check
	_, err = valKeystore.Get(publicKey, password)
	if err != nil {
		utils.Fatalf("Failed to decrypt the account: %v", err)
	}

	//fmt.Println("test2", "test 2")

	fmt.Printf("\nYour new key was generated\n\n")
	fmt.Printf("Public key:                  %s\n", publicKey.String())
	fmt.Printf("Path of the secret key file: %s\n\n", valKeystore.PathOf(publicKey))
	fmt.Printf("- You can share your public key with anyone. Others need it to validate messages from you.\n")
	fmt.Printf("- You must NEVER share the secret key with anyone! The key controls access to your validator!\n")
	fmt.Printf("- You must BACKUP your key file! Without the key, it's impossible to operate the validator!\n")
	fmt.Printf("- You must REMEMBER your password! Without the password, it's impossible to decrypt the key!\n\n")
	return nil
}

// validatorKeyConvert converts account key to validator key.
func validatorKeyConvert(ctx *cli.Context) error {
	if len(ctx.Args()) < 2 {
		utils.Fatalf("This command requires 2 arguments.")
	}
	cfg := makeAllConfigs(ctx)
	utils.SetNodeConfig(ctx, &cfg.Node)

	_, _, keydir, _ := cfg.Node.AccountConfig()

	pubkeyStr := ctx.Args().Get(1)
	pubkey, err := validatorpk.FromString(pubkeyStr)
	if err != nil {
		utils.Fatalf("Failed to decode the validator pubkey: %v", err)
	}

	var acckeypath string
	if strings.HasPrefix(ctx.Args().First(), "0x") {
		acckeypath, err = FindAccountKeypath(common.HexToAddress(ctx.Args().First()), keydir)
		if err != nil {
			utils.Fatalf("Failed to find the account: %v", err)
		}
	} else {
		acckeypath = ctx.Args().First()
	}

	valkeypath := path.Join(keydir, "validator", common.Bytes2Hex(pubkey.Bytes()))
	err = encryption.MigrateAccountToValidatorKey(acckeypath, valkeypath, pubkey)
	if err != nil {
		utils.Fatalf("Failed to migrate the account key: %v", err)
	}
	fmt.Println("\nYour key was converted and saved to " + valkeypath)
	return nil
}
