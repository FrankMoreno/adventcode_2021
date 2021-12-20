package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/FrankMoreno/adventcode_2021/helper"
)

type coordinate struct {
	x int
	y int
}

type line struct {
	point1 coordinate
	point2 coordinate
}

func main() {
	if len(os.Args) < 2 {
		log.Panic("NO FILE")
	}

	inputs := helper.ReadInput(os.Args[1], "\n")
	lines := []line{}

	for _, input := range inputs {
		coordinates := strings.Split(input, " -> ")
		lines = append(lines, line{
			point1: convertToCoordinate(coordinates[0]),
			point2: convertToCoordinate(coordinates[1]),
		})
	}

	part1(lines)
}

func part1(lines []line) int {
	seafloor := [1000][1000]int{}
	count := 0

	for _, line := range lines {
		// Vertical Line
		if line.point1.x == line.point2.x {
			for i := helper.Min(line.point1.y, line.point2.y); i <= helper.Max(line.point1.y, line.point2.y); i++ {
				seafloor[line.point1.x][i] += 1
				if seafloor[line.point1.x][i] == 2 {
					count++
				}
			}
			// Horizontal Line
		} else if line.point1.y == line.point2.y {
			for i := helper.Min(line.point1.x, line.point2.x); i <= helper.Max(line.point1.x, line.point2.x); i++ {
				seafloor[i][line.point1.y] += 1
				if seafloor[i][line.point1.y] == 2 {
					count++
				}
			}
			// Diagonal Line
		} else {
			startingPoint := line.point1
			endingPoint := line.point2
			if line.point2.x < line.point1.x {
				startingPoint = line.point2
				endingPoint = line.point1
			}

			startingY := startingPoint.y

			if startingPoint.y < endingPoint.y {
				for i := startingPoint.x; i <= endingPoint.x; i++ {
					seafloor[i][startingY] += 1
					if seafloor[i][startingY] == 2 {
						count++
					}
					startingY++
				}
			} else {
				for i := startingPoint.x; i <= endingPoint.x; i++ {
					seafloor[i][startingY] += 1
					if seafloor[i][startingY] == 2 {
						count++
					}
					startingY--
				}
			}
		}
	}

	println(count)
	return count
}

func convertToCoordinate(input string) coordinate {
	xandy := strings.Split(input, ",")
	x, _ := strconv.Atoi(xandy[0])
	y, _ := strconv.Atoi(xandy[1])

	return coordinate{
		x: x,
		y: y,
	}
}
