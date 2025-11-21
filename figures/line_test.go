package figures

import "testing"

func TestWriteHorizontalLine(t *testing.T) {
	c, er := NewCanva(5, 1)
	if er != nil {
		t.Fatal(er)
	}
	p1 := []int{1, 0}
	p2 := []int{3, 0}
	err := c.NewLine(p1, p2)
	if err != nil {
		t.Fatal(err)
	}
	cr := c.Render()
	if cr != "-***-\n" {
		t.Fatal("wrong line")
	}
}

func TestWriteVerticalLine(t *testing.T) {
	c, er := NewCanva(3, 5)
	if er != nil {
		t.Fatal(er)
	}
	p1 := []int{1, 1}
	p2 := []int{1, 3}
	err := c.NewLine(p1, p2)
	if err != nil {
		t.Fatal(err)
	}
	cr := c.Render()
	if cr != "---\n-*-\n-*-\n-*-\n---\n" {
		t.Fatal("wrong line")
	}
}

func TestWriteDiagonalLine(t *testing.T) {
	c, err := NewCanva(5, 5)
	if err != nil {
		t.Fatal(err)
	}
	p1 := []int{0, 4}
	p2 := []int{4, 0}
	c.NewLine(p1, p2)
	cr := c.Render()
	if cr != "----*\n---*-\n--*--\n-*---\n*----\n" {
		t.Fatal("wrong line")
	}
	p1 = []int{4, 0}
	p2 = []int{0, 4}
	c.NewLine(p1, p2)
	cr = c.Render()
	if cr != "*---*\n-*-*-\n--*--\n-*-*-\n*---*\n" {
		t.Fatal("wrong line")
	}
}

func TestWriteArrow(t *testing.T) {
	c, er := NewCanva(7, 11)
	if er != nil {
		t.Fatal(er)
	}
	c.NewLine([]int{2, 1}, []int{4, 1})
	c.NewLine([]int{2, 1}, []int{2, 6})
	c.NewLine([]int{4, 1}, []int{4, 6})
	c.NewLine([]int{0, 6}, []int{6, 6})
	c.NewLine([]int{3, 9}, []int{0, 6})
	c.NewLine([]int{3, 9}, []int{6, 6})
	cr := c.Render()
	if cr != "-------\n--***--\n--*-*--\n--*-*--\n--*-*--\n--*-*--\n*******\n-*---*-\n--*-*--\n---*---\n-------" {
		t.Fatal("Not an arrow")
	}

}
