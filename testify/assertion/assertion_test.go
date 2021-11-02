package assertion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertion(t *testing.T) {
	assertion := assert.New(t)
	var a int = 100
	var b int = 100
	assertion.Equal(a, b, "")
}
