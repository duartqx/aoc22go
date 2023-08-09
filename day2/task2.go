package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// A, X -> Rock
// B, Y -> Paper
// C, Z -> Scissor

// X -> Lose
// Y -> Draw
// Z -> Win

// X -> 1
// Y -> 2
// Z -> 3

func fixMatch(gnome, mine string) int {
	switch mine {
	case "X":
		if gnome == "A" {
			return 3 + 0
		} else if gnome == "B" {
			return 1 + 0
		} else {
			return 2 + 0
		}
	case "Y":
		if gnome == "A" {
			return 1 + 3
		} else if gnome == "B" {
			return 2 + 3
		} else {
			return 3 + 3
		}
	case "Z":
		if gnome == "A" {
			return 2 + 6
		} else if gnome == "B" {
			return 3 + 6
		} else {
			return 1 + 6
		}
	}
	return 0
}

func SecondTask(file string) {

	fh, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fh.Close()

	scan := bufio.NewScanner(fh)

	var my_total_score int

	for scan.Scan() {
		// Scans and removes all newline characters
		var hands []string = strings.Split(
			strings.ReplaceAll(
				scan.Text(), "\n", "",
			), " ",
		)
		my_total_score += fixMatch(hands[0], hands[1])
	}
	fmt.Println(my_total_score)
	if err := scan.Err(); err != nil {
		log.Fatal(err)
		return
	}
}
