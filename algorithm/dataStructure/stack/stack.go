package stack

import (
	"errors"
)

// 保存数据的结构体
type Stack []interface{}

// 返回栈元素个数
func (stack Stack) Len() int {
	return len(stack)
}

// 查看栈的容量
func (stack Stack) Cap() int {
	return cap(stack)
}

// Push
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

// Pop
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	theStack.checkRange()
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

func (stack Stack) Top() (interface{}, error) {
	stack.checkRange()
	value := stack[len(stack)-1]
	return value, nil
}

func (stack Stack) checkRange() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("null point exception")
	}
	return nil, nil
}

// isEmpty
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}
