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
	var n int
	read(&n)
	a := make([]int, n)
	s := math.MaxInt
	cnt := 0
	for i := range a {
		read(&a[i])
		s &= a[i]
		if s == 0 {
			s = math.MaxInt
			cnt++
		}
	}
	print(Max(cnt, 1), "\n")
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
AND always make the result smaller than SUM. So if the overall AND
is not zero, the answer is 1. Otherwise is the number of 0s in
prefix AND array.
*/