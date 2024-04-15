package server

import (
	"context"
	"errors"
	bookv1 "rpc-books/gen/book/v1"
	"rpc-books/gen/book/v1/bookv1connect"
	"rpc-books/internal/service/excel"

	"connectrpc.com/connect"
)

type ExcelHandler interface {
	GetTable() ([]map[string]excel.Cell, error)
}

type BookServer struct {
	bookv1connect.UnimplementedBookServiceHandler

	excelService ExcelHandler
}

func New(excelService ExcelHandler) *BookServer {
	
	return &BookServer{excelService: excelService}
}

func (b *BookServer) GetBooksByTitle(
	ctx context.Context,
	req *connect.Request[bookv1.GetBooksByTitleRequest]) (*connect.Response[bookv1.GetBooksByTitleResponse], error) {
	title := req.Msg.Title
	book, err := b.getBookByTitleFromDB(title)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&bookv1.GetBooksByTitleResponse{Book: book}), nil
}

func (b *BookServer) getBookByTitleFromDB(title string) (*bookv1.Book, error) {
	table, err := b.excelService.GetTable()
	if err != nil {
		return nil, err
	}
	for _, rawBook := range table {
		if rawBook["Title"].Text == title {
			return &bookv1.Book{Title: rawBook["Title"].Text, Author: rawBook["Author"].Text}, nil
		}
	}
	return nil, errors.New("book not found")
}
