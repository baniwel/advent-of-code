package day4

import (
	"advent-of-code/challenges/helpers"
	"fmt"
	"regexp"
	"strings"
)

func RunPart1(path string) {
	lines := helpers.ReadLines(path)
	cards := getCards(lines)
	totalPoints := getTotalPoints(cards)
	fmt.Println("Total points", totalPoints)
}

func RunPart2(path string) {
	lines := helpers.ReadLines(path)
	cards := getCards(lines)
	totalScratchCards := getTotalScratchcards(cards)
	fmt.Println("Total scratch cards", totalScratchCards)
}

func getTotalPoints(cards []Card) int {
	totalPoints := 0
	for _, card := range cards {
		totalPoints += card.calculatePoints()
	}
	return totalPoints
}

func getTotalScratchcards(cards []Card) int {
	for cardIndex, card := range cards {
		card.setPoints()
		if card.isWinningCard() {
			for i := cardIndex + 1; i <= cardIndex+card.countOfWinningNumbers; i++ {
				cards[i].copies = cards[i].copies + 1 + card.copies
			}
		}
	}

	totalScratchcards := 0
	for _, card := range cards {
		totalScratchcards += card.copies + 1
	}
	return totalScratchcards
}

func getCards(lines []string) []Card {
	var cards []Card
	for _, line := range lines {
		pattern := regexp.MustCompile(`Card\s+(\d+):\s+([\d+\s+]+)|\s+([\d+\s+]+)`)
		matches := pattern.FindAllStringSubmatch(line, -1)

		cards = append(cards, rowToCard(matches))
	}

	return cards
}

func rowToCard(matches [][]string) Card {
	gameNumber := helpers.ParseNumber(matches[0][1])

	var winningNumbers []int
	for _, number := range strings.Fields(matches[0][2]) {
		winningNumbers = append(winningNumbers, helpers.ParseNumber(number))
	}

	var myNumbers []int
	for _, number := range strings.Fields(matches[1][0]) {
		myNumbers = append(myNumbers, helpers.ParseNumber(number))
	}

	return Card{gameNumber, myNumbers, winningNumbers, 0, 0, 0}
}
