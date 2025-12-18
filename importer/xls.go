package importer

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

type ExcelImporter struct {
	Filename   string
	HasHeaders bool
	xls        *excelize.File
	worksheet  string
	headers    map[string]string
}

type Coords struct {
	row int
	col int
}

func (i *ExcelImporter) Import() error {
	if err := i.loadFile(); err != nil {
		return err
	}

	return nil
}

func (i *ExcelImporter) loadFile() error {
	if i.xls != nil {
		return nil
	}

	if i.Filename == "" {
		return fmt.Errorf("Please provide filename!")
	}
	var err error
	i.xls, err = excelize.OpenFile(i.Filename, excelize.Options{})
	if err != nil {
		return err
	}

	activeSheet := i.xls.GetSheetName(0)
	rows, err := i.xls.GetRows(activeSheet)

	if err != nil {
		return err
	}

	var firstData *Coords
	for r, row := range rows {
		for c, col := range row {
			if len(col) > 0 {
				firstData = &Coords{row: r, col: c}
				break
			}
		}
		if firstData != nil {
			break
		}
	}

	if firstData == nil {
		return fmt.Errorf("Sheets cotains no data")
	}

	fmt.Println(firstData)
	fmt.Println(rows[firstData.row][firstData.col])

	os.Exit(0)

	return nil
}

func (i *ExcelImporter) Filter(filterFunc func(r Record) bool) (Records, error) {
	return nil, nil
}

func (i *ExcelImporter) FilterColumns() ([]string, error) {
	return []string{}, nil
}

func (i *ExcelImporter) Worksheets() ([]string, error) {
	if err := i.loadFile(); err != nil {
		return nil, err
	}

	return i.xls.GetSheetList(), nil
}
