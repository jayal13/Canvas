package canvas

import "fmt"

func (c *Canvas) NewLine(p1, p2 []int) error {
	if len(p1) != 2 || len(p2) != 2 {
		return fmt.Errorf("this is a 2d canvas, Use a point with two coordinates")
	}
	p1[0]--
	p1[1]--
	p2[0]--
	p2[1]--
	if 0 > p1[0] || p1[0] > c.W || 0 > p2[0] || p2[0] > c.W {
		return fmt.Errorf("line out of bounds")
	}
	if p1[1] == p2[1] {
		for i := p1[0]; i <= p2[0]; i++ {
			c.Grid[p1[1]][i] = '*'
		}
		return nil
	}
	if p1[0] == p2[0] {
		for i := p1[1]; i <= p2[1]; i++ {
			c.Grid[i][p1[0]] = '*'
		}
		return nil
	}
	j := 0
	if p1[1] < p2[1] {
		j = p1[0]
		for i := p2[1]; i >= p1[1]; i-- {
			c.Grid[i][j] = '*'
			if p1[0] < p2[0] {
				j++
			} else {
				j--
			}
		}
		return nil
	}
	j = p2[0]
	for i := p2[1]; i <= p1[1]; i++ {
		c.Grid[i][j] = '*'
		if p2[0] < p1[0] {
			j++
		} else {
			j--
		}
	}
	return nil
}
