package day8

import (
	"slices"
)

type TreeContext struct {
	row     int
	column  int
	current int
	grid    *[]string
}

func (ctx *TreeContext) SetCoords(row, column int) *TreeContext {
	ctx.row = row
	ctx.column = column
	return ctx
}

func (ctx *TreeContext) SetColumn(column int) *TreeContext {
	ctx.column = column
	return ctx
}

func (ctx *TreeContext) SetRow(row int) *TreeContext {
	ctx.row = row
	return ctx
}

func (ctx *TreeContext) Build() *TreeContext {
	ctx.current = byteToInt((*ctx.grid)[ctx.row][ctx.column])
	return ctx
}

func (ctx *TreeContext) getScenicScoreTop() (score int) {
	for t := ctx.row - 1; t >= 0; t-- {
		score++
		if byteToInt((*ctx.grid)[t][ctx.column]) >= ctx.current {
			return score
		}
	}
	return score
}

func (ctx *TreeContext) getScenicScoreBottom() (score int) {
	for b := ctx.row + 1; b < len(*ctx.grid); b++ {
		score++
		if byteToInt((*ctx.grid)[b][ctx.column]) >= ctx.current {
			return score
		}
	}
	return score
}

func (ctx *TreeContext) getScenicScoreLeft() (score int) {
	for l := ctx.column - 1; l >= 0; l-- {
		score++
		if byteToInt((*ctx.grid)[ctx.row][l]) >= ctx.current {
			return score
		}
	}
	return score
}

func (ctx *TreeContext) getScenicScoreRight() (score int) {
	for r := ctx.column + 1; r < len((*ctx.grid)[0]); r++ {
		score++
		if byteToInt((*ctx.grid)[ctx.row][r]) >= ctx.current {
			return score
		}
	}
	return score
}

type ScenicScore struct {
	top    int
	bottom int
	left   int
	right  int
}

func (s *ScenicScore) getScore() int {
	return s.top * s.bottom * s.left * s.right
}

func Task2(patch_of_trees *[]string) (score int) {

	scenic_scores := []int{}

	c := make(chan int)

	go func() <-chan int {

		defer close(c)

		ctx := TreeContext{grid: patch_of_trees}

		for i := 0; i < len(*patch_of_trees); i++ {
			for j := 0; j < len((*patch_of_trees)[0]); j++ {

				ctx.SetCoords(i, j).Build()

				ss := ScenicScore{
					top:    ctx.getScenicScoreTop(),
					bottom: ctx.getScenicScoreBottom(),
					left:   ctx.getScenicScoreLeft(),
					right:  ctx.getScenicScoreRight(),
				}

				c <- ss.getScore()
			}
		}

		return c
	}()

	for s := range c {
		scenic_scores = append(scenic_scores, s)
	}

	return slices.Max(scenic_scores)
}
