package list

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushFront(t *testing.T) {
	assert := assert.New(t)
	l := List{}
	l.PushFront(12)
	assert.Equal(1, l.Len, "should have the correct length")
	if assert.NotNil(l.First) {
		assert.Equal(l.First, l.Last, "should have the same reference to first and last element")
	}
	l.PushFront(13)
	assert.Equal(2, l.Len, "should have the correct length")
	assert.NotEqual(l.First, l.Last, "should have the different references to first and last element")
}

func TestPushBack(t *testing.T) {
	assert := assert.New(t)
	l := List{}
	l.PushBack(12)
	assert.Equal(1, l.Len, "should have the correct length")
	if assert.NotNil(l.First) {
		assert.Equal(l.First, l.Last, "should have the same reference to first and last element")
	}
	l.PushBack(13)
	assert.Equal(2, l.Len, "should have the correct length")
	assert.NotEqual(l.First, l.Last, "should have the different references to first and last element")
}

func TestRemoveInNotExistedIndex(t *testing.T) {
	assert := assert.New(t)
	l := List{}
	assert.Equal(errors.New("index is put of range"), l.Remove(2), "should throw error, if they is no element with provided index")
	assert.Equal(errors.New("index is put of range"), l.Remove(-1), "should throw error, if they is no element with provided index")
}

func TestNewListWithOneElement(t *testing.T) {
	assert := assert.New(t)
	slice := make([]interface{}, 1)
	slice[0] = 12
	l := NewList(slice)
	assert.Equal(1, l.Len, "should have the correct length")
	if assert.NotNil(l.First) {
		assert.Equal(l.First, l.Last, "should have the same reference to first and last element")
	}
}

func TestNewListWithMultipleElements(t *testing.T) {
	assert := assert.New(t)
	slice := make([]interface{}, 3)
	slice[0] = 12
	slice[1] = 13
	slice[2] = 14
	l := NewList(slice)
	assert.Equal(3, l.Len, "should have the correct length")
	assert.NotNil(l.First, "should have correct First reference")
	assert.NotNil(l.Last, "should have correct Last reference")
}

func TestRemoveSingleElement(t *testing.T) {
	assert := assert.New(t)
	// Is there easier way to convert slice of ints to slice of interfaces?
	slice := make([]interface{}, 1)
	slice[0] = 12
	l := NewList(slice)
	l.Remove(0)
	assert.Nil(l.First, "should reset the First reference")
	assert.Nil(l.Last, "should reset the Last reference")
	assert.Equal(0, l.Len, "should decrease the length")
}

func TestRemoveFirstElement(t *testing.T) {
	assert := assert.New(t)
	slice := make([]interface{}, 3)
	slice[0] = 12
	slice[1] = 13
	slice[2] = 14
	l := NewList(slice)
	first := l.First
	last := l.Last
	l.Remove(0)
	assert.NotEqual(first, l.First, "should change the First ref")
	assert.Equal(last, l.Last, "should not change the Last ref")
	assert.Equal(2, l.Len, "should have the correct length")
}

func TestRemoveLastElement(t *testing.T) {
	assert := assert.New(t)
	slice := make([]interface{}, 3)
	slice[0] = 12
	slice[1] = 13
	slice[2] = 14
	l := NewList(slice)
	first := l.First
	last := l.Last
	l.Remove(2)
	assert.Equal(first, l.First, "should not change the First ref")
	assert.NotEqual(last, l.Last, "should change the Last ref")
	assert.Equal(2, l.Len, "should have the correct length")
}

func TestRemoveInnerElement(t *testing.T) {
	assert := assert.New(t)
	slice := make([]interface{}, 3)
	slice[0] = 12
	slice[1] = 13
	slice[2] = 14
	l := NewList(slice)
	first := l.First
	last := l.Last
	l.Remove(1)
	assert.Equal(first, l.First, "should not change the First ref")
	assert.Equal(last, l.Last, "should not change the Last ref")
	assert.Equal(2, l.Len, "should have the correct length")
}
