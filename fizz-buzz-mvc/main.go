package main

import (
	"fizz-buzz/view"
	"fizz-buzz/model"
)

func main() {
	execute(&view.StandardIn{}, &view.StandardOut{})
}

func execute(in view.Inputer, out view.Printer) {
	number := in.InputNumber()
	result := model.FizzBuzz(number)
	out.Println(result)
}
