package compare

import (
	"strconv"
)

type InputRep struct {
	isLeft  bool
	isRight bool
	value   int
}

func (i *InputRep) IsNumer() bool {
	return !i.isLeft && !i.isRight
}

func (i *InputRep) IsLeft() bool {
	return i.isLeft
}

func (i *InputRep) IsRight() bool {
	return i.isRight
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
	if len(second) == 0 {
		if len(first) == 0 {
			return true
		} else {
			return false
		}
	}
	if len(first) == 0 {
		if len(second) == 0 {
			return true
		} else {
			return true
		}
	}
	if len(second) == 1 && len(first) == 1 {
		if second[0].IsNumer() && first[0].IsNumer() {
			return first[0].value <= second[0].value
		}
	}
	elemensFirst := divideIntoElements(first)
	elementsSecond := divideIntoElements(second)

	for i := 0; i < min(len(elemensFirst), len(elementsSecond)); i++ {
		e1 := elemensFirst[i]
		e2 := elementsSecond[i]
		if inOrder(e1, e2) && !inOrder(e2, e1) {
			return true
		} else if inOrder(e2, e1) && !inOrder(e1, e2) {
			return false
		}
	}
	return len(elemensFirst) <= len(elementsSecond)
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
				numberBuffer = ""
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
		if equal(sequence[i], initRight()) {
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

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
