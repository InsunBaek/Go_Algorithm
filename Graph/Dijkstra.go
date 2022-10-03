// Map을 사용한 다익스트라 알고리즘

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var bw = bufio.NewWriter(os.Stdout)
var INF = math.MaxInt32
var V, E int

/* Priority Queue */
type EdgeHeap []*Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(u interface{}) { *h = append(*h, u.(*Edge)) }
func (h *EdgeHeap) Pop() interface{} {
	old := *h
	long := len(old)
	x := old[long-1]
	*h = old[0 : long-1]
	return *x
}

/* Graph */

type Edge struct {
	vtx  int
	cost int
}

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	V, E = scanInt(), scanInt()

	// Map을 사용하면 버텍스에서 연결된 버텍스를 바로 탐색가능
	Graph := make(map[int][]Edge, V+1)

	for i := 0; i < E; i++ {
		from, to, cost := scanInt(), scanInt(), scanInt()
		Graph[from] = append(Graph[from], Edge{to, cost})
	}

	start, end := scanInt(), scanInt()
	fmt.Fprintln(bw, Dijkstra(Graph, start)[end])
}

// 시작버텍스 start에서 시작하는 다익스트라 알고리즘
func Dijkstra(graph map[int][]Edge, start int) []int {

	visited := make([]bool, V+1)
	distance := make([]int, V+1)
	for i := 0; i < len(distance); i++ {
		distance[i] = INF
	}
	distance[start] = 0

	// 우선순위큐 초기화
	pq := &EdgeHeap{}
	heap.Init(pq)
	heap.Push(pq, &Edge{start, 0})

	for pq.Len() > 0 {
		curEdge := heap.Pop(pq).(Edge)
		vertex := curEdge.vtx
		if visited[vertex] {
			continue
		}
		visited[vertex] = true

		// 현재 버텍스(최소 비용 버텍스)에 연결된 버텍스들을 모두 탐색
		for i := 0; i < len(graph[vertex]); i++ {
			edge := graph[vertex][i]
			nextVtx := edge.vtx
			nextCost := edge.cost
			newCost := distance[vertex] + nextCost

			// 방문한 버텍스면 무시한다.
			if visited[nextVtx] {
				continue
			}

			// 새로운 route의 cost가 현재 저장된 cost보다 작으면 새로운 cost로 업데이트
			if distance[nextVtx] >= newCost {
				distance[nextVtx] = newCost
				heap.Push(pq, &Edge{nextVtx, distance[nextVtx]})
			}
		}
	}

	return distance
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

// Map이 아닌 구조체를 사용한 Graph
// package main

// import (
// 	"bufio"
// 	"container/heap"
// 	"fmt"
// 	"math"
// 	"os"
// 	"strconv"
// )

// var sc = bufio.NewScanner(os.Stdin)
// var bw = bufio.NewWriter(os.Stdout)
// var INF = math.MaxInt32

// /* Priority Queue */
// type CostHeap []*Cost

// func (h CostHeap) Len() int           { return len(h) }
// func (h CostHeap) Less(i, j int) bool { return h[i].cost_ < h[j].cost_ }
// func (h CostHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *CostHeap) Push(u interface{}) { *h = append(*h, u.(*Cost)) }
// func (h *CostHeap) Pop() interface{} {
// 	old := *h
// 	long := len(old)
// 	x := old[long-1]
// 	*h = old[0 : long-1]
// 	return *x
// }

// /* Graph */
// type Cost struct {
// 	to_   int
// 	cost_ int
// }

// func (cost *Cost) Print(from int) {
// 	fmt.Fprintln(bw, from, " -> ", cost.to_, cost.cost_)
// }

// type Edge struct {
// 	toList []Cost
// }

// func print(graph []Edge) {
// 	for i := 0; i < len(graph); i++ {
// 		for j := 0; j < len(graph[i].toList); j++ {
// 			fmt.Fprintln(bw, i, " -> ", graph[i].toList[j].to_, " cost : ", graph[i].toList[j].cost_)
// 		}
// 	}
// }

// func main() {
// 	defer bw.Flush()
// 	sc.Split(bufio.ScanWords)

// 	V := scanInt()
// 	E := scanInt()

// 	graph := make([]Edge, V+1)
// 	for i := 0; i < E; i++ {
// 		u, v, cost := scanInt(), scanInt(), scanInt()
// 		graph[u].toList = append(graph[u].toList, Cost{to_: v, cost_: cost})
// 	}

// 	from := scanInt()
// 	to := scanInt()

// 	visited := make([]bool, V+1)
// 	distance := make([]int, V+1)

// 	for i := 0; i < len(distance); i++ {
// 		distance[i] = INF
// 	}

// 	wentDistance := Dijkstra(graph, visited, distance, from, to)
// 	fmt.Fprintln(bw, wentDistance)
// }

// func Dijkstra(Graph []Edge, visited []bool, distance []int, start int, to int) int {
// 	distance[start] = 0

// 	pq := &CostHeap{}
// 	heap.Init(pq)
// 	heap.Push(pq, &Cost{to_: start, cost_: 0})

// 	for pq.Len() != 0 {
// 		currentCost := heap.Pop(pq).(Cost)
// 		minIndex := currentCost.to_
// 		// fmt.Fprintln(bw, "Now -> ", minIndex, currentCost.cost_)
//         if visited[minIndex] {
//             continue
//         }
// 		visited[minIndex] = true

// 		for i := 0; i < len(Graph[minIndex].toList); i++ {
// 			currentEdge := Graph[minIndex].toList[i]
// 			to := currentEdge.to_
// 			cost := currentEdge.cost_
// 			if !visited[to] && distance[to] > distance[minIndex]+cost {
// 				distance[to] = distance[minIndex] + cost
// 				// fmt.Fprintln(bw, "Push -> ", to, cost)
// 				heap.Push(pq, &Cost{to_: to, cost_: distance[to]})
// 			}
// 		}
// 	}

// 	return distance[to]

// }

// func scanInt() int {
// 	sc.Scan()
// 	v, _ := strconv.Atoi(sc.Text())
// 	return v
// }
