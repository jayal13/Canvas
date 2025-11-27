package canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFill(t *testing.T) {
	c, err := NewCanva(8, 8)
	require.NoError(t, err)
	c.NewRectangle([]int{2, 2}, 6, 6)
	c.Fill([]int{3, 3})
	expected := [][]rune{
		[]rune("--------"),
		[]rune("-******-"),
		[]rune("-*++++*-"),
		[]rune("-*++++*-"),
		[]rune("-*++++*-"),
		[]rune("-*++++*-"),
		[]rune("-******-"),
		[]rune("--------"),
	}
	assert.Equal(t, expected, c.Grid)
}

func TestFillCenterPoint(t *testing.T) {
	c, err := NewCanva(8, 8)
	require.NoError(t, err)
	c.NewRectangle([]int{2, 2}, 6, 6)
	c.Fill([]int{4, 4})
	expected := [][]rune{
		[]rune("--------"),
		[]rune("-******-"),
		[]rune("-*++++*-"),
		[]rune("-*++++*-"),
		[]rune("-*++++*-"),
		[]rune("-*++++*-"),
		[]rune("-******-"),
		[]rune("--------"),
	}
	assert.Equal(t, expected, c.Grid)
}

func TestFillOut(t *testing.T) {
	c, err := NewCanva(10, 10)
	require.NoError(t, err)

	c.NewRectangle([]int{3, 3}, 6, 6)
	c.Fill([]int{1, 1})
	expected := [][]rune{
		[]rune("++++++++++"),
		[]rune("++++++++++"),
		[]rune("++******++"),
		[]rune("++*----*++"),
		[]rune("++*----*++"),
		[]rune("++*----*++"),
		[]rune("++*----*++"),
		[]rune("++******++"),
		[]rune("++++++++++"),
		[]rune("++++++++++"),
	}
	assert.Equal(t, expected, c.Grid)
}
