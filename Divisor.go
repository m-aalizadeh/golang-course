package main

import "fmt"

func main() {

	var divisor, bound int
	fmt.Print("Please Enter divisor: ")
	fmt.Scan(&divisor)
	fmt.Println()
	fmt.Print("Please Enter bound: ")
	fmt.Scan(&bound)

	fmt.Println("Greatest Divisor is: ", maxDivisor(divisor,bound))
}

func maxDivisor(divisor, bound int) int  {

	var max = divisor
	for num := divisor + 1; num <= bound; num++ {
		if num % divisor == 0 && num > max {
			max = num
		}
	}
	return max
}


