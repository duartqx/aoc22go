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

func getStartEnd(t [2]string) (start, end int, err error) {
	start, err = strconv.Atoi(t[0])
	if err != nil {
		return start, end, err
	}
	end, err = strconv.Atoi(t[1])
	if err != nil {
		return start, end, err
	}
	return start, end, err
}

func buildSlice(start, end int) (t *[]int) {
	t = &[]int{}
	for i := start; i <= end; i++ {
		*t = append(*t, i)
	}
	return t
}

func task(input string, first_task string) {

	gnome_section_data, err := getInputData(input)
	if err != nil {
		log.Fatal(err)
	}

	var first_section, second_section *[]int

	var how_many_overlap int

	for _, pair := range *gnome_section_data {

		// Resets the sections
		first_section, second_section = &[]int{}, &[]int{}

		ts := strings.Split(pair, ",")

		first_range := strings.Split(ts[0], "-")
		second_range := strings.Split(ts[1], "-")

		// Converts and gets the start and end for the first gnome
		first_gnome_start, first_gnome_end, err := getStartEnd(
			[2]string{first_range[0], first_range[1]},
		)
		if err != nil {
			log.Fatal(err)
		}

		// Converts and gets the start and end for the second gnome
		second_gnome_start, second_gnome_end, err := getStartEnd(
			[2]string{second_range[0], second_range[1]},
		)
		if err != nil {
			log.Fatal(err)
		}

		first_section = buildSlice(first_gnome_start, first_gnome_end)
		second_section = buildSlice(second_gnome_start, second_gnome_end)

		if first_task == "first_task" {
			how_many_overlap += firstTask(first_section, second_section)
		} else {
			how_many_overlap += secondTask(first_section, second_section)
		}
	}
	log.Println(how_many_overlap)
}

func firstTask(first_section, second_section *[]int) int {
	var first_overlap, second_overlap int

	for _, s := range *first_section {
		if slices.Contains(*second_section, s) {
			first_overlap++
		}
	}

	for _, s := range *second_section {
		if slices.Contains(*first_section, s) {
			second_overlap++
		}
	}

	if first_overlap == len(*first_section) || second_overlap == len(*second_section) {
		return 1
	}
	return 0
}

func secondTask(first_section, second_section *[]int) int {
	for _, s := range *first_section {
		if slices.Contains(*second_section, s) {
			return 1
		}
	}

	for _, s := range *second_section {
		if slices.Contains(*first_section, s) {
			return 1
		}
	}
	return 0
}

func main() {
	task("./input", "first_task")  // 657
	task("./input", "second_task") // 938
}
