package main

import (
	getdata "aoc22go/get_data"
	"log"
	"strconv"
)

func getTestData() *[]string {
	return &[]string{"30373", "25512", "65332", "33549", "35390"}
}

func toInt(b byte) int {
	i, _ := strconv.Atoi(string(b))
	return i
}

func getTopVisibility(i, j, current_height int, patch_of_trees *[]string) []bool {
	is_visible_from_the_top := []bool{}
	for t := i - 1; t >= 0; t-- {
		is_visible_from_the_top = append(
			is_visible_from_the_top,
			current_height > toInt((*patch_of_trees)[t][j]),
		)
	}
	return is_visible_from_the_top
}

func getBottomVisibility(i, j, current_height int, patch_of_trees *[]string) []bool {
	is_visible_from_the_bottom := []bool{}
	for b := i + 1; b < len(*patch_of_trees); b++ {
		is_visible_from_the_bottom = append(
			is_visible_from_the_bottom,
			current_height > toInt((*patch_of_trees)[b][j]),
		)
	}
	return is_visible_from_the_bottom
}

func getLeftVisibility(i, j, current_height int, patch_of_trees *[]string) []bool {
	is_visible_from_the_left := []bool{}
	for l := j - 1; l >= 0; l-- {
		is_visible_from_the_left = append(
			is_visible_from_the_left,
			current_height > toInt((*patch_of_trees)[i][l]),
		)
	}
	return is_visible_from_the_left
}

func getRightVisibility(i, j, current_height int, patch_of_trees *[]string) []bool {
	is_visible_from_the_right := []bool{}
	for r := j + 1; r < len((*patch_of_trees)[0]); r++ {
		is_visible_from_the_right = append(
			is_visible_from_the_right,
			current_height > toInt((*patch_of_trees)[i][r]),
		)
	}
	return is_visible_from_the_right
}

func visibleCount(is_visible_from_a_side *[]bool) int {

	visibility_counter := 0
	for _, v := range *is_visible_from_a_side {
		if v {
			visibility_counter++
		}
	}
	return visibility_counter
}

func main() {

	patch_of_trees, err := getdata.GetInputData("./input")
	if err != nil {
		log.Fatal(err)
	}

	// patch_of_trees := getTestData()

	var (
		len_patch_of_trees      int
		len_row_of_trees        int
		current_tree_height     int
		number_of_visible_trees int
		visibility              []bool
	)

	len_patch_of_trees = (len(*patch_of_trees) - 2)
	len_row_of_trees = len((*patch_of_trees)[0])
	number_of_visible_trees = (len_row_of_trees * 2) + (len_patch_of_trees * 2)

	for i := 1; i < len(*patch_of_trees)-1; i++ {
		for j := 1; j < len((*patch_of_trees)[0])-1; j++ {

			current_tree_height = toInt((*patch_of_trees)[i][j])

			// Top Visibility
			visibility = getTopVisibility(
				i, j, current_tree_height, patch_of_trees,
			)
			if len(visibility) != 0 && visibleCount(&visibility) == len(visibility) {
				number_of_visible_trees++
				continue
			}

			// Bottom Visibility
			visibility = getBottomVisibility(
				i, j, current_tree_height, patch_of_trees,
			)
			if len(visibility) != 0 && visibleCount(&visibility) == len(visibility) {
				number_of_visible_trees++
				continue
			}

			// Left Visibility
			visibility = getLeftVisibility(
				i, j, current_tree_height, patch_of_trees,
			)
			if len(visibility) != 0 && visibleCount(&visibility) == len(visibility) {
				number_of_visible_trees++
				continue
			}

			// Right Visibility
			visibility = getRightVisibility(
				i, j, current_tree_height, patch_of_trees,
			)
			if len(visibility) != 0 && visibleCount(&visibility) == len(visibility) {
				number_of_visible_trees++
				continue
			}
		}
	}
	log.Println(number_of_visible_trees)
}
