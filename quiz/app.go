package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	var questions []string
	var answers []string
	var rightAnswers []bool
	var sum int = 0

	csvFile := flag.String("file", "input.csv", "csv file to be parsed")

	flag.Parse()

	// Open the file
	csvfile, err := os.Open(*csvFile)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		questions = append(questions, record[0])
		answers = append(answers, record[1])
	}

	reader := bufio.NewReader(os.Stdin)

	// Iterate through questions and ask them one by one
	for index, _ := range questions {
		fmt.Println(questions[index])
		answer, _ := reader.ReadString('\n')

		// Compare answers
		if strings.TrimRight(answer, "\n") == answers[index] {
			rightAnswers = append(rightAnswers, true)
		} else {
			rightAnswers = append(rightAnswers, false)
		}
	}

	// Count right answers
	for index, _ := range rightAnswers {
		if rightAnswers[index] == true {
			sum++
		}
	}

	fmt.Printf("%s%d%s%d%s", "You scored ", sum, " out of ", len(rightAnswers), "\n")
}
