package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (st *Stack) Push(v interface{}) {
	st.v.PushBack(v)
}

func (st *Stack) Pop() interface{} {
	top := st.v.Back()
	if top == nil {
		return nil
	}
	return st.v.Remove(top)
}

func (st *Stack) Empty() bool {
	if st.v.Len() == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	stack := *NewStack()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	for {
		if stack.Empty() {
			break
		}
		fmt.Println(stack.Pop().(int))
	}
	// 30 20 10

}
