package element

import (
	"day13/compare/inputrep"
)

type Element struct {
	symbols []inputrep.InputRep
}

func InitElement(symbols []inputrep.InputRep) Element {
	return Element{symbols: symbols}
}

func Init(symbols string) Element {
	parsedSymbols := inputrep.ParseToInputRep(symbols)
	return InitElement(parsedSymbols)
}

func (e Element) IsEmpty() bool {
	return len(e.symbols) == 0
}

func (e Element) IsNumber() bool {
	return len(e.symbols) == 1 && e.symbols[0].IsNumer()
}

func (e Element) NumberCompare(other Element) bool {
	return e.symbols[0].GetValue() <= other.symbols[0].GetValue()
}

func (e Element) IsList() bool {
	return e.symbols[0].Equal(inputrep.InitLeft()) && e.symbols[len(e.symbols)-1].Equal(inputrep.InitRight())
}

func (e *Element) UpgradeToList() {
	e.symbols = append(e.symbols, inputrep.InitRight())
	e.symbols = append([]inputrep.InputRep{inputrep.InitLeft()}, e.symbols...)
}

func (e Element) DivideIntoElements() []Element {
	if e.symbols[0].Equal(inputrep.InitLeft()) && e.symbols[len(e.symbols)-1].Equal(inputrep.InitRight()) {
		e.symbols = e.symbols[1 : len(e.symbols)-1]
	}
	elements := []Element{}
	buffer := []inputrep.InputRep{}
	bracketLevel := 0
	for i := range e.symbols {
		if e.symbols[i].Equal(inputrep.InitLeft()) {
			bracketLevel++
		}
		if e.symbols[i].Equal(inputrep.InitRight()) {
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
				buffer = []inputrep.InputRep{}
			} else {
				elements = append(elements, InitElement([]inputrep.InputRep{e.symbols[i]}))
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
		isEqual = isEqual && e.symbols[i].Equal(other.symbols[i])
	}
	return isEqual
}

func sliceElementEqual(vec1 []Element, vec2 []Element) bool {
	isEqual := true
	if len(vec1) != len(vec2) {
		return false
	}
	for i := 0; i < len(vec1); i++ {
		isEqual = isEqual && vec1[i].Equal(vec2[i])
	}
	return isEqual
}
