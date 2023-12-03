package day3

type Gear struct {
	parts []int
}

func (gear Gear) add(part int) {
	gear.parts = append(gear.parts, part)
}

func (gear Gear) gearRatio() int {
	if len(gear.parts) == 0 {
		return 0
	}

	gearRatio := gear.parts[0]

	for i := 1; i < len(gear.parts); i++ {
		gearRatio = gearRatio * gear.parts[i]
	}

	return gearRatio
}
