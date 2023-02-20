package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Message  string `json:"message"`
	Hostname string `json:"hostname"`
	Version  string `json:"version"`
	Language string `json:"language"`
}

func main() {
	router := mux.NewRouter()
	path := os.Getenv("ROUTE_PATH")
	msg := os.Getenv("MESSAGE")
	port := os.Getenv("PORT")
	hostname, _ := os.Hostname()
	version := "v1.0.1"
	language := "Go"

	if path == "" {
		path = "/api/v1/bar"
	}

	if msg == "" {
		msg = "Hello bar"
	}

	if port == "" {
		port = "3001"
	}

	message := Message{
		Message:  msg,
		Hostname: hostname,
		Version:  version,
		Language: language,
	}

	router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message)
	})
	log.Println("API is running")
	log.Printf("Hostname: %v", hostname)
	log.Printf("Path: %v", path)
	log.Printf("Port: %v", port)
	log.Printf("Message: %v", msg)
	log.Printf("Language: %v", language)
	http.ListenAndServe(":"+port, router)
}
