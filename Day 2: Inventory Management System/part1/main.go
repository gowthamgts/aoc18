package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cnt2, cnt3 := 0, 0

	for scanner.Scan() {
		text := scanner.Text()
		repeatedTwice, repeatedThrice := false, false
		countMap := make(map[rune]int)

		for _, t := range text {
			countMap[t]++
		}

		for _, cnt := range countMap {
			if cnt == 2 {
				repeatedTwice = true
			} else if cnt > 2 {
				repeatedThrice = true
			}
		}

		if repeatedTwice {
			cnt2++
		}
		if repeatedThrice {
			cnt3++
		}
	}

	fmt.Println("Twice count", cnt2)
	fmt.Println("Thrice count", cnt3)
	fmt.Println(cnt2 * cnt3)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
