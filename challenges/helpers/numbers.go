package helpers

import (
	"log"
	"strconv"
)

func IsNumber(value string) bool {
	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		return true
	}

	return false
}

func ParseNumber(value string) int {
	parsedInt, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return parsedInt
}

func ParseNumbers(values []string) []int {
	var intValues []int
	for _, value := range values {
		intValues = append(intValues, ParseNumber(value))
	}
	return intValues
}

func Sum(arr []int) int {
	var sum = 0
	for idx := 0; idx < len(arr); idx++ {
		sum += arr[idx]
	}
	return sum
}

func Min(value1 int, value2 int) int {
	if value1 > value2 {
		return value2
	} else {
		return value1
	}
}

func Max(value1 int, value2 int) int {
	if value1 > value2 {
		return value1
	} else {
		return value2
	}
}
