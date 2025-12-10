package importer

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CSVImporter struct {
	Filename   string
	HasHeaders bool
	headers    []string
	records    Records
}

func (imp *CSVImporter) Import() error {

	fd, err := os.Open(imp.Filename)
	if err != nil {
		return fmt.Errorf("Error opening CSV file %s\n", imp.Filename)
	}

	csv := csv.NewReader(fd)

	for {
		record, err := csv.Read()
		if err != nil {
			break
		}

		if imp.HasHeaders && imp.headers == nil {
			imp.headers = record
			continue
		}

		recordMap := make(map[string]string)
		for c, header := range imp.headers {
			recordMap[header] = record[c]
		}
		imp.records = append(imp.records, recordMap)

	}
	return nil
}

func (imp *CSVImporter) Filter(filterFunction func(r Record) bool) Records {
	var filteredRecords Records
	for _, r := range imp.records {
		if filterFunction(r) {
			filteredRecords = append(filteredRecords, r)
		}
	}
	return filteredRecords
}

func (imp *CSVImporter) FilterColumns() []string {
	return imp.headers
}
