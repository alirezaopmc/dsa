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

func IsPrime(n int) bool {
	if n == 2 { return true }
	for i := 2; i * i <= n; i++ {
		if n % i == 0 { return false }
	}
	return true
}

func test_case() {
	var n, m int
	read(&n, &m)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}

	if ! IsPrime(m) {
		k := 1
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				arr[i][j] = k
				k++
			}
		}
	} else if ! IsPrime(n) {
		k := 1
		for j := 0; j < m; j++ {
			for i := 0; i < n; i++ {
				arr[i][j] = k
				k++
			}
		}
	} else {
		k := 1
		s := m * ((n + 1) / 2) + 1
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if i % 2 == 0 {
					arr[i][j] = k
					k++
				} else {
					arr[i][j] = s
					s++
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			print(arr[i][j], " ")
		}
		print("\n")
	}
	print("\n")
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

/*
Let a(n) = [1, 2, .., n]
If m is not prime we can place a(i) at the i-th row.
If n is not prime we can do the same vertically.
If both n and m are prime, for i from 0 to ceil(n/2) we put a(i) at
	even indices and fill the others respectively.
*/