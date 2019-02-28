package main

import (
	"fmt"
	"github.com/gowthamgts/aoc18"
	"strings"
)

func main() {
	line := util.ReadFromInputFile("../input.txt")[0]
	minReducedLength := len(line)

	for i := 'A'; i <= 'Z'; i++ {
		replacedString := strings.ReplaceAll(line, string(i), "")
		replacedString = strings.ReplaceAll(replacedString, string(i + 32), "")
		reducedPolymerLength := util.GetReducedPolymerLength(replacedString)
		if minReducedLength > reducedPolymerLength {
			minReducedLength = reducedPolymerLength
		}
	}

	fmt.Println(minReducedLength)
}
// 6948
