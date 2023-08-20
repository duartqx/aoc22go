package src

import (
	"strconv"
)

type CathodeTube struct {
	x         int
	cycles    int
	strengths map[int]int
	prow      int
	pixels    [][]string
}

func (c *CathodeTube) checkStrength() {
	if c.cycles%40 == 20 {
		c.strengths[c.cycles] = c.cycles * c.x
	}
}

func (c *CathodeTube) Build() *CathodeTube {
	c.x = 1
	c.strengths = make(map[int]int)
	c.pixels = make([][]string, 6)
	for i := range c.pixels {
		c.pixels[i] = []string{}
	}
	return c
}

func (c *CathodeTube) GetStrengths() (strengths map[int]int, s int) {
	for key := range c.strengths {
		s += c.strengths[key]
	}
	return c.strengths, s
}

func (c *CathodeTube) drawAndCheck() {
	c.draw()
	c.cycles += 1
	c.checkStrength()
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

func (c *CathodeTube) draw() {
	if len(c.pixels[c.prow]) == 40 {
		c.prow += 1
	}

	cycle := c.cycles % 40

	if cycle == (c.x-1) || cycle == c.x || cycle == (c.x+1) {
		c.pixels[c.prow] = append(c.pixels[c.prow], "#")
	} else {
		c.pixels[c.prow] = append(c.pixels[c.prow], ".")
	}

}

func (c *CathodeTube) GetScreen() [][]string {
	return c.pixels
}
