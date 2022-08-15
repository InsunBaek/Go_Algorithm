/* 수열의 substring에 대한 XOR연산의 합
음이아닌 수열 A1, A2, A3, ..., An에서 substring A(L...R)이라고 할때
A(L) ⊕ A(L+1) ⊕ A(L+2) ⊕ ... ⊕ A(R) 을 구하고자 한다면
V0 = 0, V1 = A1, Vi = A1 ⊕ A2 ⊕ A3 ⊕...⊕ AR 이라 가정하면
결과는 V(L-1) ⊕ VR

추가적으로
x % 4 == 3 인 정수 x에 대하여 1 ⊕ 2 ⊕ ... ⊕ x -> 항상 0이다.
해당 성질을 이용하면 V(L-1), VR을 구할때 1 ~ (L-1), 1 ~ R까지 연산할 필요가 없다.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)
var bw *bufio.Writer = bufio.NewWriter(os.Stdout)

func xor(x int) {

	temp := 0
	for i := 1; i < x; i++ {
		if i%4 == 3 {

			for j := i - 3; j <= i; j++ {
				fmt.Fprintf(bw, "%b %b\n ", temp, j)
				temp ^= (j)

			}
			fmt.Fprintln(bw, i, " -> temp : ", temp)
			temp = 0
		}
	}
}

func main() {

	defer bw.Flush()
	sc.Split(bufio.ScanWords)

	x := scanInt()
	xor(x)

}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
