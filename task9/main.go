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

	var count int
	fmt.Fscanln(in, &count)

	for i := 0; i < count; i++ {

		var meCount int
		fmt.Fscanln(in, &meCount)
		var line = make([]byte, 0)
		b, _ := in.ReadBytes('\n')
		line = b

		var xCount, yCount, zCount int

		for j := 0; j < meCount; j++ {
			if line[j] == 'X' {
				xCount++
			}

			if line[j] == 'Y' {
				yCount++
			}

			if line[j] == 'Z' {
				zCount++
			}
		}

		yzCount := (meCount - xCount*2) / 2
		xzCount := zCount - yzCount

		yPos := make([]int, 0)
		xPos := make([]int, 0)

		for j := 0; j < meCount; j++ {
			if line[j] == 'Y' {
				yPos = append(yPos, j)
				continue
			} else if line[j] == 'X' {
				xPos = append(xPos, j)
			} else {
				if len(yPos) > 0 && yPos[0] < j && yzCount > 0 {
					yzCount--
					line[yPos[0]] = 'K'
					line[j] = 'K'
					yPos = yPos[1:]
				} else if len(xPos) > 0 && xPos[len(xPos)-1] < j && xzCount > 0 {
					xzCount--
					line[xPos[(len(xPos)-1)]] = 'K'
					line[j] = 'K'
					xPos = xPos[:len(xPos)-1]
				}
			}
		}

		if yzCount != 0 || xzCount != 0 {
			fmt.Fprintln(out, "No")
		} else {

			stack := 0

			for j := 0; j < meCount; j++ {
				if line[j] == 'K' {
					continue
				}

				if line[j] == 'X' {
					stack++
				} else {
					stack--
				}

				if stack < 0 {
					break
				}
			}

			if stack == 0 {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}
	}
}
