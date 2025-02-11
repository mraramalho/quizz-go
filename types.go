package main

type UserInputInterface interface {
	getUserSetup() (string, int)
	getUserAnswer(question string) (string, error)
}
