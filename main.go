package main

import (
	"adventofcode2022/day6"
	"fmt"
)

func main() {

	output1, err := day6.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 6 - Part 1 | Output: ", output1)

	output2, err := day6.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 6 - Part 2 | Output: ", output2)

}
