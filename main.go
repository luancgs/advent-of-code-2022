package main

import (
	"adventofcode2022/day3"
	"fmt"
)

func main() {

	output1, err := day3.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 3 - Part 1 | Output: ", output1)

	output2, err := day3.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 3 - Part 2 | Output: ", output2)

}
