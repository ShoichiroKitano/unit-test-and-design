package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Parallel()

	tests := []struct {
		subject  string
		number   int
		expected string
	}{
		{
			subject:  "整数が3の倍数の場合",
			number:   6,
			expected: "Fizz",
		},
		{
			subject:  "整数が5の倍数の場合",
			number:   10,
			expected: "Buzz",
		},
		{
			subject:  "整数が15の倍数の場合",
			number:   30,
			expected: "FizzBuzz",
		},
		{
			subject:  "整数が3,5,15の倍数ではない場合、入力した整数の文字列表現を返す",
			number:   1,
			expected: "1",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.subject, func(t *testing.T) {
			assert.Equal(t, tt.expected, FizzBuzz(tt.number))
		})
	}
}
