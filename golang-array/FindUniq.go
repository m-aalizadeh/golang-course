package main

import "fmt"

// time complexiy: O(n)
// extra space complexity: O(1)

func main() {
	myarr := [...]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}
	var tmp = myarr[0]

	for i := 1; i < len(myarr); i++ {
		if tmp != myarr[i] {
			if myarr[i+1] == myarr[i] {
				fmt.Println(tmp)
			} else {
				fmt.Println(myarr[i])
			}
		}
	}

}
