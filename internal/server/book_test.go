package server

import (
	"context"
	"errors"
	bookv1 "rpc-books/gen/book/v1"
	"rpc-books/internal/service/excel"
	"testing"

	"connectrpc.com/connect"
)

var (
	testBooks []map[string]excel.Cell
	testError error
	testBookService = New(TestExcelHandler{})
)

type TestExcelHandler struct {}

func (te TestExcelHandler) GetTable() ([]map[string]excel.Cell, error) {
	return testBooks, testError
}

func TestGetBookByTitleSuccess(t *testing.T) {
	testBooks = []map[string]excel.Cell{
		{
			"Title": excel.Cell{Hyperlink: "abook1.com", Text: "a book 1"},
			"Author": excel.Cell{Hyperlink: "", Text: "Author1"},
		},
		{
			"Title": excel.Cell{Hyperlink: "abook2.com", Text: "a book 2"},
			"Author": excel.Cell{Hyperlink: "", Text: "Author2"},
		},
	}

	
	res, err := testBookService.GetBooksByTitle(
		context.Background(), 
		connect.NewRequest(&bookv1.GetBooksByTitleRequest{Title: "a book 2"}))

	if err != nil {
		t.Errorf("got %s but expected nil error\n", err)
	}

	if res.Msg.Book.Title != "a book 2" {
		t.Errorf("got %s but expected %s\n", res.Msg.Book.Title, "a book 2")
	}
	testBooks = nil
}

func TestGetBookByTitleDBError(t *testing.T) {
	testError = errors.New("lol wut")

	_, err := testBookService.GetBooksByTitle(
		context.Background(), 
		connect.NewRequest(&bookv1.GetBooksByTitleRequest{Title: "a book 2"}))
	
	if err == nil {
		t.Error("expected non-nil error")
	}

	if err.Error() != "lol wut" {
		t.Errorf("got %s but expected 'lol wut'\n", err.Error())
	}
	testError = nil
}

func TestGetBookByTitleNotFound(t *testing.T) {
	testBooks = []map[string]excel.Cell{
		{
			"Title": excel.Cell{Hyperlink: "abook1.com", Text: "a book 1"},
			"Author": excel.Cell{Hyperlink: "", Text: "Author1"},
		},
		{
			"Title": excel.Cell{Hyperlink: "abook2.com", Text: "a book 2"},
			"Author": excel.Cell{Hyperlink: "", Text: "Author2"},
		},
	}

	_, err := testBookService.GetBooksByTitle(
		context.Background(), 
		connect.NewRequest(&bookv1.GetBooksByTitleRequest{Title: "a book 3"}))
	
	if err == nil {
		t.Error("expected non-nil error")
	}

	if err.Error() != "book not found" {
		t.Errorf("got %s but expected %s\n", err.Error(), "book not found")
	}
}
