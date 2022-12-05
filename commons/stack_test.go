package commons

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	stack := &Stack{}
	assert.True(t, stack.IsEmpty())

	stack.Push("test")
	assert.False(t, stack.IsEmpty())
}

func TestStack_Peek(t *testing.T) {
	stack := &Stack{}
	value, err := stack.Peek()
	assert.NotNil(t, err)

	stack.Push("1")
	value, err = stack.Peek()
	assert.Equal(t, "1", value)
	assert.Nil(t, err)

	value, err = stack.Peek()
	assert.Equal(t, "1", value)
	assert.Nil(t, err)
}

func TestStack_Pop(t *testing.T) {
	stack := &Stack{}
	value, err := stack.Pop()
	assert.NotNil(t, err)

	stack.Push("1")
	stack.Push("2")

	value, err = stack.Pop()
	assert.Equal(t, "2", value)
	assert.Nil(t, err)

	value, err = stack.Pop()
	assert.Equal(t, "1", value)
	assert.Nil(t, err)
}
