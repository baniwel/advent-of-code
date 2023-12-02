package day2

type Game struct {
	id    int
	grabs []Grab
}

type Grab struct {
	reds   int
	blues  int
	greens int
}

func (game Game) possible(red int, blue int, green int) bool {
	for _, grab := range game.grabs {
		if grab.reds != 0 && red < grab.reds {
			return false
		}

		if grab.blues != 0 && blue < grab.blues {
			return false
		}

		if grab.greens != 0 && green < grab.greens {
			return false
		}
	}

	return true
}

func (game Game) getPowerOfSet() int {
	maxRed := 0
	maxBlue := 0
	maxGreen := 0

	for _, grab := range game.grabs {
		if grab.reds > maxRed {
			maxRed = grab.reds
		}

		if grab.greens > maxGreen {
			maxGreen = grab.greens
		}

		if grab.blues > maxBlue {
			maxBlue = grab.blues
		}
	}

	return maxGreen * maxBlue * maxRed
}
