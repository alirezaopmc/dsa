package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)
 
var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)
 
func read(a ...interface{}) {
	fmt.Fscan(in, a...)
}
func print(a ...interface{}) {
	fmt.Fprint(out, a...)
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

func test_case() {
	var n, k int
	read(&n, &k)
	
	if math.Log2(float64(n+1)) < float64(k) {
		print(n+1, "\n")
	} else {
		print(Pow(2, k), "\n")
	}
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
If we had unlimited items in the cafe then we could buy in n+1 ways.
Because there are n+1 binary numbers not greater than n, where each
	of them represent one way. 0,1,...,n.
Otherwise if the items number limits us, it caps the ability of us to
	buying only 2**k items. So we just need to check if n+1>2**k.
! Note that 2**k is super heavy for large k. That's why we use logarithms.
*/