package main

import (
	"day13/parse"
	"testing"
)

func TestPart1Short(t *testing.T) {
	inputs := parse.Parse("input_short.txt")
	sum := part1(inputs)
	if sum != 13 {
		t.Errorf("Wrong result in part one short data , was %d , should be %d", sum, 13)
	}
}

func TestMain(t *testing.T){ // enbart för debugg möjligheter
	main()
}
