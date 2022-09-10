// Implementation of Queue through List

package main

import (
	"bufio"
	"container/list"
	"os"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

type Queue struct {
	v *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front == nil {
		return nil
	}
	return q.v.Remove(front)
}

func (q *Queue) Empty() bool {
	if q.v.Len() == 0 {
		return true
	}
	return false
}

func (q *Queue) Size() int {
	return q.v.Len()
}

func main() {

	// queue := NewQueue()

	// queue.Push(1)
	// queue.Push(2)
	// queue.Push(3)

	// fmt.Println(queue.Pop())
	// fmt.Println(queue.Pop())
	// fmt.Println(queue.Pop())

}
