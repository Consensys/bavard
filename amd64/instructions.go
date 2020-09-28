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

// Package amd64 contains wrapper to amd64 instructions in Go assembly.
// note that while this package is public, it is tailored for github.com/consensys/goff and github.com/consensys/gurvy
package amd64

import (
	"fmt"
	"io"
)

var writer io.Writer

func SetWriter(w io.Writer) {
	writer = w
}

func RET() {
	WriteLn("    RET")
}

func MULXQ(src, lo, hi interface{}, comment ...string) {
	writeOp(comment, "MULXQ", src, lo, hi)
}

func SUBQ(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "SUBQ", r1, r2)
}

func SBBQ(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "SBBQ", r1, r2)
}

func ADDQ(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "ADDQ", r1, r2)
}

func ADCQ(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "ADCQ", r1, r2)
}

func ADOXQ(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "ADOXQ", r1, r2)
}

func ADCXQ(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "ADCXQ", r1, r2)
}

func XORQ(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "XORQ", r1, r2)
}

func XORPS(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "XORPS", r1, r2)
}

func MOVQ(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "MOVQ", r1, r2)
}

func MOVUPS(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "MOVUPS", r1, r2)
}

func MOVNTIQ(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "MOVNTIQ", r1, r2)
}

func PUSHQ(r1 interface{}, comment ...string) {
	writeOp(comment, "PUSHQ", r1)
}

func POPQ(r1 interface{}, comment ...string) {
	writeOp(comment, "POPQ", r1)
}

func IMULQ(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "IMULQ", r1, r2)
}

func MULQ(r1 interface{}, comment ...string) {
	writeOp(comment, "MULQ", r1)
}

func CMPB(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "CMPB", r1, r2)
}

func CMPQ(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "CMPQ", r1, r2)
}

func ORQ(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "ORQ", r1, r2)
}

func TESTQ(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "TESTQ", r1, r2)
}

func CMOVQCC(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "CMOVQCC", r1, r2)
}

func CMOVQCS(r1, r2 interface{}, comment ...string) {
	writeOp(comment, "CMOVQCS", r1, r2)
}

func LABEL(l Label) {
	WriteLn(string(l) + ":")
}

func JNE(label Label, comment ...string) {
	writeOp(comment, "JNE", string(label))
}

func JCS(label Label, comment ...string) {
	writeOp(comment, "JCS", string(label))
}

func JCC(label Label, comment ...string) {
	writeOp(comment, "JCC", string(label))
}

func JMP(label Label, comment ...string) {
	writeOp(comment, "JMP", string(label))
}

func Comment(s string) {
	WriteLn("    // " + s)
}

func FnHeader(funcName string, stackSize, argSize int, reserved ...Register) Registers {
	WriteLn("")
	var header string
	if stackSize == 0 {
		header = "TEXT ·%s(SB), NOSPLIT, $%d-%d"
	} else {
		header = "TEXT ·%s(SB), $%d-%d"
	}

	WriteLn(fmt.Sprintf(header, funcName, stackSize, argSize))
	r := NewRegisters()
	for _, rr := range reserved {
		r.Remove(rr)
	}
	return r
}

func WriteLn(s string) {
	write(s + "\n")
}

func write(s string) {
	writer.Write([]byte(s))
}

func writeOp(comments []string, instruction string, r0 interface{}, r ...interface{}) {
	write(fmt.Sprintf("    %s %s", instruction, op(r0)))
	l := len(op(r0))
	for _, rn := range r {
		write(fmt.Sprintf(", %s", op(rn)))
		l += (2 + len(op(rn)))
	}
	if len(comments) == 1 {
		l = 50 - l
		for i := 0; i < l; i++ {
			write(" ")
		}
		write("// " + comments[0])
	}
	write("\n")
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
