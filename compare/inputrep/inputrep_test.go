package inputrep

import(
	"testing"
)

func TestParseToInputRep(t *testing.T) {
	a := "[1,12]"
	res := ParseToInputRep(a)
	expected := []InputRep{InitLeft(), InitNumber(1), InitNumber(12), InitRight()}
	if !sliceEqual(res, expected) {
		t.Errorf("error in simple conversion to InputRep")
	}
}

func TestComplexParseToInputRep(t *testing.T) {
	a := "[1,[2,[],23]"
	expected := []InputRep{InitLeft(), InitNumber(1), InitLeft(), InitNumber(2), InitLeft(), InitRight(), InitNumber(23), InitRight()}
	res := ParseToInputRep(a)
	if !sliceEqual(expected, res) {
		t.Errorf("error in complex conversion to InputRep")
	}
}


func sliceEqual(vec1 []InputRep, vec2 []InputRep) bool {
	isEqual := true
	maxLen := max(len(vec1), len(vec2))
	for i := 0; i < maxLen; i++ {
		isEqual = isEqual && vec1[i].Equal(vec2[i])
	}
	return isEqual
}