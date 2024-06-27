package models

import (
	"go-quiz/filemanager"
	"log"
)

type Quiz struct {
	Question string
	Answer   string
}

func LoadQuizData() ([]Quiz, error) {

	records, err := filemanager.CsvReader()
	if err != nil {
		return nil, err
	}

	var quizContent []Quiz
	for _, record := range records {
		if len(record) != 2 {
			log.Printf("skipping invalid Record: %v", record)
			continue
		}
		quiz := Quiz{
			Question: record[0],
			Answer:   record[1],
		}

		quizContent = append(quizContent, quiz)
	}

	return quizContent, nil

}
