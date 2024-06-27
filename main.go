package main

import (
	"fmt"
	"go-quiz/models"
	"log"
)

type Quiz struct {
	Question string
	Answer   string
}

func main() {

	fmt.Println("Welcome to Quiz game")
	fmt.Println("Hit Enter to start : ")
	fmt.Scan()

	quizContent, err := models.LoadQuizData()
	if err != nil {
		log.Fatalf(err.Error())
	}

	var score int64 = 0
	for index, val := range quizContent {
		var answer string
		fmt.Printf("question %v : %s ", index+1, val.Question)
		fmt.Scanln(&answer)
		if answer == val.Answer {
			score++
		}
	}

	fmt.Printf("\n your score out of %v is : %v", len(quizContent), score)

}
