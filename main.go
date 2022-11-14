package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tomoki-yamamura/practice-api/handlers"
)

type Comment struct {
	CommentID int
	ArticleID int
	Message string
	CratedAt time.Time
}

type Article struct {
	ID int
	Title string
	Contents string
	UserName string
	NiceNum int
	CommentList []Comment
	CratedAt time.Time
}

func main() {
	comment1 := Comment{
		CommentID: 1,
		ArticleID: 1,
		Message: "test comment1",
		CratedAt: time.Now(),
	}
	comment2 := Comment{
		CommentID: 2,
		ArticleID: 1,
		Message: "second comment1",
		CratedAt: time.Now(),
	}
	article := Article{
		ID: 1,
		Title: "first article",
		Contents: "This is the test article.",
		UserName: "saki",
		NiceNum: 1,
		CommentList: []Comment{comment1, comment2},
		CratedAt: time.Now(),
	}

	// fmt.Printf("%+v\n", article)
	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", jsonData)

	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/article/list", handlers.ArticleListHandler)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
