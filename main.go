package main

import (
	"adventofcode2022/day2"
	"fmt"
)

func main() {

	output1, err := day2.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 2 - Part 1 | Output: ", output1)

	output2, err := day2.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 2 - Part 2 | Output: ", output2)

}
