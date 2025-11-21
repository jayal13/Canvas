package canvas

import (
	"strings"
)

type Canvas struct {
	W, H int
	Grid [][]rune
	Bg   rune
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
