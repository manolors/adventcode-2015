package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var wires map[string]uint16

func getData(s string) {
	wires := make(map[string]uint16)

	data := strings.Split(s, " -> ")
	destinationWire := data[1]
	wires[destinationWire] = 0

	data = strings.Split(data[0], " ")

	signal, err := strconv.Atoi(data[0])

	if err != nil {
		if len(data) == 2 {
			wires[destinationWire] ^= wires[data[1]]
		} else {
			wire1, _ := wires[data[0]]
			wire2, _ := wires[data[2]]

			switch data[1] {
			case "AND":
				wires[destinationWire] = wire1 & wire2
				break
			case "OR":
				wires[destinationWire] = wire1 | wire2
				break
			case "LSHIFT":
				wires[destinationWire] = wire1 << wire2
				break
			case "RSHIFT":
				wires[destinationWire] = wire1 >> wire2
				break
			}
		}
	} else {
		wires[data[1]] = uint16(signal)
	}
}

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	line := ""

	for index := 0; index <= len(file); index++ {
		if index == len(file) || file[index] == 10 {
			getData(line)
			line = ""
		} else {
			line = line + string(file[index])
		}
	}
}
