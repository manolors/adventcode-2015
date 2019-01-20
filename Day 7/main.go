package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var wires map[string]int

func getData(s string) {
	data := strings.Split(s, " -> ")
	destinationWire := data[1]
	wires[destinationWire] = 0

	signal, err := strconv.Atoi(data[0])

	if err != nil {
		data := strings.Split(s, " ")

		if len(data) == 2 {
			operation := "NOT"
			operand1 := data[1]
		} else {
			operand1 := data[0]
			operation := data[1]
			operand2 := data[2]

		}
	}
}

func main() {
	wires := make(map[string]int)
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
