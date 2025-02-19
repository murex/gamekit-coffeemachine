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
	"github.com/murex/gamekit-coffeemachine/ref"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func sendsNotificationWhenWaterTankIsEmptyAndOrderIsCoffee() (string, func(t *testing.T, p *process.P)) {
	return assertSendsNotificationOnEmptyTank(ref.Coffee, ref.Water)
}

func sendsNotificationWhenWaterTankIsEmptyAndOrderIsTea() (string, func(t *testing.T, p *process.P)) {
	return assertSendsNotificationOnEmptyTank(ref.Tea, ref.Water)
}

func sendsNotificationWhenWaterTankIsEmptyAndOrderIsOrangeJuice() (string, func(t *testing.T, p *process.P)) {
	return assertSendsNotificationOnEmptyTank(ref.OrangeJuice, ref.Water)
}

func sendsNotificationWhenMilkTankIsEmptyAndOrderIsChocolate() (string, func(t *testing.T, p *process.P)) {
	return assertSendsNotificationOnEmptyTank(ref.Chocolate, ref.Milk)
}

func assertSendsNotificationOnEmptyTank(
	drinkType ref.Drink,
	liquid ref.Liquid,
) (string, func(t *testing.T, p *process.P)) {
	return fmt.Sprintf("sends notification when asking for a %s and %s tank is empty", drinkType.Name, liquid),
		func(t *testing.T, p *process.P) {
			_, err := p.SendMessage(process.NewSetTankMessage(liquid, ref.Empty))
			require.NoError(t, err)

			_, err2 := runBuildDrinkMakerCommand(p, drinkType)
			require.NoError(t, err2)

			mailbox, err3 := p.SendMessage(process.NewDumpMailboxMessage())
			require.NoError(t, err3)

			mails := strings.Split(mailbox, "\n")

			assert.Equal(t, 1, len(mails))
			assert.Equal(t, string(liquid)+" tank is empty", mails[0])
		}
}

func sendsNoNotificationWhenMilkTankIsFull() (string, func(t *testing.T, p *process.P)) {
	return assertSendsNoNotificationOnFullTank(ref.Chocolate, ref.Milk)
}

func sendsNoNotificationWhenWaterTankIsFull() (string, func(t *testing.T, p *process.P)) {
	return assertSendsNoNotificationOnFullTank(ref.Coffee, ref.Water)
}

func assertSendsNoNotificationOnFullTank(
	drinkType ref.Drink,
	liquid ref.Liquid,
) (string, func(t *testing.T, p *process.P)) {
	return fmt.Sprintf("sends no notification when asking for a %s and %s tank is full", drinkType.Name, liquid),
		func(t *testing.T, p *process.P) {
			_, err := p.SendMessage(process.NewSetTankMessage(liquid, ref.Full))
			require.NoError(t, err)

			_, err2 := runBuildDrinkMakerCommand(p, drinkType)
			require.NoError(t, err2)

			mailbox, err3 := p.SendMessage(process.NewDumpMailboxMessage())
			require.NoError(t, err3)

			assert.Empty(t, mailbox)
		}
}
