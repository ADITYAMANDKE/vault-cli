package vault

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower   = "abcdefghijklmnopqrstuvwxyz"
	digits  = "0123456789"
	special = "!@#$%^&*()-_+=[]{}|;:,.<>?/~"
)

func GenerateRandomPassword(length int) (string, error) {
	// Ensure minimum length is 8 characters
	if length < 8 {
		return "", fmt.Errorf("password length must be at least 8 characters")
	}

	// Combine uppercase, lowercase, digits, and special characters
	allChars := upper + lower + digits + special
	var password strings.Builder

	// Randomly pick characters from allChars
	for i := 0; i < length; i++ {
		charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		if err != nil {
			return "", err
		}
		password.WriteByte(allChars[charIndex.Int64()])
	}

	return password.String(), nil
}
