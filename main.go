package main

import (
	"adventofcode2022/day8"
	"fmt"
)

func main() {

	output1, err := day8.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 8 - Part 1 | Output: ", output1)

	output2, err := day8.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 8 - Part 2 | Output: ", output2)

}
