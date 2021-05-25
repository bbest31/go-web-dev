package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	// Use the bcrypt pkg to encrypt user passwords
	password := "mypassword123"
	bs := []byte(password)

	encrypted, err := bcrypt.GenerateFromPassword(bs, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error encrypting password!", err)
	}
	fmt.Println("[]Byte = ", bs)
	fmt.Println("Encrypted []byte: ", encrypted)
}
