package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

type Pair struct {
	first int
	second int
}

type Graph struct {
	adj		[][]Pair
}

func (g *Graph) init(n int) {
	g.adj = make([][]Pair, n)
}

func (g *Graph) addEdge(u, v, i int) {
	g.adj[u] = append(g.adj[u], Pair{v,i})
	g.adj[v] = append(g.adj[v], Pair{u,i})
}

func test_case() {
	var n, k, d int
	read(&n, &k, &d)
	p := make([]int, k)
	for i := range p {
		read(&p[i])
		p[i]--
	}

	g := &Graph{}
	g.init(n)

	for i := 0; i < n-1; i++ {
		var u, v int
		read(&u, &v)
		u--
		v--
		g.addEdge(u, v, i)
	}

	q := make([]int, 0)
	visited := make([]bool, n)
	for i := range p {
		q = append(q, p[i])
		visited[p[i]] = true
	}
	
	should_hold := make(map[int]bool)

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for _, neighbor := range g.adj[node] {
			first := neighbor.first
			second := neighbor.second

			if ! visited[first] {
				visited[first] = true
				should_hold[second] = true
				q = append(q, first)
			}
		}
	}

	roads_num := len(reflect.ValueOf(should_hold).MapKeys())


	print(n-roads_num-1, "\n")
	cnt := 0
	for i := 0; i < n; i++  {
		if cnt == k-1 {
			break
		}
		if ! should_hold[i] {
			print(i+1, " ")
			cnt++
		}
	}
	print("\n")
}

func main() {
	defer out.Flush()
	var t int = 1;
	// fmt.Scan(&t)

	for t > 0 {
		test_case()
		t--
	}
}

/*

*/