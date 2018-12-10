package main

import (
	"adventofcode2018-go/util"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"time"
)

const day = "day4"

var timeRegex *regexp.Regexp = regexp.MustCompile(`^\[(.*?)\]`)
var typeRegex *regexp.Regexp = regexp.MustCompile(`^\[.*] (\w+) `)
var guardRegex *regexp.Regexp = regexp.MustCompile(`Guard #(\d+)`)

type LineType int

const (
	UNKNOWN LineType = iota
	GUARD
	SLEEP
	WAKE
)

type LineEntry struct {
	lineType LineType
	time     time.Time
	line     string
}

type Sleep struct {
	start time.Time
	end time.Time
}

func ParseLine(line string) LineEntry {
	match := timeRegex.FindStringSubmatch(line)
	timeString := match[1]
	time, err := time.Parse("2006-01-02 15:04", timeString)
	if err != nil {
		panic(err)
	}
	match = typeRegex.FindStringSubmatch(line)
	if match == nil {
		panic("Should not happen")
	}

	typeString := match[1]
	var lineType LineType
	switch typeString {
	case "Guard":
		lineType = GUARD
	case "wakes":
		lineType = WAKE
	case "falls":
		lineType = SLEEP
	default:
		panic("Should never happen")
	}
	lineEntry := LineEntry{}
	lineEntry.lineType = lineType
	lineEntry.time = time
	lineEntry.line = line
	return lineEntry
}

func ParseGuard(line string) int {
	match := guardRegex.FindStringSubmatch(line)
	if match == nil {
		panic("Should not happen")
	}

	guard, _ := strconv.Atoi(match[1])
	return guard
}

func BucketByMinute(sleeps []Sleep) map[int]int {
	sleepBucketedByMinutes := make(map[int]int)
	for _, sleep := range sleeps {
		sleepDuration := sleep.end.Sub(sleep.start)
		for i := 0; i < int(sleepDuration.Minutes()); i++ {
			minuteDelta := time.Duration(i) * time.Minute
			minute := sleep.start.Add(minuteDelta).Minute()
			sleepBucketedByMinutes[minute]++
		}
	}
	return sleepBucketedByMinutes
}

func main() {
	lines := util.LoadInputFile(day)
	lineEntries := make([]LineEntry, len(lines))
	for i, line := range lines {
		lineEntries[i] = ParseLine(line)
	}

	sort.SliceStable(lineEntries, func(i, j int) bool {
		return lineEntries[i].time.Before(lineEntries[j].time)
	})

	var currentGuard int = 0
	var sleepTime time.Time
	var guardSleep = make(map[int][]Sleep)
	for _, lineEntry := range lineEntries {
		switch lineEntry.lineType {
		case GUARD:
			guard := ParseGuard(lineEntry.line)
			if guard != currentGuard {
				currentGuard = guard
			}
		case WAKE:
			if currentGuard == 0 {
				panic("Should not happen")
			}
			sleep := Sleep{}
			sleep.start = sleepTime
			sleep.end = lineEntry.time
			sleeps := guardSleep[currentGuard]
			if sleeps == nil {
				sleeps = make([]Sleep, 0)
			}
			sleeps = append(sleeps, sleep)
			guardSleep[currentGuard] = sleeps
		case SLEEP:
			if currentGuard == 0 {
				panic("Should not happen")
			}
			sleepTime = lineEntry.time
		}
	}

	var maxSleepGuard int
	var maxSleep time.Duration
	for guard, sleeps := range guardSleep {
		var sleepDuration time.Duration
		for _, sleep := range sleeps {
			sleepDuration += sleep.end.Sub(sleep.start)
		}
		if sleepDuration > maxSleep {
			maxSleep = sleepDuration
			maxSleepGuard = guard
		}
	}

	fmt.Printf("Maximum sleep guard is %d with duration %d seconds\n", maxSleepGuard, int(maxSleep.Seconds()))

	sleeps := guardSleep[maxSleepGuard]
	sleepBucketedByMinutes := BucketByMinute(sleeps)

	var maxMinute, maxTotalTime int
	for minute, totalTime := range sleepBucketedByMinutes {
		if totalTime > maxTotalTime {
			maxMinute = minute
			maxTotalTime = totalTime
		}
	}

	fmt.Printf("Maximum minute asleep is %d with %d total minutes\n", maxMinute, maxTotalTime)
	answerId := maxSleepGuard * maxMinute
	fmt.Printf("First star answer is %d\n", answerId)

	var secondStarGuard, maxMinuteSecondStar, maxSleepBucketValue int
	for guard, sleeps := range guardSleep {
		bucketByMinute := BucketByMinute(sleeps)
		for minute, value := range bucketByMinute {
			if value > maxSleepBucketValue {
				maxSleepBucketValue = value
				maxMinuteSecondStar = minute
				secondStarGuard = guard
			}
		}
	}

	fmt.Printf("Guard %d slept a total of %d times on minute %d\n", secondStarGuard, maxSleepBucketValue, maxMinuteSecondStar)
	secondStarAnswer := secondStarGuard * maxMinuteSecondStar
	fmt.Printf("Second star answer is %d\n", secondStarAnswer)
}
