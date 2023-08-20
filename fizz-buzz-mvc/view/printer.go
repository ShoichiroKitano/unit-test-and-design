package view

import (
	"fmt"
)

type Printer interface {
	Println(string)
}

type StandardOut struct {}

func (stdout *StandardOut) Println(str string) {
	fmt.Println(str)
}
