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

package iteration1

import (
	"github.com/murex/gamekit-coffeemachine/process"
	"github.com/murex/gamekit-coffeemachine/progress/iteration"
	"github.com/murex/gamekit-coffeemachine/ref"
)

// New returns the test runner for this iteration
func New() iteration.TestRunner {
	return iteration.New(1,
		drinkInstructionsTest,
		drinkWithNoSugarTest,
		drinkWithSugarTest,
		noMoreThan2SugarsTest,
		sugarComesWithAStickTest,
		noSugarNoStickTest,
	)
}

const (
	noExtraHot = false // Extra hot option is not supported in this iteration
)

var drinks = []ref.Drink{ref.Tea, ref.Coffee, ref.Chocolate}

func runBuildDrinkMakerCommand(p *process.P, drink ref.Drink, sugars int) (string, error) {
	// Always sending requests with drink.Price as payment
	// allows to prevent payment-related issues in this iteration
	return p.SendMessage(process.NewMakeDrinkMessage(drink.Name, sugars, drink.Price, noExtraHot))
}
