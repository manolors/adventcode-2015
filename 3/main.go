package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type Coords struct {
	x, y int
}

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	x, y := 0, 0
	m := make(map[string]int)
	m["0,0"] = 1
	for _, direction := range file {
		dir := string(direction)

		switch dir {
		case "<":
			x = x - 1
			break
		case ">":
			x = x + 1
			break
		case "^":
			y = y + 1
			break
		case "v":
			y = y - 1
			break
		}

		mapIndex := strconv.Itoa(x) + "," + strconv.Itoa(y)

		val, ok := m[mapIndex]
		if ok {
			m[mapIndex] = val + 1
		} else {
			m[mapIndex] = 1
		}
	}

	fmt.Println("Houses visited: ", len(m))

}
