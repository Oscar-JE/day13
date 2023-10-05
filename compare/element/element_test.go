package element

import "testing"

func TestDivideIntoElementsComplex(t *testing.T) {
	input := Init("[[1,2,3],[2,[4,5,6]]]")
	a := Init("[1,2,3]")
	b := Init("[2,[4,5,6]]")
	expected := []Element{}
	expected = append(expected, a, b)
	actual := input.DivideIntoElements()
	if !sliceElementEqual(expected, actual) {
		t.Errorf("error in Divide into elements complex")
	}
}


func TestDivideIntoElementsSimple(t *testing.T) {
	inputElement := Init("[1,2]")
	elementA := Init("1")
	elementB := Init("2")
	expected := []Element{
		elementA,
		elementB,
	}
	actual := inputElement.DivideIntoElements()
	if !sliceElementEqual(actual, expected) {
		t.Errorf("Error devide into elements")
	}
}
