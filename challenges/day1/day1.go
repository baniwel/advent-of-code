package day1

import (
	"advent-of-code/challenges/helpers"
	"fmt"
	"strings"
)

func RunPart1(path string) {
	lines := helpers.ReadLines(path)
	fmt.Println("Sum is", calcSumOfLines(lines))
}

func RunPart2(path string) {
	lines := helpers.ReadLines(path)

	numbersMapping := map[string]string{"one": "o1e", "two": "t2o", "three": "th3ee", "four": "f4r", "five": "f5e", "six": "s6x", "seven": "s7n", "eight": "e8t", "nine": "n9e"}
	keys := helpers.MapKeys(numbersMapping)

	for i := range lines {

		for _, k := range keys {
			lines[i] = strings.ReplaceAll(lines[i], k, numbersMapping[k])
		}
	}

	fmt.Println("Sum is", calcSumOfLines(lines))
}

func calcSumOfLines(lines []string) int {
	var numbers []int
	for _, line := range lines {
		var firstNumber, lastNumber string

		for i := 0; i < len(line); i++ {
			if helpers.IsNumber(string(line[i])) {
				firstNumber = string(line[i])
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if helpers.IsNumber(string(line[i])) {
				lastNumber = string(line[i])
				break
			}
		}

		var fullNumber = helpers.ParseNumber(firstNumber + lastNumber)
		numbers = append(numbers, fullNumber)
	}

	return helpers.Sum(numbers)
}
