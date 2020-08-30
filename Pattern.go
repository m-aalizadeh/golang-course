package main

import "fmt"

func main() {

	var number int
	fmt.Println("Please Enter your number: ")
	fmt.Scan(&number)

	if number < 1 {
		println("*")
	}else {
		for i := 1 ; i <= number ; i++ {
			for j := number ; j >= i ; j-- {
				fmt.Print(j)
			}
			fmt.Println();
		}
	}

}
