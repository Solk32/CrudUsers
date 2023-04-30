package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.HandleFunc("/solk", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello solk!")
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
