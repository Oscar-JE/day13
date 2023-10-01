package main

import (
	"day13/compare"
	"day13/parse"
	"fmt"
)

func main() {
	inputs := parse.Parse("input_short.txt")
	sum := 0
	for index, firstAndSecond := range inputs {
		if compare.PairsInOrder(firstAndSecond[0], firstAndSecond[1]) {
			sum += index + 1
		}
	}
	fmt.Println(sum)
}
