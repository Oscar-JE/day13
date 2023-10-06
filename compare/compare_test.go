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

func TestCompareSameVector(t *testing.T) {
	a := "[1,1]"
	b := "[1,1]"
	if !PairsInOrder(a, b) {
		t.Errorf("error in case same vector")
	}
}

func TestFrånCC(t *testing.T) {
	a := "[[1,1,1],2,3]"
	b := "[[1,1,1]]"

	if PairsInOrder(a, b) {
		t.Errorf("shold be in determed to be in right order should be in wrong order")
	}
}

func TestListorIListor(t *testing.T) {
	if !PairsInOrder("[[[1]],1]", "[[1],2]") {
		t.Errorf("wrong")
	}

	if PairsInOrder("[[[1]],2]", "[[1],1]") {
		t.Errorf("wrong")
	}

	if !PairsInOrder("[[1],1]", "[[[1]],2]") {
		t.Errorf("wrong")
	}

	if PairsInOrder("[[1],2]", "[[[1]],1]") {
		t.Errorf("wrong")
	}
}
