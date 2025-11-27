package canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateCanva(t *testing.T) {
	h := 10
	w := 30
	c, err := NewCanva(w, h)
	require.NoError(t, err)

	assert.Equal(t, c.H, h)
	assert.Equal(t, c.W, w)
}
