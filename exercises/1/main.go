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
	quiz := make([]Problem, len(lines))
	correct := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	fmt.Println(timer)
	for i, line := range lines {
		quiz[i] = Problem{
			Question: line[0],
			Answer: line[1],
		}
		var inputAnswer string
		fmt.Println(quiz[i].Question)
		fmt.Scanf("%s\n", &inputAnswer)
		if inputAnswer == quiz[i].Answer {
			correct++
		}
	}


	fmt.Printf("You scored %d out of %d.\n", correct, len(quiz))
}
