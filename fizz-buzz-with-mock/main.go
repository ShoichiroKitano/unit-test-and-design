package main

import (
	"strconv"
	"fizz-buzz/view"
)

func main() {
	execute(&view.StandardIn{}, &view.StandardOut{})
}

func execute(in view.Inputer, out view.Printer) {
	number := in.InputNumber()
	if number%15 == 0 {
		out.Println("FizzBuzz")
	} else if number%3 == 0 {
		out.Println("Fizz")
	} else if number%5 == 0 {
		out.Println("Buzz")
	} else {
		out.Println(strconv.Itoa(number))
	}
}
