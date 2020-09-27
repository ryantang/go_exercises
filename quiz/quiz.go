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
	quizTimer := flag.String("s", "30", "Number of seconds alloted in this quiz")
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

	fmt.Printf("Timer set to %s. Press enter to start quiz.", seconds)

	_, _ = consoleReader.ReadString('\n')

	go runTimer(seconds, &correctAnswers, totalQuestions)

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

func runTimer(waitTime time.Duration, correctAnswers *int, totalQuestions int) {
	timer := time.NewTimer(waitTime)

	<-timer.C
	fmt.Println("\nRan out of time. Corect answers: ", *correctAnswers, "/", totalQuestions)
	os.Exit(0)
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
