package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

const addr = ":8080"

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
	}

	fmt.Println("Server strating on port", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
