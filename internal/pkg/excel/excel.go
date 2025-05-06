// Package excel provides function to export Book data to excel file.
package excel

import (
	"strconv"

	excelize "github.com/xuri/excelize/v2"

	"BookShelf/internal/app/entity"
)

const (
	_sheetName = "Sheet1"
)

// Save excel file with Books content to file saveToFile.
func ExportBook(saveToFile string, books entity.Books) error {
	file := excelize.NewFile()

	// set custom width for columns
	file.SetColWidth(_sheetName, "A", "F", 30)

	// set header
	file.SetCellValue(_sheetName, "A1", "Title")
	file.SetCellValue(_sheetName, "B1", "Genre")
	file.SetCellValue(_sheetName, "C1", "Author")
	file.SetCellValue(_sheetName, "D1", "Year")
	file.SetCellValue(_sheetName, "E1", "Description")
	file.SetCellValue(_sheetName, "F1", "Status")

	// set Books data
	for idx := 0; idx < len(books); idx++ {
		rowIdx := strconv.Itoa(idx + 2)
		file.SetCellValue(_sheetName, "A"+rowIdx, books[idx].Title)
		if books[idx].Genre.Name != "" {
			file.SetCellValue(_sheetName, "B"+rowIdx, books[idx].Genre.Name)
		}
		if books[idx].Author != nil {
			file.SetCellValue(_sheetName, "C"+rowIdx, *books[idx].Author)
		}
		if books[idx].Year != nil {
			file.SetCellValue(_sheetName, "D"+rowIdx, *books[idx].Year)
		}
		if books[idx].Description != nil {
			file.SetCellValue(_sheetName, "E"+rowIdx, *books[idx].Description)
		}
		if books[idx].IsRead {
			file.SetCellValue(_sheetName, "F"+rowIdx, "Read")
		} else {
			file.SetCellValue(_sheetName, "F"+rowIdx, "Want")
		}
	}

	// save to file
	err := file.SaveAs(saveToFile)
	return err // err or nil
}
