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
	"fmt"
	"github.com/murex/gamekit-coffeemachine/process"
	"github.com/murex/gamekit-coffeemachine/progress/iteration"
	"github.com/murex/gamekit-coffeemachine/ref"
	"github.com/stretchr/testify/require"
	"testing"
)

// New returns the test runner for this iteration
func New() iteration.TestRunner {
	return iteration.New(5,
		canServeCoffeeWhenWaterTankIsFull,
		canServeTeaWhenWaterTankIsFull,
		canServeChocolateWhenMilkTankIsFull,
		canServeOrangeJuiceWhenWaterTankIsFull,
		cannotServeCoffeeWhenWaterTankIsEmpty,
		cannotServeTeaWhenWaterTankIsEmpty,
		cannotServeChocolateWhenMilkTankIsEmpty,
		cannotServeOrangeJuiceWhenWaterTankIsEmpty,
		canServeCoffeeWhenMilkTankIsEmpty,
		canServeTeaWhenMilkTankIsEmpty,
		canServeChocolateWhenWaterTankIsEmpty,
		canServeOrangeJuiceWhenMilkTankIsEmpty,
		sendsNotificationWhenWaterTankIsEmptyAndOrderIsCoffee,
		sendsNotificationWhenWaterTankIsEmptyAndOrderIsTea,
		sendsNotificationWhenWaterTankIsEmptyAndOrderIsOrangeJuice,
		sendsNotificationWhenMilkTankIsEmptyAndOrderIsChocolate,
		sendsNoNotificationWhenWaterTankIsFull,
		sendsNoNotificationWhenMilkTankIsFull,
	)
}

const (
	noSugar    = 0     // number of sugars has no impact on drink being served or not
	noExtraHot = false // Extra hot option has no impact on drink being served or not
)

func runBuildDrinkMakerCommand(p *process.P, drink ref.Drink) (string, error) {
	return p.SendMessage(process.NewMakeDrinkMessage(drink.Name, noSugar, drink.Price, noExtraHot))
}

func assertDrinkIsServed(t *testing.T, p *process.P, drink ref.Drink) {
	t.Helper()
	response, err := runBuildDrinkMakerCommand(p, drink)
	iteration.AssertDrinkIsServed(t, drink, response, err)
}

func assertDrinkIsNotServed(t *testing.T, p *process.P, drink ref.Drink) {
	t.Helper()
	cmd, err := runBuildDrinkMakerCommand(p, drink)
	iteration.AssertDrinkIsNotServed(t, drink, cmd, err)
}

func assertDrinkCanBeServed(
	drinkType ref.Drink,
	liquid ref.Liquid,
	tankStatus ref.TankStatus,
) (string, func(t *testing.T, p *process.P)) {
	return fmt.Sprintf("%s can be served when the %s tank is %s", drinkType.Name, liquid, tankStatus),
		func(t *testing.T, p *process.P) {
			_, err := p.SendMessage(process.NewSetTankMessage(liquid, tankStatus))
			require.NoError(t, err)
			assertDrinkIsServed(t, p, drinkType)
		}
}

func assertDrinkCanNotBeServed(
	drinkType ref.Drink,
	liquid ref.Liquid,
	tankStatus ref.TankStatus, // nolint: unparam
) (string, func(t *testing.T, p *process.P)) {
	return fmt.Sprintf("%s can not be served when the %s tank is %s", drinkType.Name, liquid, tankStatus),
		func(t *testing.T, p *process.P) {
			_, err := p.SendMessage(process.NewSetTankMessage(liquid, tankStatus))
			require.NoError(t, err)
			assertDrinkIsNotServed(t, p, drinkType)
		}
}
