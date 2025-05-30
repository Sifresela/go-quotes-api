package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewStore()
	router := makeRouter(store)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
