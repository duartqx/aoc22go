package src

import "strconv"

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

func (c *CathodeTube) Build() *CathodeTube {
	c.x = 1
	c.strengths = make(map[int]int)
	return c
}

func (c *CathodeTube) GetStrengths() (strengths map[int]int, s int) {
	for key := range c.strengths {
		s += c.strengths[key]
	}
	return c.strengths, s
}

func (c *CathodeTube) Cycle(instruction []string) {
	c.cycles += 1

	switch instruction[0] {
	case "noop":
		c.checkStrength()
	case "addx":
		c.checkStrength()

		c.cycles += 1

		c.checkStrength()

		y, _ := strconv.Atoi(instruction[1])
		c.x += y
	}
}
