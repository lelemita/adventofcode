package common_test

import (
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
		assert.Equal(i+1, stack.Len())
	}
	assert.False(stack.IsEmpty())

	for i := stack.Len() - 1; i >= 0; i-- {
		obj := stack.Pop()
		assert.Equal(testSlice[i], *obj)
		// fmt.Println(string(*obj))
	}
	assert.True(stack.IsEmpty())
}

func TestIntStack(t *testing.T) {
	assert := assert.New(t)
	testSlice := []int{1, 2, 3, 4, 5}

	stack := common.NewIntStack()
	for i, r := range testSlice {
		stack.Push(r)
		assert.Equal(i+1, stack.Len())
	}
	assert.False(stack.IsEmpty())

	for i := stack.Len() - 1; i >= 0; i-- {
		obj := stack.Pop()
		assert.Equal(testSlice[i], *obj)
		// fmt.Println(*obj)
	}
	assert.True(stack.IsEmpty())
}
