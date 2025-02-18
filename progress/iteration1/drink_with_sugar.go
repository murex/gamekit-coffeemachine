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
	"fmt"
	"github.com/murex/coffee-machine/progress-runner/process"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func drinkWithNoSugarTest() (string, func(t *testing.T, p *process.P)) {
	return "I want to be able to send instructions to the drink maker to add no sugar",
		func(t *testing.T, p *process.P) {
			for _, drink := range drinks {
				t.Run("with "+drink.Name, func(t *testing.T) {
					cmd, err := runBuildDrinkMakerCommand(p, drink, 0)
					require.NoError(t, err)
					pattern := regexp.MustCompile("^.*::.*$")
					assert.Regexpf(t, pattern, cmd,
						"drink maker command for %s should contain '::'",
						drink.Name)
				})
			}
		}
}

func drinkWithSugarTest() (string, func(t *testing.T, p *process.P)) {
	return "I want to be able to send instructions to the drink maker to add",
		func(t *testing.T, p *process.P) {
			for _, drink := range drinks {
				for sugars := 1; sugars <= 2; sugars++ {
					desc := fmt.Sprintf("%d sugars with %s", sugars, drink.Name)
					t.Run(desc, func(t *testing.T) {
						cmd, err := runBuildDrinkMakerCommand(p, drink, sugars)
						require.NoError(t, err)
						pattern := regexp.MustCompile(fmt.Sprintf("^.*:%d:.*$", sugars))
						assert.Regexpf(t, pattern, cmd,
							"drink maker command for %s should contain ':%d:'",
							drink.Name, sugars)
					})
				}
			}
		}
}

func noMoreThan2SugarsTest() (string, func(t *testing.T, p *process.P)) {
	return "no more than 2 sugars",
		func(t *testing.T, p *process.P) {
			for _, drink := range drinks {
				desc := fmt.Sprintf("per %s", drink.Name)
				t.Run(desc, func(t *testing.T) {
					cmd, err := runBuildDrinkMakerCommand(p, drink, 3)
					require.NoError(t, err)
					pattern := regexp.MustCompile("^.*:2:.*$")
					assert.Regexpf(t, pattern, cmd,
						"drink maker command for %s should contain ':2:'",
						drink.Name)
				})
			}
		}
}
