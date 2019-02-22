package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	lines := readFromInputFile("../input.txt")

	var matchingString string

	for index, currentLine := range lines {
		for i := index; i < len(lines); i++ {
			if mismatchIndex := checkStringMatch(currentLine, lines[i]); mismatchIndex != -1 {
				matchingString = currentLine[:mismatchIndex] + currentLine[mismatchIndex+1:]
			}
		}
	}

	fmt.Println(matchingString)
}

// readFromInputFile reads from the file line by line
// and returns the slice.
func readFromInputFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// checkStringMatch checks the two strings and returns
// true if they match except for the one character at
// the same position.
func checkStringMatch(str1 string, str2 string) int {
	// just a fail safe
	if len(str1) != len(str2) {
		return -1
	}

	diffIndex := -1
	mismatched := false

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			// check for previous mismatch
			if mismatched {
				// more than one mismatch. skip
				return -1
			}
			diffIndex = i
			mismatched = true
		}
	}

	return diffIndex
}
