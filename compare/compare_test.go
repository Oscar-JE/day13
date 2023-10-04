package compare

import (
	"testing"
)

func TestCompare1(t *testing.T) {
	a := "[1,1,3,1,1]"
	b := "[1,1,5,1,1]"
	if !PairsInOrder(a, b) {
		t.Errorf("error in case 1")
	}
}

func TestCompare2(t *testing.T) {
	a := "[[1],[2,3,4]]"
	b := "[[1],4]"
	if !PairsInOrder(a, b) {
		t.Errorf("error in case 2 ")
	}
}

func TestCompare21(t *testing.T) {
	a := "[[1],[4,3,2]]"
	b := "[[1],4]"
	if PairsInOrder(a, b) {
		t.Errorf("error in case 22 ")
	}
}

func TestCompare3(t *testing.T) {
	a := "[9]"
	b := "[[8,7,6]]"
	if PairsInOrder(a, b) {
		t.Errorf("error in case 3 ")
	}
}

func TestCompare4(t *testing.T) {
	a := "[[4,4],4,4]"
	b := "[[4,4],4,4,4]"
	if !PairsInOrder(a, b) {
		t.Errorf("error in case 4 ")
	}
}

func TestCompare5(t *testing.T) {
	a := "[7,7,7,7]"
	b := "[7,7,7]"
	if PairsInOrder(a, b) {
		t.Errorf("error in case 5")
	}
}

func TestCompare6(t *testing.T) {
	a := "[]"
	b := "[3]"
	if !PairsInOrder(a, b) {
		t.Errorf("error in case 6")
	}
}

func TestCompare7(t *testing.T) {
	a := "[[[]]]"
	b := "[[]]"
	if PairsInOrder(a, b) {
		t.Errorf("error in case 7")
	}
}

func TestCompare8(t *testing.T) {
	a := "[1,[2,[3,[4,[5,6,7]]]],8,9]"
	b := "[1,[2,[3,[4,[5,6,0]]]],8,9]"
	if PairsInOrder(a, b) {
		t.Errorf("error in case 8")
	}
}

func TestIntresstingCase(t *testing.T) {
	a := "[[1],[2,3,4]]"
	b := "[[1],2,3,4]"
	if PairsInOrder(a, b) {
		t.Errorf("Error in intressting example")
	}
}

func TestWTFCase(t *testing.T) {
	a := "[[8,[[7,10,10,5],[8,4,9]],3,5],[[[3,9,4],5,[7,5,5]],[[3,2,5],[10],[5,5],0,[8]]],[4,2,[],[[7,5,6,3,0],[4,4,10,7],6,[8,10,9]]],[[4,[],4],10,1]]"
	b := "[[[[8],[3,10],[7,6,3,7,4],1,8]]]"
	if !PairsInOrder(a, b) {
		t.Errorf("Error in wtf case")
	}
}

func TestCompareSameVector(t *testing.T) {
	a := "[1,1]"
	b := "[1,1]"
	if !PairsInOrder(a, b) {
		t.Errorf("error in case same vector")
	}
}

func TestParseToInputRep(t *testing.T) {
	a := "[1,12]"
	res := parseToInputRep(a)
	expected := []InputRep{initLeft(), initNumber(1), initNumber(12), initRight()}
	if !sliceEqual(res, expected) {
		t.Errorf("error in simple conversion to InputRep")
	}
}

func TestComplexParseToInputRep(t *testing.T) {
	a := "[1,[2,[],23]"
	expected := []InputRep{initLeft(), initNumber(1), initLeft(), initNumber(2), initLeft(), initRight(), initNumber(23), initRight()}
	res := parseToInputRep(a)
	if !sliceEqual(expected, res) {
		t.Errorf("error in complex conversion to InputRep")
	}
}

func TestDivideIntoElementsSimple(t *testing.T) {
	input := parseToInputRep("[1,2]")
	inputElement := InitElement(input)
	a := []InputRep{initNumber(1)}
	b := []InputRep{initNumber(2)}
	elementA := InitElement(a)
	elementB := InitElement(b)
	expected := []Element{
		elementA,
		elementB,
	}
	actual := inputElement.divideIntoElements()
	if !sliceElementEqual(actual, expected) {
		t.Errorf("Error devide into elements")
	}
}




func sliceEqual(vec1 []InputRep, vec2 []InputRep) bool {
	isEqual := true
	maxLen := max(len(vec1), len(vec2))
	for i := 0; i < maxLen; i++ {
		isEqual = isEqual && equal(vec1[i], vec2[i])
	}
	return isEqual
}

func sliceElementEqual(vec1 []Element, vec2 []Element) bool{
	isEqual := true
	if len(vec1) != len(vec2) {
		return false
	}
	for i := 0; i < len(vec1); i++ {
		isEqual = isEqual && vec1[i].Equal(vec2[i])
	}
	return isEqual
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
