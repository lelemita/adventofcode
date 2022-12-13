package common_test

import (
	"fmt"
	"testing"

	"github.com/lelemita/adventofcode/common"
	"github.com/stretchr/testify/assert"
)

func TestRuneStack(t *testing.T) {
	assert := assert.New(t)
	testSlice := []rune{'a', 'b', 'c', 'd', 'e'}
	stack := common.NewRuneStack()
	for i, r := range testSlice {
		stack.Push(r)
		assert.Equal(i+1, len(stack.Data))
	}
	assert.False(stack.IsEmpty())

	for i := len(stack.Data) - 1; i >= 0; i-- {
		obj := stack.Pop()
		assert.Equal(testSlice[i], *obj)
		fmt.Println(string(*obj))
	}
	assert.True(stack.IsEmpty())
}
