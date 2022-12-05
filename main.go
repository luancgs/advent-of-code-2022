package main

import (
	"adventofcode2022/day5"
	"fmt"
)

func main() {

	output1, err := day5.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 5 - Part 1 | Output: ", output1)

	output2, err := day5.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 5 - Part 2 | Output: ", output2)

}
