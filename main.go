package main

import (
	"adventofcode2022/day14"
	"adventofcode2022/utils/input"
	"fmt"
)

func main() {

	output1, err := day14.Part1(input.Reader{})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 14 - Part 1 | Output: ", output1)

	output2, err := day14.Part2(input.Reader{})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 14 - Part 2 | Output: ", output2)

}
