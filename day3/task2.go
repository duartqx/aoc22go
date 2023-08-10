package day3

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func SecondTask(file string) {
	fh, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fh.Close()

	scan := bufio.NewScanner(fh)

	var all_gnomes_sacks []string
	var priority_sum int

	for scan.Scan() {
		all_gnomes_sacks = append(all_gnomes_sacks, scan.Text())
	}

	containsOne := func(g string, c rune) bool {
		return strings.Contains(g, string(c))
	}

	for i := 0; i <= len(all_gnomes_sacks)-3; i = i + 3 {
		g1, g2, g3 := all_gnomes_sacks[i], all_gnomes_sacks[i+1], all_gnomes_sacks[i+2]
		for _, c := range g1 {
			if containsOne(g1, c) && containsOne(g2, c) && containsOne(g3, c) {
				irune := int(c)
				if irune > 96 {
					priority_sum += irune - 96
				} else {
					priority_sum += irune - 38
				}
				break
			}
		}
	}

	log.Println(priority_sum)
}
