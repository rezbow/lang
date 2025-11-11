package lang_test

import (
	"testing"

	"github.com/rezbow/lang"
)

func Test_arithmetic_expression_addition(t *testing.T) {
	tests := []struct {
		name string
		exp  string
		want int
	}{
		{name: "addition two one-digit number", exp: "2+5", want: 7},
		{name: "addition two over-one-digit number", exp: "12+5", want: 17},
		{name: "addition multiple number", exp: "12+5+1+2", want: 20},
		{name: "subtraction", exp: "10-5", want: 5},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := lang.Run(test.exp)
			if got != test.want {
				t.Errorf("got %d, wanted %d", got, test.want)
			}
		})
	}
}
