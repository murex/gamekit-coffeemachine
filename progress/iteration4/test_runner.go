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
	"github.com/murex/coffee-machine/progress-runner/progress/iteration"
	"github.com/murex/coffee-machine/progress-runner/ref"
)

// New returns the test runner for this iteration
func New() iteration.TestRunner {
	return iteration.New(4,
		totalAmountInReportTest,
		doesNotCountExtraChangeInReportTest,
		manyCoffeesInReportTest,
		manyTeasInReportTest,
		manyChocolatesInReportTest,
		manyOrangeJuicesInReportTest,
		manyDrinksInReportTest,
	)
}

func makeDrink(p *process.P, drink ref.Drink) {
	makeDrinkWithAmount(p, drink, drink.Price)
}

func makeDrinkWithAmount(p *process.P, drink ref.Drink, amount float64) {
	_, _ = p.SendMessage(process.NewMakeDrinkMessage(drink.Name, 1, amount, false))
}

func makeDrinks(p *process.P, drink ref.Drink, count int) {
	for i := 0; i < count; i++ {
		makeDrink(p, drink)
	}
}
