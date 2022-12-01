package main

import (
	"adventofcode2022/day1"
	"fmt"
)

func main() {

	output1, err := day1.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 1 - Part 1 | Output: ", output1)

	output2, err := day1.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 1 - Part 2 | Output: ", output2)

}
