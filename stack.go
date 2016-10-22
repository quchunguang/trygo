package trygo

import "container/list"

// Stack struct
type Stack struct {
	list *list.List
}

// NewStack create a new stack.
func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

// Push item to stack.
func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

// Pop item from stack.
func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

// Peak get the top item of the stack.
func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

// Len get the length of the stack.
func (stack *Stack) Len() int {
	return stack.list.Len()
}

// Empty tests if the stack is empty.
func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}
