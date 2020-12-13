package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))

	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)

	if err != nil {
		log.Fatal(err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err.Error())
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	return plaintext
}

func encryptToDst(dst string, data []byte, key string) {
	file, err := os.Create(dst)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	file.Write(encrypt(data, key))
}

func decryptToDst(dst, src, key string) {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err.Error())
	}

	decryptedData := decrypt(data, key)
	if err := ioutil.WriteFile(dst, decryptedData, 0666); err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	if len(os.Args) < 5 {
		log.Fatal("You need to provide more os arguments | format: src dst key action")
	}

	srcDir := os.Args[1]
	dstDir := os.Args[2]
	encKey := os.Args[3]
	action := os.Args[4]

	// Validate the input
	if srcDir == "" {
		log.Fatal("You need to provide a -src flag")
	}

	if dstDir == "" {
		log.Fatal("You need to provide a -dst flag")
	}

	if encKey == "" {
		log.Fatal("You need to provide an encryption key")
	}

	if !(action == "encrypt" || action == "decrypt") {
		log.Fatal("You need to provide a ")
	}

	// Handle the actions
	if action == "encrypt" {
		srcData, err := ioutil.ReadFile(srcDir)
		if err != nil {
			log.Fatal("Could not read file", err)
		}

		encryptToDst(dstDir, srcData, encKey)
	} else {
		decryptToDst(dstDir, srcDir, encKey)
	}
}
