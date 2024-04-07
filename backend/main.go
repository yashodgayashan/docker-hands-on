package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

type Book struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"published_date"`
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/book_store?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getBooks(db, w, r)
		}
	})

	// Setup CORS
	corsHandler := cors.AllowAll()
	handler := corsHandler.Handler(mux)

	fmt.Println("Server is starting on port 8080...")
	http.ListenAndServe(":8080", handler)
}

func getBooks(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var books []Book
	rows, err := db.Query("SELECT id, title, author, published_date FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.PublishedDate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
