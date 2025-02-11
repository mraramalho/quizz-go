package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type CMDUserInput struct {
	defaultFilePath  string
	defaultTimeLimit int
}

func (c CMDUserInput) getUserSetup() (string, int) {
	filePath := flag.String("file", string(c.defaultFilePath), "a string containg the quizz file path")
	timeLimit := flag.Int("limit", int(c.defaultTimeLimit), "a int containing the time limit for the quizz")
	flag.Parse()
	return *filePath, *timeLimit
}

func (c CMDUserInput) getUserAnswer(question string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s \n", question)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(userInput[:len(userInput)-1]), err
}

func checkAnswer(input string, correctAnswer string) bool {
	return input == correctAnswer
}
