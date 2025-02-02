// Copyright 2020-2024 Consensys Software Inc.
// Licensed under the Apache License, Version 2.0. See the LICENSE file for details.

// Package amd64 contains wrapper to amd64 instructions in Go assembly.
// note that while this package is public, it is tailored for github.com/consensys/goff and github.com/consensys/gurvy
package amd64

import (
	"fmt"
	"io"
)

type Amd64 struct {
	w            io.Writer
	labelCounter int
	defineMode   bool
}

func NewAmd64(w io.Writer) *Amd64 {
	return &Amd64{w: w}
}

func (amd64 *Amd64) StartDefine() {
	if amd64.defineMode {
		panic("Define cannot be nested")
	}
	amd64.defineMode = true
}

func (amd64 *Amd64) EndDefine() {
	amd64.defineMode = false
}

func (amd64 *Amd64) RET() {
	amd64.WriteLn("    RET")
}

func (amd64 *Amd64) MULXQ(src, lo, hi interface{}, comment ...string) {
	amd64.writeOp(comment, "MULXQ", src, lo, hi)
}

func (amd64 *Amd64) SUBQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "SUBQ", r1, r2)
}

func (amd64 *Amd64) SBBQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "SBBQ", r1, r2)
}

func (amd64 *Amd64) ADDQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "ADDQ", r1, r2)
}

func (amd64 *Amd64) ADCQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "ADCQ", r1, r2)
}

func (amd64 *Amd64) ADOXQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "ADOXQ", r1, r2)
}

func (amd64 *Amd64) ADCXQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "ADCXQ", r1, r2)
}

func (amd64 *Amd64) XORQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "XORQ", r1, r2)
}

func (amd64 *Amd64) XORPS(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "XORPS", r1, r2)
}

func (amd64 *Amd64) MOVQ(r1, r2 interface{}, comment ...string) {
	if op(r1) != op(r2) {
		amd64.writeOp(comment, "MOVQ", r1, r2)
	}
}

// MOVD
func (amd64 *Amd64) MOVD(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "MOVD", r1, r2)
}

func (amd64 *Amd64) BTQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "BTQ", r1, r2)
}

func (amd64 *Amd64) MOVUPS(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "MOVUPS", r1, r2)
}

func (amd64 *Amd64) ANDQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "ANDQ", r1, r2)
}

func (amd64 *Amd64) BSFQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "BSFQ", r1, r2)
}

func (amd64 *Amd64) MOVNTIQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "MOVNTIQ", r1, r2)
}

func (amd64 *Amd64) SHRQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "SHRQ", r1, r2)
}

func (amd64 *Amd64) SHLQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "SHLQ", r1, r2)
}

func (amd64 *Amd64) SHRQw(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "SHRQ", r1, r2, r3)
}

func (amd64 *Amd64) SHRDw(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "SHRD", r1, r2, r3)
}

func (amd64 *Amd64) SHRXQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "SHRXQ", r1, r2, r3)
}

func (amd64 *Amd64) TZCNTQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "TZCNTQ", r1, r2)
}

func (amd64 *Amd64) INCQ(r1 interface{}, comment ...string) {
	amd64.writeOp(comment, "INCQ", r1)
}

func (amd64 *Amd64) DECQ(r1 interface{}, comment ...string) {
	amd64.writeOp(comment, "DECQ", r1)
}

func (amd64 *Amd64) PUSHQ(r1 interface{}, comment ...string) {
	amd64.writeOp(comment, "PUSHQ", r1)
}

func (amd64 *Amd64) POPQ(r1 interface{}, comment ...string) {
	amd64.writeOp(comment, "POPQ", r1)
}

func (amd64 *Amd64) IMULQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "IMULQ", r1, r2)
}

func (amd64 *Amd64) MULQ(r1 interface{}, comment ...string) {
	amd64.writeOp(comment, "MULQ", r1)
}

func (amd64 *Amd64) CMPB(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "CMPB", r1, r2)
}

func (amd64 *Amd64) CMPQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "CMPQ", r1, r2)
}

func (amd64 *Amd64) ORQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "ORQ", r1, r2)
}

func (amd64 *Amd64) TESTQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "TESTQ", r1, r2)
}

func (amd64 *Amd64) XCHGQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "XCHGQ", r1, r2)
}

func (amd64 *Amd64) CMOVQCC(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "CMOVQCC", r1, r2)
}

func (amd64 *Amd64) CMOVQEQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "CMOVQEQ", r1, r2)
}

func (amd64 *Amd64) CMOVQCS(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "CMOVQCS", r1, r2)
}

func (amd64 *Amd64) LABEL(l Label) {
	amd64.WriteLn(string(l) + ":")
}

// JNE x86 JNZ Jump short if not zero (ZF=0).
func (amd64 *Amd64) JNE(label Label, comment ...string) {
	amd64.writeOp(comment, "JNE", string(label))
}

// JEQ: x86 JZ Jump short if zero (ZF = 1).
func (amd64 *Amd64) JEQ(label Label, comment ...string) {
	amd64.writeOp(comment, "JEQ", string(label))
}

// JCS x86 JB Jump short if below (CF=1).
func (amd64 *Amd64) JCS(label Label, comment ...string) {
	amd64.writeOp(comment, "JCS", string(label))
}

// JCC x86 JNB Jump short if not below (CF=0).
func (amd64 *Amd64) JCC(label Label, comment ...string) {
	amd64.writeOp(comment, "JCC", string(label))
}

// JGE
func (amd64 *Amd64) JGE(label Label, comment ...string) {
	amd64.writeOp(comment, "JGE", string(label))
}

func (amd64 *Amd64) JMP(label Label, comment ...string) {
	amd64.writeOp(comment, "JMP", string(label))
}

func (amd64 *Amd64) JL(label Label, comment ...string) {
	amd64.writeOp(comment, "JL", string(label))
}

func (amd64 *Amd64) Comment(s string) {
	amd64.WriteLn("    // " + s)
}

func (amd64 *Amd64) FnHeader(funcName string, stackSize, argSize int, reserved ...Register) Registers {
	var header string
	if stackSize == 0 {
		header = "TEXT ·%s(SB), NOSPLIT, $%d-%d"
	} else {
		header = "TEXT ·%s(SB), $%d-%d"
	}

	amd64.WriteLn(fmt.Sprintf(header, funcName, stackSize, argSize))
	r := NewRegisters()
	for _, rr := range reserved {
		r.Remove(rr)
	}
	return r
}

func (amd64 *Amd64) WriteLn(s string) {
	amd64.write(s + "\n")
}

func (amd64 *Amd64) write(s string) {
	// in define mode, if the last character is a newline, we insert a "\" before it
	if amd64.defineMode && len(s) > 0 && s[len(s)-1] == '\n' {
		amd64.w.Write([]byte(s[:len(s)-1] + "\\\n"))
		return
	}
	amd64.w.Write([]byte(s))
}

func (amd64 *Amd64) writeOp(comments []string, instruction string, r0 interface{}, r ...interface{}) {
	amd64.write(fmt.Sprintf("    %s %s", instruction, op(r0)))
	l := len(op(r0))
	for _, rn := range r {
		amd64.write(fmt.Sprintf(", %s", op(rn)))
		l += (2 + len(op(rn)))
	}
	if len(comments) == 1 {
		l = 50 - l
		for i := 0; i < l; i++ {
			amd64.write(" ")
		}
		amd64.write("// " + comments[0])
	}
	amd64.write("\n")
}

func op(i interface{}) string {
	switch t := i.(type) {
	case string:
		return t
	case VectorRegister:
		return string(t)
	case Register:
		return string(t)
	case MaskRegister:
		return string(t)
	case int:
		switch t {
		case 0:
			return "$0"
		case 1:
			return "$1"
		default:
			return fmt.Sprintf("$%#016x", uint64(t))
		}
	case uint64:
		switch t {
		case 0:
			return "$0"
		case 1:
			return "$1"
		default:
			return fmt.Sprintf("$%#016x", t)
		}
	}
	panic("unsupported interface type")
}
