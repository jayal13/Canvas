package canvas

import "fmt"

func (c *Canvas) Fill(p []int) error {
	if len(p) != 2 {
		return fmt.Errorf("this is a 2d canvas, use a point with two coordinates")
	}

	x := p[0] - 1
	y := p[1] - 1

	if x < 0 || x >= c.W || y < 0 || y >= c.H {
		return fmt.Errorf("point out of bounds")
	}

	target := c.Grid[y][x]

	if target == '+' {
		return nil
	}

	type pt struct {
		x, y int
	}

	queue := []pt{{x, y}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		cx, cy := cur.x, cur.y

		if cx < 0 || cx >= c.W || cy < 0 || cy >= c.H {
			continue
		}

		if c.Grid[cy][cx] != target {
			continue
		}

		c.Grid[cy][cx] = '+'

		queue = append(queue,
			pt{cx + 1, cy},
			pt{cx - 1, cy},
			pt{cx, cy + 1},
			pt{cx, cy - 1},
		)
	}

	return nil
}
