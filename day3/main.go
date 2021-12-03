package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/FrankMoreno/adventcode_2021/helper"
)

type Solution struct {
	totals      []int
	totalLines  int
	gammaRate   int
	epsilonRate int
}

func main() {
	if len(os.Args) < 2 {
		log.Panic("NO FILE")
	}
	input := helper.ReadInput(os.Args[1], "\n")

	part1(input)
	oxy, _ := strconv.ParseInt(part2_oxy(input), 2, 64)
	co2, _ := strconv.ParseInt(part2_co2(input), 2, 64)
	println(oxy * co2)
}

func part1(input []string) {
	solution := Solution{}
	solution.totals = make([]int, len(input[0]))
	for _, value := range input {
		solution.countBinary(value)
	}

	barrier := solution.totalLines / 2
	exponent := len(input[0]) - 1

	for _, value := range solution.totals {
		if value >= barrier {
			solution.gammaRate += int(math.Pow(2, float64(exponent)))
		} else {
			solution.epsilonRate += int(math.Pow(2, float64(exponent)))
		}

		exponent--
	}

	println(solution.gammaRate * solution.epsilonRate)
}

func (s *Solution) countBinary(binaryNumber string) {
	s.totalLines++
	for index, value := range binaryNumber {
		if string(value) == "1" {
			s.totals[index] += 1
		}
	}

}

func part2_oxy(inputs []string) string {
	index := 0
	for len(inputs) > 1 {
		ones, zeroes := findCharacterCounts(inputs, index)
		if ones >= zeroes {
			inputs = filterForIndexCharacter(inputs, "1", index)
		} else {
			inputs = filterForIndexCharacter(inputs, "0", index)
		}
		index += 1
	}

	return inputs[0]
}

func part2_co2(inputs []string) string {
	index := 0
	for len(inputs) > 1 {
		ones, zeroes := findCharacterCounts(inputs, index)
		if ones < zeroes {
			inputs = filterForIndexCharacter(inputs, "1", index)
		} else {
			inputs = filterForIndexCharacter(inputs, "0", index)
		}
		index += 1
	}

	return inputs[0]
}

func findCharacterCounts(inputs []string, index int) (int, int) {
	ones := 0
	zeroes := 0
	for _, input := range inputs {
		if string(input[index]) == "1" {
			ones += 1
		} else {
			zeroes += 1
		}
	}

	return ones, zeroes
}

func filterForIndexCharacter(input []string, target string, index int) []string {
	filtered := []string{}

	for _, currentString := range input {
		if string(currentString[index]) == target {
			filtered = append(filtered, currentString)
		}
	}

	return filtered
}
