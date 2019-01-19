package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func getMapIndex(dir string, x *int, y *int) string {
	switch dir {
	case "<":
		*x = *x - 1
		break
	case ">":
		*x = *x + 1
		break
	case "^":
		*y = *y + 1
		break
	case "v":
		*y = *y - 1
		break
	}
	return strconv.Itoa(*x) + "," + strconv.Itoa(*y)
}

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	xSanta, ySanta, xRobosanta, yRobosanta := 0, 0, 0, 0

	m := make(map[string]int)
	m["0,0"] = 1
	isSantaTurn := true
	mapIndex := ""
	for _, direction := range file {
		if isSantaTurn {
			mapIndex = getMapIndex(string(direction), &xSanta, &ySanta)
		} else {
			mapIndex = getMapIndex(string(direction), &xRobosanta, &yRobosanta)
		}
		isSantaTurn = !isSantaTurn

		val, ok := m[mapIndex]
		if ok {
			m[mapIndex] = val + 1
		} else {
			m[mapIndex] = 1
		}
	}

	fmt.Println("Houses visited: ", len(m))

}
