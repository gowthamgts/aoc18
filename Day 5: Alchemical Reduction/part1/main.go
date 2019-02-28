package main

import (
	"fmt"
	"github.com/gowthamgts/aoc18"
)

func main() {
	line := util.ReadFromInputFile("../input.txt")[0]

	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] - 32 || line[i] == line[i+1] + 32 {
			// stash
			line = line[:i] + line[i+2:]
			i = -1
			continue
		}
	}

	fmt.Println(len(line))
}
