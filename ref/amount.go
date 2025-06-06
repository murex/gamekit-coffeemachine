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
	"math"
)

func eurosAndCents(amount float64) (euros int, cents int) {
	e := math.Floor(amount)
	c := math.Round((amount - e) * 100)
	euros, cents = int(e), int(c)
	return euros, cents
}

// AmountRegexp returns a locale agnostic regular expression that matches the
// given amount will match both '.' and ',' as decimal separator
func AmountRegexp(amount float64) string {
	euros, cents := eurosAndCents(amount)
	switch {
	case cents == 0:
		return fmt.Sprintf("%d", euros)
	case cents%10 == 0:
		return fmt.Sprintf("%d[\\.,]%d", euros, cents/10)
	default:
		return fmt.Sprintf("%d[\\.,]%02d", euros, cents)
	}
}
