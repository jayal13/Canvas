package figures

import "testing"

func TestCreateCanva(t *testing.T) {
	h := 10
	w := 30
	c, err := NewCanva(w, h)
	if err != nil {
		t.Fatal(err)
	}

	if c.W != w || c.H != h {
		t.Fatalf("Expected %dx%d, got %dx%d", w, h, c.W, c.H)
	}
}
