package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":8080"

var Articles []Article

func main() {
	fmt.Print("Server starting on port" + PORT + "...")

	Articles = []Article{
		Article{ID: "1", Title: "Hello 1", Desc: "Article Description 1", Content: "Article Content 1"},
		Article{ID: "2", Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
		Article{ID: "3", Title: "Hello 3", Desc: "Article Description 3", Content: "Article Content 3"},
		Article{ID: "4", Title: "Hello 4", Desc: "Article Description 4", Content: "Article Content 4"},
	}

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandle)
	r.HandleFunc("/articles", getAllArticles)
	r.HandleFunc("/articles/{ID}", getArticleByID)
	r.HandleFunc("/article/create", createArticle)
	log.Fatal(http.ListenAndServe(PORT, r))

}
