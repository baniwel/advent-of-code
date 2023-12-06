package day5

import (
	"advent-of-code/challenges/helpers"
	"fmt"
	"math"
	"strings"
)

func RunPart1(path string) {
	lines := helpers.ReadLines(path)
	seeds, gardeningRangeMap := parseData(lines, part1SeedsParser)
	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		location := getSeedLocation(seed, gardeningRangeMap)

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	fmt.Println("Lowest location for part 1 is", lowestLocation)
}

func RunPart2(path string) {
	lines := helpers.ReadLines(path)
	seeds, gardeningRangeMap := parseData(lines, part2SeedsParser)
	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		location := getSeedLocation(seed, gardeningRangeMap)

		if location < lowestLocation {
			lowestLocation = location
		}
	}

	fmt.Println("Lowest location for part 2 is", lowestLocation)
}

func getSeedLocation(seed int, gardeningRangeMap map[BlockType][]GardeningRange) int {
	number := seed
	for blockType := SeedToSoil; blockType <= HumidityToLocation; blockType++ {
		for _, gardeningRanges := range gardeningRangeMap[blockType] {
			if gardeningRanges.inSourceRange(number) {
				number = gardeningRanges.getDestination(number)
				break
			}
		}
	}
	return number
}

func part1SeedsParser(seeds []int) []int {
	return seeds
}

func part2SeedsParser(seeds []int) []int {
	var finalSeeds []int
	fmt.Println("parsing seeds")
	for i := 0; i < len(seeds)-1; i = i + 2 {
		startSeedNumber := seeds[i]
		seedRange := seeds[i+1]

		for seedIndex := startSeedNumber; seedIndex < startSeedNumber+seedRange; seedIndex++ {
			finalSeeds = append(finalSeeds, seedIndex)
		}
	}

	fmt.Println("parsing seeds done")
	return finalSeeds
}

func parseData(lines []string, seedsParser func([]int) []int) ([]int, map[BlockType][]GardeningRange) {
	treating := Seeds
	previousEmpty := true
	var seeds []int
	gardeningRangeMap := make(map[BlockType][]GardeningRange)
	for _, line := range lines {
		if strings.EqualFold(line, "") {
			previousEmpty = true
			continue
		}

		if !previousEmpty {
			gardeningRangeMap[treating-1] = append(gardeningRangeMap[treating-1], parseLine(line, treating-1))
		} else {
			if Seeds == treating {
				seeds = parseSeeds(line)
			}
			treating++
		}
		previousEmpty = false
	}

	return seedsParser(seeds), gardeningRangeMap
}

func parseSeeds(line string) []int {
	var seeds []int
	for _, seed := range strings.Split(strings.Replace(line, "seeds: ", "", 1), " ") {
		seeds = append(seeds, helpers.ParseNumber(seed))
	}
	return seeds
}

func parseLine(line string, blockType BlockType) GardeningRange {
	numbers := strings.Split(line, " ")
	return GardeningRange{destination: helpers.ParseNumber(numbers[0]),
		source:    helpers.ParseNumber(numbers[1]),
		scope:     helpers.ParseNumber(numbers[2]),
		rangeType: blockType}
}
