package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

// KEY thanks to https://www.allkeysgenerator.com/Random/Security-Encryption-Key-Generator.aspx
func encryptTool(user User, db *sql.DB) {
	option := optionSelect()

	if option == 1 {
		encrypt()
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 2 {
		decrypt()
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 3 {
		time.Sleep(2 * time.Second)
		selectTool(user, db)
	} else {
		falseOptionFunc(user, db)
	}
}
func optionSelect() int {
	fmt.Println("-----------------------")
	fmt.Println("Please choose an option")
	fmt.Println("1. Encrypt")
	fmt.Println("2. Decrypt")
	fmt.Println("3. Return to the options")
	fmt.Println("You'll need a key, see ENV file")
	fmt.Println("Don't have a key, get one here: https://www.allkeysgenerator.com/Random/Security-Encryption-Key-Generator.aspx")
	fmt.Println("-----------------------")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an option: ")
	scanner.Scan()
	optionString := scanner.Text()
	option, err := strconv.Atoi(optionString)
	checkError(err)
	fmt.Println("-----------------------")
	return option
}

func encrypt() {
	fmt.Println("Please choose a file to encrypt")
	var file string
	fmt.Scanln(&file)

	fileExist := fileExist(file)
	if !fileExist {
		fmt.Println("File doesn't exist")
		time.Sleep(2 * time.Second)
	} else {

		plaintext, err := ioutil.ReadFile(file)
		checkError(err)
		key := os.Getenv("KEY")
		checkError(err)
		block, err := aes.NewCipher([]byte(key))
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
		time.Sleep(2 * time.Second)
		fmt.Println("Succesfully created ciphertext.bin")
	}
}

func decrypt() {
	fmt.Println("Now decrypting")
	fmt.Println("There has to be a file named ciphertext.bin")
	filename := "ciphertext.bin"
	fileExist := fileExist(filename)
	if fileExist {
		ciphertext, err := ioutil.ReadFile(filename)
		checkError(err)

		key := os.Getenv("KEY")
		checkError(err)
		block, err := aes.NewCipher([]byte(key))
		checkError(err)
		gcm, err := cipher.NewGCM(block)
		checkError(err)

		nonce := ciphertext[:gcm.NonceSize()]
		ciphertext = ciphertext[gcm.NonceSize():]
		plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
		checkError(err)
		err = ioutil.WriteFile("plaintext.txt", plaintext, 0777)
		checkError(err)
		time.Sleep(2 * time.Second)
		fmt.Println("Succesfully created plaintext.txt")
	} else {
		fmt.Println("File doesn't exist.")
	}
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false

	} else {
		return true
	}
}
