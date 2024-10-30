package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func createSessionToken(userID int64, email string) string {
	// Combine the userID, email, and current timestamp to create a unique token string
	data := fmt.Sprintf("%d:%s:%d", userID, email, time.Now().UnixNano())

	// Create a SHA-256 hash of the data
	hash := sha256.New()
	hash.Write([]byte(data))

	// Convert the hash to a hex string to get the final token
	token := hex.EncodeToString(hash.Sum(nil))
	return token
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return ""
	}

	return string(hash)
}

func compareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
