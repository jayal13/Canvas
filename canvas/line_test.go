package canvas

import (
	"fmt"
	"testing"
)

func TestWriteHorizontalLine(t *testing.T) {
	c, er := NewCanva(5, 1)
	if er != nil {
		t.Fatal(er)
	}
	p1 := []int{2, 1}
	p2 := []int{4, 1}
	err := c.NewLine(p1, p2)
	if err != nil {
		t.Fatal(err)
	}
	cr := c.Render()
	if cr != "-***-\n" {
		fmt.Println(cr)
		t.Fatal("wrong line")
	}
}

func TestWriteVerticalLine(t *testing.T) {
	c, er := NewCanva(3, 5)
	if er != nil {
		t.Fatal(er)
	}
	p1 := []int{2, 2}
	p2 := []int{2, 4}
	err := c.NewLine(p1, p2)
	if err != nil {
		t.Fatal(err)
	}
	cr := c.Render()
	if cr != "---\n-*-\n-*-\n-*-\n---\n" {
		fmt.Println(cr)
		t.Fatal("wrong line")
	}
}

func TestWriteDiagonalLine(t *testing.T) {
	c, err := NewCanva(5, 5)
	if err != nil {
		t.Fatal(err)
	}
	p1 := []int{1, 5}
	p2 := []int{5, 1}
	c.NewLine(p1, p2)
	cr := c.Render()
	if cr != "----*\n---*-\n--*--\n-*---\n*----\n" {
		fmt.Println(cr)
		t.Fatal("wrong line")
	}
	p1 = []int{5, 1}
	p2 = []int{1, 5}
	c.NewLine(p1, p2)
	cr = c.Render()
	if cr != "*---*\n-*-*-\n--*--\n-*-*-\n*---*\n" {
		fmt.Println(cr)
		t.Fatal("wrong line")
	}
}

func TestWriteArrow(t *testing.T) {
	c, er := NewCanva(7, 11)
	if er != nil {
		t.Fatal(er)
	}
	c.NewLine([]int{3, 2}, []int{5, 2})
	c.NewLine([]int{3, 2}, []int{3, 7})
	c.NewLine([]int{5, 2}, []int{5, 7})
	c.NewLine([]int{1, 7}, []int{7, 7})
	c.NewLine([]int{4, 10}, []int{1, 7})
	c.NewLine([]int{4, 10}, []int{7, 7})
	cr := c.Render()
	if cr != "-------\n--***--\n--*-*--\n--*-*--\n--*-*--\n--*-*--\n*******\n-*---*-\n--*-*--\n---*---\n-------\n" {
		fmt.Println(cr)
		t.Fatal("Not an arrow")
	}

}
