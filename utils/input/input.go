package input

import (
	"fmt"
	"os"
)

// Receives the day and returns the input as a string
func GetInput(day int) string {
	inputPath := fmt.Sprintf("./inputs/2022/day%02d.txt", day)
	content, err := os.ReadFile(inputPath)

	if err != nil {
		panic(err)
	}

	return string(content)
}
