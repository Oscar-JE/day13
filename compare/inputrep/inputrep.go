package inputrep

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
