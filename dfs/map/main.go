package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(a []string, y, x int, component int) {
	a[y] = a[y][:x] + string(rune('A'+component-1)) + a[y][x+1:]

	dy := []int{-1, 0, 1, 0}
	dx := []int{0, 1, 0, -1}

	for d := 0; d < len(dy); d++ {
		ty := y + dy[d]
		tx := x + dx[d]

		if 0 <= ty && ty < len(a) &&
			0 <= tx && tx < len(a[ty]) &&
			a[ty][tx] == '#' {
			dfs(a, ty, tx, component)
		}
	}

}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w int
	fmt.Fscan(in, &h, &w)

	var a []string
	a = make([]string, h)

	for i, _ := range a {
		fmt.Fscan(in, &a[i])
	}

	var visited [][]int
	visited = make([][]int, h)
	for i, _ := range visited {
		visited[i] = make([]int, w)
	}

	var compCount = 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if a[y][x] == '#' && visited[y][x] == 0 {
				compCount++
				dfs(a, y, x, compCount)
			}
		}
	}

	for _, v := range a {
		fmt.Fprintln(out, v)
	}

}
