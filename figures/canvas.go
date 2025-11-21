package figures

import (
	"fmt"
	"log"
	"strings"
)

type Canvas struct {
	W, H int
	Grid [][]rune
	Bg   rune
}

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
	log := log.Default()
	log.Println("\n", c.Render())
	return c, nil
}

func (c Canvas) Render() string {
	var b strings.Builder
	for y := 0; y < c.H; y++ {
		for x := 0; x < c.W; x++ {
			b.WriteRune(c.Grid[y][x])
		}
		b.WriteByte('\n')
	}
	return b.String()
}
