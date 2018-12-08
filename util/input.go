package util

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
)

const input string = "input.txt"

func LoadInputFile(day string) []string {
	wd, err := os.Getwd()
	filePath := filepath.Join(wd, day, input)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
