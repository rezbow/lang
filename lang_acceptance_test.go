package lang_test

import (
	"testing"

	"github.com/rezbow/lang"
)

func Test_arithmetic_expression(t *testing.T) {
	tests := []struct {
		exp  string
		want int
	}{
		{exp: "102", want: 102},
		{exp: "2+5", want: 7},
		{exp: "12+5", want: 17},
		{exp: "12+5+1+2", want: 20},
		{exp: "10-5", want: 5},
		{exp: "5+2-1+4", want: 10},
		{exp: "2*10*2", want: 40},
		{exp: "2+10*4", want: 42},
		{exp: "2*10+4", want: 24},
		{exp: "2*(10+4)", want: 28},
		{exp: "2+10*(10-5)", want: 52},
		{exp: "2+10*(10*(1+2*(2+2)))", want: 902},
		{exp: "-2*2*-4", want: 16},
		//
	}
	for _, test := range tests {
		t.Run(test.exp, func(t *testing.T) {
			got, _ := lang.Run(test.exp)
			if got != test.want {
				t.Errorf("got %d, wanted %d", got, test.want)
			}
		})
	}
}
