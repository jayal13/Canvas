package canvas

import "fmt"

func (c Canvas) Fill(p []int) error {
	if len(p) != 2 {
		return fmt.Errorf("this is a 2d canvas, Use a point with two coordinates")
	}
	x := p[0] - 1
	y := p[1] - 1
	fp := c.Grid[y][x]
	//for y
	for j := y; j < c.H; j++ {
		for i := x; i < c.W; i++ {
			c.Grid[j][i] = '+'
			if i+2 > c.W || fp != c.Grid[j][i+1] {
				break
			}
		}
		for i := x; i > 0; i-- {
			c.Grid[j][i] = '+'
			if i-1 < 0 || fp != c.Grid[j][i-1] {

				break
			}
		}
		if j+2 > c.H || fp != c.Grid[j+1][x] {
			break
		}
	}
	for j := y; j > 0; j-- {
		for i := x; i < c.W; i++ {
			c.Grid[j][i] = '+'
			if i+1 > c.W || fp != c.Grid[j][i+1] {

				break
			}
		}
		for i := x; i > 0; i-- {
			c.Grid[j][i] = '+'
			if i-1 < 0 || fp != c.Grid[j][i-1] {

				break
			}
		}
		if j-1 < 0 || fp != c.Grid[j-1][x] {
			break
		}
	}
	return nil
}
