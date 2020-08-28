package main

import "fmt"

func main() {
	fmt.Println("Hello World");

	//var i, j int
	var number int
	fmt.Println("Please Enter your number: ")
	fmt.Scan(&number)

	if number == 1 {
		println("Nothing")
	}else {
		for i := 1 ; i <= number ; i++ {
			for j := number ; j >= i ; j-- {
				fmt.Print(j)
			}
			fmt.Println();
		}
	}

}
