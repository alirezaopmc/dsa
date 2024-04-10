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

type DigitArray [6]int

func test_case() {
	var s string
	read(&s)
	n := len(s)

	charIndex := func (c uint8) int {
		return int(c - 'A');
	}

	cnt := make([]DigitArray, n)
	for i := range cnt {
		ind := charIndex(s[i])
		cnt[i][ind] = 1
		if i > 0 {
			for j := range cnt[i] {
				cnt[i][j] += cnt[i-1][j]
			}
		}
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
New numbers must be greater than the last. If we saw a smaller value we must
	change our state if needed. If the value is ok with cutting we turn on the
	cut state and we consider that afterward.
*/