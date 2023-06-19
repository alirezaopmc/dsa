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
	fmt.Fprintln(out, a...)
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

func test_case() {
	var n int
	read(&n)
	p1 := -1
	p2 := -1
	pn := -1

	var x int
	for i := 0; i < n; i++ {
		read(&x)
		if x == 1 {
			p1 = i+1
		}
		if x == 2 {
			p2 = i+1
		}
		if x == n {
			pn = i+1
		}
	}
	
	if p1 < p2 {
		if p2 < pn {
			print(p2, pn)
		} else {
			if p1 < pn {
				print(1, 1)
			} else {
				print(p1, pn)
			}
		}
	} else {
		if p1 < pn {
			print(p1, pn)
		} else {
			if p2 < pn {
				print(2, 2)
			} else {
				print(p2, pn)
			}
		}
	}

}

func main() {
	defer out.Flush()
	var t int
	read(&t)

	for t > 0 {
		test_case()
		t--
	}
}