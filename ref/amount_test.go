/*
Copyright (c) 2025 Murex

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package ref

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_euros_and_cents_parsing(t *testing.T) {
	tests := []struct {
		amount float64
		euros  int
		cents  int
	}{
		{0.0, 0, 0},
		{0.1, 0, 10},
		{2.0, 2, 0},
		{3.2, 3, 20},
		{4.02, 4, 2},
		{5.002, 5, 0},
		{45.67, 45, 67},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%.02f", test.amount), func(t *testing.T) {
			euros, cents := eurosAndCents(test.amount)
			assert.Equal(t, test.euros, euros)
			assert.Equal(t, test.cents, cents)
		})
	}
}

func Test_amount_regexp(t *testing.T) {
	tests := []struct {
		amount float64
		regexp string
	}{
		{0.0, "0"},
		{0.1, "0[\\.,]1"},
		{2.0, "2"},
		{3.2, "3[\\.,]2"},
		{4.02, "4[\\.,]02"},
		{5.002, "5"},
		{45.67, "45[\\.,]67"},
		{0.123, "0[\\.,]12"},
		{0.987, "0[\\.,]99"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%.02f", test.amount), func(t *testing.T) {
			assert.Equal(t, test.regexp, AmountRegexp(test.amount))
		})
	}
}
