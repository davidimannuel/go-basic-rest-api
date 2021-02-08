package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func getArticleByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit:getArticleByID")
	vars := mux.Vars(r)
	id := vars["ID"]
	for _, article := range Articles {
		if article.ID == id {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit:Create Article")

	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(reqBody)
}
