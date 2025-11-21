package canvas

import (
	"fmt"
	"testing"
)

func TestNewRectangle(t *testing.T) {
	c, err := NewCanva(5, 5)
	if err != nil {
		t.Fatal(err)
	}
	c.NewRectangle([]int{2, 2}, 3, 3)
	cr := c.Render()
	if cr != "-----\n-***-\n-*-*-\n-***-\n-----\n" {
		fmt.Println(cr)
		t.Fatal("Rectangle mismach")
	}
}
