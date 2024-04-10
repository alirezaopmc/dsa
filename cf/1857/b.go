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
	var s string
	read(&s)
	arr := []rune(s)
	

	i := 0
	for i < len(arr) {
		if arr[i] >= '5' {
			break
		}
		i++
	}

	if i != len(s) {
		for j := i; j < len(s); j++ {
			arr[j] = '0'
		}

		for j := i-1; j >= 0; j-- {
			if j == i-1 && arr[j] >= '5' {
				arr[j] = 0
			} else if arr[j] >= '4' {
				arr[j] = '0'
			} else {
				arr[j]++
				break
			}
		}
	}

	if arr[0] == '0' {
		print("1")
	}
	print(string(arr)+"\n")
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

*/