package main

type QuizzSetupInterface interface {
	GetQuizzSetup() (string, int)
}

type DataHanddlerInterface interface {
	GetData() error
	SetFilePath(filepath string)
	GetFilePath() string
	GetRecords() [][]string
}
