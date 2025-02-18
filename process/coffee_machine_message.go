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
	"github.com/murex/coffee-machine/progress-runner/ref"
)

// Instruction represents the different types of instructions that can be sent to the coffee machine
type Instruction string

// List of instructions that can be sent to the coffee machine
const (
	IterationInstruction   Instruction = "iteration"
	RestartInstruction     Instruction = "restart"
	ShutdownInstruction    Instruction = "shutdown"
	PrintReportInstruction Instruction = "print-report"
	MakeDrinkInstruction   Instruction = "make-drink"
	SetTankInstruction     Instruction = "set-tank"
	DumpMailboxInstruction Instruction = "dump-mailbox"
)

// SingleLineResponseMarker is the marker that indicates the end of a single line response
const SingleLineResponseMarker = ""

// SimpleMessage represents a message that contains only an instruction and no additional parameters
type SimpleMessage struct {
	instruction       Instruction
	endResponseMarker string
}

// Format returns the message formatted as a string
func (m SimpleMessage) Format() string {
	return string(m.instruction)
}

// EndResponseMarker returns the marker that indicates the end of the response for this message
func (m SimpleMessage) EndResponseMarker() string {
	return m.endResponseMarker
}

// NewIterationMessage creates a message to retrieve the current iteration of the coffee machine implementation
func NewIterationMessage() SimpleMessage {
	return SimpleMessage{
		instruction:       IterationInstruction,
		endResponseMarker: SingleLineResponseMarker,
	}
}

// NewRestartMessage creates a message to restart the coffee machine
func NewRestartMessage() Message {
	return SimpleMessage{
		instruction:       RestartInstruction,
		endResponseMarker: SingleLineResponseMarker,
	}
}

// NewShutdownMessage creates a message to shut down the coffee machine
func NewShutdownMessage() SimpleMessage {
	return SimpleMessage{
		instruction:       ShutdownInstruction,
		endResponseMarker: SingleLineResponseMarker,
	}
}

// NewPrintReportMessage creates a message to print a report
func NewPrintReportMessage() SimpleMessage {
	return SimpleMessage{
		instruction:       PrintReportInstruction,
		endResponseMarker: "END-OF-REPORT",
	}
}

// MakeDrinkMessage represents a message to make a drink
type MakeDrinkMessage struct {
	drinkType string
	sugars    int
	payment   float64
	extraHot  bool
}

// NewMakeDrinkMessage creates a message to make a drink
func NewMakeDrinkMessage(drinkType string, sugars int, payment float64, extraHot bool) MakeDrinkMessage {
	return MakeDrinkMessage{
		drinkType: drinkType,
		sugars:    sugars,
		payment:   payment,
		extraHot:  extraHot,
	}
}

// Format returns the message formatted as a string
func (m MakeDrinkMessage) Format() string {
	return fmt.Sprintf("%s %s %d %.2f %t", MakeDrinkInstruction, m.drinkType, m.sugars, m.payment, m.extraHot)
}

// EndResponseMarker returns the marker that indicates the end of the response for this message
func (MakeDrinkMessage) EndResponseMarker() string {
	return SingleLineResponseMarker
}

// SetTankMessage represents a message to set the tank status
type SetTankMessage struct {
	liquid ref.Liquid
	status ref.TankStatus
}

// NewSetTankMessage creates a message to set the tank status
func NewSetTankMessage(liquid ref.Liquid, status ref.TankStatus) Message {
	return SetTankMessage{
		liquid: liquid,
		status: status,
	}
}

// NewDumpMailboxMessage creates a message to dump the mailbox
func NewDumpMailboxMessage() Message {
	return SimpleMessage{
		instruction:       DumpMailboxInstruction,
		endResponseMarker: "END-OF-MAILBOX",
	}
}

// Format returns the message formatted as a string
func (m SetTankMessage) Format() string {
	return fmt.Sprintf("%s %s %s", SetTankInstruction, m.liquid, m.status)
}

// EndResponseMarker returns the marker that indicates the end of the response for this message
func (SetTankMessage) EndResponseMarker() string {
	return SingleLineResponseMarker
}
