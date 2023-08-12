package main

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

func getInstunctions(row string, instrunctions *[][]int) error {

	inst_keywords := []string{"move", "from", "to"}
	// If row is a list of instrunctions
	inst_ints := []int{}
	for _, inst := range slices.DeleteFunc(
		strings.Split(row, " "),
		func(s string) bool {
			return slices.Contains(inst_keywords, s)
		},
	) {
		int_inst, err := strconv.Atoi(inst)
		if err != nil {
			return err
		}
		inst_ints = append(inst_ints, int_inst)
	}

	*instrunctions = append(
		*instrunctions,
		inst_ints,
	)

	return nil
}

func main() {
	p := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	crates_stacks := make(map[int][]string)

	instrunctions := [][]int{}

	// del := func(i string) bool {
	// 	return i == ""
	// }

	skip := []string{"", " "}

	for _, row := range strings.Split(p, "\n") {
		if strings.Contains(row, "1") {
			break
		} else {
			for i, j := 1, 1; j <= (len(row) - 1); j, i = j+4, i+1 {
				if crates_stacks[i] == nil {
					crates_stacks[i] = []string{}
				}
				if !slices.Contains(skip, string(row[j])) ||
					!slices.Contains(skip, string(row[j])) {
					crates_stacks[i] = append(crates_stacks[i], string(row[j]))
				}
			}
		}
	}

	for _, row := range strings.Split(p, "\n") {

		if strings.Contains(row, "move") {
			if err := getInstunctions(row, &instrunctions); err != nil {
				log.Fatal(err)
			}
		}
	}

	for _, inst := range instrunctions {
		move, from, to := inst[0], inst[1], inst[2]
		for i := 0; i < move; i++ {

			// Gets the first element in the stack
			crate := crates_stacks[from][0]

			// Removes the first element from the stack
			crates_stacks[from] = crates_stacks[from][1:]

			// Inserts the crate to the start of the stack
			crates_stacks[to] = slices.Insert(crates_stacks[to], 0, crate)

		}
	}

	var answer string

	for i := 1; i <= len(crates_stacks); i++ {
		answer += crates_stacks[i][0]
	}

	log.Println(answer)
}
