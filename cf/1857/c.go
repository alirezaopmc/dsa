package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)
 
var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func Abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}


func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func Min(a, b int) int {
	return a + b - Max(a, b)
}

func Pow(a, n int) int {
	r := 1
	for n > 0 {
		if n & 1 == 1 {
			r *= a
		}
		a = a * a
		n >>= 1
	}
	return r
}

func read(a ...interface{}) {
	fmt.Fscan(in, a...)
}
func print(a ...interface{}) {
	fmt.Fprint(out, a...)
}

type Pair struct {
	A, B int
}

type PairHeap []Pair

func (ph PairHeap) Len() int {
	return len(ph)
}

func (ph PairHeap) Less(i, j int) bool {
	return ph[i].A < ph[j].A
}

func (ph PairHeap) Swap(i, j int) {
	ph[i], ph[j] = ph[j], ph[i]
}

func (ph *PairHeap) Push(x interface{}) {
	*ph = append(*ph, x.(Pair))
}

func (ph *PairHeap) Pop() interface{} {
	old := *ph
	n := len(old)
	x := old[n-1]
	*ph = old[0 : n-1]
	return x
}

func test_case() {
	var n int
	read(&n)
	l := n*(n-1)/2
	arr := make([]int, l)
	cnt := map[int]int{}
	var big int
	for i := range arr {
		read(&arr[i])
		cnt[arr[i]]++

		if i == 0 {
			big = arr[i]
		} else {
			big = Max(big, arr[i])
		}
	}
	ph := &PairHeap{}
	for key, value := range cnt {
		heap.Push(ph, Pair{key, value})
	}
	cur := n-1
	for cur > 0 /*ph.Len() > 0*/ {
		pair := heap.Pop(ph).(Pair)
		print(pair.A, " ")
		pair.B -= cur
		cur--
		if pair.B > 0 {
			heap.Push(ph, pair)
		}
	}
	print(big)
	print("\n")
}

func main() {
	defer out.Flush()
	var t int
	fmt.Scan(&t)

	for t > 0 {
		test_case()
		t--
	}
}

/*

*/