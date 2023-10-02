package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const addr = ":8080"

func Init() {
	r := newRouter()

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
