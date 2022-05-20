package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func encryptTool(user User, db *sql.DB) {
	option := optionSelect()

	if option == 1 {
		encrypt()
	} else if option == 2 {
		decrypt()
	} else {
		falseOptionFunc(user, db)
	}
}
func optionSelect() int {
	fmt.Println("Please choose an option")
	fmt.Println("1. Encrypt")
	fmt.Println("2. Decrypt")
	var option int
	fmt.Scanln(&option)

	return option
}

func encrypt() {
	fmt.Println("Please choose a file")
	var file string
	fmt.Scanln(&file)
	plaintext, err := ioutil.ReadFile("plaintext.txt")
	checkError(err)

	key, err := ioutil.ReadFile("key")
	checkError(err)
	block, err := aes.NewCipher(key)
	checkError(err)

	gcm, err := cipher.NewGCM(block)
	checkError(err)

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	// Save back to file
	err = ioutil.WriteFile("ciphertext.bin", ciphertext, 0777)
	checkError(err)
}

func decrypt() {
	fmt.Println("Now decrypting")
	ciphertext, err := ioutil.ReadFile("ciphertext.bin")
	checkError(err)

	key, err := ioutil.ReadFile("key")
	checkError(err)

	block, err := aes.NewCipher(key)
	checkError(err)
	gcm, err := cipher.NewGCM(block)
	checkError(err)

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	checkError(err)

	err = ioutil.WriteFile("plaintext.txt", plaintext, 0777)
	checkError(err)
}
