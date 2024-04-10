package main

import (
	"bufio"
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

func test_case() {
	var b, c, h int
	read(&b, &c, &h)

	print(Min(2*b-1, 2*(c+h)+1), "\n")
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