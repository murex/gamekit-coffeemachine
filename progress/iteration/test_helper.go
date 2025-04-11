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

package iteration

import (
	"github.com/murex/gamekit-coffeemachine/ref"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

// AssertDrinkIsServed is a helper function asserting that a drink is served
func AssertDrinkIsServed(t *testing.T, drink ref.Drink, response string, err error) {
	t.Helper()
	require.NoError(t, err)
	pattern := regexp.MustCompile("^" + drink.CommandCode + ":.*$")
	assert.Regexpf(t, pattern, response,
		"drink maker command for %s should start with '%s:'",
		drink.Name, drink.CommandCode)
}

// AssertDrinkIsNotServed is a helper function asserting that a drink is not served
func AssertDrinkIsNotServed(t *testing.T, drink ref.Drink, response string, err error) {
	t.Helper()
	require.NoError(t, err)
	pattern := regexp.MustCompile("^[^" + drink.CommandCode + "]:.*$")
	assert.Regexpf(t, pattern, response,
		"drink maker command for %s should not start with '%s:' when not enough money",
		drink.Name, drink.CommandCode)
}

// AssertMissingMoneyMessageFormat is a helper function asserting that a message indicating that
// not enough money was provided contains the expected information (starting with 'M:' and
// containing the missing amount)
// nolint:revive
func AssertMissingMoneyMessageFormat(t *testing.T, drink ref.Drink, payment float64, response string, err error) {
	t.Helper()
	require.NoError(t, err)
	missing := ref.AmountRegexp(drink.Price - payment)
	pattern := regexp.MustCompile("^M:.*" + missing + ".*$")
	assert.Regexpf(t, pattern, response,
		"drink maker command should start with 'M:' and contain missing amount '%s'",
		missing)
}
