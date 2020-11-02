package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Problem struct {
	Question string
	Answer string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "time limit for the quiz in seconds")
	flag.Parse()
	csvFile, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	
	r := csv.NewReader(csvFile)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	quiz := parseLines(lines)
	correct := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	problemLoop:
	for i, q := range quiz {
		fmt.Printf("Problem #%d: %s = \n", i+1, q.Question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			break problemLoop
		case answer := <-answerCh:
			if answer == q.Answer {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(quiz))
}


func parseLines(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))
	for i, line := range lines {
		problems[i] = Problem{
			Question: line[0],
			Answer: line[1],
		}
	}
	return problems
}