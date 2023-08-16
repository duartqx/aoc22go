package day8

import (
	"slices"
)

type TreeContext struct {
	row            int
	column         int
	current_height int
}

func getScenicScoreFromTop(ctx *TreeContext, patch_of_trees *[]string) (score int) {
	for t := ctx.row - 1; t >= 0; t-- {
		score++
		if byteToInt((*patch_of_trees)[t][ctx.column]) >= ctx.current_height {
			return score
		}
	}
	return score
}

func getScenicScoreFromBottom(ctx *TreeContext, patch_of_trees *[]string) (score int) {
	for b := ctx.row + 1; b < len(*patch_of_trees); b++ {
		score++
		if byteToInt((*patch_of_trees)[b][ctx.column]) >= ctx.current_height {
			return score
		}
	}
	return score
}

func getScenicScoreFromLeft(ctx *TreeContext, patch_of_trees *[]string) (score int) {
	for l := ctx.column - 1; l >= 0; l-- {
		score++
		if byteToInt((*patch_of_trees)[ctx.row][l]) >= ctx.current_height {
			return score
		}
	}
	return score
}

func getScenicScoreFromRight(ctx *TreeContext, patch_of_trees *[]string) (score int) {
	for r := ctx.column + 1; r < len((*patch_of_trees)[0]); r++ {
		score++
		if byteToInt((*patch_of_trees)[ctx.row][r]) >= ctx.current_height {
			return score
		}
	}
	return score
}

func Task2(patch_of_trees *[]string) (score int) {

	var (
		top_score    int
		bottom_score int
		left_score   int
		right_score  int
	)

	scenic_scores := []int{}

	for i := 0; i < len(*patch_of_trees); i++ {
		for j := 0; j < len((*patch_of_trees)[0]); j++ {

			ctx := &TreeContext{
				column: j, row: i, current_height: byteToInt((*patch_of_trees)[i][j]),
			}

			top_score = getScenicScoreFromTop(ctx, patch_of_trees)
			bottom_score = getScenicScoreFromBottom(ctx, patch_of_trees)
			left_score = getScenicScoreFromLeft(ctx, patch_of_trees)
			right_score = getScenicScoreFromRight(ctx, patch_of_trees)

			scenic_scores = append(
				scenic_scores,
				top_score*bottom_score*left_score*right_score,
			)
		}
	}

	return slices.Max(scenic_scores)
}
