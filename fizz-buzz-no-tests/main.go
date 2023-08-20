package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	fmt.Scan(&number)
	if number%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if number%3 == 0 {
		fmt.Println("Fizz")
	} else if number%5 == 0 {
		fmt.Println("Buzz")
	} else {
		s := strconv.Itoa(number)
		fmt.Println(s)
	}
}
