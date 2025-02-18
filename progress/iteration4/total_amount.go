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

package iteration4

import (
	"github.com/murex/coffee-machine/progress-runner/process"
	"github.com/murex/coffee-machine/progress-runner/ref"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func totalAmountInReportTest() (string, func(t *testing.T, p *process.P)) {
	return "I want to be able to print a report that contains the total amount of money earned so far",
		func(t *testing.T, p *process.P) {
			makeDrinks(p, ref.Tea, 3)
			makeDrinks(p, ref.Coffee, 2)

			assertTotalAmountInReport(t, p, 3*ref.Tea.Price+2*ref.Coffee.Price)
		}
}

func doesNotCountExtraChangeInReportTest() (string, func(t *testing.T, p *process.P)) {
	return "I want to be able to print a report that does not count the extra change given by the customer",
		func(t *testing.T, p *process.P) {
			makeDrinkWithAmount(p, ref.Tea, 100)

			assertTotalAmountInReport(t, p, ref.Tea.Price)
		}
}

func assertTotalAmountInReport(t *testing.T, p *process.P, expectedAmount float64) {
	t.Helper()

	response, err := p.SendMessage(process.NewPrintReportMessage())
	require.NoError(t, err)
	pattern := regexp.MustCompile(ref.AmountRegexp(expectedAmount))
	assert.Regexp(t, pattern, response)
}
