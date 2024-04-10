package main

import (
	"fmt"
	"math"
)

type SparseTable struct {
	arr		[]int
	table	[][]int
	n		int
	b		int
}

func (s *SparseTable) initialize(f func(int,int)int) {
	s.n = len(s.arr)
	s.b = int(math.Log2(float64(s.n))) + 1
	s.table = make([][]int, s.b)
	for i := 0; i < s.b; i++ {
		s.table[i] = make([]int, s.n)
	}
	for j := 0; j < s.n; j++ {
		s.table[0][j] = s.arr[j]
	}
	for i := 1; i < s.b; i++ {
		for j := 0; j + (1 << (i)) - 1 < s.n; j++ {
			s.table[i][j] = f(s.table[i-1][j], s.table[i-1][j+(1 << (i-1))])
		}
	}
}

func (s *SparseTable) String() string {
	result := ""
	for i := 0; i < s.b; i++ {
		for j := 0; j < s.n; j++ {
			result += fmt.Sprintf("%d", s.table[i][j])
			if j == s.n-1 {
				result += "\n"
			} else {
				result += "\t"
			}
		}
	}
	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Sum(a, b int) int {
	return a + b
}

func main() {
	arr := []int{1, 3, -1, 3, 2}
	s := &SparseTable{arr: arr}
	s.initialize(Sum)
	fmt.Println(s)
}

