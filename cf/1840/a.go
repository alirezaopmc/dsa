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

func test_case() {
	var n int
	read(&n)
	
	var s string
	read(&s)

	var last byte = '!'
	for i := 0; i < n; i++ {
		if last == '!' {
			last = s[i]
		} else if s[i] == last {
			print(string(s[i]))
			last = '!'
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
Iterate the string and match the duplicated saved characters.
	Skip others. Print matched chars.
*/