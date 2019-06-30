package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"bufio"
	"strings"
)

type QuizQuestion struct {
	Question string
	Answer string
}

func main() {
	//read csv
	csvfile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	//create array of questions/answers
	quiz := [12]QuizQuestion{}
	for i := 0; i < 13; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		quiz[i] = QuizQuestion{Question: record[0], Answer: record[1]};
	}

	//loop through array
	for i:=0; i < len(quiz); i++ {
		//write question to command line
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Question " + strconv.Itoa(i+1) + ": What is " + quiz[i].Question + "?\n")
		//take user input
		text, _ := reader.ReadString('\n')
		//match user input against the answer
		split := strings.Split(quiz[i].Question, "+")

		a, err := strconv.Atoi(split[0])
		b, err := strconv.Atoi(split[1])
		textInt, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))

		if err == nil {
			var correctAnswer int
			correctAnswer = a + b
			
			//return correct or incorrectx
			if (correctAnswer == textInt) {
				fmt.Println("Correct!");
			} else {
				fmt.Println("Sorry :( The correct answer is " + strconv.Itoa(correctAnswer) + ".")
			}
		} else {
			log.Fatal(err)
		}
	}
}