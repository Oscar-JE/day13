package element

import "testing"

func TestDivideIntoElementsComplex(t *testing.T) {
	input := parseToInputRep("[[1,2,3],[2,[4,5,6]]]")
	a := parseToInputRep("[1,2,3]")
	b := parseToInputRep("[2,[4,5,6]]")
	expected := []Element{}
	inA := InitElement(a)
	inB := InitElement(b)
	inputElement := InitElement(input)
	expected = append(expected, inA, inB)
	actual := inputElement.divideIntoElements()
	if !sliceElementEqual(expected, actual) {
		t.Errorf("error in Divide into elements complex")
	}
}
