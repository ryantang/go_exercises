package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	quizCSV := flag.String("f", "problems.csv", "Quiz CSV file in 'question,answer' format")
	quizTimer := flag.String("t", "5", "Number of seconds alloted in this quiz")
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

	seconds, err := time.ParseDuration(*quizTimer + "s")
	if err != nil {
		log.Fatal(err)
	}

	timer := time.NewTimer(seconds)

	go func() {
		<-timer.C
		fmt.Println("\nRan out of time. Corect answers: ", correctAnswers, "/", totalQuestions)
		os.Exit(0)
	}()

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
