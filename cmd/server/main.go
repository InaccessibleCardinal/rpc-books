package main

import (
	"log"
	"net/http"

	"rpc-books/gen/book/v1/bookv1connect"
	"rpc-books/internal/server"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	bookServer := &server.BookServer{}
	mux := http.NewServeMux()
	path, handler := bookv1connect.NewBookServiceHandler(bookServer)
	mux.Handle(path, handler)
	if err := http.ListenAndServe("localhost:9999", h2c.NewHandler(mux, &http2.Server{})); err != nil {
		log.Fatalf("failed to serve with error %s\n", err)
	}
}