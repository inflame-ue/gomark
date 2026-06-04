package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/inflame-ue/gomark/internal/db"
	"github.com/inflame-ue/gomark/internal/grammar"
	"github.com/inflame-ue/gomark/internal/notes"
	"github.com/joho/godotenv"
)

func databaseMiddleWare(db *db.DB, handler func(w http.ResponseWriter, r *http.Request, db *db.DB),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, db)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the .env file")
	}

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /grammar", grammar.HandleGrammarCheck)
	mux.HandleFunc("POST /notes", databaseMiddleWare(db, notes.HandlePostNote))

	port := os.Getenv("SERVER_PORT")
	serv := http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Printf("server started on %v\n", port)
	err = serv.ListenAndServe()
	log.Fatal(err)
}
