package day5

type GardeningRange struct {
	destination int
	source      int
	scope       int
	rangeType   BlockType
}

func (gardeningRange GardeningRange) inSourceRange(val int) bool {
	return val >= gardeningRange.source && val < gardeningRange.source+gardeningRange.scope
}

func (gardeningRange GardeningRange) inDestinationRange(val int) bool {
	return val >= gardeningRange.destination && val < gardeningRange.destination+gardeningRange.scope
}

func (gardeningRange GardeningRange) getDestination(val int) int {
	return gardeningRange.destination + (val - gardeningRange.source)
}

type BlockType int64

const (
	Seeds BlockType = iota
	SeedToSoil
	SoilToFertilizer
	FertilizerToWater
	WaterToLight
	LightToTemperature
	TemperatureToHumidity
	HumidityToLocation
)
