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
			stack = stack[0: len(stack) - 1]
			continue
		}
		stack = append(stack, rightRune)
	}
	return string(stack)
}

func lowerAlpha() string {
	characters := make([]byte, 26)
	for i := range characters {
		characters[i] = 'a' + byte(i)
	}
	return string(characters)
}

func main() {
	lines := util.LoadInputFile(day)
	if len(lines) > 1 {
		panic("Should not happen")
	}

	fmt.Printf("There are originally %d units\n", len(lines[0]))
	line := lines[0]
	reactedLine := react(line)
	fmt.Printf("The remaining line has %d units (first star)\n", len(reactedLine))

	alphas := lowerAlpha()
	mostReactedCharacter := string(alphas[0])
	smallestResultingLineSize := len(line)
	for _, character := range alphas {
		s := string(character)
		copiedLine := line
		copiedLine = strings.Replace(copiedLine, s, "", -1)
		copiedLine = strings.Replace(copiedLine, strings.ToUpper(s), "", -1)
		length := len(react(copiedLine))
		if length < smallestResultingLineSize {
			smallestResultingLineSize = length
			mostReactedCharacter = s
		}
	}
	fmt.Printf("The most reacted character is %s with length %d\n", mostReactedCharacter, smallestResultingLineSize)
}
