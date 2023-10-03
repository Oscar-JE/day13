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

func equal(a InputRep, b InputRep) bool {
	return a.isLeft == b.isLeft && a.isRight == b.isRight && a.value == b.value
}

type Element struct {
	symbols []InputRep
}

func InitElement(symbols []InputRep) Element {
	return Element{symbols: symbols}
}

func (e Element) isEmpty() bool {
	return len(e.symbols) == 0
}

func (e Element) isNumber() bool {
	return len(e.symbols) == 1 && e.symbols[0].IsNumer()
}

func (e Element) NumberCompare(other Element) bool {
	return e.symbols[0].value <= other.symbols[0].value
}

func (e Element) isList() bool {
	return equal(e.symbols[0], initLeft()) && equal(e.symbols[len(e.symbols)-1], initRight())
}

func (e *Element) upgradeToList() {
	e.symbols = append(e.symbols, initRight())
	e.symbols = append([]InputRep{initLeft()}, e.symbols...)
}

func (e Element) divideIntoElements() []Element {
	if equal(e.symbols[0], initLeft()) && equal(e.symbols[len(e.symbols)-1], initRight()) {
		e.symbols = e.symbols[1 : len(e.symbols)-1]
	}
	elements := []Element{}
	buffer := []InputRep{}
	bracketLevel := 0
	for i := range e.symbols {
		if equal(e.symbols[i], initLeft()) {
			bracketLevel++
		}
		if equal(e.symbols[i], initRight()) {
			bracketLevel--
			if bracketLevel < 0 {
				panic("inposible bracketLevel")
			}
		}
		if bracketLevel > 0 {
			buffer = append(buffer, e.symbols[i])
		} else {
			if len(buffer) != 0 {
				buffer = append(buffer, e.symbols[i])
				newElement := InitElement(buffer)
				elements = append(elements, newElement)
				buffer = []InputRep{}
			} else {
				elements = append(elements, InitElement([]InputRep{e.symbols[i]}))
			}
		}
	}
	return elements
}

func (e Element) Equal(other Element) bool {
	isEqual := true
	if len(e.symbols) != len(other.symbols) {
		return false
	}
	for i := 0; i < len(e.symbols); i++ {
		isEqual = isEqual && equal(e.symbols[i], other.symbols[i])
	}
	return isEqual
}

func PairsInOrder(first string, second string) bool {
	symbolsFirst := parseToInputRep(first)
	symbolsSecond := parseToInputRep(second)
	firstElement := InitElement(symbolsFirst)
	secondElement := InitElement(symbolsSecond)
	return inOrder(firstElement, secondElement)
}

func inOrder(first Element, second Element) bool {
	if second.isEmpty() {
		if first.isEmpty() {
			return true
		} else {
			return false
		}
	}
	if first.isEmpty() {
		if second.isEmpty() {
			return true
		} else {
			return true
		}
	}
	if first.isNumber() && second.isNumber() {
		return first.NumberCompare(second)
	}

	if first.isList() && second.isNumber() {
		second.upgradeToList()
		return inOrder(first, second)
	}

	if first.isNumber() && second.isList() {
		first.upgradeToList()
		return inOrder(first, second)
	}

	subElemensFirst := first.divideIntoElements()
	subElementsSecond := second.divideIntoElements()

	for i := 0; i < min(len(subElemensFirst), len(subElementsSecond)); i++ {
		e1 := subElemensFirst[i]
		e2 := subElementsSecond[i]
		if inOrder(e1, e2) && !inOrder(e2, e1) {
			return true
		} else if inOrder(e2, e1) && !inOrder(e1, e2) {
			return false
		}
	}
	return len(subElemensFirst) <= len(subElementsSecond)
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

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
