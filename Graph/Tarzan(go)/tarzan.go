/*
SCC(강한연결요소) -> 코사라주 알고리즘, 타잔 알고리즘
해당 파일은 타잔 알고리즘을 다룬다.
ING
*/

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

type SCC struct {
	sccList []int
}

type Edge struct {
	edges []int
}

/* Stack */
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

func min(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

// Forward Graph DFS search
var stack Stack = *NewStack()
var finished []int
var node []int
var graph []Edge
var scc []SCC
var id int = 0

func Tarzan(now int) int {
	id += 1
	node[now] = id
	stack.Push(now)
	fmt.Println("푸쉬 : ", now, " id : ", id)
	parent := node[now]
	for i := 0; i < len(graph[now].edges); i++ {
		next := graph[now].edges[i]
		fmt.Println("next : ", next)
		if node[next] == 0 {
			parent = min(parent, Tarzan(next))
			fmt.Println(now, " 의 부모 ", parent)
		} else if finished[next] == 0 {
			parent = min(parent, node[next])
			fmt.Println(now, " 의 부모 2", parent)
		}
	}

	fmt.Println(now, " parent :  ", parent, " node : ", node[now])
	if parent == node[now] {
		partScc := make([]int, 0)
		for {
			top := stack.Pop().(int)
			fmt.Println("팝 : ", top, " now : ", now)
			partScc = append(partScc, top)
			finished[top] = 1
			if top == now {
				break
			}
		}
		print("partScc : ")
		scc = append(scc, SCC{sccList: partScc})
	}

	return parent
}

func printStack(st Stack) {
	fmt.Print("top -> ")
	for {
		if st.Empty() {
			break
		}
		fmt.Print(st.Pop().(int), " ")
	}
}

func main() {

	br.Scan()
	line := strings.Split(br.Text(), " ")

	vertex, _ := strconv.Atoi(line[0])
	nodes, _ := strconv.Atoi(line[1])

	finished = make([]int, vertex+1)
	node = make([]int, vertex+1)
	graph = make([]Edge, vertex+1)
	scc = make([]SCC, 0)

	for i := 0; i < vertex; i++ {
		graph[i].edges = make([]int, 0)
	}

	for i := 0; i < nodes; i++ {
		br.Scan()
		line := strings.Split(br.Text(), " ")
		from, _ := strconv.Atoi(line[0])
		to, _ := strconv.Atoi(line[1])
		graph[from].edges = append(graph[from].edges, to)
	}

	// for i := 0; i < len(graph); i++ {
	// 	fmt.Print(i, " 의 리프 : ")
	// 	for j := 0; j < len(graph[i].edges); j++ {
	// 		fmt.Print(graph[i].edges[j], " ")
	// 	}
	// 	fmt.Println()
	// }

	for i := 1; i <= vertex; i++ {
		if node[i] == 0 {
			Tarzan(i)
		}
	}
	fmt.Println("-----------")
	for i := 0; i < len(scc); i++ {
		for j := 0; j < len(scc[i].sccList); j++ {
			fmt.Print(scc[i].sccList[j], " ")
		}
		fmt.Println()
	}

	bw.Flush()

}
