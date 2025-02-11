package main

import (
	"encoding/csv"
	"os"
)

type CSVDataHanddler struct {
	FilePath string
	Records  [][]string
}

func (c *CSVDataHanddler) GetData() error {
	file, err := os.Open(c.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	c.Records = records
	return nil
}

func (c *CSVDataHanddler) SetDataLocation(dataLocation string) {
	c.FilePath = dataLocation
}

func (c *CSVDataHanddler) GetDataLocation() string {
	return c.FilePath
}

func (c *CSVDataHanddler) GetRecords() [][]string {
	return c.Records
}
