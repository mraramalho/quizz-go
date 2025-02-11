package main

import "fmt"

type Score struct {
	totalQuestions int
	correctAnswers int
}

func (s *Score) incrementScore() {
	s.correctAnswers++
}

func (s *Score) showScore() {
	fmt.Printf("You got %d out of %d questions\n", s.correctAnswers, s.totalQuestions)
}
