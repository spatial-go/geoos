package utils

import (
	"container/list"
	"sync"
)

// Stack ...
type Stack struct {
	list *list.List
	lock *sync.RWMutex
}

// NewStack ...
func NewStack() *Stack {
	_list := list.New()
	l := &sync.RWMutex{}
	return &Stack{_list, l}
}

// Push ...
func (stack *Stack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.list.PushBack(value)
}

// Pop ...
func (stack *Stack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

// Peak ...
func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

// Len ...
func (stack *Stack) Len() int {
	return stack.list.Len()
}

// Empty ...
func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}
