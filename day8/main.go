package main

import (
	"log"
	"strconv"
)

func getTestData() []string {
	return []string{"30373", "25512", "65332", "33549", "35390"}
}

func toInt(b byte) int {
	i, _ := strconv.Atoi(string(b))
	return i
}

func main() {

	patch_of_trees := getTestData()

	len_patch_of_trees := (len(patch_of_trees) - 2)

	len_row_of_trees := len(patch_of_trees[0])

	number_of_visible_trees := (len_row_of_trees * 2) + (len_patch_of_trees * 2)

	for i := 1; i < len(patch_of_trees)-1; i++ {
		for j := 1; j < len(patch_of_trees[0])-1; j++ {

			seen_by_all_sides := true
			current_tree_height := toInt(patch_of_trees[i][j])

			tblr := []int{
				toInt(patch_of_trees[i-1][j]),
				toInt(patch_of_trees[i+1][j]),
				toInt(patch_of_trees[i][j-1]),
				toInt(patch_of_trees[i][j+1]),
			}

			for other_tree := range tblr {
				if current_tree_height <= other_tree {
					seen_by_all_sides = false
				}

			}
			if seen_by_all_sides {
				number_of_visible_trees += 1
			}
		}
	}
	log.Println(number_of_visible_trees)
}
