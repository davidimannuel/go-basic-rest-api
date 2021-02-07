package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Server starting...")

	http.HandleFunc("/", HomeHandle)
	http.HandleFunc("/articles", GetAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
