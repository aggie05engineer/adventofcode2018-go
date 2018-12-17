package main

import (
	"adventofcode2018-go/util"
	"fmt"
	"strings"
)

const day string = "day5"

func react(line string) string {
	var leftIndex int = 0
	for rightIndex := 1; rightIndex < len(line); {
		leftRune := line[leftIndex]
		rightRune := line[rightIndex]
		fmt.Printf("[%s,%s]\n", string(leftRune), string(rightRune))
		if strings.ToLower(string(leftRune)) == strings.ToLower(string(rightRune)) && leftRune != rightRune {
			fmt.Printf("Stripping before %s\n", line)
			line = line[0:leftIndex] + line[rightIndex+1:]
			fmt.Printf("Stripping after  %s\n", line)
			if leftIndex > 0 {
				leftIndex--
			} else {
				leftIndex = 0
			}
			if rightIndex > 1 {
				rightIndex--
			} else {
				rightIndex = 1
			}
			continue
		}
		fmt.Printf("Moving right [%d,%d]\n", leftIndex, rightIndex)
		leftIndex++
		rightIndex++
	}
	return line
}

func main() {
	lines := util.LoadInputFile(day)
	if len(lines) > 1 {
		panic("Should not happen")
	}

	fmt.Printf("There are originally %d units\n", len(lines[0]))
	line := lines[0]
	reactedLine := react(line)
	fmt.Printf("The remaining line has %d units", len(reactedLine))
}
