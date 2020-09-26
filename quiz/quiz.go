package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	quizCSV := flag.String("file", "problems.csv", "Quiz CSV file in 'question,answer' format")
	flag.Parse()

	records := loadCSV(*quizCSV)
	// assume records structure:
	// [
	//  ["question1, "answer1"],
	//  ["question2, "answer2"],
	//  ...
	//  ]

	correctAnswers := 0
	totalQuestions := len(records)
	consoleReader := bufio.NewReader(os.Stdin)

	for _, s := range records {
		fmt.Print(s[0], ": ")

		answer, _ := consoleReader.ReadString('\n')
		answer = strings.ReplaceAll(answer, "\n", "")

		if answer == s[1] {
			correctAnswers++
		}
	}

	fmt.Print("Correct answers: ", correctAnswers, "/", totalQuestions, "\n")
}

func loadCSV(csvFile string) [][]string {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return records
}
