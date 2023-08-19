package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SecondTask(file_str string) {
	file, err := os.Open(file_str)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	var (
		all_gnomes_calories          []int
		current_gnome_total_calories int
	)

	for scan.Scan() {
		// Scans and removes all newline characters
		var strcalory string = strings.ReplaceAll(scan.Text(), "\n", "")
		// Converts string to integer and sets the current calory
		var calory int
		calory, err := strconv.Atoi(strcalory)
		if err != nil {
			all_gnomes_calories = append(all_gnomes_calories, current_gnome_total_calories)
			current_gnome_total_calories = 0
		} else {
			// Sums the value
			current_gnome_total_calories += calory
		}
	}

	// Prints the result highest_amount_of_calories
	sort.Ints(all_gnomes_calories)

	biggest_three_carried_calories := all_gnomes_calories[len(all_gnomes_calories)-3:]
	fmt.Println(biggest_three_carried_calories)

	var total_calories int
	for _, calory := range biggest_three_carried_calories {
		total_calories += calory
	}

	fmt.Println(total_calories)

	if err := scan.Err(); err != nil {
		log.Fatal(err)
		return
	}
}
