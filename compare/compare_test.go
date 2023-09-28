package compare

import (
	"testing"
)

func TestCompare1(t *testing.T) {
	a := "[1,1,3,1,1]"
	b := "[1,1,5,1,1]"
	if !pairsInOrder(a, b) {
		t.Errorf("error in case 1")
	}
}

func TestCompare2(t *testing.T) {
	a := "[[1],[2,3,4]]"
	b := "[[1],4]"
	if !pairsInOrder(a, b) {
		t.Errorf("error in case 2 ")
	}
}

func TestCompare3(t *testing.T) {
	a := "[9]"
	b := "[[8,7,6]]"
	if pairsInOrder(a, b) {
		t.Errorf("error in case 3 ")
	}
}

func TestCompare4(t *testing.T) {
	a := "[[4,4],4,4]"
	b := "[[4,4],4,4,4]"
	if !pairsInOrder(a, b) {
		t.Errorf("error in case 4 ")
	}
}

func TestCompare5(t *testing.T) {
	a := "[7,7,7,7]"
	b := "[7,7,7]"
	if pairsInOrder(a, b) {
		t.Errorf("error in case 5")
	}
}

func TestCompare6(t *testing.T) {
	a := "[]"
	b := "[3]"
	if !pairsInOrder(a, b) {
		t.Errorf("error in case 6")
	}
}

func TestCompare7(t *testing.T) {
	a := "[[[]]]"
	b := "[[]]"
	if pairsInOrder(a, b) {
		t.Errorf("error in case 7")
	}
}

func TestCompare8(t *testing.T) {
	a := "[1,[2,[3,[4,[5,6,7]]]],8,9]"
	b := "[1,[2,[3,[4,[5,6,0]]]],8,9]"
	if pairsInOrder(a, b) {
		t.Errorf("error in case 8")
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

func sliceEqual(vec1 []InputRep, vec2 []InputRep) bool {
	isEqual := true
	maxLen := max(len(vec1), len(vec2))
	for i := 0; i < maxLen; i++ {
		isEqual = isEqual && equal(vec1[i], vec2[i])
	}
	return isEqual
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
