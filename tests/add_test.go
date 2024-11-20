package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}

	assert.Equal(t, Add(1, 2), 3)
}
