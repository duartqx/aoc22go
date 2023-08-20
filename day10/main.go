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

type CathodeTube struct {
	x         int
	cycles    int
	strengths map[int]int
}

func (c *CathodeTube) checkStrength() {
	if c.cycles%40 == 20 {
		c.strengths[c.cycles] = c.cycles * c.x
	}
}

func (c *CathodeTube) getStrengthSum() (s int) {
	for key := range c.strengths {
		s += c.strengths[key]
	}
	return s
}

func main() {

	data := src.GetTestData()

	c := CathodeTube{x: 1, strengths: make(map[int]int)}

	for line := range data {

		c.cycles += 1

		switch line[0] {
		case "noop":
			c.checkStrength()
		case "addx":
			c.checkStrength()

			c.cycles += 1

			c.checkStrength()

			y, _ := strconv.Atoi(line[1])
			c.x += y
		}
	}
	log.Println(c.strengths, c.getStrengthSum())
}
