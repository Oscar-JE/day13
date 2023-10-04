package compare

import (
	"strconv"
	"day13/compare/inputrep"
)


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
