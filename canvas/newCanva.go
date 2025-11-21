package canvas

import "fmt"

func NewCanva(w, h int) (*Canvas, error) {
	if w <= 0 || h <= 0 {
		return nil, fmt.Errorf("invalid canvas size")
	}
	g := make([][]rune, h)
	for i := range g {
		row := make([]rune, w)
		for j := range row {
			row[j] = '-'
		}
		g[i] = row
	}
	c := &Canvas{
		W:    w,
		H:    h,
		Grid: g,
		Bg:   '-',
	}
	return c, nil
}
