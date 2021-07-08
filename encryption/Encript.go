package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"log"
)

/*
Accepts input data in the form of a byte array and a key string, which is typically a secret passphrase.
Returns the encrypted data.
*/
func encrypt(data []byte, key string) (resp []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return resp, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return resp, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return resp, err
	}
	return gcm.Seal(resp, nonce, data, []byte("test")), nil
}

/*
Decrypt data. It accept the encrypted data in the form of a byte array and the passphrase as a string.
Return the decrypted data.
*/
func decrypt(data []byte, key string) (resp []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return resp, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return resp, err
	}
	ciphertext := data[gcm.NonceSize():]
	nonce := data[:gcm.NonceSize()]
	resp, err = gcm.Open(nil, nonce, ciphertext, []byte("test"))
	if err != nil {
		return resp, fmt.Errorf("error decrypting data: %v", err)
	}
	return resp, nil
}

func main() {
	const key = "mysecurepassword"
	encrypted, err := encrypt([]byte("Hello World!"), key)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Encrypted Text: ", string(encrypted))

	// raise exception: "cipher: message authentication failed" - why?
	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Decrypted Text: ", string(decrypted))
}
