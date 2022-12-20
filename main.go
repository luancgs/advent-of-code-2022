package main

import (
	"adventofcode2022/day14"
	"fmt"
)

func main() {

	output1, err := day14.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 14 - Part 1 | Output: ", output1)

	output2, err := day14.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 14 - Part 2 | Output: ", output2)

}
