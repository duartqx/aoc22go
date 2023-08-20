package src

import (
	"strconv"
)

type CathodeTube struct {
	x         int
	cycles    int
	strengths map[int]int
	screen    [][]string
	prow      int
}

func (c *CathodeTube) checkStrength() {
	if c.cycles%40 == 20 {
		c.strengths[c.cycles] = c.cycles * c.x
	}
}

func (c *CathodeTube) draw() {
	if len(c.screen[c.prow]) == 40 {
		c.prow += 1
	}

	var (
		cycle int       = c.cycles % 40
		pixel *[]string = &(c.screen[c.prow])
	)

	if cycle == (c.x-1) || cycle == c.x || cycle == (c.x+1) {
		*pixel = append(*pixel, "#")
	} else {
		*pixel = append(*pixel, ".")
	}
}

func (c *CathodeTube) drawAndCheck() {
	c.draw()
	c.cycles += 1
	c.checkStrength()
}

func (c *CathodeTube) Build() *CathodeTube {
	c.x = 1
	c.strengths = make(map[int]int)
	c.screen = make([][]string, 6)
	for i := range c.screen {
		c.screen[i] = []string{}
	}
	return c
}

func (c *CathodeTube) Cycle(instruction []string) {

	c.drawAndCheck()

	switch instruction[0] {
	case "noop":
	case "addx":
		c.drawAndCheck()
		y, _ := strconv.Atoi(instruction[1])
		c.x += y
	}
}

func (c *CathodeTube) GetStrengths() (s int) {
	for key := range c.strengths {
		s += c.strengths[key]
	}
	return s
}

func (c *CathodeTube) GetScreen() [][]string {
	return c.screen
}
