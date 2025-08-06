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

package iteration3

import (
	"regexp"
	"testing"

	"github.com/murex/gamekit-coffeemachine/process"
	"github.com/murex/gamekit-coffeemachine/progress/iteration"
	"github.com/murex/gamekit-coffeemachine/ref"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// New returns the test runner for this iteration
func New() iteration.TestRunner {
	return iteration.New(3,
		orangeJuiceTest,
		noOrangeJuiceIfNotEnoughMoneyTest,
		noSugarsInOrangeJuiceTest,
		extraHotCoffeeTest,
		extraHotTeaTest,
		extraHotChocolateTest,
		noExtraHotOrangeJuiceTest,
	)
}

// nolint: revive
func runBuildDrinkMakerCommand(p *process.P, drink ref.Drink, sugars int, payment float64, extraHot bool) (string, error) {
	return p.SendMessage(process.NewMakeDrinkMessage(drink.Name, sugars, payment, extraHot))
}

func assertExtraHotCommand(t *testing.T, p *process.P, drink ref.Drink) {
	t.Helper()
	cmd, err := runBuildDrinkMakerCommand(p, drink, 0, drink.Price, true)
	require.NoError(t, err)
	commandCode := drink.CommandCode + ref.ExtraHotCommandFlag
	pattern := regexp.MustCompile("^" + commandCode + ":.*$")
	assert.Regexpf(t, pattern, cmd,
		"Drink maker command for extra hot %s should start with '%s:'",
		drink.Name, commandCode)
}

func assertNoExtraHotCommand(t *testing.T, p *process.P, drink ref.Drink) {
	t.Helper()
	cmd, err := runBuildDrinkMakerCommand(p, drink, 0, drink.Price, true)
	require.NoError(t, err)
	pattern := regexp.MustCompile("^" + drink.CommandCode + ":.*$")
	assert.Regexpf(t, pattern, cmd,
		"Drink maker command for %s should start with '%s:'",
		drink.Name, drink.CommandCode)
}
