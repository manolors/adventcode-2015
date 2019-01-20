package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func inc(i *int) {
	*i = *i + 1
}
func checkNiceVowels(s string) bool {
	matches := 0
	for _, rune := range s {
		if strings.ContainsAny(string(rune), "a & e & i & o & u") {
			inc(&matches)
		}
	}
	return matches >= 3
}

func checkDuplicateLetter(s string) bool {
	prevRune := s[0]
	for index := 1; index < len(s); index++ {
		if s[index] == prevRune {
			return true
		}
		prevRune = s[index]
	}
	return false
}
func checkBlacklist(s string) bool {
	blacklist := []string{"ab", "cd", "pq", "xy"}
	for _, rune := range blacklist {
		if strings.Contains(s, rune) {
			return false
		}
	}
	return true
}

func isNiceString(s string) bool {
	// checkVowels := checkNiceVowels(s)
	// checkDupes := checkDuplicateLetter(s)
	// checkBlack := checkBlacklist(s)
	checkDupes2 := checkDuplicatedPair(s)
	return checkDupes2
}

func checkDuplicatedPair(s string) bool {
	pairs := make(map[string]int)
	prevRune := string(s[0])
	first := true
	for _, rune := range s {
		if first {
			first = false
			continue
		}
		_, ok := pairs[prevRune+string(rune)]

		if !ok {
			pairs[prevRune+string(rune)] = 1
			prevRune = string(rune)
		} else {
			return true
		}
	}
	return false
}

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	line := ""
	nice := 0
	for index := 0; index <= len(file); index++ {
		if index == len(file) || file[index] == 10 {
			if isNiceString(line) {
				inc(&nice)
			}
			line = ""
		} else {
			line = line + string(file[index])
		}
	}

	fmt.Println("Nice strings:", nice)
}
