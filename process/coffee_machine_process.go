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
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var (
	debugLog  = log.New(os.Stderr, "DEBUG: ", log.LstdFlags)
	infoLog   = log.New(os.Stderr, "INFO: ", log.Ltime)
	errorLog  = log.New(os.Stderr, "ERROR: ", log.Ltime)
	stdinLog  = log.New(os.Stderr, "> ", 0)
	stdoutLog = log.New(os.Stderr, "< ", 0)
)

var errTimeout = fmt.Errorf("timeout while reading response from coffee machine")

// LangImplPathKey is the environment variable key used to specify the path to the language implementation to run
const LangImplPathKey = "LANG_IMPL_PATH"

const responseTimeout = 5 * time.Second

// Message is the interface that all messages sent to the coffee machine must implement
type Message interface {
	Format() string
	EndResponseMarker() string
}

// P is the process running the coffee machine implementation.
// It is responsible for the whole lifecycle management of this process.
// This is done mainly through sending messages to its stdin and receiving
// responses from its stdout.
// The short name P is chosen on purpose to make it easier to use in test cases
// in a similar fashion to what is used with testing.T
type P struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
}

// NewCoffeeMachineProcess creates (but does not start) a new process P
func NewCoffeeMachineProcess() (*P, error) {
	infoLog.Println("creating Coffee Machine Process")

	langImplPath := os.Getenv(LangImplPathKey)

	if langImplPath == "" {
		errorEnv := fmt.Errorf("%s environment variable is not set", LangImplPathKey)
		errorLog.Println(errorEnv.Error())
		return nil, errorEnv
	}

	language := filepath.Base(langImplPath)
	infoLog.Println("running", language, "implementation of coffee machine")

	cmd := exec.Command("bash", "-c", "./run.sh")
	cmd.Dir = langImplPath

	stdin, errStdinPipe := cmd.StdinPipe()
	if errStdinPipe != nil {
		errorLog.Println(errStdinPipe)
		return nil, errStdinPipe
	}

	stdout, errStdoutPipe := cmd.StdoutPipe()
	if errStdoutPipe != nil {
		errorLog.Println(errStdoutPipe)
		return nil, errStdoutPipe
	}

	cmd.Stderr = os.Stderr

	return &P{cmd: cmd, stdin: stdin, stdout: stdout}, nil
}

// Run actually starts the process P, and returns once it's ready to receive messages
func (p *P) Run() error {
	var ready = make(chan bool)
	go func() {
		debugLog.Println("p.cmd.Start()")
		if errStart := p.cmd.Start(); errStart != nil {
			errorLog.Println(errStart)
		}
		ready <- true

		debugLog.Println("before p.cmd.Wait()")
		if errWait := p.cmd.Wait(); errWait != nil {
			errorLog.Println(errWait)
		}
		debugLog.Println("after p.cmd.Wait()")
	}()
	<-ready
	return nil
}

// SendMessage sends a message to the coffee machine implementation and returns its response.
// Communication is done through the standard input and output of the process
func (p *P) SendMessage(msg Message) (string, error) {
	if msg == nil {
		return "", nil
	}

	str := msg.Format()
	stdinLog.Println(str)
	_, err := io.WriteString(p.stdin, str+"\n")
	if err != nil {
		errorLog.Println(err)
		return "", err
	}

	var scannerFunc func(scanner *bufio.Scanner) string
	switch msg.EndResponseMarker() {
	case SingleLineResponseMarker:
		scannerFunc = scanSingleLine
	default:
		scannerFunc = func(scanner *bufio.Scanner) string {
			return scanMultipleLines(scanner, msg.EndResponseMarker())
		}
	}
	return p.scanResponse(scannerFunc)
}

// scanResponse reads and returns the response from the coffee machine process.
func (p *P) scanResponse(scannerFunc func(scanner *bufio.Scanner) string) (string, error) {
	done := make(chan bool)
	timeoutCtx, cancel := context.WithTimeout(context.Background(), responseTimeout)
	defer cancel()

	scanner := bufio.NewScanner(p.stdout)
	var response string
	go func() {
		response = scannerFunc(scanner)
		done <- true
	}()

	select {
	case <-done:
		if errScanner := scanner.Err(); errScanner != nil {
			errorLog.Println(errScanner)
			return "", errScanner
		}
	case <-timeoutCtx.Done():
		errorLog.Println(errTimeout)
		errKill := p.cmd.Process.Signal(os.Kill)
		if errKill != nil {
			errorLog.Println(errKill)
			return "", errKill
		}
		return "", errTimeout
	}
	return response, nil
}

func scanSingleLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	stdoutLog.Println(scanner.Text())
	return scanner.Text()
}

func scanMultipleLines(scanner *bufio.Scanner, endMarker string) string {
	var outBuffer strings.Builder
	for scanner.Scan() {
		if scanner.Text() == endMarker {
			break
		}
		stdoutLog.Println(scanner.Text())
		_, _ = outBuffer.WriteString(scanner.Text())
		_, _ = outBuffer.WriteRune('\n')
	}
	return strings.TrimSpace(outBuffer.String())
}
