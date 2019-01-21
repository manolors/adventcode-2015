package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var wires = map[string]uint16{}

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
	wire        string
}

// 100 -> z
type WireAssignment struct {
	signalWire      string
	destinationWire string
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

	data := strings.Split(s, " -> ")

	_, err := strconv.Atoi(data[0])

	if err != nil {
		return "WireAssignment"
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

func getSignalAssignment(s string, w string) SignalAssignment {
	signalValue, _ := strconv.Atoi(s)
	return SignalAssignment{uint16(signalValue), w}
}

func getWireAssignment(s string, w string) WireAssignment {
	return WireAssignment{strings.Split(s, " -> ")[0], w}
}

func getNegateSignal(s string) NegateSignal {
	return NegateSignal{strings.TrimPrefix(strings.Split(s, " -> ")[0], "NOT ")}
}

func getSignalsFromWires(wire1 string, wire2 string) (uint16, uint16) {
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
	signal1, signal2 := getSignalsFromWires(lo.wire1, lo.wire2)
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

func (wa WireAssignment) apply() uint16 {
	return wires[wa.signalWire]
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

func applyOperation(s string) {
	data := strings.Split(s, " -> ")
	destinationWire := data[1]

	_, ok := wires[destinationWire]
	// initialize all destination wires
	if !ok {
		wires[destinationWire] = 0
	}
	operationType := getOperationType(s)
	switch operationType {
	case "LogicOperation":
		wires[destinationWire] = getLogicOperation(data[0]).apply()
	case "ShiftOperation":
		wires[destinationWire] = getShiftOperation(data[0]).apply()
	case "NegateSignal":
		wires[destinationWire] = getNegateSignal(data[0]).apply()
	case "SignalAssignment":
		wires[destinationWire] = getSignalAssignment(data[0], destinationWire).apply()
	case "WireAssignment":
		wires[destinationWire] = getWireAssignment(data[0], destinationWire).apply()
	}
}

func (lo LogicOperation) areWiresConnected() bool {
	// wire could be a number!
	_, err := strconv.Atoi(lo.wire1)
	ok1 := true
	ok2 := true
	if err != nil {
		_, ok1 = wires[lo.wire1]
	}
	_, err2 := strconv.Atoi(lo.wire2)

	if err2 != nil {
		_, ok2 = wires[lo.wire2]
	}

	return ok1 && ok2
}

func (so ShiftOperation) isWireConnected() bool {
	_, ok := wires[so.wire]
	return ok
}

func (ns NegateSignal) isWireConnected() bool {
	_, ok := wires[ns.wire]
	return ok
}

func (wa WireAssignment) isWireConnected() bool {
	_, ok := wires[wa.signalWire]
	return ok
}

func canApplyOperation(s string) bool {
	operationType := getOperationType(s)

	switch operationType {
	case "LogicOperation":
		return getLogicOperation(s).areWiresConnected()
	case "ShiftOperation":
		return getShiftOperation(s).isWireConnected()
	case "WireAssignment":
		return getWireAssignment(s, strings.Split(s, " -> ")[1]).isWireConnected()
	case "NegateSignal":
		return getNegateSignal(s).isWireConnected()
	case "SignalAssignment":
		return true
	}
	return false
}

func main() {
	file, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	operations := strings.Split(string(file), "\n")
	pendingOperations := []string{}
	loops := 0
	compares := 0

	// TODO - REFACTOR!! FUCKING INEFFICIENT!!
	// 113 Loops and 19191 compares...
	for len(operations) > 0 {
		loops++
		size := len(operations)
		for i := 0; i < size; i++ {
			compares++
			if canApplyOperation(operations[i]) {
				applyOperation(operations[i])
			} else {
				pendingOperations = append(pendingOperations, operations[i])
			}
		}
		operations = pendingOperations
		pendingOperations = nil
	}

	fmt.Println("Wire a:", wires["a"])
	fmt.Println("Loops:", loops)
	fmt.Println("Compares:", compares)
}
