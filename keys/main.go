package main

import (
	"crypto/rand"
	"encoding/hex"

	// "encoding/base64"
	"fmt"
)

func GenerateKey(length int) string {
	buf := make([]byte, length/2)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err) // out of randomness, should never happen
	}

	//return fmt.Sprintf("%x", buf)
	return hex.EncodeToString(buf)
	// or base64.StdEncoding.EncodeToString(buf)
}

func main() {
	key := GenerateKey(3)
	fmt.Println(key)
}
