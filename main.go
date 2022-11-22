package main

import (
	"bufio"
	"calc/calculator"
	"calc/parser"
	"fmt"
	"os"
)

func main() {
	calculator := calculator.NewDefaultCalculator()

	scanner := bufio.NewScanner(os.Stdin)

	parsers := []parser.Parser{
		&parser.ArabicParser{},
		&parser.RomanParser{},
	}

	for {
		var (
			operation string
			num       float64
			inputStr  string
		)

		fmt.Print("\033[H\033[2J") // clear console output
		fmt.Printf("Result: %v\n", calculator.GetResult())
		fmt.Printf("Enter operation and number: ")

		if scanner.Scan() {
			inputStr = scanner.Text()
		}

		switch inputStr {
		case "restart":
			calculator.Restart()
		case "exit":
			return
		default:
			for i := 0; i < len(parsers); i++ {
				if err := parsers[i].Parse(inputStr, &operation, &num); err == nil {
					break
				}
			}
		}

		calculator.Evaluate(operation, num)
	}
}
