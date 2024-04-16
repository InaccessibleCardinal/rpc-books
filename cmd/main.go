package main

import (
	"fmt"
	"os"
	"rpc-books/internal/env"
	"rpc-books/internal/service/excel"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	env.LoadEnv(f)

	RunExcelThings()
}

func RunExcelThings() {
	opener := excelize.OpenFile
	svc := excel.New(os.Getenv("BOOKS_SOURCE"), os.Getenv("BOOKS_SHEET"), opener)

	table, err := svc.GetTable()
	if err != nil {
		panic(err)
	}

	fmt.Println(len(table))
}
