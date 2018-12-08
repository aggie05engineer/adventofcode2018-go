package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"unicode/utf8"
)

const day string = "day2"
const fileName string = "input.txt"

func getFile() *os.File {
	wd, err := os.Getwd()
	filePath := filepath.Join(wd, day, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return file
}

func StringMatch(str string, count int) bool {
	characters := make(map[int32]int8)
	for _, r := range str {
		characters[r]++
	}

	for _, c := range characters {
		if int(c) == count {
			return true
		}
	}

	return false
}

func firstStar() int {
	file := getFile()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	file.Close()

	scanner := bufio.NewScanner(bytes.NewReader(b))
	var count2 int
	var count3 int
	for scanner.Scan() {
		text := scanner.Text()
		if StringMatch(text, 2) {
			fmt.Printf("%s matches 2\n", text)
			count2++
		}
		if StringMatch(text, 3) {
			fmt.Printf("%s matches 3\n", text)
			count3++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return count2 * count3
}

func StringCompare(a string, b string) bool {
	var minLen int
	var minString *string
	var maxString *string
	if utf8.RuneCountInString(a) < utf8.RuneCountInString(b) {
		minLen = utf8.RuneCountInString(a)
		minString = &a
		maxString = &b
	} else {
		minLen = utf8.RuneCountInString(b)
		minString = &b
		maxString = &a
	}

	minStringSlice := []rune(*minString)
	maxStringSlice := []rune(*maxString)
	var diff int
	for i := 0; i < minLen; i++ {
		min := minStringSlice[i]
		max := maxStringSlice[i]
		if min != max {
			diff++
		}

		if diff > 2 {
			break
		}
	}

	return diff <= 1
}

func secondStar() string {
	file := getFile()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var found bool
	var left string
	var right string
	for i := 0; i < len(lines); i++ {
		left = lines[i]
		for j := i + 1; j < len(lines); j++ {
			right = lines[j]
			if StringCompare(left, right) {
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	if !found {
		panic("Could not find matching lines")
	}

	fmt.Printf("%s and %s\n", left, right)
	var same []rune
	for i, char := range []rune(left) {
		if char == []rune(right)[i] {
			same = append(same, char)
		}
	}

	return string(same)
}

func main() {
	checksum := firstStar()
	fmt.Printf("First checksum is %d\n", checksum)
	commonChars := secondStar()
	fmt.Printf("Second common characters is %s\n", commonChars)
}
