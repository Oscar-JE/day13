package compare

import (
	"day13/compare/element"
)


func PairsInOrder(first string, second string) bool {
	firstElement := element.Init(first)
	secondElement := element.Init(second)
	return inOrder(firstElement, secondElement)
}

func inOrder(first element.Element , second element.Element) bool {
	if second.IsEmpty() { // kan problemt ligga ilist jämförelsen ? 
		if first.IsEmpty() {
			return true
		} else {
			return false
		}
	}
	if first.IsEmpty() {
		return true
	}
	if first.IsNumber() && second.IsNumber() {
		return first.NumberCompare(second)
	}

	if first.IsList() && second.IsNumber() {
		second.UpgradeToList()
		return inOrder(first, second)
	}

	if first.IsNumber() && second.IsList() {
		first.UpgradeToList()
		return inOrder(first, second)
	}

	subElemensFirst := first.DivideIntoElements() // kan felet ligga här ? 
	subElementsSecond := second.DivideIntoElements()

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

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
