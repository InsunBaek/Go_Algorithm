package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

type Unit struct {
	value    int
	priority int
}

type UnitHeap []*Unit

func (h UnitHeap) Len() int           { return len(h) }
func (h UnitHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (h UnitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *UnitHeap) Push(u interface{}) { *h = append(*h, u.(*Unit)) }
func (h *UnitHeap) Pop() interface{} {
	old := *h
	long := len(old)
	x := old[long-1]
	*h = old[0 : long-1]
	return *x
}

// 우선순위 작은 순으로 Unit 정렬
func main() {
	defer bw.Flush()
	h := &UnitHeap{}
	heap.Init(h)

	br.Scan()
	try, _ := strconv.Atoi(br.Text())

	for i := 0; i < try; i++ {
		fmt.Print("Input value, priority : ")
		br.Scan()
		line := strings.Split(br.Text(), " ")
		value_, _ := strconv.Atoi(line[0])
		priority_, _ := strconv.Atoi(line[1])

		heap.Push(h, &Unit{value: value_, priority: priority_})
	}

	for h.Len() != 0 {
		unit := heap.Pop(h).(Unit)
		fmt.Println(unit.value, unit.priority)
	}
}
