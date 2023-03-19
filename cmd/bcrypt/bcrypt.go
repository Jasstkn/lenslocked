package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	switch os.Args[1] {
	case "compare":
		err := Compare(os.Args[2], os.Args[3])
		if err != nil {
			panic(err)
		}
		fmt.Println("password matches hash")
	case "hash":
		hash, err := Generate(os.Args[2])
		if err != nil {
			panic(err)
		}
		fmt.Println(hash)
	default:
		fmt.Printf("Invalid command: %v\n", os.Args[1])
		fmt.Printf("Usage: %v [compare|hash] [password|hash]\n", os.Args[0])
		os.Exit(1)
	}
}

func Generate(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Compare(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
