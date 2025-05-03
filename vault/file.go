package vault

import (
	"encoding/json"
	"errors"
	"os"
)

func LoadVault(path string, key []byte) (Vault, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return NewVault(), nil // File doesn't exist → return empty vault
		}
		return nil, err
	}
	decrypted, err := Decrypt(string(data), key)
	if err != nil {
		return nil, err
	}
	var v Vault
	err = json.Unmarshal([]byte(decrypted), &v)
	if err != nil {
		return nil, err
	}
	return v, err
}
func SaveVault(v Vault, path string, key []byte) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	encrypted, err := Encrypt(data, key)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(encrypted), 0644)
}
