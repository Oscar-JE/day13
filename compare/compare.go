package compare

import (
	"strconv"
)

type InputRep struct {
	isLeft  bool
	isRight bool
	value   int
}

func initLeft() InputRep {
	return InputRep{true, false, -1}
}

func initRight() InputRep {
	return InputRep{false, true, -1}
}

func initNumber(number int) InputRep {
	return InputRep{false, false, number}
}

func PairsInOrder(first string, second string) bool {
	a := parseToInputRep(first)
	b := parseToInputRep(second)
	return inOrder(a, b)
}

func inOrder(first []InputRep, second []InputRep) bool {
	return false
}

func equal(a InputRep, b InputRep) bool {
	return a.isLeft == b.isLeft && a.isRight == b.isRight && a.value == b.value
}

func parseToInputRep(line string) []InputRep {
	representation := []InputRep{}
	numberBuffer := ""
	for index := range line {
		if line[index] == '[' {
			representation = append(representation, initLeft())
		} else if line[index] == ']' {
			if numberBuffer != "" {
				value, errorConv := strconv.Atoi(numberBuffer)
				if errorConv != nil {
					panic("parsing error")
				}
				representation = append(representation, initNumber(value))
			}
			representation = append(representation, initRight())
		} else if line[index] == ',' {
			if numberBuffer != "" {
				value, errorConv := strconv.Atoi(numberBuffer)
				if errorConv != nil {
					panic("parsing error")
				}
				representation = append(representation, initNumber(value))
				numberBuffer = ""
			}
		} else {
			numberBuffer += string(line[index])
		}
	}
	return representation
}

func divideIntoElements(sequence []InputRep) [][]InputRep {
	if equal(sequence[0], initLeft()) && equal(sequence[len(sequence)-1], initRight()) {
		sequence = sequence[1 : len(sequence)-1]
	}
	elements := [][]InputRep{}
	buffer := []InputRep{}
	bracketLevel := 0
	for i := range sequence {
		if equal(sequence[i], initLeft()) {
			bracketLevel++
		}
		if equal(sequence[i], initLeft()) {
			bracketLevel--
			if bracketLevel < 0 {
				panic("inposible bracketLevel")
			}
		}
		if bracketLevel > 0 {
			buffer = append(buffer, sequence[i])
		} else {
			if len(buffer) != 0 {
				buffer = append(buffer, sequence[i])
				elements = append(elements, buffer)
				buffer = []InputRep{}
			} else {
				elements = append(elements, []InputRep{sequence[i]})
			}
		}
	}
	return elements
}
