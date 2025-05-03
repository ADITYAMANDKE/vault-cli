package vault

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

func Encrypt(data []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}
	stream := cipher.NewCBCEncrypter(block, iv)
	paddedData := pad(data)
	ciphertext := make([]byte, len(paddedData))
	stream.CryptBlocks(ciphertext, paddedData)
	encryptedData := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(encryptedData), nil

}
func Decrypt(encryptedData string, key []byte) ([]byte, error) {
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	if len(encryptedBytes) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := encryptedBytes[:aes.BlockSize]
	ciphertext := encryptedBytes[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCBCDecrypter(block, iv)
	decryptedData := make([]byte, len(ciphertext))
	stream.CryptBlocks(decryptedData, ciphertext)

	// Remove padding
	return unpad(decryptedData), nil
}

func pad(data []byte) []byte {
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func unpad(data []byte) []byte {
	padding := data[len(data)-1]
	return data[:len(data)-int(padding)]
}
