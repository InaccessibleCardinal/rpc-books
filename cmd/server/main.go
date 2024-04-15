package main

import (
	"log"
	"net/http"
	"os"

	"rpc-books/gen/book/v1/bookv1connect"
	"rpc-books/internal/env"
	"rpc-books/internal/server"
	"rpc-books/internal/service/excel"

	"github.com/xuri/excelize/v2"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var (
	addr = "localhost:9999"
)
func init() {
	f, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	env.LoadEnv(f)
}

func main() {
	opener := excelize.OpenFile
	excelHandler := excel.New(os.Getenv("BOOKS_SOURCE"), os.Getenv("BOOKS_SHEET"), opener)

	bookServer := server.New(excelHandler)
	mux := http.NewServeMux()
	path, handler := bookv1connect.NewBookServiceHandler(bookServer)
	mux.Handle(path, handler)
	log.Printf("starting rpc server on %s...\n", addr)
	if err := http.ListenAndServe(addr, h2c.NewHandler(mux, &http2.Server{})); err != nil {
		log.Fatalf("failed to serve with error %s\n", err)
	}
}