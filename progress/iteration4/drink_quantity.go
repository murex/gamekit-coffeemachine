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
	"fmt"
	"regexp"
	"testing"

	"github.com/murex/gamekit-coffeemachine/process"
	"github.com/murex/gamekit-coffeemachine/ref"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func manyCoffeesInReportTest() (string, func(t *testing.T, p *process.P)) {
	return "I want to be able to print a report that contains many coffees",
		func(t *testing.T, p *process.P) {
			makeDrinks(p, ref.Coffee, 14)

			report, err := p.SendMessage(process.NewPrintReportMessage())
			require.NoError(t, err)

			assertDrinkQuantityInReport(t, report, ref.Coffee, 14)
		}
}

func manyTeasInReportTest() (string, func(t *testing.T, p *process.P)) {
	return "I want to be able to print a report that contains many teas",
		func(t *testing.T, p *process.P) {
			makeDrinks(p, ref.Tea, 9)

			report, err := p.SendMessage(process.NewPrintReportMessage())
			require.NoError(t, err)

			assertDrinkQuantityInReport(t, report, ref.Tea, 9)
		}
}

func manyChocolatesInReportTest() (string, func(t *testing.T, p *process.P)) {
	return "I want to be able to print a report that contains many chocolates",
		func(t *testing.T, p *process.P) {
			makeDrinks(p, ref.Chocolate, 11)

			report, err := p.SendMessage(process.NewPrintReportMessage())
			require.NoError(t, err)

			assertDrinkQuantityInReport(t, report, ref.Chocolate, 11)
		}
}

func manyOrangeJuicesInReportTest() (string, func(t *testing.T, p *process.P)) {
	return "I want to be able to print a report that contains many orange juices",
		func(t *testing.T, p *process.P) {
			makeDrinks(p, ref.OrangeJuice, 27)

			report, err := p.SendMessage(process.NewPrintReportMessage())
			require.NoError(t, err)

			assertDrinkQuantityInReport(t, report, ref.OrangeJuice, 27)
		}
}

func manyDrinksInReportTest() (string, func(t *testing.T, p *process.P)) {
	return "I want to be able to print a report that contains many different drinks",
		func(t *testing.T, p *process.P) {
			makeDrinks(p, ref.Coffee, 35)
			makeDrinks(p, ref.Tea, 4)
			makeDrinks(p, ref.Chocolate, 10)
			makeDrinks(p, ref.OrangeJuice, 13)

			report, err := p.SendMessage(process.NewPrintReportMessage())
			require.NoError(t, err)

			assertDrinkQuantityInReport(t, report, ref.Coffee, 35)
			assertDrinkQuantityInReport(t, report, ref.Tea, 4)
			assertDrinkQuantityInReport(t, report, ref.Chocolate, 10)
			assertDrinkQuantityInReport(t, report, ref.OrangeJuice, 13)
		}
}

func assertDrinkQuantityInReport(t *testing.T, report string, drink ref.Drink, expectedCount int) {
	pattern := regexp.MustCompile(fmt.Sprintf("(?i)%s.*\\b%d\\b", drink.ReportKeyword, expectedCount))
	assert.Regexp(t, pattern, report)
}
