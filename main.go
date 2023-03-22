package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

type TextEntry struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/easy")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/api/texts", textsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func textsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rows, err := db.Query("SELECT * FROM texts")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var texts []TextEntry
		for rows.Next() {
			var t TextEntry
			if err := rows.Scan(&t.ID, &t.Content); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			texts = append(texts, t)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(texts)

	case "POST":
		var t TextEntry
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := db.Exec("INSERT INTO texts (content) VALUES (?)", t.Content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := res.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int64{"id": id})
	}
}
