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
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_process_init_should_fail_when_lang_impl_path_environment_variable_is_not_set(t *testing.T) {
	_ = os.Unsetenv(LangImplPathKey)
	p, err := NewCoffeeMachineProcess()
	assert.Nil(t, p)
	assert.Error(t, err)
}

func Test_process_init_should_fail_when_lang_impl_path_environment_variable_is_set_but_empty(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Unsetenv(LangImplPathKey)
	})
	_ = os.Setenv(LangImplPathKey, "")
	p, err := NewCoffeeMachineProcess()
	assert.Nil(t, p)
	assert.Error(t, err)
}

func Test_process_init_should_pass_when_lang_impl_path_environment_variable_is_set(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Unsetenv(LangImplPathKey)
	})
	_ = os.Setenv(LangImplPathKey, "path/to/language/implementation")
	p, err := NewCoffeeMachineProcess()
	assert.NotNil(t, p)
	assert.NoError(t, err)
}

func Test_scan_single_line_response(t *testing.T) {
	payload := "some single-line response"
	var stdin bytes.Buffer
	scanner := bufio.NewScanner(bufio.NewReader(&stdin))
	stdin.Write([]byte(payload))
	assert.Equal(t, payload, scanSingleLine(scanner))
}

func Test_scan_multiline_response(t *testing.T) {
	payload := "line 1\nline 2\nline 3"
	endMarker := "END-OF-RESPONSE"
	var stdin bytes.Buffer
	scanner := bufio.NewScanner(bufio.NewReader(&stdin))
	stdin.Write([]byte(payload + "\n" + endMarker))
	assert.Equal(t, payload, scanMultipleLines(scanner, endMarker))
}

func Test_run_and_message_smoke_test(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Unsetenv(LangImplPathKey)
	})
	_ = os.Setenv(LangImplPathKey, "dummy")
	p, err := NewCoffeeMachineProcess()
	require.NoError(t, err)

	// Replace the real call to a language's run.sh with a simple call to tee command
	// which allows replicating stdin to stdout.
	// This is to verify that the process is correctly started
	// and that we can send a message to its stdin and capture its response on stdout
	tee := exec.Command("tee")
	p.cmd.Path, p.cmd.Args, p.cmd.Dir = tee.Path, tee.Args, tee.Dir

	errRun := p.Run()
	assert.NoError(t, errRun)

	instruction := "dummy"
	response, errMsg := p.SendMessage(SimpleMessage{
		instruction:       Instruction(instruction),
		endResponseMarker: SingleLineResponseMarker,
	})
	assert.NoError(t, errMsg)
	assert.Equal(t, instruction, response)
}
