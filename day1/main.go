package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

const day string = "day1"
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

func firstStar() int {
	file := getFile()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	file.Close()

	scanner := bufio.NewScanner(bytes.NewReader(b))
	var freq int = 0
	for scanner.Scan() {
		text := scanner.Text()
		delta, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		freq += delta
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return freq
}

func secondStar() int {
	file := getFile()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	file.Close()

	found := false
	var freq int = 0
	freqsSeen := make(map[int]bool)
	for i := 0; i < 10000 && !found; i++ {
		scanner := bufio.NewScanner(bytes.NewReader(b))
		for scanner.Scan() {
			text := scanner.Text()
			delta, err := strconv.Atoi(text)
			if err != nil {
				panic(err)
			}
			freq += delta
			if _, ok := freqsSeen[freq]; ok {
				found = true
				break
			}
			freqsSeen[freq] = true
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}

	if found != true {
		panic("did not find repeated freq")
	}

	return freq
}

func main() {
	freq := firstStar()
	fmt.Printf("First freq is %d\n", freq)
	freq = secondStar()
	fmt.Printf("Second freq is %d\n", freq)
}
