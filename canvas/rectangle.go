package canvas

import "fmt"

func (c *Canvas) NewRectangle(p []int, w int, h int) error {
	if len(p) > 2 {
		return fmt.Errorf("this is a 2d canvas, Use a point with two coordinates")
	}
	if p[0] > c.W || p[1] > c.H {
		return fmt.Errorf("line out of bounds")
	}
	w--
	h--
	if w < 0 || h < 0 {
		return fmt.Errorf("invalid rectangle size")
	}
	c.NewLine([]int{p[0], p[1]}, []int{p[0] + w, p[1]})
	c.NewLine([]int{p[0], p[1]}, []int{p[0], p[1] + h})
	c.NewLine([]int{p[0], p[1] + h}, []int{p[0] + w, p[1] + h})
	c.NewLine([]int{p[0] + w, p[1]}, []int{p[0] + w, p[1] + h})
	return nil
}
