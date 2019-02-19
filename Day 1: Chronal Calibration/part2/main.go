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

	sum := 0
	var sums []int

	scanner := bufio.NewScanner(file)
	for {
		file.Seek(0, 0)
		scanner.Scan()

		sums = append(sums, sum)

		t, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		sum += t

		for _, s := range sums {
			if s == sum {
				fmt.Println("Repeated:", s)
				return
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
