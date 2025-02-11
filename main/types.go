package main

type QuizzSetupInterface interface {
	GetQuizzSetup() (string, int)
}

type DataHanddlerInterface interface {
	GetData() error
	SetDataLocation(filepath string)
	GetDataLocation() string
	GetRecords() [][]string
}
