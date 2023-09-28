package compare

import (
	"strconv"
)

type InputRep struct {
	isLeft  bool
	isRight bool
	value   int
}

func initLeft() InputRep {
	return InputRep{true, false, -1}
}

func initRight() InputRep {
	return InputRep{false, true, -1}
}

func initNumber(number int) InputRep {
	return InputRep{false, false, number}
}

func pairsInOrder(first string, second string) bool {
	return false
}

func equal(a InputRep, b InputRep) bool {
	return a.isLeft == b.isLeft && a.isRight == b.isRight && a.value == b.value
}

func parseToInputRep(line string) []InputRep {
	retValue := []InputRep{}
	numberBuffer := ""
	for index := range line {
		if line[index] == '[' {
			retValue = append(retValue, initLeft())
		} else if line[index] == ']' {
			if numberBuffer != "" {
				value, errorConv := strconv.Atoi(numberBuffer)
				if errorConv != nil {
					panic("parsing error")
				}
				retValue = append(retValue, initNumber(value))
			}
			retValue = append(retValue, initRight())
		} else if line[index] == ',' {
			if numberBuffer != "" {
				value, errorConv := strconv.Atoi(numberBuffer)
				if errorConv != nil {
					panic("parsing error")
				}
				retValue = append(retValue, initNumber(value))
				numberBuffer = ""
			}
		} else {
			numberBuffer += string(line[index])
		}
	}
	return retValue
}
