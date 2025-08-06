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
	"testing"

	"github.com/murex/gamekit-coffeemachine/process"
	"github.com/murex/gamekit-coffeemachine/ref"
)

func extraHotCoffeeTest() (string, func(t *testing.T, p *process.P)) {
	drink := ref.Coffee
	return "I want to be able to have my " + drink.Name + " extra hot",
		func(t *testing.T, p *process.P) {
			assertExtraHotCommand(t, p, drink)
		}
}

func extraHotTeaTest() (string, func(t *testing.T, p *process.P)) {
	drink := ref.Tea
	return "I want to be able to have my " + drink.Name + " extra hot",
		func(t *testing.T, p *process.P) {
			assertExtraHotCommand(t, p, drink)
		}
}

func extraHotChocolateTest() (string, func(t *testing.T, p *process.P)) {
	drink := ref.Chocolate
	return "I want to be able to have my " + drink.Name + " extra hot",
		func(t *testing.T, p *process.P) {
			assertExtraHotCommand(t, p, drink)
		}
}

func noExtraHotOrangeJuiceTest() (string, func(t *testing.T, p *process.P)) {
	drink := ref.OrangeJuice
	return "my " + drink.Name + " cannot be extra hot",
		func(t *testing.T, p *process.P) {
			assertNoExtraHotCommand(t, p, drink)
		}
}
