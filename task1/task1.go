package task1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Task1(file_str string) {
	file, err := os.Open(file_str)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	var (
		highest_amount_of_calories   int
		current_gnome_total_calories int
	)

	for scan.Scan() {
		// Scans and removes all newline characters
		var strcalory string = strings.ReplaceAll(scan.Text(), "\n", "")
		// Converts string to integer and sets the current calory
		var calory int
		calory, err := strconv.Atoi(strcalory)
		if err != nil {
			// Atoi could not parse strcalory so -> strcalory == ""
			if current_gnome_total_calories >= highest_amount_of_calories {
				// Checks if current total is bigger than highest_calories
				highest_amount_of_calories = current_gnome_total_calories
			}
			current_gnome_total_calories = 0
		} else {
			// Sums the value
			current_gnome_total_calories += calory
		}
	}

	// Prints the result highest_amount_of_calories
	fmt.Println(highest_amount_of_calories)

	if err := scan.Err(); err != nil {
		log.Fatal(err)
		return
	}
}
