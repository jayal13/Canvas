package canvas

import (
	"fmt"
	"testing"
)

func TestCreateCanva(t *testing.T) {
	h := 10
	w := 30
	c, err := NewCanva(w, h)
	if err != nil {
		t.Fatal(err)
	}

	if c.W != w || c.H != h {
		fmt.Println(c.Render())
		t.Fatalf("Expected %dx%d, got %dx%d", w, h, c.W, c.H)
	}
}
