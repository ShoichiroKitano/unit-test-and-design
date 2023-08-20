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

	tests := []struct {
		subject string
		inputNumber int
		expectedOutput string
	} {
		{
			subject: "整数が3の倍数の場合",
			inputNumber: 6,
			expectedOutput: "Fizz",
		},
		{
			subject:  "整数が5の倍数の場合",
			inputNumber: 10,
			expectedOutput: "Buzz",
		},
		{
			subject:  "整数が15の倍数の場合",
			inputNumber: 30,
			expectedOutput: "FizzBuzz",
		},
		{
			subject:  "整数が3,5,15の倍数ではない場合、入力した整数の文字列表現を出力する",
			inputNumber: 30,
			expectedOutput: "FizzBuzz",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.subject, func(t *testing.T) {
			t.Parallel()

			printer := &PrinterMock{}
			execute(&InputerMock{number: tt.inputNumber}, printer)
			assert.Equal(t, tt.expectedOutput, printer.result)
		})
	}
}
