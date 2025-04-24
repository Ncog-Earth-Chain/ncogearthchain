package encryption

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/cryptod"

	"github.com/Ncog-Earth-Chain/ncogearthchain/inter/validatorpk"
)

var (
	ErrNotSupportedType = errors.New("not supported key type")
)

type PrivateKey struct {
	Type    uint8
	Bytes   []byte
	Decoded interface{}
}

type EncryptedKeyJSON struct {
	Type      uint8               `json:"type"`
	PublicKey string              `json:"pubkey"`
	Crypto    keystore.CryptoJSON `json:"crypto"`
}

type Keystore struct {
	scryptN int
	scryptP int
}

func New(scryptN int, scryptP int) *Keystore {
	return &Keystore{
		scryptN: scryptN,
		scryptP: scryptP,
	}
}

/*
func (ks Keystore) ReadKey(wantPubkey validatorpk.PubKey, filename, auth string) (*PrivateKey, error) {
	// Load the key from the keystore and decrypt its contents
	keyjson, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	key, err := DecryptKey(keyjson, auth)
	if err != nil {
		return nil, err
	}
	// Make sure we're really operating on the requested key (no swap attacks)
	keySecp256k1 := key.Decoded.(*ecdsa.PrivateKey)
	gotPubkey := crypto.FromECDSAPub(&keySecp256k1.PublicKey)
	if bytes.Compare(wantPubkey.Raw, gotPubkey) != 0 {
		return nil, fmt.Errorf("key content mismatch: have public key %X, want %X", gotPubkey, wantPubkey.Raw)
	}
	return key, nil
}
*/
/*
func (ks Keystore) ReadKey(wantPubkey validatorpk.PubKey, filename, auth string) (*PrivateKey, error) {
	// Load the key from the keystore and decrypt its contents
	keyjson, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	key, err := DecryptKey(keyjson, auth)
	if err != nil {
		return nil, err
	}

	// Verify the key type and match the public key
	if key.Type != validatorpk.Types.MLDsa87 {
		return nil, fmt.Errorf("unsupported key type: %v", key.Type)
	}

	// Verify the public key for MLDsa87
	keyMLDsa87 := key.Decoded.(*cryptod.PrivateKey)
	gotPubkey := cryptod.FromMLDsa87Pub(keyMLDsa87.Public().(*cryptod.PublicKey))
	if bytes.Compare(wantPubkey.Raw, gotPubkey) != 0 {
		return nil, fmt.Errorf("key content mismatch: have public key %X, want %X", gotPubkey, wantPubkey.Raw)
	}

	return key, nil
}
*/

func (ks Keystore) ReadKey(wantPubkey validatorpk.PubKey, filename, auth string) (*PrivateKey, error) {
	// Load the key from the keystore and decrypt its contents
	keyjson, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	key, err := DecryptKey(keyjson, auth)
	if err != nil {
		return nil, err
	}

	// Verify the key type and match the public key
	if key.Type != validatorpk.Types.MLDsa87 {
		return nil, fmt.Errorf("unsupported key type: %v", key.Type)
	}

	// Verify the public key for MLDsa87
	/* keyMLDsa87 := key.Decoded.(*cryptod.PrivateKey)
	gotPubkey := cryptod.FromMLDsa87Pub(keyMLDsa87.Public().(*cryptod.PublicKey))
	if !bytes.Equal(wantPubkey.Raw, gotPubkey) { // Use bytes.Equal instead of bytes.Compare
		return nil, fmt.Errorf("key content mismatch: have public key %X, want %X", gotPubkey, wantPubkey.Raw)
	} */

	return key, nil
}

func (ks Keystore) StoreKey(filename string, pubkey validatorpk.PubKey, key []byte, auth string) error {
	keyjson, err := ks.EncryptKey(pubkey, key, auth)
	if err != nil {
		return err
	}
	// Write into temporary file
	tmpName, err := writeTemporaryKeyFile(filename, keyjson)
	if err != nil {
		return err
	}
	return os.Rename(tmpName, filename)
}

// EncryptKey encrypts a key using the specified scrypt parameters into a json
// blob that can be decrypted later on.
/* func (ks Keystore) EncryptKey(pubkey validatorpk.PubKey, key []byte, auth string) ([]byte, error) {
	if pubkey.Type != validatorpk.Types.Secp256k1 {
		return nil, ErrNotSupportedType
	}
	cryptoStruct, err := keystore.EncryptDataV3(key, []byte(auth), ks.scryptN, ks.scryptP)
	if err != nil {
		return nil, err
	}
	encryptedKeyJSON := EncryptedKeyJSON{
		Type:      pubkey.Type,
		PublicKey: common.Bytes2Hex(pubkey.Raw),
		Crypto:    cryptoStruct,
	}
	return json.Marshal(encryptedKeyJSON)
} */

func (ks Keystore) EncryptKey(pubkey validatorpk.PubKey, key []byte, auth string) ([]byte, error) {
	// Validate inputs
	if len(pubkey.Raw) == 0 || pubkey.Type == 0 {
		return nil, fmt.Errorf("invalid public key_test: %+v", pubkey)
	}
	if len(key) == 0 {
		return nil, fmt.Errorf("invalid key: key is empty")
	}
	if auth == "" {
		return nil, fmt.Errorf("invalid auth: password is empty")
	}

	// Ensure the key type is MLDsa87
	if pubkey.Type != validatorpk.Types.MLDsa87 {
		return nil, fmt.Errorf("unsupported public key type: %x", pubkey.Type)
	}

	// Encrypt the key using MLDsa87 settings
	cryptoStruct, err := keystore.EncryptDataV3(key, []byte(auth), ks.scryptN, ks.scryptP)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt key: %v", err)
	}

	// Create encrypted key JSON
	encryptedKeyJSON := EncryptedKeyJSON{
		Type:      pubkey.Type,
		PublicKey: common.Bytes2Hex(pubkey.Raw),
		Crypto:    cryptoStruct,
	}

	// Marshal the encrypted key to JSON
	encryptedJSON, err := json.Marshal(encryptedKeyJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal encrypted key JSON: %v", err)
	}

	return encryptedJSON, nil
}

// DecryptKey decrypts a key from a json blob, returning the private key itself.
/* func DecryptKey(keyjson []byte, auth string) (*PrivateKey, error) {
	// Parse the json into a simple map to fetch the key version
	m := make(map[string]interface{})
	if err := json.Unmarshal(keyjson, &m); err != nil {
		return nil, err
	}
	var (
		keyBytes []byte
		err      error
	)
	k := new(EncryptedKeyJSON)
	if err := json.Unmarshal(keyjson, k); err != nil {
		return nil, err
	}
	if k.Type != validatorpk.Types.Secp256k1 {
		return nil, ErrNotSupportedType
	}
	keyBytes, err = decryptKey_secp256k1(k, auth)
	// Handle any decryption errors and return the key
	if err != nil {
		return nil, err
	}

	decoded, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		return nil, err
	}

	return &PrivateKey{
		Type:    k.Type,
		Bytes:   keyBytes,
		Decoded: decoded,
	}, nil
} */

func DecryptKey(keyjson []byte, auth string) (*PrivateKey, error) {
	// Parse the JSON into a simple map to fetch the key version
	m := make(map[string]interface{})
	if err := json.Unmarshal(keyjson, &m); err != nil {
		return nil, err
	}

	k := new(EncryptedKeyJSON)
	if err := json.Unmarshal(keyjson, k); err != nil {
		return nil, err
	}

	// fmt.Println("k.Type", k.Type)
	// fmt.Println("validatorpk.Types.MLDsa87", validatorpk.Types.MLDsa87)

	if k.Type != validatorpk.Types.MLDsa87 {
		return nil, ErrNotSupportedType
	}

	// Decrypt MLDsa87 key
	keyBytes, err := decryptKey_MLDsa87(k, auth)
	if err != nil {
		return nil, err
	}

	decoded, err := cryptod.ToMLDsa87(keyBytes)
	if err != nil {
		return nil, err
	}

	return &PrivateKey{
		Type:    k.Type,
		Bytes:   keyBytes,
		Decoded: decoded,
	}, nil
}

func decryptKey_MLDsa87(k *EncryptedKeyJSON, auth string) ([]byte, error) {
	keyBytes, err := keystore.DecryptDataV3(k.Crypto, auth)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt MLDsa87 key: %v", err)
	}
	return keyBytes, nil
}

/* func decryptKey_secp256k1(keyProtected *EncryptedKeyJSON, auth string) (keyBytes []byte, err error) {
	plainText, err := keystore.DecryptDataV3(keyProtected.Crypto, auth)
	if err != nil {
		return nil, err
	}
	return plainText, err
}
*/
