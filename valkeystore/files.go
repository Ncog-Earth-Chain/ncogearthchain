package valkeystore

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Ncog-Earth-Chain/go-ncogearthchain/inter/validatorpk"
	"github.com/Ncog-Earth-Chain/go-ncogearthchain/valkeystore/encryption"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrNotFound      = errors.New("key is not found")
	ErrAlreadyExists = errors.New("key already exists")
)

type FileKeystore struct {
	enc *encryption.Keystore
	dir string
}

func NewFileKeystore(dir string, enc *encryption.Keystore) *FileKeystore {
	return &FileKeystore{
		enc: enc,
		dir: dir,
	}
}

func (f *FileKeystore) Has(pubkey validatorpk.PubKey) bool {
	filePath := f.PathOf(pubkey)
	exists := fileExists(filePath)
	//fmt.Printf("Checking existence of file: %s, Exists: %v\n", filePath, exists)
	return exists
}

func (f *FileKeystore) Add(pubkey validatorpk.PubKey, key []byte, auth string) error {
	if f.Has(pubkey) {
		return ErrAlreadyExists
	}
	return f.enc.StoreKey(f.PathOf(pubkey), pubkey, key, auth)
}

/* func (f *FileKeystore) Get(pubkey validatorpk.PubKey, auth string) (*encryption.PrivateKey, error) {
	if !f.Has(pubkey) {
		return nil, ErrNotFound
	}

	fmt.Println("Get Pubkey", pubkey)

	// Convert the static string to a validatorpk.PubKey
	staticPubKey := validatorpk.PubKey{
		Raw: []byte("5aa2ebfffaed1884d0ead7903b5b5a97214034fcca7b2ddefc6c05ff09a083a202bed01e4f250f324e82896b088d71572053132b3f14530464abb3314ef4a1e8"),
	}

	return f.enc.ReadKey(pubkey, f.PathOf(staticPubKey), auth)
} */

// func (f *FileKeystore) PathOf(publicKey validatorpk.PubKey) string {
// 	// Generate a shorter file name
// 	fileName := generateFileName(publicKey.Raw)
// 	fmt.Println("PathOf", path.Join(f.dir, fileName))
// 	return path.Join(f.dir, fileName)
// }

// func (f *FileKeystore) PathOf(pubkey validatorpk.PubKey) string {
// 	fmt.Println("common.Bytes2Hex(pubkey.Bytes())", common.Bytes2Hex(pubkey.Bytes()))
// 	return path.Join(f.dir, common.Bytes2Hex(pubkey.Bytes()))
// }

/* func generateFileName(publicKey []byte) string {
	// Use a hash of the public key to create a shorter and unique file name
	hash := sha256.Sum256(publicKey)
	return fmt.Sprintf("%x", hash[:16]) // Use the first 16 bytes of the hash for brevity
} */
/* func generateFileName(publicKey []byte) string {
	hash := cryptod.Keccak512Hash(publicKey)
	return fmt.Sprintf("%x", hash[:]) // Use 32 bytes (256 bits)
} */

/* func generateFileName(publicKey []byte) string {
	// Use a hash of the public key to create a shorter and unique file name
	hash := cryptod.Keccak512(publicKey)
	return fmt.Sprintf("%x", hash[:64]) // Use the first 64 bytes of the hash for brevity
} */

func generateFileName(publicKey []byte) string {
	return common.Bytes2Hex(publicKey[:32]) // Use the first 32 bytes for brevity
}

/* func generateFileName(publicKey []byte) string {
	// Return a static hex string as the file name
	return "5aa2ebfffaed1884d0ead7903b5b5a97214034fcca7b2ddefc6c05ff09a083a202bed01e4f250f324e82896b088d71572053132b3f14530464abb3314ef4a1e8"
} */

/* func (f *FileKeystore) PathOf(publicKey validatorpk.PubKey) string {
	// Generate a shorter file name
	fileName := generateFileName(publicKey.Raw)
	fmt.Println("PathOf", path.Join(f.dir, fileName))
	return path.Join(f.dir, fileName)
} */

////////

/* func (f *FileKeystore) Get(pubkey validatorpk.PubKey, auth string) (*encryption.PrivateKey, error) {
	if !f.Has(pubkey) {
		return nil, ErrNotFound
	}
	fmt.Println("Get Pubkey", pubkey)

	return f.enc.ReadKey(pubkey, f.PathOf(pubkey), auth)
} */

func (f *FileKeystore) Get(pubkey validatorpk.PubKey, auth string) (*encryption.PrivateKey, error) {
	if !f.Has(pubkey) {
		return nil, ErrNotFound
	}
	return f.enc.ReadKey(pubkey, f.PathOf(pubkey), auth)
}

/* func (f *FileKeystore) PathOf(pubkey validatorpk.PubKey) string {

	fmt.Println("PathOfaaa", path.Join(f.dir, common.Bytes2Hex(pubkey.Bytes())))

	return path.Join(f.dir, common.Bytes2Hex(pubkey.Bytes()))
}
*/

func (f *FileKeystore) PathOf(pubkey validatorpk.PubKey) string {
	if pubkey.Type == validatorpk.Types.MLDsa87 {
		return path.Join(f.dir, generateFileName(pubkey.Raw))
	}
	return path.Join(f.dir, common.Bytes2Hex(pubkey.Bytes()))
}

////////

/* func (f *FileKeystore) PathOf(pubkey validatorpk.PubKey) string {
	// Compute a SHA-256 hash of the public key bytes
	hash := cryptod.Keccak512(pubkey.Bytes())

	// Use the hash as the filename
	return path.Join(f.dir, fmt.Sprintf("%x", hash))
} */

/* func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
} */

func fileExists(filename string) bool {
	if filename == "" {
		fmt.Println("Error: fileExists called with empty filename")
		return false
	}
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Printf("File does not exist: %s\n", filename)
		return false
	}
	if err != nil {
		fmt.Printf("Error checking file: %s, error: %v\n", filename, err)
		return false
	}
	return !info.IsDir()
}
