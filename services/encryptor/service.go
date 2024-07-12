package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"tech/domain/services"
)

const key = "285c00252c7d392c80f9cdcd448c2dac"

type service struct {
}

func Make() services.Encryptor {
	return &service{}
}

// Encrypt - encrypts a string with AES and returns the result as a hexadecimal string
func (s *service) Encrypt(in string) (string, error) {
	keyBytes := []byte(key)
	plaintextBytes := []byte(in)

	// Create a new AES cipher using the given key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Create a new AES-GCM cipher with a random nonce
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Generate a random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	// Encrypt the plaintext with the GCM cipher
	ciphertext := gcm.Seal(nonce, nonce, plaintextBytes, nil)

	// Return the ciphertext as a hexadecimal string
	return hex.EncodeToString(ciphertext), nil
}

func (s *service) Decrypt(in string) (string, error) {
	keyBytes := []byte(key)

	// Decode the hexadecimal-encoded ciphertext
	ciphertextBytes, err := hex.DecodeString(in)
	if err != nil {
		return "", err
	}

	// Create a new AES cipher using the given key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Create a new AES-GCM cipher with the same nonce and key as the original encryption
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Split the ciphertext into the nonce and encrypted data
	nonceSize := gcm.NonceSize()
	if len(ciphertextBytes) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	nonce, ciphertextBytes := ciphertextBytes[:nonceSize], ciphertextBytes[nonceSize:]

	// Decrypt the ciphertext with the GCM cipher
	plaintextBytes, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}

	// Return the plaintext as a string
	return string(plaintextBytes), nil
}
