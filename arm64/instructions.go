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

func (arm64 *Arm64) CBZ(label string, comment ...string) {
	arm64.writeOp(comment, "BLE", label)
}

func (arm64 *Arm64) LDP(address string, x, y interface{}, comment ...string) {
	arm64.writeOp(comment, "LDP", address, toTuple(x, y))
}

func (arm64 *Arm64) LDPP(offset int, src, x, y interface{}, comment ...string) {
	src = fmt.Sprintf("%d(%s)", offset, Operand(src))
	arm64.writeOp(comment, "LDP.P", src, toTuple(x, y))
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

func (arm64 *Arm64) MOVDP(offset int, src, dst interface{}, comment ...string) {
	src = fmt.Sprintf("%d(%s)", offset, Operand(src))
	arm64.writeOp(comment, "MOVD.P", src, dst)
}

func (arm64 *Arm64) MUL(op1, op2, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "MUL", op1, op2, dst)
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
		header = "TEXT ·%s(SB), NOSPLIT, $%d-%d"
	} else {
		header = "TEXT ·%s(SB), $%d-%d"
	}

	arm64.WriteLn(fmt.Sprintf(header, funcName, stackSize, argSize))
	r := NewRegisters()
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
