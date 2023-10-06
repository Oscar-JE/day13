package parse

import (
	"bufio"
	"os"
)

func Parse(inputPath string) [][2]string {

	lines := SimpleRead(inputPath)
	return combine(lines)
}

func SimpleRead(inputPath string) []string {
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
	return lines
}

func combine(lines []string) [][2]string {
	inputs := [][2]string{}
	for i := 0; i < len(lines); i += 2 {
		first := lines[i]
		second := lines[i+1]
		inputs = append(inputs, [2]string{first, second})
	}
	return inputs
}
