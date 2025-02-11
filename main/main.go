package main

import (
	"fmt"
	"os"
	"strings"
)

var quizzSetup QuizzSetupInterface = &CMDQuizzSetup{
	DefaultFilePath:  "./padrao.csv",
	DefaultTimeLimit: 30,
}

var dataHandler DataHanddlerInterface = &CSVDataHanddler{
	FilePath: "",
	Records:  [][]string{},
}

var score Score = Score{}

func main() {
	dataLocation, timeLimit := quizzSetup.GetQuizzSetup()
	dataHandler.SetFilePath(dataLocation)
	fmt.Println(dataHandler.GetFilePath())
	err := dataHandler.GetData()
	check(err)
	records := dataHandler.GetRecords()
	score.totalQuestions = len(records)

	fmt.Println("Welcome to QuizzGame!\nAnswer the questions below to score points:")
	go timer(timeLimit, &score)
	for _, record := range records {
		question := record[0]
		answer := strings.TrimSpace(record[1])
		userInput, err := getUserAnswer(question)
		check(err)

		if checkAnswer(userInput, answer) {
			score.incrementScore()
		}
	}
	score.showScore()
}

func check(e error) {
	if e != nil {
		fmt.Printf("Erro: %v\n", e)
		os.Exit(1)
	}
}
