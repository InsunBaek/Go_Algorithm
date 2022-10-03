/*
Z Algorithm
can make Z array while O(n)

1. i > r : 현재 구하고자 하는 Z[i]보다 공통 최대 접두사의 오른쪽 끝 index(r)이 왼쪽에 있기 때문에 사용할 수 있는 정보가 없습니다, 따라서 l = r = i로 설정하고 S[l...] 비교해 주면서 같다면 r을 늘리면서 구해줍니다.
2. i ≤ r && i+Z[i−l]≤ r : S[l... r) = S[0... r-l） 이므로 기존에 구한 Z 배열 값을 재사용 할 수 있으며 Z[i] = Z[i – l]
3. i ≤ r && i+Z[i−l] > r : 이때는 i+Z[i-l]이 r을 초과하기 때문에 Z[i]이 r-i보다 크거나 같다는 건 알 수 있지만 r을 초과하는 경우는 직접 비교해 주어야 합니다.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

var br *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func ZArray(str string) []int {
	Z := make([]int, len(str))
	l, r := 0, 0
	n := len(str)
	Z[0] = n

	for i := 1; i < n; i++ {
		// case 1 -> (i > r)
		if i > r {
			l = i
			r = i
			for r < len(str) && str[r] == str[r-l] {
				r++
			}
			Z[i] = r - l
			r--
		} else {
			// case 2-1 -> (i <= r) && (i + Z[i - l] <= r)
			if i+Z[i-l] <= r {
				Z[i] = Z[i-l]
			} else {
				// cout << "case 2-2 -> (i <= r) && (i + Z[i - l] > r)
				l = i
				for r < len(str) && str[r] == str[r-l] {
					r++
				}
				Z[i] = r - l
				r--
			}
		}
	}
	return Z
}
func main() {

	br.Scan()
	text := br.Text()
	z := ZArray(text)

	fmt.Println(z)

	// Input = ABCABCABAB -> Z array : [10 0 0 5 0 0 2 0 2 0]
}
