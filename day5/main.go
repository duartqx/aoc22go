package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getInputData(filename string) (data *[]string, err error) {

	data = &[]string{}

	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		*data = append(*data, scan.Text())
	}

	return data, nil
}

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

func crateMover9000(instrunctions [][]int, crates_stacks *map[int][]byte) {

	for _, inst := range instrunctions {
		move, from, to := inst[0], inst[1], inst[2]
		for i := 0; i < move; i++ {
			// Gets the first element in the stack
			crate := (*crates_stacks)[from][0]

			// Removes the first element from the stack
			(*crates_stacks)[from] = (*crates_stacks)[from][1:]

			// Inserts the crate to the start of the stack
			(*crates_stacks)[to] = slices.Insert((*crates_stacks)[to], 0, crate)
		}
	}
}

func crateMover9001(instrunctions [][]int, crates_stacks *map[int][]byte) {

	for _, inst := range instrunctions {
		move, from, to := inst[0], inst[1], inst[2]
		// Gets the first element in the stack
		crates := (*crates_stacks)[from][:move]

		// Removes the first element from the stack
		(*crates_stacks)[from] = (*crates_stacks)[from][move:]

		// Inserts the crate to the start of the stack
		for i := len(crates) - 1; i >= 0; i-- {
			(*crates_stacks)[to] = slices.Insert((*crates_stacks)[to], 0, crates[i])
		}
	}
}

func task(input string, crate_type string) {

	data, err := getInputData(input)
	if err != nil {
		log.Fatal(err)
	}

	crates_stacks := make(map[int][]byte)

	instrunctions := [][]int{}

	for _, row := range *data {
		if strings.Trim(row, " ") == "" || string(row[1]) == "1" {
			continue
		} else if strings.Contains(row, "move") {
			if err := getInstunctions(row, &instrunctions); err != nil {
				log.Fatal(err)
			}
		} else {
			for i, j := 1, 1; j <= (len(row) - 1); j, i = j+4, i+1 {
				if crates_stacks[i] == nil {
					crates_stacks[i] = []byte{}
				}
				if strings.Trim(string(row[j]), " ") != "" {
					crates_stacks[i] = append(crates_stacks[i], row[j])
				}
			}
		}
	}

	if crate_type == "CrateMover 9000" {
		crateMover9000(instrunctions, &crates_stacks)
	} else {
		crateMover9001(instrunctions, &crates_stacks)
	}

	var answer string

	for i := 1; i <= len(crates_stacks); i++ {
		answer += string(crates_stacks[i][0])
	}

	log.Println(answer)
}

func main() {
	input := "./input"
	task(input, "CrateMover 9000") // TBVFVDZPN
	task(input, "CrateMover 9001") // VLCWHTDSZ
}
