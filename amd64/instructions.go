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

func (amd64 *Amd64) MOVL(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "MOVL", r1, r2)
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

func (amd64 *Amd64) LEAL(offset, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "LEAL", fmt.Sprintf("%s(%s)", op(offset), op(r1)), r2)
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

func (amd64 *Amd64) IMUL3Q(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "IMUL3Q", r1, r2, r3)
}

func (amd64 *Amd64) XCHGL(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "XCHGL", r1, r2)
}

func (amd64 *Amd64) IMUL3L(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "IMUL3L", r1, r2, r3)
}

func (amd64 *Amd64) MULQ(r1 interface{}, comment ...string) {
	amd64.writeOp(comment, "MULQ", r1)
}

func (amd64 *Amd64) CMPB(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "CMPB", r1, r2)
}

func (amd64 *Amd64) CMPL(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "CMPL", r1, r2)
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

func (amd64 *Amd64) CMOVLCC(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "CMOVLCC", r1, r2)
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

func (amd64 *Amd64) TESTB(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "TESTB", r1, r2)
}

func (amd64 *Amd64) JNZ(label interface{}, comment ...string) {
	amd64.writeOp(comment, "JNZ", label)
}

// -----------------------------------------------------------------------------
// Prefetch Instructions
// -----------------------------------------------------------------------------

// PREFETCHT0 prefetches data into all cache levels (temporal data).
// Fetches the line of data from memory that contains the byte specified
// with the source operand to a location in the cache hierarchy specified
// by a locality hint (T0 = all cache levels).
//
// Forms:
//
//	PREFETCHT0 m8
//
// Example:
//
//	PREFETCHT0 2048(AX) // prefetch data 2KB ahead
func (amd64 *Amd64) PREFETCHT0(r1 interface{}, comment ...string) {
	amd64.writeOp(comment, "PREFETCHT0", r1)
}

// PREFETCHT1 prefetches data into L2 cache and higher (temporal data).
// Fetches the line of data from memory that contains the byte specified
// with the source operand to a location in the cache hierarchy specified
// by a locality hint (T1 = L2 and higher).
//
// Forms:
//
//	PREFETCHT1 m8
func (amd64 *Amd64) PREFETCHT1(r1 interface{}, comment ...string) {
	amd64.writeOp(comment, "PREFETCHT1", r1)
}

// PREFETCHT2 prefetches data into L3 cache and higher (temporal data).
// Fetches the line of data from memory that contains the byte specified
// with the source operand to a location in the cache hierarchy specified
// by a locality hint (T2 = L3 and higher).
//
// Forms:
//
//	PREFETCHT2 m8
func (amd64 *Amd64) PREFETCHT2(r1 interface{}, comment ...string) {
	amd64.writeOp(comment, "PREFETCHT2", r1)
}

// PREFETCHNTA prefetches data using non-temporal hint (minimizes cache pollution).
// Fetches the line of data from memory that contains the byte specified
// with the source operand to a location in the cache hierarchy specified
// by a locality hint (NTA = non-temporal, use streaming stores).
//
// Forms:
//
//	PREFETCHNTA m8
func (amd64 *Amd64) PREFETCHNTA(r1 interface{}, comment ...string) {
	amd64.writeOp(comment, "PREFETCHNTA", r1)
}

// -----------------------------------------------------------------------------
// Function Call Instructions
// -----------------------------------------------------------------------------

// CALL calls a procedure (function).
// Saves procedure linking information on the stack and branches to the
// procedure specified with the destination operand.
//
// Forms:
//
//	CALL rel32   // relative call
//	CALL r/m64   // indirect call
//
// Example:
//
//	CALL ·myFunction(SB)     // call Go function
//	CALL ·_mulGeneric(SB)    // call fallback implementation
func (amd64 *Amd64) CALL(target interface{}, comment ...string) {
	amd64.writeOp(comment, "CALL", target)
}

// -----------------------------------------------------------------------------
// Assembly Directives for Data Sections
// -----------------------------------------------------------------------------

// DATA defines a data constant in the data section.
// Used to define initialized data that will be embedded in the binary.
//
// Format:
//
//	DATA symbol<>+offset(SB)/width, value
//
// Example:
//
//	DATA ·myConst<>+0(SB)/8, $0x123456789ABCDEF0
//	DATA ·permuteIdx<>+0(SB)/8, $0
//
// Parameters:
//   - symbol: the symbol name (e.g., "·myConst<>")
//   - offset: byte offset from symbol start
//   - width: size of the data in bytes (1, 2, 4, or 8)
//   - value: the constant value
func (amd64 *Amd64) DATA(symbol string, offset int, width int, value interface{}, comment ...string) {
	amd64.writeOp(comment, "DATA", fmt.Sprintf("%s+%d(SB)/%d", symbol, offset, width), value)
}

// GLOBL declares a global symbol with the specified size and flags.
// Used to make symbols visible for linking and to specify their attributes.
//
// Common flags:
//   - RODATA: read-only data section
//   - NOPTR: data contains no pointers (helps GC)
//   - DUPOK: allow duplicate symbols
//
// Format:
//
//	GLOBL symbol(SB), flags, $size
//
// Example:
//
//	GLOBL ·myConst<>(SB), RODATA|NOPTR, $64
//
// Parameters:
//   - symbol: the symbol name (e.g., "·myConst<>")
//   - flags: symbol attributes as string (e.g., "RODATA|NOPTR")
//   - size: total size of the symbol in bytes
func (amd64 *Amd64) GLOBL(symbol string, flags string, size int, comment ...string) {
	amd64.writeOp(comment, "GLOBL", fmt.Sprintf("%s(SB)", symbol), flags, fmt.Sprintf("$%d", size))
}

// NO_LOCAL_POINTERS is an assembly directive that indicates the function
// does not have any local variables that contain pointers.
// This is a hint to the Go runtime garbage collector.
//
// This directive should be placed at the start of a function body,
// after the TEXT directive.
func (amd64 *Amd64) NO_LOCAL_POINTERS() {
	amd64.WriteLn("    NO_LOCAL_POINTERS")
}

// -----------------------------------------------------------------------------
// Additional Integer Instructions
// -----------------------------------------------------------------------------

// NEGQ negates a 64-bit value (two's complement negation).
// Replaces the value with its two's complement (negative value).
//
// Forms:
//
//	NEGQ r64
//	NEGQ m64
func (amd64 *Amd64) NEGQ(r1 interface{}, comment ...string) {
	amd64.writeOp(comment, "NEGQ", r1)
}

// NOTQ performs bitwise NOT on a 64-bit value.
// Inverts every bit of the operand (one's complement).
//
// Forms:
//
//	NOTQ r64
//	NOTQ m64
func (amd64 *Amd64) NOTQ(r1 interface{}, comment ...string) {
	amd64.writeOp(comment, "NOTQ", r1)
}

// LEAQ computes the effective address and stores it.
// Useful for computing offsets and pointer arithmetic without memory access.
//
// Forms:
//
//	LEAQ m64, r64
//
// Example:
//
//	LEAQ 8(AX)(BX*8), CX  // CX = AX + BX*8 + 8
func (amd64 *Amd64) LEAQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "LEAQ", r1, r2)
}

// ROLQ rotates 64-bit value left by specified count.
//
// Forms:
//
//	ROLQ imm8, r64
//	ROLQ CL, r64
func (amd64 *Amd64) ROLQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "ROLQ", r1, r2)
}

// RORQ rotates 64-bit value right by specified count.
//
// Forms:
//
//	RORQ imm8, r64
//	RORQ CL, r64
func (amd64 *Amd64) RORQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "RORQ", r1, r2)
}

// SARQ shifts 64-bit value right arithmetically (preserves sign).
//
// Forms:
//
//	SARQ imm8, r64
//	SARQ CL, r64
func (amd64 *Amd64) SARQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "SARQ", r1, r2)
}
