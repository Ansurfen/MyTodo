package utils_test

import (
	"MyTodo/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	assert.Equal(t, len(utils.RandString(8)), 8)
}
