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
			return 3
		} else if mine == "Y" {
			return 6
		} else {
			return 0
		}
	case "B":
		if mine == "X" {
			return 0
		} else if mine == "Y" {
			return 3
		} else {
			return 6
		}
	case "C":
		if mine == "X" {
			return 6
		} else if mine == "Y" {
			return 0
		} else {
			return 3
		}
	}
	return 0
}

func FirstTask(file string) {

	points_per_hand := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

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
		my_total_score += points_per_hand[hands[1]]
		my_total_score += compareHands(hands[0], hands[1])
	}
	fmt.Println(my_total_score)
	if err := scan.Err(); err != nil {
		log.Fatal(err)
		return
	}
}
