package main

import (
	"fmt"
	"go-quiz/models"
	"log"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

func main() {

	fmt.Println("Welcome to Quiz game")
	fmt.Print("Hit Enter to start... ")
	fmt.Scanln()

	quizContent, err := models.LoadQuizData()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("you have in total 10 seconds time for all questions ")

	answerCh := make(chan string)

	timer := time.NewTimer(10 * time.Second)

	// go routine to handle user input
	go func() {
		for index, val := range quizContent {
			var answer string
			fmt.Printf("question %v : %s ", index+1, val.Question)
			fmt.Scanln(&answer)
			answerCh <- answer
		}
	}()

	// Loop to process user answers and handle timer expiration
	var score int64
	for i := 0; i < len(quizContent); i++ {
		select {
		case <-timer.C:
			fmt.Println("\n Time is Up ...")
			fmt.Printf("Your score out of %v is : %v\n", len(quizContent), score)
			return
		case answer := <-answerCh:
			if answer == quizContent[i].Answer {
				score++
			}
		}
	}

	// Stop the timer explicitly if all questions are answered within time
	if !timer.Stop() {
		<-timer.C
	}

	fmt.Printf("\n your score out of %v is : %v", len(quizContent), score)

}
