package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	floor := 0
	for _, rune := range file {
		switch string(rune) {
		case "(":
			floor = floor + 1
			break
		case ")":
			floor = floor - 1
			break
		}
	}

	fmt.Println("Floor ", floor)
}
