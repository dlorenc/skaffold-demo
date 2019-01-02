package main

import (
	"fmt"
	"net/http"
)

const (
	startingEmoji = '🐰'
	addr          = ":8080"
)

func main() {
	i := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%c HELLO WORLD", startingEmoji+i)
		i = i + 1
		fmt.Printf("Next emoji: %c\n", startingEmoji+i)
	})

	fmt.Printf("Starting emoji server on %s\n", addr)
	http.ListenAndServe(addr, nil)
}
