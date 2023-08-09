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

	var highest_calories int
	var sum int

	for scan.Scan() {
		// Scans and removes all newline characters
		line := strings.ReplaceAll(scan.Text(), "\n", "")
		// Converts string to integer
		line_int, err := strconv.Atoi(line)
		if err != nil {
			// Reached a empty line
			if sum >= highest_calories {
				// Checks if sum is bigger than highest_calories
				highest_calories = sum
			}
			sum = 0
		} else {
			// Sums the value
			sum += line_int
		}
	}

	fmt.Println(highest_calories)

	if err := scan.Err(); err != nil {
		log.Fatal(err)
		return
	}
}
