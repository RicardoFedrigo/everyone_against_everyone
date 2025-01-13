package main

import (
	"fmt"
)

type Answer struct {
	ID   int
	Text string
}

func findAnswerByID(answers []Answer, id int) *Answer {
	for _, answer := range answers {
		if answer.ID == id {
			return &answer
		}
	}
	return nil
}

func main() {
	answers := []Answer{
		{ID: 1, Text: "Yes"},
		{ID: 2, Text: "No"},
		{ID: 3, Text: "Maybe"},
	}

	answer := findAnswerByID(answers, 2)
	if answer != nil {
		fmt.Printf("Found answer: %s\n", answer.Text)
	} else {
		fmt.Println("Answer not found")
	}
}
