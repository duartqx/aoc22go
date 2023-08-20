package main

import (
	"aoc22go/day10/src"
	"aoc22go/getdata"
	"log"
	"strings"
)

// DURING cycle -> check strength
// noop: DOES NOT, X += 1
// addx: AFTER 2 cycles -> SUM value to X

// addx V takes two cycles to complete. After two cycles, the X register is
// increased by the value V. (V can be negative.)

// noop takes one cycle to complete. It has no other effect

func main() {

	// data := src.GetTestData()
	data, err := getdata.GetInputChannel("./day10/input")
	if err != nil {
		log.Fatal(err)
	}

	c := new(src.CathodeTube)

	c.Build()

	for d := range data {
		instrunction := strings.Split(d, " ")
		c.Cycle(instrunction)
		// c.Cycle(d)
	}

	log.Println(c.GetStrengths()) // 17180
	for _, row := range c.GetScreen() {
		log.Println(strings.Join(row, ""))
	} // rehprlub
}
