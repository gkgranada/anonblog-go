package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	//"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
)

var database *sqlx.DB

func main() {

	database, _ = sqlx.Open("sqlite3", "./anonblog.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY, postedat DATETIME DEFAULT CURRENT_TIMESTAMP, postbody TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO posts (postbody) values (?)")
	statement.Exec("this is message 1")
	statement.Exec("this is message 2")
	statement.Exec("this is message 3")

	router := mux.NewRouter()
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	postCollection := []Post{}
	retrievedPost := Post{}
	rows, err := database.Queryx("SELECT * from posts")
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		err := rows.StructScan(&retrievedPost)
		if err != nil {
			log.Fatalln(err)
		}
		postCollection = append(postCollection, retrievedPost)
		fmt.Printf("%#v\n", retrievedPost)
	}

	json.NewEncoder(w).Encode(postCollection)
}

type Post struct {
	ID			string		`json:"id,omitempty"`
	PostedAt	time.Time	`json:"timestamp,omitempty"`
	PostBody	string		`json:"body,omitempty"`
}

var posts []Post
