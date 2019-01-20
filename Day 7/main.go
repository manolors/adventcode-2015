package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var wires = map[string]uint16{}

// x AND y -> z
type LogicOperation struct {
	operation string // AND - OR
	wire1     string
	wire2     string
}

// x RSHIFT 2 -> z
type ShiftOperation struct {
	operation   string // RSHIFT LSHIFT
	wire        string
	signalValue uint16
}

// 100 -> z
type SignalAssignment struct {
	signalValue uint16
}

// NOT x -> z
type NegateSignal struct {
	wire string
}

func getOperationType(s string) string {
	logicOperationKeyWords := []string{"AND", "OR"}
	for _, operation := range logicOperationKeyWords {
		if strings.Contains(s, operation) {
			return "LogicOperation"
		}
	}

	ShiftOperationKeyWords := []string{"RSHIFT", "LSHIFT"}
	for _, operation := range ShiftOperationKeyWords {
		if strings.Contains(s, operation) {
			return "ShiftOperation"
		}
	}

	if strings.HasPrefix(s, "NOT ") {
		return "NegateSignal"
	}

	return "SignalAssignment"
}

func getLogicOperation(s string) LogicOperation {
	values := strings.Split(s, " ")
	return LogicOperation{values[1], values[0], values[2]}
}

func getShiftOperation(s string) ShiftOperation {
	values := strings.Split(s, " ")
	shiftValue, _ := strconv.Atoi(values[2])
	return ShiftOperation{values[1], values[0], uint16(shiftValue)}
}

func getSignalAssignment(s string) SignalAssignment {
	shiftValue, _ := strconv.Atoi(s)
	return SignalAssignment{uint16(shiftValue)}
}

func getNegateSignal(s string) NegateSignal {
	return NegateSignal{strings.TrimPrefix(s, "NOT ")}
}

func getSignals(wire1 string, wire2 string) (uint16, uint16) {
	var signal int
	var signal1, signal2 uint16
	var err error
	signal, err = strconv.Atoi(wire1)
	if err != nil {
		signal1 = wires[wire1]
	} else {
		signal1 = uint16(signal)
	}
	signal, err = strconv.Atoi(wire2)
	if err != nil {
		signal2 = wires[wire2]
	} else {
		signal2 = uint16(signal)
	}
	return signal1, signal2
}

func (lo LogicOperation) apply() uint16 {
	signal1, signal2 := getSignals(lo.wire1, lo.wire2)
	switch lo.operation {
	case "AND":
		return signal1 & signal2
	case "OR":
		return signal1 | signal2
	}
	panic("Invalid LogicOperation operation")
}

func (sa SignalAssignment) apply() uint16 {
	return sa.signalValue
}

func (so ShiftOperation) apply() uint16 {
	switch so.operation {
	case "RSHIFT":
		return wires[so.wire] >> so.signalValue
	case "LSHIFT":
		return wires[so.wire] << so.signalValue
	}
	panic("Invalid ShiftOperation operation")
}

func (ns NegateSignal) apply() uint16 {
	return ^wires[ns.wire]
}

func getWireValue(s string) uint16 {
	switch getOperationType(s) {
	case "LogicOperation":
		return getLogicOperation(s).apply()
	case "ShiftOperation":
		return getShiftOperation(s).apply()
	case "NegateSignal":
		return getNegateSignal(s).apply()
	case "SignalAssignment":
		return getSignalAssignment(s).apply()
	}
	panic("Invalid Operation")
}

func applyOperation(s string) {
	data := strings.Split(s, " -> ")
	destinationWire := data[1]
	operation := data[0]

	_, ok := wires[destinationWire]

	if !ok {
		wires[destinationWire] = 0
	}

	wires[destinationWire] = getWireValue(operation)

}

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	line := ""

	for index := 0; index <= len(file); index++ {
		if index == len(file) || file[index] == 10 {
			applyOperation(line)
			line = ""
		} else {
			line = line + string(file[index])
		}
	}

	fmt.Println("Wire a:", wires["a"])
}
