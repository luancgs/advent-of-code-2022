package main

import (
	"adventofcode2022/day9"
	"fmt"
)

func main() {

	output1, err := day9.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 9 - Part 1 | Output: ", output1)

	output2, err := day9.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 9 - Part 2 | Output: ", output2)

}
