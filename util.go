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
// and return id and x, y, width, height positions
// Used for day 3 problems.
func ParseRegexForFabric(line string) (int, int, int, int, int) {
	re := regexp.MustCompile(`(?m)(\d+)\s@\s(\d+),(\d+):\s(\d+)x(\d+)`)

	result := re.FindStringSubmatch(line)
	id, _ := strconv.Atoi(result[1])
	x, _ := strconv.Atoi(result[2])
	y, _ := strconv.Atoi(result[3])
	width, _ := strconv.Atoi(result[4])
	height, _ := strconv.Atoi(result[5])

	return id, x, y, width, height
}

// ParseRegexForGuard is used to parse input for guard problems (4.1 and 4.2)
func ParseRegexForGuard(line string) (year, month, day, hour, minute int, message string) {
	re := regexp.MustCompile(`\[(\d+)-(\d+)-(\d+)\s(\d+):(\d+)\]\s(.+)`)

	result := re.FindStringSubmatch(line)
	year, _ = strconv.Atoi(result[1])
	month, _ = strconv.Atoi(result[2])
	day, _ = strconv.Atoi(result[3])
	hour, _ = strconv.Atoi(result[4])
	minute, _ = strconv.Atoi(result[5])

	return year, month, day, hour, minute, result[6]
}

// GetGuardId returns the Guard Id from the string
func ParseRegexForGuardId(message string) int {
	re := regexp.MustCompile(`Guard\s#(\d+)`)

	result := re.FindStringSubmatch(message)
	guardId, _ := strconv.Atoi(result[1])

	return guardId
}

// GetReducedPolymerLength reduces the polymer and returns
// the length. Used by 5.1 and 5.2
func GetReducedPolymerLength(line string) int {
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1]-32 || line[i] == line[i+1]+32 {
			// stash
			line = line[:i] + line[i+2:]
			i = -1
			continue
		}
	}

	return len(line)
}

// ParseRegexForCoordinate will parse the input string
// and will give (x, y) location
func ParseRegexForCoordinate(line string) (int, int) {
	re := regexp.MustCompile(`(\d+),\s(\d+)`)

	result := re.FindStringSubmatch(line)
	x, _ := strconv.Atoi(result[1])
	y, _ := strconv.Atoi(result[2])

	return x, y
}

func ManhattanDistance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
