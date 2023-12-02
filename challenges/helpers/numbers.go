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

func Sum(arr []int64) int64 {
	var sum int64 = 0
	for idx := 0; idx < len(arr); idx++ {
		sum += arr[idx]
	}
	return sum
}
