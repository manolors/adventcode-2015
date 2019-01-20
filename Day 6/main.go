package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const Size = 1000

var lights [Size][Size]int

type Action struct {
	action string
	start  Coords
	end    Coords
}

type Coords struct {
	x int
	y int
}

func getAction(s string) string {
	actions := []string{"turn on", "turn off", "toggle"}
	for _, action := range actions {
		if strings.Contains(s, action) {
			return action
		}
	}
	return ""
}

func getCoords(s string) Coords {
	coords := strings.Split(s, ",")
	var c Coords
	c.x, _ = strconv.Atoi(coords[0])
	c.y, _ = strconv.Atoi(coords[1])
	return c
}

func getData(s string) Action {
	action := getAction(s)
	runes := strings.Split(strings.TrimPrefix(s, action+" "), " through ")
	return Action{action, getCoords(runes[0]), getCoords(runes[1])}
}

func handleLights(s string) int {
	intensities := map[string]int{
		"turn on":  1,
		"turn off": -1,
		"toggle":   2,
	}
	return intensities[s]
}

func processAction(a Action) bool {
	for i := a.start.x; i <= a.end.x; i++ {
		for j := a.start.y; j <= a.end.y; j++ {
			lights[i][j] = lights[i][j] + handleLights(a.action)
			if lights[i][j] < 0 {
				lights[i][j] = 0
			}
		}
	}
	return true
}

func countLitLights() int {
	count := 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			count = count + lights[i][j]
		}
	}
	return count
}

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	line := ""

	for index := 0; index <= len(file); index++ {
		if index == len(file) || file[index] == 10 {
			a := getData(line)
			processAction(a)
			line = ""
		} else {
			line = line + string(file[index])
		}
	}

	fmt.Println("Total intensity:", countLitLights())
}
