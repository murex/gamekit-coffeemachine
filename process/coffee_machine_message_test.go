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

package process

import (
	"fmt"
	"github.com/murex/gamekit-coffeemachine/ref"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_iteration_message_format(t *testing.T) {
	assert.Equal(t, "iteration", NewIterationMessage().Format())
}

func Test_iteration_message_response_end_marker(t *testing.T) {
	assert.Equal(t, SingleLineResponseMarker, NewIterationMessage().EndResponseMarker())
}

func Test_restart_message_format(t *testing.T) {
	assert.Equal(t, "restart", NewRestartMessage().Format())
}

func Test_restart_message_response_end_marker(t *testing.T) {
	assert.Equal(t, SingleLineResponseMarker, NewRestartMessage().EndResponseMarker())
}

func Test_shutdown_message_format(t *testing.T) {
	assert.Equal(t, "shutdown", NewShutdownMessage().Format())
}

func Test_shutdown_message_response_end_marker(t *testing.T) {
	assert.Equal(t, SingleLineResponseMarker, NewShutdownMessage().EndResponseMarker())
}

func Test_print_report_format(t *testing.T) {
	assert.Equal(t, "print-report", NewPrintReportMessage().Format())
}

func Test_print_report_response_end_marker(t *testing.T) {
	assert.Equal(t, "END-OF-REPORT", NewPrintReportMessage().EndResponseMarker())
}

func Test_make_drink_message_format(t *testing.T) {
	tests := []struct {
		drink    ref.Drink
		sugars   int
		payment  float64
		extraHot bool
		expected string
	}{
		{ref.Coffee, 0, 0.0, false, "make-drink coffee 0 0.00 false"},
		{ref.Tea, 1, 2.0, false, "make-drink tea 1 2.00 false"},
		{ref.Chocolate, 2, 1.5, true, "make-drink chocolate 2 1.50 true"},
		{ref.OrangeJuice, 0, 1.0, false, "make-drink orange-juice 0 1.00 false"},
	}
	for _, test := range tests {
		desc := fmt.Sprintf("%s with %d sugars %.2f payment and extra hot %t",
			test.drink.Name, test.sugars, test.payment, test.extraHot)
		t.Run(desc, func(t *testing.T) {
			msg := NewMakeDrinkMessage(test.drink.Name, test.sugars, test.payment, test.extraHot)
			assert.Equal(t, test.expected, msg.Format())
		})
	}
}

func Test_make_drink_message_response_end_marker(t *testing.T) {
	msg := NewMakeDrinkMessage(ref.Coffee.Name, 0, 0.0, false)
	assert.Equal(t, SingleLineResponseMarker, msg.EndResponseMarker())
}

func Test_set_tank_message_format(t *testing.T) {
	tests := []struct {
		liquid   ref.Liquid
		status   ref.TankStatus
		expected string
	}{
		{ref.Water, ref.Empty, "set-tank water empty"},
		{ref.Milk, ref.Full, "set-tank milk full"},
	}
	for _, test := range tests {
		desc := fmt.Sprintf("%s tank with %s status", test.liquid, test.status)
		t.Run(desc, func(t *testing.T) {
			msg := NewSetTankMessage(test.liquid, test.status)
			assert.Equal(t, test.expected, msg.Format())
		})
	}
}

func Test_set_tank_message_response_end_marker(t *testing.T) {
	msg := NewSetTankMessage(ref.Water, ref.Empty)
	assert.Equal(t, SingleLineResponseMarker, msg.EndResponseMarker())
}

func Test_dump_mailbox_format(t *testing.T) {
	assert.Equal(t, "dump-mailbox", NewDumpMailboxMessage().Format())
}

func Test_dump_mailbox_response_end_marker(t *testing.T) {
	assert.Equal(t, "END-OF-MAILBOX", NewDumpMailboxMessage().EndResponseMarker())
}
