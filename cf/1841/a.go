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
	
	if n >= 5 {
		print("Alice\n")
	} else {
		print("Bob\n")
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
For n >= 5 Alice can do {n-2,2} and wins. For smaller n Bob wins.
*/