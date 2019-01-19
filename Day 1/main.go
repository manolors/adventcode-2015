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
	position := -1
	for i := 1; i <= len(file); i++ {
		switch string(file[i-1]) {
		case "(":
			floor = floor + 1
			break
		case ")":
			floor = floor - 1
			break
		}

		if floor < 0 && position < 0 {
			position = i
		}
	}

	fmt.Println("Floor ", floor)
	fmt.Println("Position ", position)
}
