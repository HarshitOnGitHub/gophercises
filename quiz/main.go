package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	ques string
	answer string
}

func (p problem) ask() (bool) {
	fmt.Printf("%s ", p.ques)
	var answer string
	fmt.Scanf("%s\n", &answer)
	answer = strings.TrimSpace(answer)
	if answer != p.answer {
		return false
	}
	return true
}

func parseCSV(filename string) (ret []problem) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error opening problems file.")
		os.Exit(1)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	questions, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	}

	ret = make([]problem, len(questions))
	for index, question := range(questions) {
		if len(question) != 2 {
			fmt.Printf("Invalid record.")
			os.Exit(1)
		}
		ret[index].ques = question[0]
		ret[index].answer = strings.TrimSpace(question[1])
	}
	return
}

func main() {
	questions := parseCSV("problems.csv")
	total := len(questions)
	correct := 0
	for _, question := range(questions) {
		if match := question.ask(); match == true {
			correct++
		}
	}
	fmt.Printf("Total: %d Correct: %d\n", total, correct)
}
