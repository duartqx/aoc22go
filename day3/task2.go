package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SecondTask(file string) {
	fh, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fh.Close()

	scan := bufio.NewScanner(fh)

	var priority_sum int
	var gnome_group []string

	for scan.Scan() {
		t := scan.Text()

		if len(gnome_group) != 3 {
			gnome_group = append(gnome_group, t)
			fmt.Println(gnome_group)
		} else {

			uniq_1 := map[rune]string{}
			uniq_2 := map[rune]string{}
			uniq_3 := map[rune]string{}

			for _, c := range gnome_group[0] {
				uniq_1[c] = string(c)
			}
			for _, c := range gnome_group[1] {
				uniq_2[c] = string(c)
			}
			for _, c := range gnome_group[2] {
				uniq_3[c] = string(c)
			}

			fmt.Println(uniq_1)
			fmt.Println(uniq_2)
			fmt.Println(uniq_3)

			for key := range uniq_1 {
				v2, uniq_2_has_key := uniq_2[key]
				v3, uniq_3_has_key := uniq_3[key]

				fmt.Println(uniq_2_has_key, uniq_3_has_key)

				if uniq_2_has_key && uniq_3_has_key {
					fmt.Println(uniq_1[key], v2, v3)
					i := int(key)
					if i > 96 {
						priority_sum += (int(key) - 96)
					} else {
						priority_sum += (int(key) - 38)
					}
				}
			}
			fmt.Println()

			gnome_group = []string{}
		}
	}
	fmt.Println(priority_sum)
}
