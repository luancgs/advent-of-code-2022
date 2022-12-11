package main

import (
	"adventofcode2022/day10"
	"fmt"
)

func main() {

	output1, err := day10.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 10 - Part 1 | Output: ", output1)

	output2, err := day10.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 10 - Part 2 | Output: ", output2)

}
