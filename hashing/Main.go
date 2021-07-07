package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
)

func getHash(str string, hashType string) string {
	switch hashType {
	case "MD5":
		return fmt.Sprintf("%x", md5.Sum([]byte(str)))
	case "SHA256":
		return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
	case "SHA512":
		return fmt.Sprintf("%x", sha512.Sum512([]byte(str)))
		// from "golang.org/x/crypto/blake2s"
	case "BLAKE2s_256":
		return fmt.Sprintf("%x", blake2s.Sum256([]byte(str)))
		// from "golang.org/x/crypto/blake2b"
	case "BLAKE2b_512":
		return fmt.Sprintf("%x", blake2b.Sum512([]byte(str)))
	default:
		return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
	}
}

func main() {
	fmt.Println("MD5:", getHash("Hello World!", "MD5"))
	fmt.Println("SHA256:", getHash("Hello World!", "SHA256"))
	fmt.Println("SHA512:", getHash("Hello World!", "SHA512"))
	fmt.Println("SHA3_512:", getHash("Hello World!", "SHA3_512"))
	fmt.Println("BLAKE2s_256:", getHash("Hello World!", "BLAKE2s_256"))
	fmt.Println("BLAKE2b_512:", getHash("Hello World!", "BLAKE2b_512"))
}
