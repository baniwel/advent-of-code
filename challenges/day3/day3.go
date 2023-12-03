package day3

import (
	"advent-of-code/challenges/helpers"
	"fmt"
	"strings"
)

func RunPart1(path string) {
	lines := helpers.ReadLines(path)
	engineNumbers := getEngineNumbers(lines)
	sumOfEngineNumbers := helpers.Sum(engineNumbers)
	fmt.Println("Sum of engine numbers is", sumOfEngineNumbers)
}

func RunPart2(path string) {
	lines := helpers.ReadLines(path)
	gears := getGears(lines)

	gearsRatioTotal := 0
	for _, gear := range gears {
		gearsRatioTotal += gear.gearRatio()
	}
	fmt.Println("Gears ratio addition", gearsRatioTotal)
}

func getGears(lines []string) []Gear {
	var gears []Gear

	for i, line := range lines {
		chars := []rune(line)

		for j, char := range chars {
			if char != '*' {
				continue
			}

			var lineNumbers []LineNumber
			lineNumbers = append(lineNumbers, extractNumbersFromLine(lines[i-1])...)
			lineNumbers = append(lineNumbers, extractNumbersFromLine(line)...)
			lineNumbers = append(lineNumbers, extractNumbersFromLine(lines[i+1])...)

			var aroundNumbers []LineNumber
			for _, lineNumber := range lineNumbers {
				if lineNumber.isAround(j) {
					aroundNumbers = append(aroundNumbers, lineNumber)
				}
			}

			if len(aroundNumbers) != 2 {
				continue
			}

			var gear = Gear{}
			for _, lineNumber := range aroundNumbers {
				gear.parts = append(gear.parts, lineNumber.number)
			}
			gears = append(gears, gear)
		}
	}

	return gears
}

func extractNumbersFromLine(line string) []LineNumber {
	var lineNumbers []LineNumber

	chars := []rune(line)
	var number string
	for i, char := range chars {
		if charIsNumber(char) {
			number = number + string(char)

			if i+1 == len(chars) {
				lineNumbers = append(
					lineNumbers,
					LineNumber{number: helpers.ParseNumber(number), length: len(number), startingIndex: i - len(number) + 1},
				)
				number = ""
			}
		} else if strings.Compare(number, "") != 0 {
			lineNumbers = append(
				lineNumbers,
				LineNumber{number: helpers.ParseNumber(number), length: len(number), startingIndex: i - len(number)},
			)
			number = ""
		}
	}

	return lineNumbers
}

func getEngineNumbers(lines []string) []int {
	var engineNumbers []int

	for i, line := range lines {
		chars := []rune(line)

		var number = ""
		for j, char := range chars {
			charIsNumber := charIsNumber(char)
			if charIsNumber {
				number = number + string(char)
			}

			shouldCheck := strings.Compare(number, "") != 0 && (char < '0' || char > '9')
			shouldCheck = shouldCheck || len(chars) == j+1
			if shouldCheck {
				leftIndex := j - len(number)
				if !charIsNumber {
					leftIndex--
				}
				leftIndex = helpers.Max(leftIndex, 0)
				rightIndex := j

				isEngineNumber := false

				if isSymbol(chars[rightIndex]) || isSymbol(chars[leftIndex]) {
					isEngineNumber = true
				} else {
					if i > 0 {
						isEngineNumber = checkEngineNumber([]rune(lines[i-1]), leftIndex, rightIndex)
					}

					if !isEngineNumber && i < len(lines)-1 {
						isEngineNumber = checkEngineNumber([]rune(lines[i+1]), leftIndex, rightIndex)
					}
				}

				if isEngineNumber {
					engineNumbers = append(engineNumbers, helpers.ParseNumber(number))
				}
				number = ""
			}
		}
	}

	return engineNumbers
}

func charIsNumber(char rune) bool {
	return char >= '0' && char <= '9'
}

func checkEngineNumber(otherLine []rune, leftIndex int, rightIndex int) bool {
	for i := leftIndex; i <= rightIndex; i++ {
		if isSymbol(otherLine[i]) {
			return true
		}
	}

	return false
}

func isSymbol(char rune) bool {
	return (char < '0' || char > '9') && char != '.'
}
