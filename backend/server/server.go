package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

const addr = ":8080"

var server *http.Server

func Serve() {
	r := newRouter()

	server = &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: 3 * time.Second,
	}

	fmt.Println("Server strating on port", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Shutdown(ctx context.Context) error {
	err := server.Shutdown(ctx)

	if err != nil {
		return err
	}

	return nil
}
