package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type InputerMock struct {
	number int
}

func (mock *InputerMock) InputNumber() int {
	return mock.number
}

type PrinterMock struct {
	result string
}

func (mock *PrinterMock) Println(result string) {
	mock.result = result
}

func TestExecute(t *testing.T) {
	t.Parallel()

	t.Run("ユーザが入力した整数に応じた結果を表示できること", func(t *testing.T) {
		printer := &PrinterMock{}
		execute(&InputerMock{number: 6}, printer)
		assert.Equal(t, "Fizz", printer.result)
	})
}
