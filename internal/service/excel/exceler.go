package excel

import "github.com/xuri/excelize/v2"


type ExcelOpener func(filename string, opts ...excelize.Options) (*excelize.File, error)

type ExcelFile interface {
	GetRows(sheet string, opts ...excelize.Options) ([][]string, error)
}