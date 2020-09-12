package main

import (
	"fmt"
	"strings"
)

func main() {

	var myMap = make(map[string]int)
	var expression string
	fmt.Println("please Enter your expresssion:")
	fmt.Scan(&expression)

	arr := strings.Split(strings.ToLower(expression), "")

	for i := 0; i < len(arr); i++ {
		var count = myMap[arr[i]]
		//		fmt.Println(myMap[arr[i]])
		if count == 0 {
			myMap[arr[i]] = 1
		} else {
			myMap[arr[i]] = count + 1
		}

	}

	for character, counter := range myMap {
		if counter > 1 {
			fmt.Printf("%s: %d", character, counter)
			fmt.Println()
		}

	}
	//	fmt.Println(myMap)
}
