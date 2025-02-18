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

const (
	IterationInstruction   Instruction = "iteration"
	RestartInstruction     Instruction = "restart"
	ShutdownInstruction    Instruction = "shutdown"
	PrintReportInstruction Instruction = "print-report"
	MakeDrinkInstruction   Instruction = "make-drink"
	SetTankInstruction     Instruction = "set-tank"
	DumpMailboxInstruction Instruction = "dump-mailbox"
)

const SingleLineResponseMarker = ""

type SimpleMessage struct {
	instruction       Instruction
	endResponseMarker string
}

func (m SimpleMessage) Format() string {
	return string(m.instruction)
}

func (m SimpleMessage) EndResponseMarker() string {
	return m.endResponseMarker
}

func NewIterationMessage() SimpleMessage {
	return SimpleMessage{
		instruction:       IterationInstruction,
		endResponseMarker: SingleLineResponseMarker,
	}
}

func NewRestartMessage() Message {
	return SimpleMessage{
		instruction:       RestartInstruction,
		endResponseMarker: SingleLineResponseMarker,
	}
}

func NewShutdownMessage() SimpleMessage {
	return SimpleMessage{
		instruction:       ShutdownInstruction,
		endResponseMarker: SingleLineResponseMarker,
	}
}

func NewPrintReportMessage() SimpleMessage {
	return SimpleMessage{
		instruction:       PrintReportInstruction,
		endResponseMarker: "END-OF-REPORT",
	}
}

type MakeDrinkMessage struct {
	drinkType string
	sugars    int
	payment   float64
	extraHot  bool
}

func NewMakeDrinkMessage(drinkType string, sugars int, payment float64, extraHot bool) MakeDrinkMessage {
	return MakeDrinkMessage{
		drinkType: drinkType,
		sugars:    sugars,
		payment:   payment,
		extraHot:  extraHot,
	}
}

func (m MakeDrinkMessage) Format() string {
	return fmt.Sprintf("%s %s %d %.2f %t", MakeDrinkInstruction, m.drinkType, m.sugars, m.payment, m.extraHot)
}
func (m MakeDrinkMessage) EndResponseMarker() string {
	return SingleLineResponseMarker
}

type SetTankMessage struct {
	liquid ref.Liquid
	status ref.TankStatus
}

func NewSetTankMessage(liquid ref.Liquid, status ref.TankStatus) Message {
	return SetTankMessage{
		liquid: liquid,
		status: status,
	}
}

func NewDumpMailboxMessage() Message {
	return SimpleMessage{
		instruction:       DumpMailboxInstruction,
		endResponseMarker: "END-OF-MAILBOX",
	}
}

func (m SetTankMessage) Format() string {
	return fmt.Sprintf("%s %s %s", SetTankInstruction, m.liquid, m.status)
}

func (m SetTankMessage) EndResponseMarker() string {
	return SingleLineResponseMarker
}
