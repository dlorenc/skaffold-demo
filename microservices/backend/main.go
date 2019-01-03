package main

import (
	"fmt"
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
