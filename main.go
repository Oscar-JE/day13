package main

import (
	"day13/compare"
	"day13/parse"
	"fmt"
)

func main() {
	inputs := parse.Parse("input_short.txt")
	sum := 0
	for _, firstAndSecond := range inputs {
		if compare.FirstLargerThanSecond(firstAndSecond[0], firstAndSecond[1]) {
			sum++
		}
	}
	fmt.Println(sum)
}
