package main

import "fmt"

func main() {

	var number int
	var sum int = 0
	fmt.Println("Please Enter your number: ")
	fmt.Scan(&number)

	if number == 1 {
		fmt.Println("One")
	}else {
		for i := 1 ; i <= number ; i++ {
			sum += i
		}
		fmt.Println("Sum of numbers is = ", sum)
	}
}
