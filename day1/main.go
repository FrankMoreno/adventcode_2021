package main

import (
	"log"
	"os"

	"github.com/FrankMoreno/adventcode_2021/helper"
)

func main() {
	if len(os.Args) < 2 {
		log.Panic("NO FILE")
	}

	input := helper.ReadInput(os.Args[1], "\n")
	depths := helper.StringtoIntSlice(input)
	println(part1(depths))
	println(part2(depths))
}

func part1(input []int) int {
	count := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			count++
		}
	}

	return count
}

func part2(input []int) int {
	prevCount := 0
	currentCount := 0
	count := 0

	for i := 1; i < len(input)-2; i++ {
		prevCount = input[i-1] + input[i] + input[i+1]
		currentCount = input[i] + input[i+1] + input[i+2]

		if currentCount > prevCount {
			count++
		}
	}

	return count
}
