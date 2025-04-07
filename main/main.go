package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq" // postgres
	"log"
	"net/http"
	"os"
	"rest-api/config"
)

type post struct {
	Id     int
	Author string
	Text   string
}

type Match struct {
	Id          uint
	Game        string
	Description string
}

func main() {
	fmt.Println("Server start")

	config := config.GetConfig()
	mux := http.NewServeMux()
	db := psqlConnect(config.DB)

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		getHttpMethodHandlerFunction(w, r, db)
	})

	log.Printf("Server now running on port: %v", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.PORT), mux))
}

func handleGet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	data, err := db.Query("select * from matches")

	var matches []Match
	for data.Next() {
		var match Match
		err = data.Scan(&match.Id, &match.Game, &match.Description)

		if err != nil {
			fmt.Printf("Error Looping data, and %s", err)
		}

		matches = append(matches, match)
	}

	if err != nil {
		fmt.Println("Cant query for some reason")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
	fmt.Println("The endpoint %m", r.Method)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	newPost := post{
		Id:     1,
		Author: "Me",
		Text:   "Doing a post",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPost)
	fmt.Println("The endpoint %m", r.Method)
}

func getHttpMethodHandlerFunction(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r, db)
	case http.MethodPost:
		handlePost(w, r)
	case http.MethodPut:
	case http.MethodDelete:
	default:
	}
}

func psqlConnect(connString string) *sql.DB {
	result, err := sql.Open("postgres", connString)

	if err != nil {
		log.Fatal("Can't connect to postgress", err)
	}

	return result
}
