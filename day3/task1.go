package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func FirstTask(file string) {
	fh, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fh.Close()

	scan := bufio.NewScanner(fh)

	var priority_sum int

	for scan.Scan() {
		t := scan.Text()
		l := len(t) / 2

		uniques_in_left := []rune{}
		uniques_in_right := []rune{}

		for _, c := range t[:l] {
			if !slices.Contains(uniques_in_left, c) {
				uniques_in_left = append(uniques_in_left, c)
			}
		}
		for _, c := range t[l:] {
			if !slices.Contains(uniques_in_right, c) {
				uniques_in_right = append(uniques_in_right, c)
			}
		}

		for _, c := range uniques_in_left {
			if slices.Contains(uniques_in_right, c) {
				i := int(c)
				if i > 96 {
					priority_sum += (int(c) - 96)
				} else {
					priority_sum += (int(c) - 38)
				}
			}
		}
	}
	fmt.Println(priority_sum)
}
