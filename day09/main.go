package main

import (
	"aoc22go/getdata"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Command struct {
	dir   string
	steps int
}

type Knot struct {
	x      int
	y      int
	pmoves [][]int
}

func (k *Knot) Move(c *Command) {
	switch c.dir {
	case "R":
		k.x++
	case "L":
		k.x--
	case "U":
		k.y++
	case "D":
		k.y--
	}
	k.storeMove()
}

func (k *Knot) Follow(h *Knot) {
	ox := h.x - k.x
	oy := h.y - k.y

	dx := int(math.Abs(float64(ox)))
	dy := int(math.Abs(float64(oy)))

	if dx > 1 || (dx > 0 && dy > 1) {
		k.x += 1 * (ox / dx)
	}

	if dy > 1 || (dy > 0 && dx > 1) {
		k.y += 1 * (oy / dy)
	}

	k.storeMove()
}

func (k *Knot) storeMove() {
	if len(k.pmoves) == 0 {
		k.pmoves = [][]int{{0, 0}}
	}
	k.pmoves = append(k.pmoves, []int{k.x, k.y})
}

func (k *Knot) getUniqueMoves() *map[string]int {
	uniq := make(map[string]int)
	for _, move := range k.pmoves {
		key := fmt.Sprintf("[ %d %d ]", move[0], move[1])
		uniq[key] += 1
	}
	return &uniq
}

func getTestInputData() *[]Command {
	return &[]Command{
		{dir: "R", steps: 4},
		{dir: "U", steps: 4},
		{dir: "L", steps: 3},
		{dir: "D", steps: 1},
		{dir: "R", steps: 4},
		{dir: "D", steps: 1},
		{dir: "L", steps: 5},
		{dir: "R", steps: 2},
	}
}

func getInputData(input string) (*[]Command, error) {
	commands := []Command{}
	data, err := getdata.GetInputData(input)
	if err != nil {
		return nil, err
	}
	for _, d := range *data {
		c := strings.Split(d, " ")
		step, err := strconv.Atoi(c[1])
		if err != nil {
			return nil, err
		}
		commands = append(commands, Command{dir: c[0], steps: step})
	}
	return &commands, nil
}

func main() {
	h := new(Knot)
	t := new(Knot)

	data, err := getInputData("./input")
	if err != nil {
		fmt.Println(err)
		return
	}
	// data := getTestInputData()
	len_data := len(*data)
	for i := 0; i < len_data; i++ {
		command := (*data)[i]
		for j := 0; j < command.steps; j++ {
			h.Move(&command)
			t.Follow(h)
		}
	}

	uniq := t.getUniqueMoves()
	fmt.Println(len(*uniq))
}

// 6149
// 6026
