package day8

type ScenicScore struct {
	row     int
	column  int
	current int
	grid    *[]string
}

func (s *ScenicScore) SetCoords(row, column int) *ScenicScore {
	return s.SetRow(row).SetColumn(column).SetCurrent()
}

func (s *ScenicScore) SetRow(row int) *ScenicScore {
	s.row = row
	return s
}

func (s *ScenicScore) SetColumn(column int) *ScenicScore {
	s.column = column
	return s
}

func (s *ScenicScore) SetCurrent() *ScenicScore {
	s.current = byteToInt((*s.grid)[s.row][s.column])
	return s
}

func (s *ScenicScore) scoreTop() (score int) {
	for t := s.row - 1; t >= 0; t-- {
		score++
		if byteToInt((*s.grid)[t][s.column]) >= s.current {
			return score
		}
	}
	return score
}

func (s *ScenicScore) scoreBottom() (score int) {
	for b := s.row + 1; b < len(*s.grid); b++ {
		score++
		if byteToInt((*s.grid)[b][s.column]) >= s.current {
			return score
		}
	}
	return score
}

func (s *ScenicScore) scoreLeft() (score int) {
	for l := s.column - 1; l >= 0; l-- {
		score++
		if byteToInt((*s.grid)[s.row][l]) >= s.current {
			return score
		}
	}
	return score
}

func (s *ScenicScore) scoreRight() (score int) {
	for r := s.column + 1; r < len((*s.grid)[0]); r++ {
		score++
		if byteToInt((*s.grid)[s.row][r]) >= s.current {
			return score
		}
	}
	return score
}

func (s *ScenicScore) GetScore() int {
	return s.scoreTop() * s.scoreBottom() * s.scoreLeft() * s.scoreRight()
}

func Task2(patch_of_trees *[]string) (score int) {

	c := make(chan int)

	go func() <-chan int {

		defer close(c)

		s := ScenicScore{grid: patch_of_trees}

		for i := 0; i < len(*patch_of_trees); i++ {
			for j := 0; j < len((*patch_of_trees)[0]); j++ {

				c <- s.SetCoords(i, j).GetScore()

			}
		}

		return c
	}()

	for s := range c {
		if score < s {
			score = s
		}
	}

	return score
}
