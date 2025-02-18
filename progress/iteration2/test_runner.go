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

package iteration2

import (
	"github.com/murex/coffee-machine/progress-runner/process"
	"github.com/murex/coffee-machine/progress-runner/progress/iteration"
	"github.com/murex/coffee-machine/progress-runner/ref"
	"testing"
)

// New returns the test runner for this iteration
func New() iteration.TestRunner {
	return iteration.New(2,
		noMoneyNoDrinkTest,
		exactAmountForCoffeeTest,
		exactAmountForTeaTest,
		exactAmountForChocolateTest,
		noCoffeeIfNotEnoughMoneyTest,
		noTeaIfNotEnoughMoneyTest,
		noChocolateIfNotEnoughMoneyTest,
		moreMoneyThanNeededTest,
	)
}

const (
	noSugar    = 0     // number of sugars has no impact on drink price
	noExtraHot = false // Extra hot option is not supported in this iteration
)

func runBuildDrinkMakerCommand(p *process.P, drink ref.Drink, payment float64) (string, error) {
	return p.SendMessage(process.NewMakeDrinkMessage(drink.Name, noSugar, payment, noExtraHot))
}

func assertDrinkIsServed(t *testing.T, p *process.P, drink ref.Drink, money float64) {
	t.Helper()
	response, err := runBuildDrinkMakerCommand(p, drink, money)
	iteration.AssertDrinkIsServed(t, drink, response, err)
}

func assertDrinkIsNotServed(t *testing.T, p *process.P, drink ref.Drink, money float64) {
	t.Helper()
	cmd, err := runBuildDrinkMakerCommand(p, drink, money)
	iteration.AssertDrinkIsNotServed(t, drink, cmd, err)
}
