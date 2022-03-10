package main

import "fmt"

func main() {
	var result = SumDivisibleBy(5, 999) + SumDivisibleBy(3, 999) - SumDivisibleBy(15, 999)
	fmt.Println(result)
}

//Return the sum of all the numbers divisibles by -number- bellow a -target- number.

func SumDivisibleBy(number int, target int) int {
	var p = target / number
	return number * (p * (p + 1)) / 2
}
