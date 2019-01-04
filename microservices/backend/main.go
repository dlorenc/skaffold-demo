package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

const (
	addr = ":8000"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()
		msg := v.Get("message")

		fmt.Printf("Encrypting message: %s\n", msg)
		result := encrypt(msg)
		fmt.Printf("Got: %s\n", result)
		fmt.Fprintln(w, result)
	})

	fmt.Printf("Starting super-secure encryption server on %s\n", addr)
	http.ListenAndServe(addr, nil)
}

// This is not a real encryption function. It is a joke. Please do not use in production.
func encrypt(msg string) string {
	result := ""
	for i := 0; i < len(msg); i++ {
		result = string(msg[i]) + result
	}
	return result
}

var key = []byte("secretthathappenstobe32charslong")

func reallyEncrypt(msg string) string {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	encrypted := gcm.Seal(nonce, nonce, []byte(msg), nil)
	return string(base64.StdEncoding.EncodeToString(encrypted))
}
