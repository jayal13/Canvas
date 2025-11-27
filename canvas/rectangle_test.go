package canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRectangle(t *testing.T) {
	c, err := NewCanva(5, 5)
	require.NoError(t, err)
	c.NewRectangle([]int{2, 2}, 3, 3)
	expected := [][]rune{
		[]rune("-----"),
		[]rune("-***-"),
		[]rune("-*-*-"),
		[]rune("-***-"),
		[]rune("-----"),
	}
	assert.Equal(t, expected, c.Grid)
}
