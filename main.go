package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

func main() {
	router := mux.NewRouter()
	path := os.Getenv("BAR_PATH")
	title := os.Getenv("TITLE")
	port := os.Getenv("PORT")

	if path == "" {
		path = "bar"
	}

	if title == "" {
		title = "Hello bar2"
	}

	if port == "" {
		port = "3001"
	}

	message := Message{
		Title:   title,
		Version: "v1.0.0",
	}

	router.HandleFunc("/"+path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message)
	})
	log.Println("API is running")
	log.Printf("Path: %v", path)
	log.Printf("Port: %v", port)
	log.Printf("Title: %v", title)
	http.ListenAndServe(":"+port, router)
}
