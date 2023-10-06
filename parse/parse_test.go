package parse

import(
	"testing"
)


func TestParseSimpleExample(t *testing.T){
	inputFile := "test_file.txt"
	inputs := Parse(inputFile)
	if inputs[0][0] != "[]" ||inputs[0][1] != "[]" ||inputs[1][0] != "[1]" ||inputs[1][1] != "[1]"{
		t.Errorf(" faulty pars")
	}
}