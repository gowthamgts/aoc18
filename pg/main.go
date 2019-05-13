package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//numbers := []int{1, 0, 10, 0, 5, 0, 0, 100, 11, 0, 20}
	//fmt.Println(numbers)
	//fmt.Println(arrangeNumbers(numbers))
	//
	//powerSetNumbers := []int{1, 2, 3}
	//fmt.Println(findPowerSets(powerSetNumbers))

	numbers := []int{1, 3, 6, 1, 0, 9}

	fmt.Println(FindJumpsNeeded(numbers))

}

func arrangeNumbers(numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; numbers[i] != 0 && j < len(numbers); j++ {
			if numbers[j] == 0 {
				// push to the front
				numbers = Prepend(numbers, j)
			}
		}
	}

	return numbers
}

func Prepend(numbers []int, position int) []int {
	if position == 0 {
		return numbers
	}
	arr := make([]int, 0)
	arr = append(arr, numbers[position])
	arr = append(arr, numbers[0:position]...)
	arr = append(arr, numbers[position+1:]...)

	return arr
}

func findPowerSets(numbers []int) []string {
	totalSize := 1 << uint(len(numbers))
	powerSet := make([]string, 0)

	for i := 0; i < totalSize; i++ {
		result := ""
		for j := 0; j < len(numbers); j++ {
			if i&(1<<uint(len(numbers)-1-j)) > 0 {
				result = result + strconv.Itoa(numbers[j])
			}
		}
		powerSet = append(powerSet, result)
	}

	return powerSet
}

func FindJumpsNeeded(numbers []int) int {
	jumps := make([]int, len(numbers))

	jumps[0] = 0

	if numbers[0] == 0 {
		return -1
	}

	for i := 1; i < len(numbers); i++ {
		jumps[i] = math.MaxInt32

		for j := 0; j < i; j++ {
			if numbers[j] 
		}

	}
}
