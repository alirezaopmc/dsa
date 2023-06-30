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
	
	var x int
	b := make([]bool, n)
	var cut bool
	var last int
	var first int
	for i := range b {
		read(&x)
		b[i] = true
		if i == 0 {
			first = x
			last = x
		} else if !cut && x >= last {
			last = x
		} else if !cut && x <= first {
			cut = true
			last = x
		} else if cut && x <= first && x >= last {
			last = x
		} else {
			b[i] = false
		}
	}

	for i := range b {
		if b[i] {
			print("1")
		} else {
			print("0")
		}
	}
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
New numbers must be greater than the last. If we saw a smaller value we must
	change our state if needed. If the value is ok with cutting we turn on the
	cut state and we consider that afterward.
*/