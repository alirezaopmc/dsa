package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
 
func read(a ...interface{}) {
	fmt.Fscan(in, a...)
}
func print(a ...interface{}) {
	fmt.Fprint(out, a...)
}

func test_case() {
	var n, k int
	read(&n, &k)
	a := make([]int, n)
	for i := range a {
		read(&a[i])
	}
	if k == n {
		print(0, "\n")
		return
	}
	b := make([]int, n-1)
	s := 0
	for i := range a {
		if i > 0 {
			b[i-1]=Abs(a[i]-a[i-1])
			s += b[i-1]
		}
	}
	sort.Ints(b)
	for i := 0; i < k-1; i++ {
		s -= b[n-i-1-1];
	}
	print(s, "\n")
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
Subtract k-1 largest values.
*/