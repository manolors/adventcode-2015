package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type GiftSize struct {
	l int
	w int
	h int
}

func calcDimensions(g GiftSize) int {
	return 2*g.l*g.w + 2*g.w*g.h + 2*g.h*g.l
}

func getSmallestSide(g GiftSize) int {
	if g.l > g.w && g.l > g.h {
		return g.w * g.h
	}

	if g.w > g.h {
		return g.l * g.h
	}

	return g.l * g.w
}

func getRibbonSize(g GiftSize) int {
	if g.l > g.w && g.l > g.h {
		return g.w*2 + g.h*2
	}

	if g.w > g.h {
		return g.l*2 + g.h*2
	}

	return g.l*2 + g.w*2
}

func getCubicVolume(g GiftSize) int {
	return g.l * g.w * g.h
}

func lineToDimensions(s string) GiftSize {
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
	return GiftSize{l, w, h}
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
			gift := lineToDimensions(line)
			totalPaper = totalPaper + calcDimensions(gift) + getSmallestSide(gift)
			ribbonSize = ribbonSize + getSmallestSide(gift) + getCubicVolume(gift)
			line = ""
		} else {
			line = line + string(file[index])
		}
	}

	fmt.Println("Total Paper: ", totalPaper)
	fmt.Println("Ribbon Size: ", ribbonSize)

}
