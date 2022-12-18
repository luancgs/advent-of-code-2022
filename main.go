package main

import (
	"adventofcode2022/day13"
	"fmt"
)

func main() {

	output1, err := day13.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 13 - Part 1 | Output: ", output1)

	output2, err := day13.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 13 - Part 2 | Output: ", output2)

}
