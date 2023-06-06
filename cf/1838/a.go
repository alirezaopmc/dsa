package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func test_case() {
	var n int
	read(&n)
	
	arr := make([]int, n)
	for i := range arr {
		read(&arr[i])
	}

	sort.Ints(arr)

	if arr[0] < 0 {
		print(arr[0])
	} else {
		print(arr[n-1])
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