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

	var str, replaceStr string
	var count, a, b int

	fmt.Fscan(in, &str)
	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {
		fmt.Fscan(in, &a, &b, &replaceStr)
		str = str[:a-1] + replaceStr + str[b:]
	}

	fmt.Fprintln(out, str)

}
