package main

import (
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

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

func buildSlice(start, end int) *[]int {
	t := &[]int{}
	for i := start; i <= end; i++ {
		*t = append(*t, i)
	}
	return t
}

func main() {
	test := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}

	var first_section, second_section *[]int

	var how_many_overlap int

	for _, pair := range test {

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
			how_many_overlap++
		}
	}
	log.Println(how_many_overlap)
}
