package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	filePtr := flag.String("file", "problems.csv", "filename to use for csv file")
	timePtr := flag.Int("time", 30, "time to complete quiz")
	flag.Parse()

	file, err := os.Open(*filePtr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	scanner := bufio.NewScanner(os.Stdin)
	numCorrectAnswers := 0
	numQuestions := 0

	fmt.Println("Press ENTER to start quiz or 'q' to quit")
	scanner.Scan()
	if scanner.Text() == "q" {
		return
	}

	results, _ := reader.ReadAll()
	numQuestions = len(results)

	timer := time.NewTimer(time.Second * time.Duration(*timePtr))
	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Println("\nTime's up!\n\nFinal score:", numCorrectAnswers, "out of", numQuestions)
		os.Exit(0)
	}()

	for i := range results {
		question := results[i][0]
		correctAnswer := results[i][1]

		fmt.Println("What is", question, "?")
		scanner.Scan()
		answer := scanner.Text()

		if answer == correctAnswer {
			numCorrectAnswers++
		}
	}

	fmt.Println("\nFinal score:", numCorrectAnswers, "out of", numQuestions)
}
