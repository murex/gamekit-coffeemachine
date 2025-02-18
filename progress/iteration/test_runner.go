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
	"github.com/murex/coffee-machine/progress-runner/process"
	"log"
	"os"
	"testing"
)

// TestCase is a convenience type used for defining a test case
// that can be passed directly to t.Run()
type TestCase func() (string, func(t *testing.T, p *process.P))

var (
	infoLog = log.New(os.Stderr, "INFO: ", log.Ltime)
	//warningLog = log.New(os.Stderr, "WARNING: ", log.Ltime)
)

// TestRunner provides the test case implementation for this iteration
type TestRunner struct {
	IterationIndex int        // the iteration this test runner is for
	TestCases      []TestCase // the test cases to run for this iteration
}

// New creates a new TestRunner for the given iteration, associated with the
// given test cases
func New(iterationIndex int, testCases ...TestCase) TestRunner {
	return TestRunner{
		IterationIndex: iterationIndex,
		TestCases:      testCases,
	}
}

// TestMain is the main test function for the iteration test runner.
// It runs sequentially all the test cases defined for this iteration.
func (tr TestRunner) TestMain(p *process.P) func(t *testing.T) {
	return func(t *testing.T) {
		tr.setupIterationTestSuite()
		for _, testCase := range tr.TestCases {
			desc, testFunc := testCase()
			t.Run(desc, func(t *testing.T) {
				setupTestCase(t, p)
				testFunc(t, p)
				teardownTestCase(t, p)
			})
		}
		tr.teardownIterationTestSuite()
	}
}

// setupIterationTestSuite is called before running the test cases for this iteration.
// currently it only logs a message
func (tr TestRunner) setupIterationTestSuite() {
	infoLog.Println("setting up progress tests for iteration", tr.IterationIndex)
}

// teardownIterationTestSuite is called after running the test cases for this iteration.
// currently it only logs a message
func (tr TestRunner) teardownIterationTestSuite() {
	infoLog.Println("tearing down progress tests for iteration", tr.IterationIndex)
}

// setupTestCase is called before running each test case.
// this is where the "restart" message is sent to the implementation process
func setupTestCase(t *testing.T, p *process.P) {
	t.Log("setting up progress test")
	_, _ = p.SendMessage(process.NewRestartMessage())
}

// teardownTestCase is called after running each test case.
// currently it only logs a message
func teardownTestCase(t *testing.T, _ *process.P) {
	t.Log("tearing down progress test")
}
