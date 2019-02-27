package main

import (
	"fmt"
	"github.com/gowthamgts/aoc18"
	"sort"
	"time"
)

type Guard struct {
	id                   int
	minuteMap            []int
	totalSleepingMinutes int
}

func main() {
	lines := util.ReadFromInputFile("../input.txt")
	sort.Strings(lines)

	var guards = make(map[int]*Guard)

	var guardInDuty *Guard
	var lastEventTimeStamp time.Time

	for _, val := range lines {
		year, month, day, hour, minute, message := util.ParseRegexForGuard(val)

		currentEventTimestamp := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)

		if message[0] == 'G' {
			// new Guard incoming
			guardId := util.ParseRegexForGuardId(message)
			if _, present := guards[guardId]; !present {
				guardInDuty = &Guard{id: guardId, minuteMap: make([]int, 60)}
				guards[guardId] = guardInDuty
			} else {
				guardInDuty = guards[guardId]
			}
		} else if message[0] == 'w' {
			// wakes up
			for i := lastEventTimeStamp.Minute(); i < currentEventTimestamp.Minute(); i++ {
				guardInDuty.minuteMap[i]++
				guardInDuty.totalSleepingMinutes++
			}
		}
		lastEventTimeStamp = currentEventTimestamp
	}

	// finding the sleepy guard
	var sleepyGuardId = guardInDuty.id
	for _, val := range guards {
		if val.totalSleepingMinutes > guards[sleepyGuardId].totalSleepingMinutes {
			sleepyGuardId = val.id
		}
	}

	// finding the max intersecting minute
	var intersectingMinute int
	for index, val := range guards[sleepyGuardId].minuteMap {
		if val > guards[sleepyGuardId].minuteMap[intersectingMinute] {
			intersectingMinute = index
		}
	}

	fmt.Println(sleepyGuardId, intersectingMinute)
	fmt.Println(sleepyGuardId * intersectingMinute)
}
