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
	"github.com/murex/coffee-machine/progress-runner/process"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func drinkInstructionsTest() (string, func(t *testing.T, p *process.P)) {
	return "the drink maker should receive correct instructions",
		func(t *testing.T, p *process.P) {
			for _, drink := range drinks {
				t.Run("for "+drink.Name, func(t *testing.T) {
					cmd, err := runBuildDrinkMakerCommand(p, drink, 0)
					require.NoError(t, err)
					pattern := regexp.MustCompile("^" + drink.CommandCode + ":.*$")
					assert.Regexpf(t, pattern, cmd,
						"drink maker command for %s should start with '%s:'",
						drink.Name, drink.CommandCode)
				})
			}
		}
}
