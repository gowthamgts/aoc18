package main

import (
	"fmt"
	"github.com/gowthamgts/aoc18"
)

func main() {
	var fabric [1000][1000]int

	lines := util.ReadFromInputFile("../input.txt")

	var resultingId int
	var intersectionFlags = make([]bool, len(lines) + 1)

	for _, currentLine := range lines {

		id, x, y, width, height := util.ParseRegexForFabric(currentLine)

		for i := x; i < x+width; i++ {
			for j := y; j < y+height; j++ {

				if fabric[i][j] > 0 {
					// mark as intersected
					intersectionFlags[fabric[i][j]] = true
					intersectionFlags[id] = true
				} else {
					// set the id
					fabric[i][j] = id
				}
			}
		}

	}

	for i := 1; i < len(intersectionFlags); i++ {
		if !intersectionFlags[i] {
			resultingId = i
			break
		}
	}

	fmt.Println("Resulting ID:", resultingId)
}

// Resulting ID: 1097