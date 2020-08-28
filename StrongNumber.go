package main

import "fmt"

func main() {
	var number int
	fmt.Print("Please Enter your number: ")
	fmt.Scan(&number)
	var temp = number
	var sum = 0
	for temp > 0 {
		rem := temp % 10
		temp = temp / 10
		sum += factorial(rem)
		//fmt.Println(rem,number,sum)
	}
	if sum == number {
		fmt.Printf("%d is a strong number.", number)
	}else {
		fmt.Printf("%d is not a strong number.", number)
	}
	//fmt.Println(sum)
}

func factorial(number int) int {
	if number == 1 {
		return 1
	}else {
		return factorial(number-1) * number
	}
}
