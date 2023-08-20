package view

import (
	"fmt"
)

type Inputer interface {
	InputNumber() int
}

type StandardIn struct {
}

func (stdin *StandardIn) InputNumber() int {
	var number int
	fmt.Scan(&number)
	return number
}
