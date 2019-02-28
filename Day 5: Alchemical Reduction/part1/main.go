package main

import (
	"fmt"
	"github.com/gowthamgts/aoc18"
)

func main() {
	line := util.ReadFromInputFile("../input.txt")[0]

	fmt.Println(util.GetReducedPolymerLength(line))
}
