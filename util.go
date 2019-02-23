package util

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

// readFromInputFile reads from the file line by line
// and returns the slice.
func ReadFromInputFile(path string) []string {
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

// ParseRegexForFabric is used to parse the input string
// and return x, y, width and height positions.
// Used for day 3 problems.
func ParseRegexForFabric(line string) (int, int, int, int) {
	re := regexp.MustCompile(`(?m)\s(\d+),(\d+):\s(\d+)x(\d+)`)

	result := re.FindStringSubmatch(line)
	x, _ := strconv.Atoi(result[1])
	y, _ := strconv.Atoi(result[2])
	width, _ := strconv.Atoi(result[3])
	height, _ := strconv.Atoi(result[4])

	return x, y, width, height
}
