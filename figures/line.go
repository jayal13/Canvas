package figures

import "fmt"

func (c Canvas) NewLine(p1, p2 []int) error {
	if p1[1] == p2[1] {
		for i := p1[0]; i <= p2[0]; i++ {
			c.Grid[p1[1]][i] = '*'
		}
		fmt.Println(c.Render())
		return nil
	}
	if p1[0] == p2[0] {
		for i := p1[1]; i <= p2[1]; i++ {
			c.Grid[i][p1[0]] = '*'
		}
		fmt.Println(c.Render())
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
		fmt.Println(c.Render())
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
	fmt.Println(c.Render())
	return nil
}
