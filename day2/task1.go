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

func compareHands(gnome, mine string) int {
	switch gnome {
	case "A":
		if mine == "X" {
			return 1 + 3
		} else if mine == "Y" {
			return 2 + 6
		} else {
			return 3 + 0
		}
	case "B":
		if mine == "X" {
			return 1 + 0
		} else if mine == "Y" {
			return 2 + 3
		} else {
			return 3 + 6
		}
	case "C":
		if mine == "X" {
			return 1 + 6
		} else if mine == "Y" {
			return 2 + 0
		} else {
			return 3 + 3
		}
	}
	return 0
}

func FirstTask(file string) {

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
		my_total_score += compareHands(hands[0], hands[1])
	}
	fmt.Println(my_total_score)
	if err := scan.Err(); err != nil {
		log.Fatal(err)
		return
	}
}
