package filemanager

import (
	"encoding/csv"
	"os"
)

func CsvReader() ([][]string, error) {

	file, err := os.Open("quizfile/quiz.csv")
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil

}
