package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"math/big"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "-g":
		password := os.Args[2]
		hash := GenerarHash(password)
		fmt.Printf("String: %s\nHash: %s\n", password, hash)

	case "-v":
		if len(os.Args) < 4 {
			printUsage()
			os.Exit(1)
		}
		password := os.Args[2]
		hash := os.Args[3]
		if ValidarPassword(password, hash) {
			fmt.Println("✓ Valid Password")
		} else {
			fmt.Println("✗ Invalid Password")
			os.Exit(1)
		}

	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  hashpass -g <password>")
	fmt.Println("  hashpass -v <password> <hash>")
}

func GenerarHash(password string) string {
	return GetHash(password, nil)
}

func GetHash(password string, salt []byte) string {
	if salt == nil {
		minSaltSize := 4
		maxSaltSize := 8
		saltSizeBig, _ := rand.Int(rand.Reader, big.NewInt(int64(maxSaltSize-minSaltSize)))
		saltSize := int(saltSizeBig.Int64()) + minSaltSize
		salt = make([]byte, saltSize)
		for i := range salt {
			for salt[i] == 0 {
				b := make([]byte, 1)
				rand.Read(b)
				salt[i] = b[0]
			}
		}
	}

	passwordBytes := []byte(password)
	passwordSalt := append(passwordBytes, salt...)

	hash := sha512.Sum512(passwordSalt)
	hashSalt := append(hash[:], salt...)

	return base64.StdEncoding.EncodeToString(hashSalt)
}

func ValidarPassword(password, hash string) bool {
	hashSaltBytes, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false
	}

	hashSize := 512 / 8
	if len(hashSaltBytes) < hashSize {
		return false
	}

	salt := hashSaltBytes[hashSize:]
	expectedHash := GetHash(password, salt)

	return hash == expectedHash
}
