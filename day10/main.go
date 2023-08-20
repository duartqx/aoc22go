package main

import (
	"aoc22go/day10/src"
	"log"
	"strconv"
)

// DURING cycle -> check strength
// noop: DOES NOT, X += 1
// addx: AFTER 2 cycles -> SUM value to X

// addx V takes two cycles to complete. After two cycles, the X register is
// increased by the value V. (V can be negative.)

// noop takes one cycle to complete. It has no other effect

func checkStrength(cycles, x int) (str int) {
	if cycles%40 == 20 {
		str = cycles * x
	}
	return str
}

func main() {
	c := src.GetTestData()

	x := 1
	cycles := 0

	strengths := make(map[int]int)

	for line := range c {

		cycles += 1

		switch line[0] {
		case "noop":
			strength := checkStrength(cycles, x)
			if strength != 0 {
				log.Println(line[0], cycles, x)
				strengths[cycles] = strength
			}
		case "addx":
			strength := checkStrength(cycles, x)
			if strength != 0 {
				log.Println(line[0], cycles, x)
				strengths[cycles] = strength
			}

			cycles += 1

			strength = checkStrength(cycles, x)
			if strength != 0 {
				log.Println(line[0], cycles, x)
				strengths[cycles] = strength
			}

			y, _ := strconv.Atoi(line[1])
			x += y
		}
	}
	log.Println(strengths)
}
