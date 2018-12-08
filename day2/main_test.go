package main

import "testing"

func TestStringMatchShouldFail(t *testing.T) {
	const line string = "abc"
	result := StringMatch(line, 2)
	if result {
		t.Errorf("Should have failed")
	}
}

func TestStringMatchShouldPassWithCount2(t *testing.T) {
	const line string = "aabc"
	result := StringMatch(line, 2)
	if !result {
		t.Errorf("Should have passed")
	}
}

func TestStringMatchShouldFailWithCount2(t *testing.T) {
	const line string = "aaabc"
	result := StringMatch(line, 2)
	if result {
		t.Errorf("Should have failed")
	}
}

func TestStringMatchShouldPassWithCount3(t *testing.T) {
	const line string = "aaabc"
	result := StringMatch(line, 3)
	if !result {
		t.Errorf("Should have passed")
	}
}

func TestStringMatchTwoMatchesShouldPassWithCount2(t *testing.T) {
	const line string = "aabbc"
	result := StringMatch(line, 2)
	if !result {
		t.Errorf("Should have passed")
	}
}

func TestStringCompare(t *testing.T) {
	const string1 string = "abcde"
	const string2 string = "abZde"
	result := StringCompare(string1, string2)
	if !result {
		t.Errorf("Should have passed")
	}
}

func TestStringMatchIndexFail(t *testing.T) {
	const string1 string = "abcde"
	const string2 string = "abZdF"
	result := StringCompare(string1, string2)
	if result {
		t.Errorf("Should have failed")
	}
}
