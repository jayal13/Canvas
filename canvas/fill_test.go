package canvas

import (
	"fmt"
	"testing"
)

func TestFill(t *testing.T) {
	c, err := NewCanva(8, 8)
	if err != nil {
		t.Fatal(err)
	}
	c.NewRectangle([]int{2, 2}, 6, 6)
	c.Fill([]int{3, 3})
	cr := c.Render()
	if cr != "--------\n-******-\n-*++++*-\n-*++++*-\n-*++++*-\n-*++++*-\n-******-\n--------\n" {
		fmt.Println(cr)
		t.Fatal("Bad filling")
	}
}

func TestFillCenterPoint(t *testing.T) {
	c, err := NewCanva(8, 8)
	if err != nil {
		t.Fatal(err)
	}
	c.NewRectangle([]int{2, 2}, 6, 6)
	c.Fill([]int{4, 4})
	cr := c.Render()
	if cr != "--------\n-******-\n-*++++*-\n-*++++*-\n-*++++*-\n-*++++*-\n-******-\n--------\n" {
		fmt.Println(cr)
		t.Fatal("Bad filling")
	}
}

func TestFillOut(t *testing.T) {
	c, err := NewCanva(10, 10)
	if err != nil {
		t.Fatal(err)
	}
	c.NewRectangle([]int{3, 3}, 6, 6)
	c.Fill([]int{1, 1})
	cr := c.Render()
	if cr != "--------\n-******-\n-*++++*-\n-*++++*-\n-*++++*-\n-*++++*-\n-******-\n--------\n" {
		fmt.Println(cr)
		t.Fatal("Bad filling")
	}
}
