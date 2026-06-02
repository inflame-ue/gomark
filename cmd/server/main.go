package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the .env file")
	}

	mux := http.NewServeMux()

	port := os.Getenv("PORT")
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
