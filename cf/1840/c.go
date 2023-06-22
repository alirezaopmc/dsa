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
	var n, k, q int
	read(&n, &k, &q)
	var x int
	
	ans, cur := 0, 0
	for i := 0; i < n; i++ {
		read(&x)
		if x <= q {
			cur++
		} else {
			if cur >= k {
				ans += (cur - k + 2) * (cur - k + 1) / 2
			}
			cur = 0
		}
		/* My first wrong attempt
		if x <= q {
			cur++
		} else if cur >= k {
			ans += ...
			cur = 0
		}

		The problem is that cur=0 must not be done only when cur>=k
		It must be done when x > q.
		*/
	}

	if cur >= k {
		ans += (cur - k + 2) * (cur - k + 1) / 2
	}

	print(ans, "\n")
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
Iterate and count the number of consecutive good days.
At the end of each intervals, count the number of possible ways.
*/