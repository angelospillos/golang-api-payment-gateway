package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	// Database Configuration
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Routes Configuration
	router := setupRouter(db)

	// Server Configuration
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
