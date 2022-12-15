package main

import (
	"adventofcode2022/day11"
	"fmt"
)

func main() {

	output1, err := day11.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 11 - Part 1 | Output: ", output1)

	output2, err := day11.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 11 - Part 2 | Output: ", output2)

}
