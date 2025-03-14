package main

import "fmt"

type Pair struct {
	X int
	Y int
}

type PairSet map[Pair]bool

func exist(board [][]byte, word string) bool {
	vis := make(PairSet)
	for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[1]); j++ {
					if check(board, i, j, word, vis) {
							return true
					}
			}
	}
	return false
}

func check(board [][]byte, i, j int, word string, vis PairSet) bool {
	return dfs(board, i, j, word, vis)
}

func dfs(
	board [][]byte,
	i, j int,
	word string,
	vis PairSet,
) bool {
	n, m := len(board), len(board[0])
	
	// If the word is searched completely
	if len(word) == 0 {
			return true
	}
	
	// If index is out of range
	if i < 0 || i >= n || j < 0 || j >= m {
			return false
	}
	
	// If the cell is already visited
	if vis[Pair{i, j}] {
			return false
	}
	
	// If the character is not the same
	if board[i][j] != word[0] {
			return false
	}
	
	// Mark current cell as visited
	vis[Pair{i, j}] = true
	
	// Backtrack for neighbors
	result := dfs(board, i+1, j, word[1:], vis) || 
						dfs(board, i-1, j, word[1:], vis) || 
						dfs(board, i, j+1, word[1:], vis) || 
						dfs(board, i, j-1, word[1:], vis)
	
	// Backtrack by marking current cell as unvisited
	vis[Pair{i, j}] = false
	
	return result
}

func main() {
	tests := []struct {
		board [][]byte
		word string
		expected bool
	}{
		{
			board: [][]byte{
				{'A','B','C','E'},
				{'S','F','C','S'},
				{'A','D','E','E'},
			},
			word: "ABCCED",
			expected: true,
		},
		{
			board: [][]byte{
				{'A','B','C','E'},
				{'S','F','C','S'},
				{'A','D','E','E'},
			},
			word: "SEE",
			expected: true,
		},
		{
			board: [][]byte{
				{'A','B','C','E'},
				{'S','F','C','S'},
				{'A','D','E','E'},
			},
			word: "ABCB",
			expected: false,
		},
	}

	for i, t := range tests {
		actual := exist(t.board, t.word)
		if actual != t.expected {
			fmt.Printf("Error in test %d: expected %v, actual %v\n", i+1, t.expected, actual)
		}
	}
}