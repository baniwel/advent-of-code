package day6

import (
	"advent-of-code/challenges/helpers"
	"fmt"
	"regexp"
	"strings"
)

func RunPart1(path string) {
	lines := helpers.ReadLines(path)
	races := parseInputPart1(lines)

	multiplication := 1
	for _, race := range races {
		multiplication = multiplication * race.getNumberOfWaysToWin()
	}

	fmt.Println("Final number is", multiplication)
}

func RunPart2(path string) {
	lines := helpers.ReadLines(path)
	race := parseInputPart2(lines)
	println("Number of ways to win is", race.getNumberOfWaysToWin())
}

func parseInputPart1(lines []string) []Race {
	times := helpers.ParseNumbers(
		strings.Fields(
			strings.Trim(
				strings.ReplaceAll(lines[0], "Time:", ""),
				" ",
			),
		),
	)
	distances := helpers.ParseNumbers(
		strings.Fields(
			strings.Trim(
				strings.ReplaceAll(lines[1], "Distance:", ""),
				" ",
			),
		),
	)

	var races []Race
	for i := 0; i < len(times); i++ {
		races = append(races, Race{time: times[i], distance: distances[i]})
	}

	return races
}

func parseInputPart2(lines []string) Race {
	re := regexp.MustCompile(`\s+`)

	time := helpers.ParseNumber(
		re.ReplaceAllString(
			strings.ReplaceAll(lines[0], "Time:", ""),
			"",
		),
	)
	distance := helpers.ParseNumber(
		re.ReplaceAllString(
			strings.ReplaceAll(lines[1], "Distance:", ""),
			"",
		),
	)

	race := Race{time, distance}
	return race
}
