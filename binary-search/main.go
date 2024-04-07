package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	count, err := strconv.Atoi(line)
	handleError(err)

	for i := 0; i < count; i++ {

		scanner.Scan()
		numberToFindStr := scanner.Text()
		numberToFind, err := strconv.Atoi(numberToFindStr)
		handleError(err)

		scanner.Scan()
		str := scanner.Text()
		arrayStr := strings.Fields(str)

		array := make([]int, 0)

		for _, value := range arrayStr {
			elem, err := strconv.Atoi(value)
			handleError(err)
			array = append(array, elem)
		}

		res := hasElem(array, numberToFind)

		fmt.Printf("Has %s in %s? ", numberToFindStr, str)
		fmt.Println(res)
	}
}

func hasElem(arr []int, searchNum int) bool {
	if len(arr) == 0 {
		return false
	}

	ind := len(arr) / 2
	center := arr[ind]
	if center == searchNum {
		return true
	} else if center < searchNum {
		return hasElem(arr[ind+1:], searchNum)
	} else {
		return hasElem(arr[:ind-1], searchNum)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
