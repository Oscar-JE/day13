package main

import (
	"day13/compare"
	"day13/parse"
	"fmt"
)

func main() {
	//inputs := parse.Parse("input_short.txt")

	//inputs := parse.Parse("inputCC.txt")
	sum := part2("input.txt")
	fmt.Println(sum)
}

func part1File(fileName string) int {
	inputs := parse.Parse(fileName)
	sum := 0
	for index, firstAndSecond := range inputs {
		//fmt.Println(firstAndSecond)
		inOrder := compare.PairsInOrder(firstAndSecond[0], firstAndSecond[1])
		//fmt.Println(inOrder)
		if inOrder {
			sum += index + 1
		}
	}
	return sum
}

func part1(inputs [][2]string) int {
	sum := 0
	for index, firstAndSecond := range inputs {
		//fmt.Println(firstAndSecond)
		inOrder := compare.PairsInOrder(firstAndSecond[0], firstAndSecond[1])
		//fmt.Println(inOrder)
		if inOrder {
			sum += index + 1
		}
	}
	return sum
}

func part2(fileName string) int {
	lines := parse.SimpleRead(fileName)
	lines = append(lines, "[[2]]")
	lines = append(lines, "[[6]]")
	lines = bubbleSort(lines)
	index2 := 0
	index6 := 0
	for i := range lines {
		if lines[i] == "[[2]]" {
			index2 = i + 1
		}
		if lines[i] == "[[6]]" {
			index6 = i + 1
		}
	}
	return index2 * index6
}

func bubbleSort(inputs []string) []string {
	isInOrder := false
	for !isInOrder {
		isInOrder = true
		i := 0
		for i < len(inputs)-1 {
			if !compare.PairsInOrder(inputs[i], inputs[i+1]) {
				swapHelp := inputs[i]
				inputs[i] = inputs[i+1]
				inputs[i+1] = swapHelp
				isInOrder = false
			}
			i++
		}
	}
	return inputs
}
