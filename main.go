package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var userInput UserInputInterface = CMDUserInput{
		defaultFilePath:  "./padrao.csv",
		defaultTimeLimit: 30,
	}
	filePath, timeLimit := userInput.getUserSetup()

	dataHandler := CSVDataHanddler{
		filePath: filePath,
		records:  [][]string{},
	}
	err := dataHandler.getData()
	check(err)

	score := Score{
		totalQuestions: len(dataHandler.records),
		correctAnswers: 0,
	}

	fmt.Println("Bem vindo ao quizz!")
	fmt.Println("Responda as questões:")
	go timer(timeLimit, &score)
	for _, record := range dataHandler.records {
		question := record[0]
		answer := strings.TrimSpace(record[1])
		userInput, err := userInput.getUserAnswer(question)
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

func timer(timeLimit int, score *Score) {
	time.Sleep(time.Duration(timeLimit) * time.Second)
	fmt.Println("Tempo esgotado!")
	score.showScore()
	os.Exit(1)
}
