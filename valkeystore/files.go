package valkeystore

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Ncog-Earth-Chain/go-ncogearthchain/inter/validatorpk"
	"github.com/Ncog-Earth-Chain/go-ncogearthchain/valkeystore/encryption"
	"github.com/ethereum/go-ethereum/cryptod"
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
	return fileExists(f.PathOf(pubkey))
}

func (f *FileKeystore) Add(pubkey validatorpk.PubKey, key []byte, auth string) error {
	if f.Has(pubkey) {
		return ErrAlreadyExists
	}
	return f.enc.StoreKey(f.PathOf(pubkey), pubkey, key, auth)
}

func (f *FileKeystore) Get(pubkey validatorpk.PubKey, auth string) (*encryption.PrivateKey, error) {
	if !f.Has(pubkey) {
		return nil, ErrNotFound
	}
	return f.enc.ReadKey(pubkey, f.PathOf(pubkey), auth)
}

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

func generateFileName(publicKey []byte) string {
	// Use a hash of the public key to create a shorter and unique file name
	hash := cryptod.Keccak512(publicKey)
	return fmt.Sprintf("%x", hash[:64]) // Use the first 64 bytes of the hash for brevity
}

func (f *FileKeystore) PathOf(publicKey validatorpk.PubKey) string {
	// Generate a shorter file name
	fileName := generateFileName(publicKey.Raw)
	return path.Join(f.dir, fileName)
}

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
		return false
	}
	if err != nil {
		fmt.Printf("Error checking file: %s, error: %v\n", filename, err)
		return false
	}
	return !info.IsDir()
}
