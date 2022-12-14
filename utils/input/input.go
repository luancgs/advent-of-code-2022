package input

import (
	"fmt"
	"os"
)

type InputReader interface {
	GetInput(day int, isTest ...bool) string
}

type Reader struct{}

// Receives the day and returns the input as a string
func (Reader) GetInput(day int, isTest ...bool) string {
	var inputPath string

	if len(isTest) > 0 && isTest[0] {
		inputPath = fmt.Sprintf("../inputs/2022/sample/day%02d.txt", day)
	} else {
		inputPath = fmt.Sprintf("./inputs/2022/day%02d.txt", day)
	}

	content, err := os.ReadFile(inputPath)

	if err != nil {
		panic(err)
	}

	return string(content)
}
