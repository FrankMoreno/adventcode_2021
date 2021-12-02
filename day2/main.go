package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/FrankMoreno/adventcode_2021/helper"
)

type command struct {
	direction string
	magnitude int
}

func main() {
	if len(os.Args) < 2 {
		log.Panic("NO FILE")
	}

	input := helper.ReadInput(os.Args[1], "\n")
	commands := convertToCommands(input)
	println(part1(commands))
	println(part2(commands))
}

func part1(commands []command) int {
	position := 0
	depth := 0

	for _, currentCommand := range commands {
		switch currentCommand.direction {
		case "up":
			depth -= currentCommand.magnitude
		case "down":
			depth += currentCommand.magnitude
		case "forward":
			position += currentCommand.magnitude
		}
	}

	return position * depth
}

func part2(commands []command) int {
	position := 0
	depth := 0
	aim := 0

	for _, currentCommand := range commands {
		switch currentCommand.direction {
		case "up":
			aim -= currentCommand.magnitude
		case "down":
			aim += currentCommand.magnitude
		case "forward":
			position += currentCommand.magnitude
			depth += aim * currentCommand.magnitude
		}
	}

	return position * depth
}

func convertToCommands(input []string) []command {
	commands := make([]command, len(input))

	for index, value := range input {
		fields := strings.Fields(value)
		direction := fields[0]
		magnitude, _ := strconv.Atoi(fields[1])

		commands[index] = command{
			direction: direction,
			magnitude: magnitude,
		}
	}

	return commands
}
