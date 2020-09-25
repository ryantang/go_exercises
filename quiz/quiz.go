package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	records := loadQuizCSV("problems.csv")
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
		answer = strings.Replace(answer, "\n", "", -1)

		if answer == s[1] {
			correctAnswers = correctAnswers + 1
		}
	}

	fmt.Print("Correct answers: ", correctAnswers, "/", totalQuestions, "\n")
}

func loadQuizCSV(csvFile string) [][]string {
	data, err := ioutil.ReadFile("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(data)))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}
