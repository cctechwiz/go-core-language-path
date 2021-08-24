package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!")) //Writer is expecting a slice of bytes
	})

	err := http.ListenAndServe(":3000", nil) //nil specifies to use the built-in handler
	if err != nil {
		log.Fatal(err)
	}
}
