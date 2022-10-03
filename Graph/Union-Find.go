// Union-find

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

var parent []int

func main() {
	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	N, query := scanInt(), scanInt()
	parent = make([]int, N)
	for i := 0; i < N; i++ {
		parent[i] = -1
	}

	for i := 0; i < query; i++ {
		vtx1, vtx2 := scanInt(), scanInt()
		p1, p2 := find(vtx1), find(vtx2)
		if p1 == p2 {
			fmt.Println(i + 1)
			os.Exit(0)
		}
		union(vtx1, vtx2)
	}
}

// weighting Rule
func union(x_, y_ int) {
	x := find(x_)
	y := find(y_)

	// root가 같다면 사이클 형성
	if x == y {
		return
	}

	// 적은 node를 leaf로 가진 node를 아래에 붙임
	if parent[x] < parent[y] {
		parent[x] = parent[x] + parent[y]
		parent[y] = x
	} else {
		parent[y] = parent[y] + parent[x]
		parent[x] = y
	}
}

// Collapsing Rule
func find(node int) int {
	var root, trail, front int
	
	// node(left)에서 출발해서 parent값이 음수인(뿌리)노드를 탐색
	for root = node; parent[root] >= 0; {
		root = parent[root]
	}
	// 뿌리까지 올라가는데 중간에 만나는 노드들을 이미 알고있는 뿌리로 다이렉트로 연결
	for trail = node; trail != root; trail = front {
		front = parent[trail]
		parent[trail] = root
	}
	return root
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
