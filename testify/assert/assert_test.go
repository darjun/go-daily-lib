package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	var a = 100
	var b = 200
	assert.Equal(t, a, b, "")
}
