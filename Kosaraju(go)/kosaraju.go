/*
SCC(강한연결요소) -> 코사라주 알고리즘, 타잔 알고리즘
해당 파일은 코사라주 알고리즘을 다룬다.

Graph information을 받았을 때 Forward, Reverse Graph 두 가지 형태로 저장
첫 번째 DFS 탐색에서 깊은 node부터 스택에 저장히고 Reverse Graph에서 top부터
탐색을 시작하여 도달할 수 있는 node는 SCC로 판단 가능하다.
*/

package main

import (
	"bufio"
	"container/list"
	"os"
	"sort"
	"strconv"
	"strings"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

var stack Stack = *NewStack()
var Graph []Edge
var InvGraph []Edge
var opt []Edge
var visited []int
var cnt int = 0

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

// Forward Graph DFS search
func DFS(x int) {
	if visited[x] == 1 { // 방문한 버텍스라면 무시한다.
		return
	}
	visited[x] = 1 // 방문 처리
	for i := 0; i < len(Graph[x].edges); i++ {
		temp := Graph[x].edges[i]
		DFS(temp)
	}

	stack.Push(x)
}

// Reverse Graph DFS search
func InvDFS(x int, index int) {

	visited[x] = 1
	opt[index].edges = append(opt[index].edges, x)

	for i := 0; i < len(InvGraph[x].edges); i++ {
		next := InvGraph[x].edges[i]
		if visited[next] == 0 {
			InvDFS(next, index) // next node -> SCC에 포함
		}
	}
}

func main() {

	br.Scan()
	line := strings.Split(br.Text(), " ")
	N, _ := strconv.Atoi(line[0])
	M, _ := strconv.Atoi(line[1])

	Graph = make([]Edge, N+1)
	InvGraph = make([]Edge, N+1)
	opt = make([]Edge, N+1)
	visited = make([]int, N+1)

	for i := 0; i < N; i++ {
		Graph[i].edges = make([]int, 0)
	}
	for i := 0; i < N; i++ {
		InvGraph[i].edges = make([]int, 0)
	}
	for i := 0; i < N; i++ {
		opt[i].edges = make([]int, 0)
	}

	for i := 0; i < M; i++ {
		br.Scan()
		line := strings.Split(br.Text(), " ")
		from, _ := strconv.Atoi(line[0])
		to, _ := strconv.Atoi(line[1])

		// Forward Graph
		Graph[from].edges = append(Graph[from].edges, to)
		// Reverse Graph
		InvGraph[to].edges = append(InvGraph[to].edges, from)
	}

	for i := 1; i <= N; i++ {
		if visited[i] == 0 {
			DFS(i)
		}
	}

	visited = make([]int, N+1)
	for {
		if stack.Empty() {
			break
		}

		top := stack.Pop().(int)
		if visited[top] == 1 {
			continue
		}
		cnt++
		InvDFS(top, top)
	}

	bw.WriteString(strconv.Itoa(cnt) + "\n")

	// for print
	tempSlice := make([]Edge, 0)
	for i := 0; i < len(opt); i++ {
		slice := opt[i].edges
		if len(slice) == 0 {
			continue
		}
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
		tempSlice = append(tempSlice, Edge{slice})
	}

	sort.Slice(tempSlice, func(i, j int) bool {
		return tempSlice[i].edges[0] < tempSlice[j].edges[0]
	})

	for i := 0; i < len(tempSlice); i++ {
		slice := tempSlice[i].edges

		for i := 0; i < len(slice); i++ {
			bw.WriteString(strconv.Itoa(slice[i]) + " ")
		}
		bw.WriteString(" end\n")
	}

	bw.Flush()

}
