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
	"github.com/murex/gamekit-coffeemachine/process"
	"github.com/murex/gamekit-coffeemachine/ref"
	"testing"
)

func noMoneyNoDrinkTest() (string, func(t *testing.T, p *process.P)) {
	drink := ref.Tea
	return "the drink maker should not make the drink when no money is given",
		func(t *testing.T, p *process.P) {
			assertDrinkIsNotServed(t, p, drink, 0.00)
		}
}

func exactAmountForCoffeeTest() (string, func(t *testing.T, p *process.P)) {
	drink := ref.Coffee
	return "the drink maker should make " + drink.Name + " if exact amount is given",
		func(t *testing.T, p *process.P) {
			assertDrinkIsServed(t, p, drink, drink.Price)
		}
}

func exactAmountForTeaTest() (string, func(t *testing.T, p *process.P)) {
	drink := ref.Tea
	return "the drink maker should make " + drink.Name + " if exact amount is given",
		func(t *testing.T, p *process.P) {
			assertDrinkIsServed(t, p, drink, drink.Price)
		}
}

func exactAmountForChocolateTest() (string, func(t *testing.T, p *process.P)) {
	drink := ref.Chocolate
	return "the drink maker should make " + drink.Name + " if exact amount is given",
		func(t *testing.T, p *process.P) {
			assertDrinkIsServed(t, p, drink, drink.Price)
		}
}

func noCoffeeIfNotEnoughMoneyTest() (string, func(t *testing.T, p *process.P)) {
	drink := ref.Coffee
	return "the drink maker should not make " + drink.Name + " if not enough money is given",
		func(t *testing.T, p *process.P) {
			assertDrinkIsNotServed(t, p, drink, drink.Price-0.01)
		}
}

func noTeaIfNotEnoughMoneyTest() (string, func(t *testing.T, p *process.P)) {
	drink := ref.Tea
	return "the drink maker should not make " + drink.Name + " if not enough money is given",
		func(t *testing.T, p *process.P) {
			assertDrinkIsNotServed(t, p, drink, drink.Price-0.01)
		}
}

func noChocolateIfNotEnoughMoneyTest() (string, func(t *testing.T, p *process.P)) {
	drink := ref.Chocolate
	return "the drink maker should not make " + drink.Name + " if not enough money is given",
		func(t *testing.T, p *process.P) {
			assertDrinkIsNotServed(t, p, drink, drink.Price-0.01)
		}
}

func moreMoneyThanNeededTest() (string, func(t *testing.T, p *process.P)) {
	drink := ref.Tea
	return "the drink maker should make the drink if more money than needed is given",
		func(t *testing.T, p *process.P) {
			assertDrinkIsServed(t, p, drink, drink.Price+1.00)
		}
}
