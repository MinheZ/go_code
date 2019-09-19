package stack

import (
	"fmt"
	"testing"
)

func TestStack_Push(t *testing.T) {
	var mStack Stack
	mStack.Push(3)
	if mStack.Len() == 1 {
		t.Log("Pass Stack.Push")
	} else {
		t.Error("Failed Stack.Push")
	}
}

func TestStack_Cap(t *testing.T) {
	myStack := make(Stack, 3)
	if myStack.Cap() == 3 {
		t.Log("Pass Stack.Cap")
	} else {
		t.Error("Failed Stack.Cap")
	}
}

func TestStack_Pop(t *testing.T) {
	var mStack Stack
	mStack.Push(1)
	mStack.Push("abc")
	m1, _ := mStack.Pop()
	m2, _ := mStack.Pop()
	fmt.Println(m1, m2)
}
