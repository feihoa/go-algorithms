package main

import (
	"bufio"
	"fmt"
	"os"
)

// список смежности
func dfs(graph *[][]int, v int, visited *[]int) {
	(*visited)[v] = 1

	for _, to := range (*graph)[v] {
		if (*visited)[to] != 1 {
			dfs(graph, to, visited)
		}
	}
}

// матрица
// func dfs(graph *[][]int, v int, visited *[]int) {
// 	(*visited)[v] = 1

// 	for i, to := range *graph {
// 		if (*graph)[v][i] != 0 && (*visited)[to] != 1 {
// 			dfs(graph, to, visited)
// 		}
// 	}

// }

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var vertexCount, edgeCount int
	fmt.Fscan(in, &vertexCount, &edgeCount)

	var graph [][]int
	graph = make([][]int, vertexCount)

	fmt.Println(vertexCount, edgeCount)

	for i := 0; i < edgeCount; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	var start, finish int
	fmt.Fscan(in, &start, &finish)
	start--
	finish--

	visited := make([]int, vertexCount)
	dfs(&graph, start, &visited)

	fmt.Println(visited, start, finish)

	if visited[finish] == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
