package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func getUserAnswer(question string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s \n>>> ", question)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(userInput[:len(userInput)-1]), err
}

func checkAnswer(input string, correctAnswer string) bool {
	return input == correctAnswer
}

func timer(timeLimit int, score *Score) {
	time.Sleep(time.Duration(timeLimit) * time.Second)
	fmt.Println("Tempo esgotado!")
	score.showScore()
	os.Exit(0)
}
