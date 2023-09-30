package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const addr = ":8080"

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	hub := newHub()
	go hub.run()
	r.GET("/", func(c *gin.Context) {
		serveWs(hub, c)
	})

	server := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: 3 * time.Second,
	}

	fmt.Println("Server strating on port", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
