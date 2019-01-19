package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type Coords struct {
	x int
	y int
}

func getMapIndex(dir string, coord *Coords) string {
	switch dir {
	case "<":
		coord.x = coord.x - 1
		break
	case ">":
		coord.x = coord.x + 1
		break
	case "^":
		coord.y = coord.y + 1
		break
	case "v":
		coord.y = coord.y - 1
		break
	}
	return strconv.Itoa(coord.x) + "," + strconv.Itoa(coord.y)
}

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	coordSanta, coordRoboSanta := Coords{0, 0}, Coords{0, 0}

	m := make(map[string]int)
	m["0,0"] = 1
	isSantaTurn := true
	mapIndex := ""
	for _, direction := range file {
		if isSantaTurn {
			mapIndex = getMapIndex(string(direction), &coordSanta)
		} else {
			mapIndex = getMapIndex(string(direction), &coordRoboSanta)
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
