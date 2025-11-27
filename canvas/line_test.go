package canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteHorizontalLine(t *testing.T) {
	c, err := NewCanva(5, 1)
	require.NoError(t, err)
	p1 := []int{2, 1}
	p2 := []int{4, 1}
	c.NewLine(p1, p2)
	if err != nil {
		t.Fatal(err)
	}
	expected := [][]rune{
		[]rune("-***-"),
	}
	assert.Equal(t, expected, c.Grid)
}

func TestWriteVerticalLine(t *testing.T) {
	c, err := NewCanva(3, 5)
	require.NoError(t, err)
	p1 := []int{2, 2}
	p2 := []int{2, 4}
	c.NewLine(p1, p2)
	if err != nil {
		t.Fatal(err)
	}
	expected := [][]rune{
		[]rune("---"),
		[]rune("-*-"),
		[]rune("-*-"),
		[]rune("-*-"),
		[]rune("---"),
	}
	assert.Equal(t, expected, c.Grid)
}

func TestWriteDiagonalLine(t *testing.T) {
	c, err := NewCanva(5, 5)
	if err != nil {
		t.Fatal(err)
	}
	p1 := []int{1, 5}
	p2 := []int{5, 1}
	c.NewLine(p1, p2)
	expected := [][]rune{
		[]rune("----*"),
		[]rune("---*-"),
		[]rune("--*--"),
		[]rune("-*---"),
		[]rune("*----"),
	}
	assert.Equal(t, expected, c.Grid)
	p1 = []int{5, 1}
	p2 = []int{1, 5}
	c.NewLine(p1, p2)
	expected = [][]rune{
		[]rune("*---*"),
		[]rune("-*-*-"),
		[]rune("--*--"),
		[]rune("-*-*-"),
		[]rune("*---*"),
	}
	assert.Equal(t, expected, c.Grid)
}

func TestWriteArrow(t *testing.T) {
	c, err := NewCanva(7, 11)
	require.NoError(t, err)
	c.NewLine([]int{3, 2}, []int{5, 2})
	c.NewLine([]int{3, 2}, []int{3, 7})
	c.NewLine([]int{5, 2}, []int{5, 7})
	c.NewLine([]int{1, 7}, []int{7, 7})
	c.NewLine([]int{4, 10}, []int{1, 7})
	c.NewLine([]int{4, 10}, []int{7, 7})
	expected := [][]rune{
		[]rune("-------"),
		[]rune("--***--"),
		[]rune("--*-*--"),
		[]rune("--*-*--"),
		[]rune("--*-*--"),
		[]rune("--*-*--"),
		[]rune("*******"),
		[]rune("-*---*-"),
		[]rune("--*-*--"),
		[]rune("---*---"),
		[]rune("-------"),
	}
	assert.Equal(t, expected, c.Grid)
}
