package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateHMACKey() ([]byte, error) {
	key := make([]byte, 32) // Adjust the key size as per your requirements
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func main() {
	key, err := generateHMACKey()
	if err != nil {
		fmt.Println("Error generating HMAC key:", err)
		return
	}

	encodedKey := base64.StdEncoding.EncodeToString(key)
	fmt.Println("Generated HMAC key:", encodedKey)
}
