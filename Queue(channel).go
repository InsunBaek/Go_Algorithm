// Implementation of Queue through Go Channel

package main

import (
	"bufio"
	"os"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

type Queue struct {
	item chan interface{}
	cnt  int
}

func (q *Queue) Push(val interface{}) {
	q.item <- val
	q.cnt += 1
}

func (q *Queue) Pop() interface{} {
	q.cnt -= 1
	return <-q.item
}

func (q *Queue) Size() int {
	return q.cnt
}

func (q *Queue) Empty() bool {
	if q.cnt == 0 {
		return true
	}
	return false
}

func main() {

	// queue := Queue{item: make(chan interface{}, 5), cnt: 0}

	// queue.Push(1)
	// queue.Push(2)
	// queue.Push(3)

	// fmt.Println(queue.Pop())
	// fmt.Println(queue.Pop())
	// fmt.Println(queue.Pop())

}
