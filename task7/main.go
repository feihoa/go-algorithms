package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Patient struct {
	Id    int
	Place int
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
		var win, patientsCount int
		fmt.Fscan(in, &win, &patientsCount)
		patients := make([]Patient, 0)

		for k := 0; k < patientsCount; k++ {
			elem := Patient{
				Id:    k + 1,
				Place: 0,
			}
			fmt.Fscan(in, &elem.Place)
			patients = append(patients, elem)
		}

		sort.Slice(patients, func(i, j int) (less bool) {
			return patients[i].Place < patients[j].Place
		})

		allWinsUsed := make(map[int]int, patientsCount)

		x := false
		var result = make([]byte, patientsCount)

		for _, v := range patients {
			place := v.Place
			if place-1 > 0 && allWinsUsed[place-1] == 0 {
				allWinsUsed[place-1] = v.Id
				result[v.Id-1] = '-'
			} else if allWinsUsed[place] == 0 {
				allWinsUsed[place] = v.Id
				result[v.Id-1] = '0'
			} else if place+1 <= win && allWinsUsed[place+1] == 0 {
				allWinsUsed[place+1] = v.Id
				result[v.Id-1] = '+'
			} else {
				x = true
				break
			}

		}

		if x {
			fmt.Fprintln(out, "x")
			continue
		} else {
			fmt.Fprintln(out, string(result))
		}

	}

}
