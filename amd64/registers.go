// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package amd64

import "fmt"

const DX = Register("DX")
const AX = Register("AX")

type Label string
type Register string

type Registers struct {
	registers []Register
}

func (r *Register) At(wordOffset int) string {
	return fmt.Sprintf("%d(%s)", wordOffset*8, string(*r))
}

func (r *Registers) Available() int {
	return len(r.registers)
}

func (r *Registers) Pop() Register {
	toReturn := r.registers[0]
	r.registers = r.registers[1:]
	return toReturn
}

func (r *Registers) PopN(n int) []Register {
	toReturn := make([]Register, n)
	for i := 0; i < n; i++ {
		toReturn[i] = r.Pop()
	}
	return toReturn
}

func (r *Registers) Remove(toRemove Register) {
	for j := 0; j < len(r.registers); j++ {
		if r.registers[j] == toRemove {
			r.registers[j] = r.registers[len(r.registers)-1]
			r.registers = r.registers[:len(r.registers)-1]
			return
		}
	}
	panic("register not found")
}

func (r *Registers) Push(rIn ...Register) {
	r.registers = append(r.registers, rIn...)
}

func NewRegisters() Registers {
	r := Registers{
		registers: make([]Register, len(registers)),
	}
	copy(r.registers, registers)
	return r
}

var registers = []Register{
	"AX",
	"DX",
	"CX",
	"BX",
	"BP",
	"SI",
	"DI",
	"R8",
	"R9",
	"R10",
	"R11",
	"R12",
	"R13",
	"R14",
	"R15",
}

func (amd64 *Amd64) NewLabel() Label {
	amd64.labelCounter++
	return Label(fmt.Sprintf("l%d", amd64.labelCounter))
}
