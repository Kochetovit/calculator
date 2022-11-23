package main

import (
	"bufio"
	"calc/calculator"
	"calc/parser"
	"fmt"
	"os"
)

const colorClear = "\033[H\033[2J"

func main() {
	calculator := calculator.NewDefaultCalculator()

	scanner := bufio.NewScanner(os.Stdin)

	parsers := []parser.Parser{
		&parser.ArabicParser{},
		&parser.RomanParser{},
	}

	for {
		var (
			operation, inputStr string
			num                 float64
		)

		fmt.Print(colorClear) // clear console output
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
