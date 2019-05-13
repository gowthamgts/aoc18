package main

import (
	"fmt"
	"github.com/gowthamgts/aoc18"
	"math"
	"strings"
)

type Point struct {
	x, y int
}

type Area struct {
	point  *Point
	area   int
	finite bool
}

var maxX, maxY int

func main() {
	lines := util.ReadFromInputFile("/Users/gts/area51/gopath/src/github.com/gowthamgts/aoc18/Day 6: Chronal Coordinates/input.txt")

	var graph [][]string
	parsedPoints := map[string]*Point{}

	// plot the graph
	for index, line := range lines {
		x, y := util.ParseRegexForCoordinate(line)
		graph[x][y] = string('A' + index)
		parsedPoints[graph[x][y]] = &Point{x, y}

		if maxX < x {
			maxX = x
		}

		if maxY < y {
			maxY = y
		}
	}

	graph = make([][]string, maxX + 1)
	for i := range graph {
		graph[i] = make([]string, maxY + 1)
	}

	// visualize and paint the grid
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			if graph[j][i] == "" {
				graph[j][i] = findClosestPoint(j, i, parsedPoints)
			}
			fmt.Printf("%4v", graph[j][i])
		}
		fmt.Println("")
	}

	// find the area
	fmt.Println("Max Area:", findArea(graph, parsedPoints))
}

// findClosestPoint will calculate and return the closest point
func findClosestPoint(x, y int, parsedPoints map[string]*Point) string {
	var closestPoints []string
	var closestDistance int = math.MaxInt32
	for key, val := range parsedPoints {
		distance := util.ManhattanDistance(x, y, val.x, val.y)
		if distance <= closestDistance {
			if distance == closestDistance {
				closestPoints = append(closestPoints, key)
			} else {
				closestPoints = append(closestPoints[:0], key)
			}
			closestDistance = distance
		}
	}

	if len(closestPoints) > 1 {
		return "."
	} else {
		return strings.ToLower(closestPoints[0])
	}
}

// findArea will calculate the area
func findArea(graph [][]string, parsedPoints map[string]*Point) int {
	areaMap := map[string]*Area{}

	for key, value := range parsedPoints {
		areaMap[key] = &Area{point: value, area: 0, finite: true}
	}

	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			if graph[j][i] == "." {
				// nothing to be done here
				continue
			}

			key := strings.ToUpper(graph[j][i])

			areaMap[key].area++
			if i == 0 || j == 0 || i == maxX - 1 || j == maxY - 1 {
				areaMap[key].finite = false
			}
		}
	}

	var maxArea string
	for key, val := range areaMap {
		if !val.finite {
			continue
		}
		// assign the first one
		if maxArea == "" && val.finite {
			maxArea = key
			continue
		}
		if val.area > areaMap[maxArea].area {
			// assign the current one
			maxArea = key
		}
	}

	return areaMap[maxArea].area
}