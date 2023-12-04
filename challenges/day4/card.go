package day4

import (
	"math"
)

type Card struct {
	number                int
	myNumbers             []int
	winningNumbers        []int
	totalPoints           int
	countOfWinningNumbers int
	copies                int
}

func (card Card) isWinningCard() bool {
	return card.totalPoints > 0
}

func (card *Card) setPoints() {
	card.totalPoints = card.calculatePoints()
}

func (card *Card) calculatePoints() int {
	matches := 0

	for _, myNumber := range card.myNumbers {
		for _, winningNumber := range card.winningNumbers {
			if myNumber == winningNumber {
				matches++
			}
		}
	}

	card.countOfWinningNumbers = matches
	if matches < 2 {
		return matches
	} else {
		return int(math.Pow(2, float64(matches-1)))
	}
}
