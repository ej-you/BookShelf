package excel

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	excelize "github.com/xuri/excelize/v2"
)

const (
	_defaultSheetName = "Sheet1"
	_rowStart         = 1
	_colStart         = 1
	_colWidth         = 30
)

// Contains struct field name and index.
type structField struct {
	Name string
	Idx  int
}

// FillAndSave is equivalent to FillSheetAndSave("Sheet1", filepath, data).
// func FillAndSave(filepath string, s any) error {
func FillAndSave[T any](filepath string, data []T) error {
	return FillSheetAndSave(_defaultSheetName, filepath, data)
}

// Fills the sheetName Excel sheet with given data and saves
// the file to the given filepath.
// The filepath must end with a file name with the xlsx extension.
// Given data must be a slice of structs. Unexported fields of struct will be skipped.
// func FillSheetAndSave(sheetName, filepath string, s any) error {
func FillSheetAndSave[T any](sheetName, filepath string, data []T) error {
	if len(data) == 0 {
		return errors.New("data slice is empty")
	}
	if !strings.HasSuffix(filepath, ".xlsx") {
		return errors.New("file must have the xlsx extension")
	}
	elemType := reflect.TypeOf(data[0])
	if elemType.Kind() != reflect.Struct {
		return errors.New("data slice must be a slice of structs")
	}

	// slice of exported data elem fields
	fields := make([]structField, 0, elemType.NumField())
	for fieldIdx := range cap(fields) {
		field := elemType.Field(fieldIdx)
		// skip unexported fields
		if field.PkgPath != "" {
			continue
		}
		fields = append(fields, structField{field.Name, fieldIdx})
	}

	file := excelize.NewFile()
	// set custom width for columns
	err := file.SetColWidth(sheetName,
		numberToExcelColumn(_colStart),
		numberToExcelColumn(len(fields)),
		_colWidth)
	if err != nil {
		return fmt.Errorf("set columns width: %w", err)
	}

	// define row index for header
	rowIdx := strconv.Itoa(_rowStart)
	// sets table header according to struct fields
	for colIdx, structField := range fields {
		setCellValue(file, sheetName,
			numberToExcelColumn(_colStart+colIdx)+rowIdx,
			structField.Name)
	}

	var elemValue reflect.Value
	var value any
	// fill table with data
	for idx, elem := range data {
		elemValue = reflect.ValueOf(elem)
		rowIdx = strconv.Itoa(_rowStart + 1 + idx)

		for colIdx, structField := range fields {
			value = elemValue.Field(structField.Idx).Interface()

			setCellValue(file, sheetName,
				numberToExcelColumn(_colStart+colIdx)+rowIdx,
				value)
		}
	}

	// save to file
	err = file.SaveAs(filepath)
	return err // err or nil
}

// setCellValue is a wrapping over file.SetCellValue with error logging.
func setCellValue(file *excelize.File, sheetName, cell string, value any) {
	err := file.SetCellValue(sheetName, cell, value)
	if err != nil {
		log.Printf("[ERROR] Set %s cell value %v: %v", cell, value, err)
	}
}
