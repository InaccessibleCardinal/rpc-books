package main

import (
	"context"
	"net/http"
	bookv1 "rpc-books/gen/book/v1"
	"rpc-books/gen/book/v1/bookv1connect"
	"testing"

	"connectrpc.com/connect"
)

var (
	addr = "http://localhost:9999"
)

func TestServer(t *testing.T) {
	title := "The Russian Revolution: A New History"

	client := bookv1connect.NewBookServiceClient(http.DefaultClient, addr)

	res, err := client.GetBooksByTitle(
		context.Background(), connect.NewRequest(&bookv1.GetBooksByTitleRequest{Title: title}))
	
	if err != nil {
		t.Errorf("got %s but expected nil error\n", err)
	}

	actual :=  res.Msg.Book.Title
	if actual != title {
		t.Errorf("got %s but expected %s\n", actual, title)
	}
}