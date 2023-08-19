package day8

import (
	"strconv"
)

type VisibilityFunc func(i, j, ch int, patch_of_trees *[]string) bool

func byteToInt(b byte) int {
	i, _ := strconv.Atoi(string(b))
	return i
}

func isVisibleFromTop(i, j, current_height int, patch_of_trees *[]string) bool {
	for t := i - 1; t >= 0; t-- {
		if current_height <= byteToInt((*patch_of_trees)[t][j]) {
			return false
		}
	}
	return true
}

func isVisibleFromBottom(i, j, current_height int, patch_of_trees *[]string) bool {
	for b := i + 1; b < len(*patch_of_trees); b++ {
		if current_height <= byteToInt((*patch_of_trees)[b][j]) {
			return false
		}
	}
	return true
}

func isVisibleFromLeft(i, j, current_height int, patch_of_trees *[]string) bool {
	for l := j - 1; l >= 0; l-- {
		if current_height <= byteToInt((*patch_of_trees)[i][l]) {
			return false
		}
	}
	return true
}

func isVisibleFromRight(i, j, current_height int, patch_of_trees *[]string) bool {
	for r := j + 1; r < len((*patch_of_trees)[0]); r++ {
		if current_height <= byteToInt((*patch_of_trees)[i][r]) {
			return false
		}
	}
	return true
}

func Task1(patch_of_trees *[]string) int {

	var (
		len_patch_of_trees      int
		len_row_of_trees        int
		current_tree_height     int
		number_of_visible_trees int
	)

	len_patch_of_trees = (len(*patch_of_trees) - 2)
	len_row_of_trees = len((*patch_of_trees)[0])
	number_of_visible_trees = (len_row_of_trees * 2) + (len_patch_of_trees * 2)

	visibility_functions := []VisibilityFunc{
		isVisibleFromTop,
		isVisibleFromBottom,
		isVisibleFromLeft,
		isVisibleFromRight,
	}

	for i := 1; i < len(*patch_of_trees)-1; i++ {
		for j := 1; j < len_row_of_trees-1; j++ {

			current_tree_height = byteToInt((*patch_of_trees)[i][j])

			for _, isVisibleFunc := range visibility_functions {
				if isVisibleFunc(i, j, current_tree_height, patch_of_trees) {
					number_of_visible_trees++
					break
				}
			}
		}
	}
	return number_of_visible_trees
}
