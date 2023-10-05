package inputrep

import (
	"strconv"
)

type InputRep struct {
	isLeft  bool
	isRight bool
	value   int
}

func (i InputRep) IsNumer() bool {
	return !i.isLeft && !i.isRight
}

func (i InputRep) IsLeft() bool {
	return i.isLeft
}

func (i InputRep) IsRight() bool {
	return i.isRight
}

func (i InputRep) GetValue() int {
	return i.value
}

func InitLeft() InputRep {
	return InputRep{true, false, -1}
}

func InitRight() InputRep {
	return InputRep{false, true, -1}
}

func InitNumber(number int) InputRep {
	return InputRep{false, false, number}
}

func (i InputRep) Equal(other InputRep) bool {
	return i.isLeft == other.isLeft && i.isRight == other.isRight && i.value == other.value
}

func ParseToInputRep(line string) []InputRep {
	representation := []InputRep{}
	numberBuffer := ""
	for index := range line {
		if line[index] == '[' {
			representation = append(representation, InitLeft())
		} else if line[index] == ']' {
			if numberBuffer != "" {
				value, errorConv := strconv.Atoi(numberBuffer)
				if errorConv != nil {
					panic("parsing error")
				}
				representation = append(representation, InitNumber(value))
				numberBuffer = ""
			}
			representation = append(representation, InitRight())
		} else if line[index] == ',' {
			if numberBuffer != "" {
				representation = appendNumber(representation, numberBuffer)
				numberBuffer = ""
			}
		} else { // här har vi ett fel
			numberBuffer += string(line[index])
		}
	}
	if numberBuffer != "" {
		representation = appendNumber(representation, numberBuffer)
		numberBuffer = ""
	}
	return representation
}

func appendNumber(representation []InputRep, numberBuffer string) []InputRep {
	value, errorConv := strconv.Atoi(numberBuffer)
	if errorConv != nil {
		panic("parsing error")
	}
	representation = append(representation, InitNumber(value))
	return representation
}
