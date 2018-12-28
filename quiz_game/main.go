package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("./problems.csv")

	if err != nil {
		panic(err)
	}

	r := csv.NewReader(strings.NewReader(string(dat)))

	var correct int
	var incorrect int

	fmt.Println("Starting Quiz!")

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Printf("The question is: %s", record[0])

		userInput := bufio.NewReader(os.Stdin)
		fmt.Println("")
		fmt.Print("Your Answer is: ")
		text, _ := userInput.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if record[1] == text {
			correct++
		} else {
			incorrect++
		}
	}

	fmt.Println("------------------------")
	fmt.Println("Final Results")
	fmt.Printf("Correct: %d", correct)
	fmt.Println("")
	fmt.Printf("Incorrect: %d", incorrect)
	fmt.Println("")
	fmt.Println("------------------------")
}
