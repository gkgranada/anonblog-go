package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

func main() {

	posts = append(posts, Post{ID:"1", PostedAt:time.Now(), PostBody:"this is jodel1"})

	router := mux.NewRouter()
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}

type Post struct {
	ID			string		`json:"id,omitempty"`
	PostedAt	time.Time	`json:"timestamp,omitempty"`
	PostBody	string		`json:"body,omitempty"`
}

var posts []Post
