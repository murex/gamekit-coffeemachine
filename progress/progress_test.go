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

package progress

import (
	"fmt"
	"github.com/murex/gamekit-coffeemachine/process"
	"github.com/murex/gamekit-coffeemachine/progress/iteration"
	"github.com/murex/gamekit-coffeemachine/progress/iteration1"
	"github.com/murex/gamekit-coffeemachine/progress/iteration2"
	"github.com/murex/gamekit-coffeemachine/progress/iteration3"
	"github.com/murex/gamekit-coffeemachine/progress/iteration4"
	"github.com/murex/gamekit-coffeemachine/progress/iteration5"
	"log"
	"os"
	"strconv"
	"testing"
)

var (
	infoLog    = log.New(os.Stderr, "INFO: ", log.Ltime)
	warningLog = log.New(os.Stderr, "WARNING: ", log.Ltime)
)

// NoIteration is a constant used to indicate that the ctx
// can run with any test iteration
const NoIteration = 0 // when iteration cannot be retrieved, all test iterations can be executed

// Context holds the context information for the current implementation process
type Context struct {
	iteration int
	process   *process.P
}

var ctx Context

// IterationTestRunner provides the interface to be implemented by each iteration test runner
//type IterationTestRunner interface {
//	TestMain(p *process.P) func(t *testing.T)
//}

// TestMain is the main function when running go test
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	proc := startImplementationProcess()
	ctx = Context{
		process:   proc,
		iteration: retrieveImplementationIteration(proc),
	}
}

func teardown() {
	stopImplementationProcess(ctx.process)
}

func startImplementationProcess() *process.P {
	infoLog.Println("starting implementation process")
	proc, errNew := process.NewCoffeeMachineProcess()
	if errNew != nil {
		warningLog.Println(errNew)
		return nil
	}
	errRun := proc.Run()
	if errRun != nil {
		warningLog.Println(errRun)
	}
	return proc
}

func retrieveImplementationIteration(proc *process.P) int {
	if proc == nil {
		warningLog.Println("no process to retrieve iteration from")
		return NoIteration
	}
	result, errMessage := proc.SendMessage(process.NewIterationMessage())
	// If there is any error, return NoIteration, e.g. all test iterations will be skipped
	if errMessage != nil {
		warningLog.Println(errMessage)
		return NoIteration
	}
	iter, errAtoi := strconv.Atoi(result)
	if errAtoi != nil {
		warningLog.Println(errAtoi)
		return NoIteration
	}
	return iter
}

func stopImplementationProcess(proc *process.P) {
	if proc != nil {
		infoLog.Println("stopping implementation process")
		_, _ = proc.SendMessage(process.NewShutdownMessage())
	}
}

// Test_Progress is the entry point for running all tests for all iterations
func Test_Progress(t *testing.T) {
	t.Log("testing implementation progress on coffee machine")
	if ctx.iteration == NoIteration {
		t.Skipf("skipping progress tests (implementation not ready)")
	}

	for _, it := range []iteration.TestRunner{
		iteration1.New(),
		iteration2.New(),
		iteration3.New(),
		iteration4.New(),
		iteration5.New(),
	} {
		t.Run(fmt.Sprintf("iteration %d", it.IterationIndex), func(t *testing.T) {
			if it.IterationIndex > ctx.iteration {
				t.Skipf("skipping iteration %d tests (implementation at iteration %d)",
					it.IterationIndex, ctx.iteration)
			}
			it.TestMain(ctx.process)(t)
		})
	}
}
