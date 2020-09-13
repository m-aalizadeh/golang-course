package main

import (
	"fmt"
	"strconv"
)

func main() {

	myarr := []int{4, 5, 1, 8}
	target := 6
	var myMap = make(map[string]int)

	for i := 0; i < len(myarr); i++ {
		value, ok := myMap[strconv.Itoa(target-myarr[i])]
		if !ok {
			myMap[strconv.Itoa(myarr[i])] = myarr[i]
		} else {
			fmt.Printf("%d, %d", IndexOf(value, myarr), myarr[i])
		}
	}
	fmt.Println()
	fmt.Println(myMap)
}
