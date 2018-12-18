package main

import (
	"adventofcode2018-go/util"
	"fmt"
	"strings"
)

const day string = "day5"

func react(line string) string {
	var stack []rune = make([]rune, 0)
	stack = append(stack, rune(line[0]))
	for i:= 1; i < len(line); i++ {
		rightRune := rune(line[i])
		if len(stack) == 0 {
			stack = append(stack, rightRune)
			continue
		}
		leftRune := stack[len(stack) - 1]
		if strings.ToLower(string(leftRune)) == strings.ToLower(string(rightRune)) && leftRune != rightRune {
			//fmt.Printf("Stripping before %v\n", stack)
			stack = stack[0: len(stack) - 1]
			//fmt.Printf("Stripping after  %v\n", stack)
			continue
		}
		stack = append(stack, rightRune)
	}
	return string(stack)
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
