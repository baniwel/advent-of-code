package day3

type LineNumber struct {
	number        int
	length        int
	startingIndex int
}

func (lineNumber LineNumber) isAround(index int) bool {
	endingIndex := lineNumber.startingIndex + lineNumber.length - 1
	return endingIndex == index || endingIndex == index-1 ||
		lineNumber.startingIndex == index || lineNumber.startingIndex == index+1 ||
		(index > lineNumber.startingIndex && index < endingIndex)
}
