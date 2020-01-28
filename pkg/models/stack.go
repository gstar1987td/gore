package models

import (
	"github.com/emirpasic/gods/stacks/arraystack"
	"sync"
)

type Stack interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)

	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
	Close()
}

type ArrayStack struct {
	*arraystack.Stack
}

func NewArrayStack() Stack {
	stk := StackPool.Get().(*ArrayStack)
	stk.Clear()
	return stk
}

func (stk *ArrayStack) Close() {
	StackPool.Put(stk)
}

var StackPool = sync.Pool{
	New: func() interface{} {
		return &ArrayStack{arraystack.New()}
	},
}
