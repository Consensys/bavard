// Copyright 2020 ConsenSys AG
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

package bavard

import (
	"fmt"
	"io"
)

const DX = "DX"
const AX = "AX"

type Register string

func (r *Register) At(wordOffset int) string {
	return fmt.Sprintf("%d(%s)", wordOffset*8, string(*r))
}

type Assembly struct {
	writer    io.Writer
	registers []Register
}

func NewAssembly(w io.Writer) *Assembly {
	b := &Assembly{
		writer:    w,
		registers: make([]Register, len(staticRegisters)),
	}
	copy(b.registers, staticRegisters)
	return b
}
func (builder *Assembly) Reset() {
	builder.registers = make([]Register, len(staticRegisters))
	copy(builder.registers, staticRegisters)
}

func (builder *Assembly) AvailableRegisters() int {
	return len(builder.registers)
}

func (builder *Assembly) PopRegister() Register {
	r := builder.registers[0]
	builder.registers = builder.registers[1:]
	return r
}

func (builder *Assembly) PushRegister(r ...Register) {
	builder.registers = append(builder.registers, r...)
}

func (builder *Assembly) Comment(s string) {
	builder.WriteLn("    // " + s)
}

func (builder *Assembly) FuncHeader(funcName string, pSize int) {
	builder.WriteLn("")
	header := "TEXT ·%s(SB), NOSPLIT, $0-%d"
	builder.WriteLn(fmt.Sprintf(header, funcName, pSize))
}

func (builder *Assembly) WriteLn(s string) {
	builder.Write(s + "\n")
}

func (builder *Assembly) Write(s string) {
	builder.writer.Write([]byte(s))
}

func (builder *Assembly) RET() {
	builder.WriteLn("    RET")
}

func (builder *Assembly) MULXQ(src, lo, hi interface{}, comment ...string) {
	builder.writeOp(comment, "MULXQ", src, lo, hi)
}

func (builder *Assembly) SUBQ(r1, r2 interface{}, comment ...string) {
	builder.writeOp(comment, "SUBQ", r1, r2)
}

func (builder *Assembly) SBBQ(r1, r2 interface{}, comment ...string) {
	builder.writeOp(comment, "SBBQ", r1, r2)
}

func (builder *Assembly) ADDQ(r1, r2 interface{}, comment ...string) {
	builder.writeOp(comment, "ADDQ", r1, r2)
}

func (builder *Assembly) ADCQ(r1, r2 interface{}, comment ...string) {
	builder.writeOp(comment, "ADCQ", r1, r2)
}

func (builder *Assembly) ADOXQ(r1, r2 interface{}, comment ...string) {
	builder.writeOp(comment, "ADOXQ", r1, r2)
}

func (builder *Assembly) ADCXQ(r1, r2 interface{}, comment ...string) {
	builder.writeOp(comment, "ADCXQ", r1, r2)
}

func (builder *Assembly) XORQ(r1, r2 interface{}, comment ...string) {
	builder.writeOp(comment, "XORQ", r1, r2)
}

func (builder *Assembly) MOVQ(r1, r2 interface{}, comment ...string) {
	builder.writeOp(comment, "MOVQ", r1, r2)
}

func (builder *Assembly) IMULQ(r1, r2 interface{}, comment ...string) {
	builder.writeOp(comment, "IMULQ", r1, r2)
}

func (builder *Assembly) MULQ(r1 interface{}, comment ...string) {
	builder.writeOp(comment, "MULQ", r1)
}

func (builder *Assembly) CMPB(r1, r2 interface{}, comment ...string) {
	builder.writeOp(comment, "CMPB", r1, r2)
}

func (builder *Assembly) CMPQ(r1, r2 interface{}, comment ...string) {
	builder.writeOp(comment, "CMPQ", r1, r2)
}

func (builder *Assembly) CMOVQCC(r1, r2 interface{}, comment ...string) {
	builder.writeOp(comment, "CMOVQCC", r1, r2)
}

func (builder *Assembly) CMOVQCS(r1, r2 interface{}, comment ...string) {
	builder.writeOp(comment, "CMOVQCS", r1, r2)
}

func (builder *Assembly) JNE(label string, comment ...string) {
	builder.writeOp(comment, "JNE", label)
}

func (builder *Assembly) JCS(label string, comment ...string) {
	builder.writeOp(comment, "JCS", label)
}

func (builder *Assembly) JPS(label string, comment ...string) {
	builder.writeOp(comment, "JCC", label)
}

func (builder *Assembly) JMP(label string, comment ...string) {
	builder.writeOp(comment, "JMP", label)
}

func (builder *Assembly) writeOp(comments []string, instruction string, r0 interface{}, r ...interface{}) {
	builder.Write(fmt.Sprintf("    %s %s", instruction, op(r0)))
	l := len(op(r0))
	for _, rn := range r {
		builder.Write(fmt.Sprintf(", %s", op(rn)))
		l += (2 + len(op(rn)))
	}
	if len(comments) == 1 {
		l = 50 - l
		for i := 0; i < l; i++ {
			builder.Write(" ")
		}
		builder.Write("// " + comments[0])
	}
	builder.Write("\n")
}

func op(i interface{}) string {
	switch t := i.(type) {
	case string:
		return t
	case Register:
		return string(t)
	case int:
		return fmt.Sprintf("$%#016x", uint64(t))
	case uint64:
		return fmt.Sprintf("$%#016x", t)
	}
	panic("unsupported interface type")
}

var staticRegisters = []Register{ // AX and DX are reserved
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
