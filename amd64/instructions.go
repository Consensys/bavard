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

// AVX 512 instructions

// VSHUFI64X2: Shuffle 128-Bit Packed Quadword Integer Values
func (amd64 *Amd64) VSHUFI64X2(r1, r2, r3, r4 interface{}, comment ...string) {
	amd64.writeOp(comment, "VSHUFI64X2", r1, r2, r3, r4)
}

// VPERMQ: Permute Quadword Integers
func (amd64 *Amd64) VPERMQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPERMQ", r1, r2, r3)
}

// VPERMD: Permute Doubleword Integers
func (amd64 *Amd64) VPERMD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPERMD", r1, r2, r3)
}

// VPERMD_BCST_Z: Permute Doubleword Integers (Broadcast, Zeroing Masking)
func (amd64 *Amd64) VPERMD_BCST_Z(r1, r2, k, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPERMD.BCST.Z", r1, r2, k, r3)
}

// KMOVW Move 16-bit Mask
func (amd64 *Amd64) KMOVW(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KMOVW", r1, r2)
}

// KMOVD Move 32-bit Mask
func (amd64 *Amd64) KMOVD(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KMOVD", r1, r2)
}

// KSHIFTLW Shift 16-bit Mask Left
func (amd64 *Amd64) KSHIFTLW(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KSHIFTLW", r1, r2, r3)
}

// KADDW Add 16-bit Masks
func (amd64 *Amd64) KADDW(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KADDW", r1, r2, r3)
}

// VXORPS Bitwise Logical XOR
func (amd64 *Amd64) VXORPS(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VXORPS", r1, r2, r3)
}

// VPXORQ Bitwise Logical Exclusive OR of Packed Quadword Integers
func (amd64 *Amd64) VPXORQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPXORQ", r1, r2, r3)
}

// VMOVDQA64 Move Aligned Quadword Values
func (amd64 *Amd64) VMOVDQA64(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVDQA64", r1, r2)
}

// VMOVDQA64_Z Move Aligned Quadword Values  (Zeroing Masking).
func (amd64 *Amd64) VMOVDQA64_Z(r1, k, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVDQA64.Z", r1, k, r2)
}

// VPMOVZXDQ Move Packed Doubleword Integers to Quadword Integers
func (amd64 *Amd64) VPMOVZXDQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMOVZXDQ", r1, r2)
}

// VMOVDQU64 Move Unaligned Quadword Values
func (amd64 *Amd64) VMOVDQU64(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVDQU64", r1, r2)
}

// VPADDQ Add Packed Quadword Integers
func (amd64 *Amd64) VPADDQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPADDQ", r1, r2, r3)
}

// VPMULUDQ Multiply Packed Unsigned Doubleword Integers
func (amd64 *Amd64) VPMULUDQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMULUDQ", r1, r2, r3)
}

// VPMULUDQ_BCST Multiply Packed Unsigned Doubleword Integers (Broadcast).
func (amd64 *Amd64) VPMULUDQ_BCST(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMULUDQ.BCST", r1, r2, r3)
}

// VPANDQ Bitwise Logical AND of Packed Quadword Integers
func (amd64 *Amd64) VPANDQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPANDQ", r1, r2, r3)
}

// VPSRLQ Shift Packed Quadword Data Right Logical
func (amd64 *Amd64) VPSRLQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSRLQ", r1, r2, r3)
}

// VPEXTRQ Extract Quadword
func (amd64 *Amd64) VPEXTRQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPEXTRQ", r1, r2, r3)
}

// VALIGND Align Doubleword Vectors
func (amd64 *Amd64) VALIGND(r1, r2, r3, k, r4 interface{}, comment ...string) {
	amd64.writeOp(comment, "VALIGND", r1, r2, r3, k, r4)
}

// VALIGND_Z Align Doubleword Vectors (Zeroing Masking).
func (amd64 *Amd64) VALIGND_Z(r1, r2, r3, k, r4 interface{}, comment ...string) {
	amd64.writeOp(comment, "VALIGND.Z", r1, r2, r3, k, r4)
}

// VALIGNQ Align Quadword Vectors
func (amd64 *Amd64) VALIGNQ(r1, r2, r3, r4 interface{}, comment ...string) {
	amd64.writeOp(comment, "VALIGNQ", r1, r2, r3, r4)
}

// VMOVQ Move Quadword
func (amd64 *Amd64) VMOVQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVQ", r1, r2)
}

// VMOVD Move Doubleword
func (amd64 *Amd64) VMOVD(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVD", r1, r2)
}

// VPCMPEQB Compare Packed Byte Data for Equality
func (amd64 *Amd64) VPCMPEQB(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPCMPEQB", r1, r2, r3)
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
	case Register:
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
