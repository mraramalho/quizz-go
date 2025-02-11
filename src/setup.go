package main

import (
	"flag"
)

type CMDQuizzSetup struct {
	DefaultFilePath  string
	DefaultTimeLimit int
}

func (c *CMDQuizzSetup) GetQuizzSetup() (string, int) {
	filePath := flag.String("file", string(c.DefaultFilePath), "a string containg the quizz file path")
	timeLimit := flag.Int("limit", int(c.DefaultTimeLimit), "a int containing the time limit for the quizz")
	flag.Parse()
	return *filePath, *timeLimit
}
