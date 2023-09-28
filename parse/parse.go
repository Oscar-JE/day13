package parse

import (
	"bufio"
	"os"
)

func Parse(inputPath string) [][2]string {
	file, fileError := os.Open(inputPath)
	if fileError != nil {
		panic("file not found")
	}
	reader := bufio.NewScanner(file)
	lines := []string{}
	for reader.Scan() {
		line := reader.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	inputs := [][2]string{}
	loopInput := [2]string{}
	for index := range lines {
		interalIndex := index % 2
		if index != 0 && interalIndex == 0 {
			inputs = append(inputs, loopInput)
			loopInput = [2]string{}
		}
		loopInput[interalIndex] = lines[index]
	}
	return inputs
}
