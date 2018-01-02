package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	database, _ := sql.Open("sqlite3", "./jodelgo.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY, postedat DATETIME DEFAULT CURRENT_TIMESTAMP, postbody TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO posts (postbody) values (?)")
	statement.Exec("this is message 1")
	statement.Exec("this is message 2")
	statement.Exec("this is message 3")
	
	comments = append(comments, Comment{ID:"1", PostID:"1", CommentedAt:time.Now(), CommentBody:"cool commment"})
	comments = append(comments, Comment{ID:"2", PostID:"3", CommentedAt:time.Now(), CommentBody:"upvoting"})
	comments = append(comments, Comment{ID:"3", PostID:"3", CommentedAt:time.Now(), CommentBody:"spam commment"})

	router := mux.NewRouter()
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	//router.HandleFunc("/posts/{id}/comments", posts.showHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}

//func (posts *Post) showHandler(w http.ResponseWriter, r *http.Request) {
//	ideally this would retrieve items from a database


//}

type Post struct {
	ID			string		`json:"id,omitempty"`
	PostedAt	time.Time	`json:"timestamp,omitempty"`
	PostBody	string		`json:"body,omitempty"`
}

var posts []Post

type Comment struct {
	ID			string		`json:"id,omitempty"`
	PostID		string		`json:"postid,omitempty"`
	CommentedAt	time.Time	`json:"timestamp,omitempty"`
	CommentBody	string		`json:"body,omitempty"`
}

var comments []Comment