package arm64

import (
	"fmt"
	"io"
)

type Arm64 struct {
	w            io.Writer
	labelCounter int
	defineMode   bool
}

func NewArm64(w io.Writer) *Arm64 {
	return &Arm64{w: w}
}

func (arm64 *Arm64) StartDefine() {
	if arm64.defineMode {
		panic("Define cannot be nested")
	}
	arm64.defineMode = true
}

func (arm64 *Arm64) EndDefine() {
	arm64.defineMode = false
}

func (arm64 *Arm64) CBZ(r interface{}, label Label, comment ...string) {
	arm64.writeOp(comment, "CBZ", r, string(label))
}

func (arm64 *Arm64) BLE(label Label, comment ...string) {
	arm64.writeOp(comment, "BLE", string(label))
}

func (arm64 *Arm64) LDP(address string, x, y interface{}, comment ...string) {
	arm64.writeOp(comment, "LDP", address, toTuple(x, y))
}

func (arm64 *Arm64) LDPP(offset int, src, x, y interface{}, comment ...string) {
	src = fmt.Sprintf("%d(%s)", offset, Operand(src))
	arm64.writeOp(comment, "LDP.P", src, toTuple(x, y))
}

func (arm64 *Arm64) LDPW(offset int, src, x, y interface{}, comment ...string) {
	src = fmt.Sprintf("%d(%s)", offset, Operand(src))
	arm64.writeOp(comment, "LDP.W", src, toTuple(x, y))
}

func (arm64 *Arm64) STP(x, y interface{}, address string, comment ...string) {
	arm64.writeOp(comment, "STP", toTuple(x, y), address)
	//arm64.WriteLn(fmt.Sprintf("STP (R%d, R%d), %s", uint64(x), uint64(y), address))
}

func (arm64 *Arm64) ADDS(op1, op2, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "ADDS", op1, op2, dst)
}

func (arm64 *Arm64) ADD(op1, op2, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "ADD", op1, op2, dst)
}

func (arm64 *Arm64) ADCS(op1, op2, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "ADCS", op1, op2, dst)
}

func (arm64 *Arm64) ADC(op1, op2, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "ADC", op1, op2, dst)
}

func (arm64 *Arm64) SUB(op1, op2, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "SUB", op1, op2, dst)
}

func (arm64 *Arm64) SUBS(subtrahend, minuend, difference interface{}, comment ...string) {
	arm64.writeOp(comment, "SUBS", subtrahend, minuend, difference)
}

func (arm64 *Arm64) SBCS(subtrahend, minuend, difference interface{}, comment ...string) {
	arm64.writeOp(comment, "SBCS", subtrahend, minuend, difference)
}

func (arm64 *Arm64) ORR(op1, op2, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "ORR", op1, op2, dst)
}

func (arm64 *Arm64) MOVD(src, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "MOVD", src, dst)
}

// JMP
func (arm64 *Arm64) JMP(label Label, comment ...string) {
	arm64.writeOp(comment, "JMP", string(label))
}

// VLD1
func (arm64 *Arm64) VLD1(offset int, src any, dst VectorRegister, comment ...string) {
	src = fmt.Sprintf("%d(%s)", offset, Operand(src))
	arm64.writeOp(comment, "VLD1", src, dst.MemString())
}

// VADDV: Add all vector elements to produce a scalar result
func (arm64 *Arm64) VADDV(src, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VADDV", src, dst)
}

func (arm64 *Arm64) VLD1_P(offset int, src any, dst VectorRegister, comment ...string) {
	src = fmt.Sprintf("%d(%s)", offset, Operand(src))
	arm64.writeOp(comment, "VLD1.P", src, dst.MemString())
}

func (arm64 *Arm64) VLD2_P(offset int, src any, dst1, dst2 VectorRegister, comment ...string) {
	src = fmt.Sprintf("%d(%s)", offset, Operand(src))
	dst := VectorRegister(string(dst1) + ", " + string(dst2))
	arm64.writeOp(comment, "VLD2.P", src, dst.MemString())
}

// VSHL
func (arm64 *Arm64) VSHL(offset any, src, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VSHL", offset, src, dst)
}

// VMOV
func (arm64 *Arm64) VMOV(src, dst any, comment ...string) {
	arm64.writeOp(comment, "VMOV", src, dst)
}

func (arm64 *Arm64) VMOVI(value any, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VMOVI", value, dst)
}

func (arm64 *Arm64) VMOVS(value any, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VMOVS", value, dst)
}

func (arm64 *Arm64) VUSHLL(offset any, src, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VUSHLL", offset, src, dst)
}

func (arm64 *Arm64) VUSHLL2(offset any, src, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VUSHLL2", offset, src, dst)
}

// VDUP
func (arm64 *Arm64) VDUP(src any, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VDUP", src, dst)
}

// VUSHR
func (arm64 *Arm64) VUSHR(offset any, src, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VUSHR", offset, src, dst)
}

// SHRN
func (arm64 *Arm64) SHRN(immediate any, src, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "SHRN", immediate, src, dst)
}

func (arm64 *Arm64) VUSRA(immediate any, src, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VUSRA", immediate, src, dst)
}

// VST1_P
func (arm64 *Arm64) VST1_P(src VectorRegister, dst interface{}, offset int, comment ...string) {
	dst = fmt.Sprintf("%d(%s)", offset, Operand(dst))
	arm64.writeOp(comment, "VST1.P", src.MemString(), dst)
}

// VST2_P
func (arm64 *Arm64) VST2_P(src1, src2 VectorRegister, dst interface{}, offset int, comment ...string) {
	dst = fmt.Sprintf("%d(%s)", offset, Operand(dst))
	src := VectorRegister(string(src1) + ", " + string(src2))
	arm64.writeOp(comment, "VST2.P", src.MemString(), dst)
}

// VST2
func (arm64 *Arm64) VST2(src1, src2 VectorRegister, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "VST2", src1, src2, dst)
}

func (arm64 *Arm64) VMOVQ_cst(c1, c2 any, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VMOVQ", c1, c2, dst)
}

// VEOR
func (arm64 *Arm64) VEOR(op1, op2, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VEOR", op1, op2, dst)
}

//   "VREV16: Reverse byte order within 16-bit half-words.",
//   "VREV32: Reverse byte order within 32-bit words.",
//   "VREV64: Reverse byte order within 64-bit doublewords.",

// VREV16
func (arm64 *Arm64) VREV16(src, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VREV16", src, dst)
}

// VREV32
func (arm64 *Arm64) VREV32(src, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VREV32", src, dst)
}

// VREV64
func (arm64 *Arm64) VREV64(src, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VREV64", src, dst)
}

// VORR
func (arm64 *Arm64) VORR(op1, op2, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VORR", op1, op2, dst)
}

func (arm64 *Arm64) VEXT(n any, src, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VEXT", n, src, dst)
}

// VADD
func (arm64 *Arm64) VADD(op1, op2, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VADD", op1, op2, dst)
}

// VUADDW
func (arm64 *Arm64) VUADDW(op1, op2, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VUADDW", op1, op2, dst)
}

// UMULL
func (arm64 *Arm64) UMULL(op1, op2, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "UMULL", op1, op2, dst)
}

// VPMULL
func (arm64 *Arm64) VPMULL(op1, op2, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VPMULL", op1, op2, dst)
}

// VPMULL2
func (arm64 *Arm64) VPMULL2(op1, op2, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VPMULL2", op1, op2, dst)
}

// VAND
func (arm64 *Arm64) VAND(op1, op2, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VAND", op1, op2, dst)
}

// VSUB
func (arm64 *Arm64) VSUB(op1, op2, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VSUB", op1, op2, dst)
}

// VUMIN
func (arm64 *Arm64) VUMIN(op1, op2, dst VectorRegister, comment ...string) {
	arm64.writeOp(comment, "VUMIN", op1, op2, dst)
}

// MOVWU
func (arm64 *Arm64) MOVWU(src, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "MOVWU", src, dst)
}

// MOVWUP_Load
func (arm64 *Arm64) MOVWUP_Load(offset int, src, dst interface{}, comment ...string) {
	src = fmt.Sprintf("%d(%s)", offset, Operand(src))
	arm64.writeOp(comment, "MOVWU.P", src, dst)
}

// MOVWUP_Store
func (arm64 *Arm64) MOVWUP_Store(src, dst interface{}, offset int, comment ...string) {
	dst = fmt.Sprintf("%d(%s)", offset, Operand(dst))
	arm64.writeOp(comment, "MOVWU.P", src, dst)
}

// MOVW
func (arm64 *Arm64) MOVW(src, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "MOVW", src, dst)
}

func (arm64 *Arm64) MOVDP(offset int, src, dst interface{}, comment ...string) {
	src = fmt.Sprintf("%d(%s)", offset, Operand(src))
	arm64.writeOp(comment, "MOVD.P", src, dst)
}

func (arm64 *Arm64) MUL(op1, op2, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "MUL", op1, op2, dst)
}

func (arm64 *Arm64) MULW(op1, op2, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "MULW", op1, op2, dst)
}

func (arm64 *Arm64) UMULH(op1, op2, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "UMULH", op1, op2, dst)
}

func (arm64 *Arm64) CSEL(condition string, ifTrue, ifFalse, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "CSEL", condition, ifTrue, ifFalse, dst)
}

func (arm64 *Arm64) TST(a, b interface{}, comment ...string) {
	arm64.writeOp(comment, "TST", a, b)
}

func (arm64 *Arm64) CMP(a, b interface{}, comment ...string) {
	arm64.writeOp(comment, "CMP", a, b)
}

func (arm64 *Arm64) BEQ(label string, comment ...string) {
	arm64.writeOp(comment, "BEQ", label)
}

func (arm64 *Arm64) BLT(label Label, comment ...string) {
	arm64.writeOp(comment, "BLT", string(label))
}

func toTuple(x, y interface{}) string {
	return fmt.Sprintf("(%s, %s)", Operand(x), Operand(y))
}

//<copy paste> TODO: Super class?

type Label string

func (arm64 *Arm64) LABEL(l Label) {
	arm64.WriteLn(string(l) + ":")
}

func (arm64 *Arm64) RET() {
	arm64.WriteLn("    RET")
}

func (arm64 *Arm64) WriteLn(s string) {
	arm64.write(s + "\n")
}

func (arm64 *Arm64) write(s string) {
	// in define mode, if the last character is a newline, we insert a "\" before it
	if arm64.defineMode && len(s) > 0 && s[len(s)-1] == '\n' {
		arm64.w.Write([]byte(s[:len(s)-1] + "\\\n"))
		return
	}
	arm64.w.Write([]byte(s))
}

func (arm64 *Arm64) Comment(s string) {
	arm64.WriteLn("    // " + s)
}

func (arm64 *Arm64) FnHeader(funcName string, stackSize, argSize int, reserved ...Register) Registers {
	var header string
	if stackSize == 0 {
		header = "TEXT ·%s(SB), NOFRAME|NOSPLIT, $%d-%d"
	} else {
		header = "TEXT ·%s(SB), $%d-%d"
	}

	arm64.WriteLn(fmt.Sprintf(header, funcName, stackSize, argSize))
	r := NewRegisters(arm64)
	for _, rr := range reserved {
		r.Remove(rr)
	}
	return r
}

func Operand(i interface{}) string {
	switch t := i.(type) {
	case string:
		return t
	case Register:
		return string(t)
	case VectorRegister:
		return string(t)
	case int:
		switch t {
		case 0:
			return "$0"
		case 1:
			return "$1"
		default:
			return fmt.Sprintf("$0x%x", t)
		}
	case uint64:
		switch t {
		case 0:
			return "$0"
		case 1:
			return "$1"
		default:
			return fmt.Sprintf("$0x%x", t)
		}
	}
	panic("unsupported interface type")
}

func (arm64 *Arm64) writeOp(comments []string, instruction string, r0 interface{}, r ...interface{}) {
	arm64.write(fmt.Sprintf("    %s %s", instruction, Operand(r0)))
	l := len(Operand(r0))
	for _, rn := range r {
		arm64.write(fmt.Sprintf(", %s", Operand(rn)))
		l += 2 + len(Operand(rn))
	}
	if len(comments) == 1 {
		l = 50 - l
		for i := 0; i < l; i++ {
			arm64.write(" ")
		}
		arm64.write("// " + comments[0])
	}
	arm64.write("\n")
}

// -----------------------------------------------------------------------------
// NEON instructions not directly supported by the Go assembler
// These are encoded as raw WORD instructions
// -----------------------------------------------------------------------------

// vRegNum extracts the numeric register ID from a VectorRegister (V0 -> 0, V31 -> 31)
func vRegNum(v VectorRegister) uint32 {
	s := string(v)
	// Remove any suffix like .S4, .D2, etc.
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	// Parse "Vn" where n is 0-31
	if len(s) < 2 || s[0] != 'V' {
		panic("invalid vector register: " + string(v))
	}
	var n uint32
	for i := 1; i < len(s); i++ {
		n = n*10 + uint32(s[i]-'0')
	}
	return n
}

// VUMULL performs unsigned multiply long on the lower halves of two vectors
// UMULL Vd.2D, Vn.2S, Vm.2S - multiplies 2 pairs of 32-bit elements to produce 2 64-bit results
func (arm64 *Arm64) VUMULL(src1, src2, dst VectorRegister, comment ...string) {
	// Encoding: 0 01 01110 10 1 Rm 1100 00 Rn Rd
	// 0x2ea0c000 | (Rm << 16) | (Rn << 5) | Rd
	n := vRegNum(src1)
	m := vRegNum(src2)
	d := vRegNum(dst)
	encoding := uint32(0x2ea0c000) | (m << 16) | (n << 5) | d
	arm64.writeWordOp(encoding, fmt.Sprintf("UMULL %s.2D, %s.2S, %s.2S", baseReg(dst), baseReg(src1), baseReg(src2)), comment...)
}

// VUMULL2 performs unsigned multiply long on the upper halves of two vectors
// UMULL2 Vd.2D, Vn.4S, Vm.4S - multiplies 2 pairs of upper 32-bit elements to produce 2 64-bit results
func (arm64 *Arm64) VUMULL2(src1, src2, dst VectorRegister, comment ...string) {
	// Encoding: 0 10 01110 10 1 Rm 1100 00 Rn Rd
	// 0x6ea0c000 | (Rm << 16) | (Rn << 5) | Rd
	n := vRegNum(src1)
	m := vRegNum(src2)
	d := vRegNum(dst)
	encoding := uint32(0x6ea0c000) | (m << 16) | (n << 5) | d
	arm64.writeWordOp(encoding, fmt.Sprintf("UMULL2 %s.2D, %s.4S, %s.4S", baseReg(dst), baseReg(src1), baseReg(src2)), comment...)
}

// VUMLSL performs unsigned multiply-subtract long on the lower halves of two vectors
// UMLSL Vd.2D, Vn.2S, Vm.2S - multiplies 2 pairs of 32-bit elements to produce 2 64-bit results and subtracts from accumulator
func (arm64 *Arm64) VUMLSL(src1, src2, dst VectorRegister, comment ...string) {
	// Encoding: 0 01 01110 10 1 Rm 1010 00 Rn Rd
	// 0x2ea0a000 | (Rm << 16) | (Rn << 5) | Rd
	n := vRegNum(src1)
	m := vRegNum(src2)
	d := vRegNum(dst)
	encoding := uint32(0x2ea0a000) | (m << 16) | (n << 5) | d
	arm64.writeWordOp(encoding, fmt.Sprintf("UMLSL %s.2D, %s.2S, %s.2S", baseReg(dst), baseReg(src1), baseReg(src2)), comment...)
}

// VUMLSL2 performs unsigned multiply-subtract long on the upper halves of two vectors
// UMLSL2 Vd.2D, Vn.4S, Vm.4S - multiplies 2 pairs of upper 32-bit elements to produce 2 64-bit results and subtracts from accumulator
func (arm64 *Arm64) VUMLSL2(src1, src2, dst VectorRegister, comment ...string) {
	// Encoding: 0 10 01110 10 1 Rm 1010 00 Rn Rd
	// 0x6ea0a000 | (Rm << 16) | (Rn << 5) | Rd
	n := vRegNum(src1)
	m := vRegNum(src2)
	d := vRegNum(dst)
	encoding := uint32(0x6ea0a000) | (m << 16) | (n << 5) | d
	arm64.writeWordOp(encoding, fmt.Sprintf("UMLSL2 %s.2D, %s.4S, %s.4S", baseReg(dst), baseReg(src1), baseReg(src2)), comment...)
}

// VMUL_S4 performs 32-bit integer multiply on vectors (4 lanes)
// MUL Vd.4S, Vn.4S, Vm.4S
func (arm64 *Arm64) VMUL_S4(src1, src2, dst VectorRegister, comment ...string) {
	// Encoding: 0 1 0 01110 10 1 Rm 10011 1 Rn Rd
	// 0x4ea09c00 | (Rm << 16) | (Rn << 5) | Rd
	n := vRegNum(src1)
	m := vRegNum(src2)
	d := vRegNum(dst)
	encoding := uint32(0x4ea09c00) | (m << 16) | (n << 5) | d
	arm64.writeWordOp(encoding, fmt.Sprintf("MUL %s.4S, %s.4S, %s.4S", baseReg(dst), baseReg(src1), baseReg(src2)), comment...)
}

// VUZP2 deinterleaves the odd elements from two vectors
// UZP2 Vd.4S, Vn.4S, Vm.4S
func (arm64 *Arm64) VUZP2(src1, src2, dst VectorRegister, comment ...string) {
	// Encoding: 0 1 0 01110 10 0 Rm 0 101 10 Rn Rd
	// 0x4e805800 | (Rm << 16) | (Rn << 5) | Rd
	n := vRegNum(src1)
	m := vRegNum(src2)
	d := vRegNum(dst)
	encoding := uint32(0x4e805800) | (m << 16) | (n << 5) | d
	arm64.writeWordOp(encoding, fmt.Sprintf("UZP2 %s.4S, %s.4S, %s.4S", baseReg(dst), baseReg(src1), baseReg(src2)), comment...)
}

// VCMGT performs signed greater-than comparison
// CMGT Vd.4S, Vn.4S, Vm.4S - sets each element of Vd to all 1s if Vn > Vm, else all 0s
func (arm64 *Arm64) VCMGT(src1, src2, dst VectorRegister, comment ...string) {
	// Encoding: 0 1 0 01110 10 1 Rm 0011 01 Rn Rd
	// 0x4ea03400 | (Rm << 16) | (Rn << 5) | Rd
	n := vRegNum(src1)
	m := vRegNum(src2)
	d := vRegNum(dst)
	encoding := uint32(0x4ea03400) | (m << 16) | (n << 5) | d
	arm64.writeWordOp(encoding, fmt.Sprintf("CMGT %s.4S, %s.4S, %s.4S", baseReg(dst), baseReg(src1), baseReg(src2)), comment...)
}

// VLD1_P_Multi loads multiple registers with post-increment
// VLD1.P offset(src), [Vt1.4S, Vt2.4S, ...]
func (arm64 *Arm64) VLD1_P_Multi(offset int, src interface{}, dsts ...VectorRegister) {
	srcStr := fmt.Sprintf("%d(%s)", offset, Operand(src))
	var dstParts []string
	for _, d := range dsts {
		dstParts = append(dstParts, string(d.S4()))
	}
	arm64.write(fmt.Sprintf("    VLD1.P %s, [%s]\n", srcStr, join(dstParts, ", ")))
}

// VST1_P_Multi stores multiple registers with post-increment
// VST1.P [Vt1.4S, Vt2.4S, ...], offset(dst)
func (arm64 *Arm64) VST1_P_Multi(offset int, dst interface{}, srcs ...VectorRegister) {
	dstStr := fmt.Sprintf("%d(%s)", offset, Operand(dst))
	var srcParts []string
	for _, s := range srcs {
		srcParts = append(srcParts, string(s.S4()))
	}
	arm64.write(fmt.Sprintf("    VST1.P [%s], %s\n", join(srcParts, ", "), dstStr))
}

func join(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result += sep + parts[i]
	}
	return result
}

func baseReg(v VectorRegister) string {
	s := string(v)
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return s[:i]
		}
	}
	return s
}

func (arm64 *Arm64) writeWordOp(encoding uint32, asmComment string, comment ...string) {
	line := fmt.Sprintf("    WORD $0x%08x", encoding)
	if asmComment != "" {
		line += " // " + asmComment
	}
	if len(comment) == 1 {
		if asmComment != "" {
			line += " - "
		} else {
			line += " // "
		}
		line += comment[0]
	}
	arm64.write(line + "\n")
}
