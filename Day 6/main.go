package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const Size = 1000

var lights [Size][Size]bool

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

func handleLights(s string, b bool) bool {
	switch s {
	case "turn on":
		return true
	case "turn off":
		return false
	case "toggle":
		return !b
	}
	return true
}

func processAction(a Action) bool {
	for i := a.start.x; i <= a.end.x; i++ {
		for j := a.start.y; j <= a.end.y; j++ {
			lights[i][j] = handleLights(a.action, lights[i][j])
		}
	}
	return true
}

func countLitLights() int {
	count := 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if lights[i][j] {
				count = count + 1
			}
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

	fmt.Println("Lit Lights:", countLitLights())
}
