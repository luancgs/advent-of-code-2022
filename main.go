package main

import (
	"adventofcode2022/day7"
	"fmt"
)

func main() {

	output1, err := day7.Part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 7 - Part 1 | Output: ", output1)

	output2, err := day7.Part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 7 - Part 2 | Output: ", output2)

}
