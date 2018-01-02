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

	router := mux.NewRouter()
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	router.HandleFunc("/posts", CreatePost).Methods("POST")
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
		//fmt.Printf("%#v\n", retrievedPost)
	}

	json.NewEncoder(w).Encode(postCollection)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p Post
	err := decoder.Decode(&p)

	if err != nil {
		log.Fatalln(err)
	}
	statement, _ := database.Prepare("INSERT INTO posts (postbody) values(?)")
	statement.Exec(p.PostBody)
	fmt.Printf("%#v\n", p)
}

type Post struct {
	ID			string		`json:"id,omitempty"`
	PostedAt	time.Time	`json:"timestamp,omitempty"`
	PostBody	string		`json:"body,omitempty"`
}

var posts []Post
