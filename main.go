package main

import (
	"adventofcode2022/day4"
	"fmt"
)

func main() {

	output1, err := day4.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 4 - Part 1 | Output: ", output1)

	output2, err := day4.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 4 - Part 2 | Output: ", output2)

}
