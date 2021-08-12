package equal_values

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyInt int

func TestEqual(t *testing.T) {
	var a = 100
	var b MyInt = 100
	assert.Equal(t, a, b, "")
	assert.EqualValues(t, a, b, "")
}
