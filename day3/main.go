package main

import (
	"adventofcode2018-go/util"
	"fmt"
	"regexp"
	"strconv"
)

const day string = "day3"

type Claim struct {
	id     string
	x      int
	y      int
	width  int
	height int
}

func buildClaims(lines *[]string) []Claim {
	// #1 @ 829,837: 11x22
	var re = regexp.MustCompile(`^#(?P<id>\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<width>\d+)x(?P<height>\d+)$`)
	claims := make([]Claim, 0)
	for _, line := range *lines {
		if !re.MatchString(line) {
			panic(fmt.Sprintf("%s does not match", line))
		}
		match := re.FindStringSubmatch(line)
		result := make(map[string]string)
		for i, name := range match {
			result[re.SubexpNames()[i]] = name
		}
		claim := Claim{}
		claim.id = result["id"]
		claim.x, _ = strconv.Atoi(result["x"])
		claim.y, _ = strconv.Atoi(result["y"])
		claim.width, _ = strconv.Atoi(result["width"])
		claim.height, _ = strconv.Atoi(result["height"])
		claims = append(claims, claim)
	}
	return claims
}

func squareInchesCovered(claims []Claim) int {
	var fabric [1000][1000]int
	for _, claim := range claims {
		for x := claim.x; x < claim.x+claim.width; x++ {
			for y := claim.y; y < claim.y+claim.height; y++ {
				fabric[x][y]++
			}
		}
	}

	var squareInchesCovered int
	for i := 0; i < len(fabric); i++ {
		for j := 0; j < len(fabric); j++ {
			if fabric[i][j] > 1 {
				squareInchesCovered++
			}
		}
	}
	return squareInchesCovered
}

func findNonOverlappedClaim(claims []Claim) string {
	const overlapped = "overlapped"
	var badClaims = make(map[string]bool)
	var fabric [1000][1000]string
	for _, claim := range claims {
		for x := claim.x; x < claim.x+claim.width; x++ {
			for y := claim.y; y < claim.y+claim.height; y++ {
				id := fabric[x][y]
				if id == "" {
					fabric[x][y] = claim.id
				} else {
					badClaims[id] = true
					badClaims[claim.id] = true
					fabric[x][y] = overlapped
				}
			}
		}
	}

	nonOverlappedClaims := make(map[string]bool)
	for i := 0; i < len(fabric); i++ {
		for j := 0; j < len(fabric); j++ {
			id := fabric[i][j]
			if id == "" {
				continue
			}
			if id == overlapped {
				continue
			}
			if !badClaims[id] {
				nonOverlappedClaims[id] = true
			}
		}
	}

	if len(nonOverlappedClaims) > 1 {
		fmt.Printf("%v\n", nonOverlappedClaims)
		fmt.Printf("%d overlapped\n", len(nonOverlappedClaims))
		panic("Should not happen per puzzle")
	}
	for id := range nonOverlappedClaims {
		return id
	}
	panic("should not happen")
}

func main() {
	lines := util.LoadInputFile(day)
	claims := buildClaims(&lines)
	squareInchesCovered := squareInchesCovered(claims)
	fmt.Printf("There are %d square inches covered", squareInchesCovered)
	nonOverlappedClaim := findNonOverlappedClaim(claims)
	fmt.Printf("The only claim not overlapped is %s", nonOverlappedClaim)
}
