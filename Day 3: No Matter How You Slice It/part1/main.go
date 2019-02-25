package main

import (
	"fmt"
	"github.com/gowthamgts/aoc18"
)

func main() {
	var fabric [1000][1000]int

	lines := util.ReadFromInputFile("../input.txt")

	for _, currentLine := range lines {
		_, x, y, width, height := util.ParseRegexForFabric(currentLine)
		for i := x; i < x+width; i++ {
			for j := y; j < y+height; j++ {
				fabric[i][j]++
			}
		}

	}

	resultingClaims := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] > 1 {
				resultingClaims++
			}
		}
	}

	fmt.Println("Claims that more than 1:", resultingClaims)
}
