package main

import (
	"day13/compare"
	"day13/parse"
	"fmt"
)

func main() {
	//inputs := parse.Parse("input_short.txt")
	inputs := parse.Parse("input.txt")
	//inputs := parse.Parse("redditorsInput.txt")
	sum := part1(inputs)
	fmt.Println(sum)
}

func part1(inputs [][2]string) int {
	sum := 0
	for index, firstAndSecond := range inputs {
		inOrder := compare.PairsInOrder(firstAndSecond[0], firstAndSecond[1])
		//fmt.Printf("Pair nr : %d \n", index+1)
		//fmt.Println(inOrder)
		//fmt.Println("")
		if inOrder {
			sum += index + 1
		}
	}
	return sum
}
