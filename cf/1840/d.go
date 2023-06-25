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

func check(a []int, k int) bool {
	/*
	Greedy approach here. It can be optimized by binary search.
		Looking for the next value greater than last checkpoint.
	*/
	checkpoint := a[0] + 2 * k
	cnt := 1
	for i := range a {
		if a[i] <= checkpoint {
			continue
		} else {
			cnt++
			checkpoint = a[i] + 2 * k
		}
	}
	return cnt <= 3
}

func test_case() {
	var n int
	read(&n)
	a := make([]int, n)
	for i := range a {
		read(&a[i])
	}
	sort.Ints(a)

	ans := a[n-1]
	left, right := 0, a[n-1]
	for left <= right {
		mid := left + (right - left) / 2
		if check(a, mid) {
			ans = mid
			right = mid-1
		} else {
			left = mid+1
		}
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
Binary search on k.
Check whether the solution is k or not.
Check for each k:
At first sort the array a. Select the first pattern a[0] + k.
So then all toys from a[0] to the a[0] + 2k can be done in <=k.
Select the next pattern as a[x] + k. Where x is the first toy after
	a[0] + 2k. So on ...
If it was only three steps return true. O.w. return false.
*/