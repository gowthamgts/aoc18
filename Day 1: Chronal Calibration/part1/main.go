package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// result var
	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Unable to parse string", scanner.Text())
		}

		result += t
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", result)
}
