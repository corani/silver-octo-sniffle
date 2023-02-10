package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken_Range(t *testing.T) {
	assert.Panics(t, func() {
		NewRange(1, 2, 3)
	})
}
