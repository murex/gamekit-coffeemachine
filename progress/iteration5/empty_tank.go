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

package iteration5

import (
	"github.com/murex/gamekit-coffeemachine/process"
	"github.com/murex/gamekit-coffeemachine/ref"
	"testing"
)

func canServeCoffeeWhenWaterTankIsFull() (string, func(t *testing.T, p *process.P)) {
	return assertDrinkCanBeServed(ref.Coffee, ref.Water, ref.Full)
}

func canServeTeaWhenWaterTankIsFull() (string, func(t *testing.T, p *process.P)) {
	return assertDrinkCanBeServed(ref.Tea, ref.Water, ref.Full)
}

func canServeChocolateWhenMilkTankIsFull() (string, func(t *testing.T, p *process.P)) {
	return assertDrinkCanBeServed(ref.Chocolate, ref.Milk, ref.Full)
}

func canServeOrangeJuiceWhenWaterTankIsFull() (string, func(t *testing.T, p *process.P)) {
	return assertDrinkCanBeServed(ref.OrangeJuice, ref.Water, ref.Full)
}

func cannotServeTeaWhenWaterTankIsEmpty() (string, func(t *testing.T, p *process.P)) {
	return assertDrinkCanNotBeServed(ref.Tea, ref.Water, ref.Empty)
}

func cannotServeCoffeeWhenWaterTankIsEmpty() (string, func(t *testing.T, p *process.P)) {
	return assertDrinkCanNotBeServed(ref.Coffee, ref.Water, ref.Empty)
}

func cannotServeChocolateWhenMilkTankIsEmpty() (string, func(t *testing.T, p *process.P)) {
	return assertDrinkCanNotBeServed(ref.Chocolate, ref.Milk, ref.Empty)
}

func cannotServeOrangeJuiceWhenWaterTankIsEmpty() (string, func(t *testing.T, p *process.P)) {
	return assertDrinkCanNotBeServed(ref.OrangeJuice, ref.Water, ref.Empty)
}

func canServeCoffeeWhenMilkTankIsEmpty() (string, func(t *testing.T, p *process.P)) {
	return assertDrinkCanBeServed(ref.Coffee, ref.Milk, ref.Empty)
}

func canServeTeaWhenMilkTankIsEmpty() (string, func(t *testing.T, p *process.P)) {
	return assertDrinkCanBeServed(ref.Tea, ref.Milk, ref.Empty)
}

func canServeChocolateWhenWaterTankIsEmpty() (string, func(t *testing.T, p *process.P)) {
	return assertDrinkCanBeServed(ref.Chocolate, ref.Water, ref.Empty)
}

func canServeOrangeJuiceWhenMilkTankIsEmpty() (string, func(t *testing.T, p *process.P)) {
	return assertDrinkCanBeServed(ref.OrangeJuice, ref.Milk, ref.Empty)
}
