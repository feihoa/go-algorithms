package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Ren struct {
	Id     int
	Result int
	Place  int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {
		var renCount int
		var arr []Ren
		var res int

		fmt.Fscan(in, &renCount)

		arr = make([]Ren, 0)

		for k := 0; k < renCount; k++ {
			fmt.Fscan(in, &res)
			elem := Ren{
				Id:     k,
				Result: res,
				Place:  0,
			}
			arr = append(arr, elem)
		}

		sort.Slice(arr, func(i, j int) (less bool) {
			return arr[i].Result < arr[j].Result
		})

		var ind int

		for k := 0; k < len(arr); k++ {
			if !(k > 0 && (arr[k].Result == arr[k-1].Result || arr[k].Result == arr[k-1].Result+1)) {
				ind = k
			}

			arr[k].Place = ind
		}

		var resArr []string
		resArr = make([]string, len(arr))

		sort.Slice(arr, func(i, j int) (less bool) {
			return arr[i].Id < arr[j].Id
		})

		for k := 0; k < len(arr); k++ {
			resArr[k] = strconv.Itoa(arr[k].Place + 1)
		}

		fmt.Fprintln(out, strings.Join(resArr, " "))
	}
}
