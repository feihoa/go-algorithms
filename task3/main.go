package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	var t, id int
	m := make(map[int]int)

	messageCount := 0
	messageAllCount := 0

	fmt.Fscan(in, &n, &q)

	for i := 0; i < q; i++ {
		fmt.Fscan(in, &t, &id)
		if t == 1 {
			messageCount++
			if id != 0 {
				m[id] = messageCount
			} else {
				for key := range m {
					m[key] = messageCount
				}
				messageAllCount = messageCount
			}
		} else {
			var res int
			if m[id] != 0 {
				res = m[id]
			} else {
				res = messageAllCount
			}
			fmt.Fprintln(out, res)
		}

	}

}
