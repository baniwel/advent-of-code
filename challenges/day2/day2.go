package day2

import (
	"advent-of-code/challenges/helpers"
	"fmt"
	"regexp"
	"strings"
)

func RunPart1(path string, red int, green int, blue int) {
	lines := helpers.ReadLines(path)
	games := parseGame(lines)

	validGames := getValidGames(games, red, blue, green)

	sum := 0
	for _, validGame := range validGames {
		sum += validGame.id
	}
	fmt.Println("Total is", sum)
}

func RunPart2(path string) {
	lines := helpers.ReadLines(path)
	games := parseGame(lines)

	sumOfPowers := 0
	for _, game := range games {
		sumOfPowers += game.getPowerOfSet()
	}
	fmt.Println("Power of Sets ", sumOfPowers)
}

func getValidGames(games []Game, red int, blue int, green int) []Game {
	var validGames []Game
	for _, game := range games {
		if game.possible(red, blue, green) {
			validGames = append(validGames, game)
		}
	}
	return validGames
}

func parseGame(gameLines []string) []Game {
	var games []Game
	for _, gameLine := range gameLines {
		pattern := regexp.MustCompile(`Game (\d+):((?: \d+ (?:red|blue|green)[,;]*)+)`)
		matches := pattern.FindAllStringSubmatch(gameLine, -1)

		for _, match := range matches {
			gameNumber := match[1]

			colorsAndNumbers := match[2]

			grabsStrings := strings.Split(colorsAndNumbers, ";")
			var grabs []Grab
			for _, grabString := range grabsStrings {
				subPattern := regexp.MustCompile(`(\d+) (red|blue|green)`)
				subMatches := subPattern.FindAllStringSubmatch(grabString, -1)

				grab := Grab{}
				for _, subMatch := range subMatches {
					number := helpers.ParseNumber(subMatch[1])
					color := subMatch[2]

					if strings.Compare(color, "red") == 0 {
						grab.reds = number
					} else if strings.Compare(color, "blue") == 0 {
						grab.blues = number
					} else {
						grab.greens = number
					}
				}
				grabs = append(grabs, grab)
			}

			games = append(games, Game{id: helpers.ParseNumber(gameNumber), grabs: grabs})
		}
	}

	return games
}
