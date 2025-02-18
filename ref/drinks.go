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

package ref

// Drink represents a drink that can be served by the coffee machine
type Drink struct {
	Name          string
	Price         float64
	CommandCode   string
	ReportKeyword string
}

// List of drinks that can be served by the coffee machine
var (
	Coffee      = Drink{Name: "coffee", Price: 0.60, CommandCode: "C", ReportKeyword: "coffee"}
	Tea         = Drink{Name: "tea", Price: 0.40, CommandCode: "T", ReportKeyword: "tea"}
	Chocolate   = Drink{Name: "chocolate", Price: 0.50, CommandCode: "H", ReportKeyword: "chocolate"}
	OrangeJuice = Drink{Name: "orange-juice", Price: 0.60, CommandCode: "O", ReportKeyword: "orange"}
)

// ExtraHotCommandFlag is the command flag to indicate that the drink should be served extra hot
const ExtraHotCommandFlag = "h"
