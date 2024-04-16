package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Cell struct {
	Text string
	Hyperlink string
}

type ExcelService struct{
	Opener ExcelOpener
	FilePath string
	Sheet string
	table []map[string]Cell
}

func New(filePath string, sheet string, opener ExcelOpener) *ExcelService {
	return &ExcelService{Opener: opener, FilePath: filePath, Sheet: sheet, table: make([]map[string]Cell, 0)}
}

func (ex *ExcelService) GetTable() ([]map[string]Cell, error) {
	if len(ex.table) == 0 {
		f, err := ex.Opener(ex.FilePath)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		rows, err := f.GetRows(ex.Sheet)
		if err != nil {
			return nil, err
		}
		table := ex.makeTable(f, rows)
		return table, nil
	}
	return ex.table, nil
}

func (ex *ExcelService) makeTable(f *excelize.File, rows [][]string) []map[string]Cell {
	headers := rows[0]
	var table []map[string]Cell
	for i, row := range rows[1:] {
		mp := make(map[string]Cell)
		for j, h := range headers {
			cell := Cell{Text: row[j]}
			cellName := fmt.Sprintf("%s%d", columns[j], i+2)
			hyperlink := ex.getCellHyperlink(f, cellName)
			if hyperlink != "" {
				cell.Hyperlink = hyperlink
			}
			mp[h] = cell
		}
		table = append(table, mp)
	}
	ex.table = table
	return table
}

func findBookIndex(books []map[string]Cell, title string) int {
	for i, book := range books {
		if book["Title"].Text == title {
			return i
		}
	}
	return -1
}


func (ex *ExcelService) getCellHyperlink(f *excelize.File, cell string) string {
	ok, h, err := f.GetCellHyperLink(ex.Sheet, cell)
	if err != nil {
		return ""
	}
	if ok {
		return h
	}
	return ""
}

func (ex *ExcelService) AddBook(book map[string]Cell) int {
	ex.table = append(ex.table, book)
	return len(ex.table)
}

func (ex *ExcelService) DeleteBook(title string) (didSucceed bool) {
	didSucceed = true
	booksTable, err := ex.GetTable()
	if err != nil {
		return !didSucceed
	}
	idx := findBookIndex(booksTable, title)
	if idx < 0 {
		return !didSucceed
	}
	books1 := booksTable[:idx]
	books1 = append(books1, booksTable[idx+1:]...)
	ex.table = books1
	return didSucceed
}