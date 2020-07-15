package scraper

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type xlsx struct {
	f            *excelize.File
	currentRow   int
	currentSheet string
}

func newXlsx() *xlsx {
	f := excelize.NewFile()
	return &xlsx{
		f:            f,
		currentRow:   1,
		currentSheet: f.GetSheetName(f.GetActiveSheetIndex()),
	}
}

func (x *xlsx) WriteNextRow(values ...interface{}) error {
	for i, value := range values {
		col := i + 1
		axis, err := excelize.CoordinatesToCellName(col, x.currentRow)
		if err != nil {
			return fmt.Errorf("covert cell name error %v", err)
		}

		err = x.f.SetCellValue(x.currentSheet, axis, value)
		if err != nil {
			return fmt.Errorf("set cell value error %v", err)
		}
	}

	x.currentRow++
	return nil
}

func (x *xlsx) Save(filename string) error {
	err := x.f.SaveAs(filename)
	if err != nil {
		return fmt.Errorf("save as error %v", err)
	}

	return nil
}
