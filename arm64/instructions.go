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

// </ copy paste>
