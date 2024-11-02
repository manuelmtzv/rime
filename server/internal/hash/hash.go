package hash

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	memory := uint32(64 * 1024)
	time := uint32(3)
	threads := uint8(2)
	keyLen := uint32(32)

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	saltHash := append(salt, hash...)
	encodedHash := base64.RawStdEncoding.EncodeToString(saltHash)

	return encodedHash, nil
}

func VerifyPassword(password, encodedHash string) (bool, error) {
	saltHash, err := base64.RawStdEncoding.DecodeString(encodedHash)
	if err != nil {
		return false, err
	}

	if len(saltHash) < 16 {
		return false, errors.New("decoded hash too short")
	}

	salt, hash := saltHash[:16], saltHash[16:]

	memory := uint32(64 * 1024)
	time := uint32(3)
	threads := uint8(2)
	keyLen := uint32(len(hash))

	computedHash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	return sha256.Sum256(hash) == sha256.Sum256(computedHash), nil
}
