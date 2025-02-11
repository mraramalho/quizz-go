package main

import (
	"encoding/csv"
	"os"
)

type CSVDataHanddler struct {
	filePath string
	records  [][]string
}

func (c *CSVDataHanddler) getData() error {
	file, err := os.Open(c.filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	c.records = records
	return nil
}