package excel

import (
	"errors"
	"reflect"
	"testing"

	"github.com/xuri/excelize/v2"
)

func testOpener(filename string, opts ...excelize.Options) (*excelize.File, error) {
	switch filename {
	case "good.xlsx":
		return createExcelFile(), nil
	case "bad.xlsx":
		return nil, errors.New("lol")
	case "books.xlsx":
		return createBooksFile(), nil
	default:
		return nil, errors.New("lol")
	}
}

func createExcelFile() *excelize.File {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "name")
	f.SetCellValue("Sheet1", "B1", "job")
	f.SetCellValue("Sheet1", "A2", "ken")
	f.SetCellValue("Sheet1", "B2", "programmer geek")
	f.SetCellValue("Sheet1", "A3", "jen")
	f.SetCellValue("Sheet1", "B3", "qa")
	display, tooltip := "https://github.com", "GitHub"
	f.SetCellHyperLink("Sheet1", "B2", display, "External", excelize.HyperlinkOpts{Display: &display, Tooltip: &tooltip})
	return f
}

func createBooksFile() *excelize.File {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Title")
	f.SetCellValue("Sheet1", "B1", "Author")
	f.SetCellValue("Sheet1", "A2", "some book")
	f.SetCellValue("Sheet1", "B2", "some author")
	f.SetCellValue("Sheet1", "A3", "some other book")
	f.SetCellValue("Sheet1", "B3", "some author")
	return f
}

func TestGetTableError(t *testing.T) {
	exSvc := New("bad.xlsx", "Sheet1", testOpener)
	_, err := exSvc.GetTable()
	if err == nil {
		t.Error("expected non-nil error")
	}
	if err.Error() != "lol" {
		t.Errorf("got %s but expected 'lol'\n", err.Error())
	}
}

func TestGetTableSuccess(t *testing.T) {
	exSvc := New("good.xlsx", "Sheet1", testOpener)
	table, err := exSvc.GetTable()
	if err != nil {
		t.Errorf("expected nil error but got %s\n", err)
	}
	actualRow := table[0]
	if actualRow["name"].Text != "ken" {
		t.Errorf("got %v but expected %v\n", actualRow["name"].Text, "ken")
	}
	if actualRow["job"].Hyperlink != "https://github.com" {
		t.Errorf("got %v but expected %v\n", actualRow["job"].Hyperlink, "https://github.com")
	}
}

func Test_getCellHyperlink(t *testing.T) {
	exSvc := New("good.xlsx", "Sheet1", testOpener)

	testFile := createExcelFile()

	value := exSvc.getCellHyperlink(testFile, "lol")
	if value != "" {
		t.Errorf("got %s but expected an empty string", value)
	}
}

func TestAddBook(t *testing.T) {
	exSvc := New("good.xlsx", "Sheet1", testOpener)
	expectedBook := map[string]Cell{
		"Title": {Hyperlink: "test.com", Text: "test book"},
		"Author": {Hyperlink: "", Text: "test author"}}
	exSvc.GetTable()
	result := exSvc.AddBook(expectedBook)
	if result != 3 {
		t.Errorf("got %d but expected 3\n", result)
	}

	booksTable, err := exSvc.GetTable()
	if err != nil {
		t.Errorf("got %s but expected a nil error\n", err)
	}

	actualBook := booksTable[len(booksTable)-1]
	if !reflect.DeepEqual(actualBook, expectedBook) {
		t.Errorf("got actual: %v, expected: %v\n", actualBook, expectedBook)
	}
}

func TestDeleteBook(t *testing.T) {
	exSvc := New("books.xlsx", "Sheet1", testOpener)
	results, _ := exSvc.GetTable()

	if len(results) != 2 {
		t.Errorf("got %d but expected 2\n", len(results))
	}
	titleToDelete := "some other book"

	didSucceed := exSvc.DeleteBook(titleToDelete)
	if !didSucceed {
		t.Error("expected success")
	}

	results, _ = exSvc.GetTable()

	if len(results) != 1 {
		t.Errorf("got %d but expected 1\n", len(results))
	}

	titleCantDelete := "lol wut"

	didSucceed = exSvc.DeleteBook(titleCantDelete)

	if didSucceed {
		t.Error("expected failure")
	}

}