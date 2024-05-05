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
	Id           int
	CardsCount   int
	CardsPresent int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countFriends, countCards int
	fmt.Fscan(in, &countFriends, &countCards)

	friendsCardsArr := make([]Ren, 0)

	for i := 0; i < countFriends; i++ {
		elem := Ren{
			Id:           i,
			CardsCount:   0,
			CardsPresent: 0,
		}
		fmt.Fscan(in, &elem.CardsCount)
		friendsCardsArr = append(friendsCardsArr, elem)
	}

	sort.Slice(friendsCardsArr, func(i, j int) (less bool) {
		return friendsCardsArr[i].CardsCount < friendsCardsArr[j].CardsCount
	})

	for i := countFriends - 1; i >= 0; i-- {
		if friendsCardsArr[i].CardsCount >= countCards {
			fmt.Fprintln(out, -1)
			return
		}
		friendsCardsArr[i].CardsPresent = countCards
		countCards--
	}

	sort.Slice(friendsCardsArr, func(i, j int) (less bool) {
		return friendsCardsArr[i].Id < friendsCardsArr[j].Id
	})

	var resArr []string
	resArr = make([]string, len(friendsCardsArr))

	for k := 0; k < len(friendsCardsArr); k++ {
		resArr[k] = strconv.Itoa(friendsCardsArr[k].CardsPresent)
	}

	fmt.Fprintln(out, strings.Join(resArr, " "))

}
