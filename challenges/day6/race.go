package day6

type Race struct {
	time     int
	distance int
}

func (race Race) getNumberOfWaysToWin() int {
	numberOfWaysToWin := 0

	for i := 1; i <= race.time; i++ {
		hold := i
		restOfMilliSeconds := race.time - i

		if hold*restOfMilliSeconds > race.distance {
			numberOfWaysToWin++
		}
	}

	return numberOfWaysToWin
}
