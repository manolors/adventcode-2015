package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func calcDimensions(l int, w int, h int) int {
	return 2*l*w + 2*w*h + 2*h*l
}

func getSmallestSide(l int, w int, h int) int {
	if l > w && l > h {
		return w * h
	}

	if w > h {
		return l * h
	}

	return l * w
}

func getRibbonSize(l int, w int, h int) int {
	if l > w && l > h {
		return w*2 + h*2
	}

	if w > h {
		return l*2 + h*2
	}

	return l*2 + w*2
}

func getCubicVolume(l int, w int, h int) int {
	return l * w * h
}

func lineToDimensions(s string) (int, int, int) {
	dimensions := strings.Split(s, "x")
	l, err := strconv.Atoi(dimensions[0])
	if err != nil {
		panic(err)
	}
	w, err := strconv.Atoi(dimensions[1])
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(dimensions[2])
	if err != nil {
		panic(err)
	}
	return l, w, h
}

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}
	line := ""
	totalPaper := 0
	ribbonSize := 0
	for index := 0; index <= len(file); index++ {
		if index == len(file) || file[index] == 10 {
			l, w, h := lineToDimensions(line)
			totalPaper = totalPaper + calcDimensions(l, w, h) + getSmallestSide(l, w, h)
			ribbonSize = ribbonSize + getRibbonSize(l, w, h) + getCubicVolume(l, w, h)
			line = ""
		} else {
			line = line + string(file[index])
		}
	}

	fmt.Println("Total Paper: ", totalPaper)
	fmt.Println("Ribbon Size: ", ribbonSize)

}
