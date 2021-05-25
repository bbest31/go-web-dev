package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getCode("test@example.com")
	fmt.Println(c)
	c = getCode("test@example.com")
	fmt.Println(c)
}

func getCode(s string) string {
	//hmac creates a hash based on user info.
	// It creates the same hash if given identical info
	// This allows us to verify that he information underneath has not been altered.
	h := hmac.New(sha256.New, []byte("ourkey")) // uses private key to hash
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
