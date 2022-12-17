package main

import (
	"adventofcode2022/day12"
	"fmt"
)

func main() {

	output1, err := day12.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 12 - Part 1 | Output: ", output1)

	output2, err := day12.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 12 - Part 2 | Output: ", output2)

}
