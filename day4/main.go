package main

import (
	getdata "aoc22go/get_data"
	"log"
	"slices"
	"strconv"
	"strings"
)

type tsk func(input string) (int, error)

type Result struct {
	s1  *[]int
	s2  *[]int
	err error
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

func buildSlice(start, end int) (t []int) {
	t = []int{}
	for i := start; i <= end; i++ {
		t = append(t, i)
	}
	return t
}

func task(input string) <-chan Result {

	c := make(chan Result)

	gnome_section_data, err := getdata.GetInputData(input)
	if err != nil {
		c <- Result{s1: nil, s2: nil, err: err}
		return c
	}

	go func() {

		defer close(c)

		for _, pair := range *gnome_section_data {

			// Resets the sections
			first_section, second_section := []int{}, []int{}

			ts := strings.Split(pair, ",")

			first_range := strings.Split(ts[0], "-")
			second_range := strings.Split(ts[1], "-")

			// Converts and gets the start and end for the first gnome
			first_gnome_start, first_gnome_end, err := getStartEnd(
				[2]string{first_range[0], first_range[1]},
			)
			if err != nil {
				c <- Result{s1: nil, s2: nil, err: err}
			}

			// Converts and gets the start and end for the second gnome
			second_gnome_start, second_gnome_end, err := getStartEnd(
				[2]string{second_range[0], second_range[1]},
			)
			if err != nil {
				c <- Result{s1: nil, s2: nil, err: err}
			}

			first_section = buildSlice(first_gnome_start, first_gnome_end)
			second_section = buildSlice(second_gnome_start, second_gnome_end)

			c <- Result{s1: &first_section, s2: &second_section, err: nil}
		}
	}()
	return c
}

func firstTask(input string) (int, error) {

	var how_many_overlap int

	iterAndSum := func(iter, check_contains *[]int) int {
		var overlap int

		for _, s := range *iter {
			if slices.Contains(*check_contains, s) {
				overlap++
			}
		}
		return overlap
	}

	for res := range task(input) {
		if res.err != nil {
			return 0, res.err
		}

		if iterAndSum(res.s1, res.s2) == len(*res.s1) ||
			iterAndSum(res.s2, res.s1) == len(*res.s2) {
			how_many_overlap++
		}
	}
	return how_many_overlap, nil
}

func secondTask(input string) (int, error) {

	var how_many_overlap int

	iterAndSum := func(iter, check_contains *[]int) bool {
		for _, s := range *iter {
			if slices.Contains(*check_contains, s) {
				return true
			}
		}
		return false
	}

	for res := range task(input) {
		if res.err != nil {
			return 0, res.err
		}

		if iterAndSum(res.s1, res.s2) || iterAndSum(res.s2, res.s1) {
			how_many_overlap++
		}
	}
	return how_many_overlap, nil
}

func main() {

	input := "./input"

	fs := [2]tsk{firstTask, secondTask}

	for i := 0; i <= 1; i++ {

		res, err := fs[i](input)

		if err != nil {
			log.Fatal(err)
		}

		log.Println(res) // 657 938
	}
}
