package main

import (
	"fmt"
	"log"
	"net/http"
)

func gameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Da Vinci Code game server is up and running!")
}

func main() {
	http.HandleFunc("/game", gameHandler)
	log.Println("Starting Da Vinci Code game server on :8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
