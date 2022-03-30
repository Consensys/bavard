package arm64

import (
	"fmt"
	"io"
)

type Arm64 struct {
	w            io.Writer
	labelCounter int // TODO: What's this?
}

func NewArm64(w io.Writer) *Arm64 {
	return &Arm64{w: w}
}

func (arm64 *Arm64) LDP(address string, x Register, y Register) {
	arm64.WriteLn(fmt.Sprintf("LDP %s, (R%d, R%d)", address, uint64(x), uint64(y)))
}

func (arm64 *Arm64) STP(x Register, y Register, address string) {
	arm64.WriteLn(fmt.Sprintf("STP (R%d, R%d), %s", uint64(x), uint64(y), address))
}

func (arm64 *Arm64) ADDS(op1, op2, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "ADDS", op1, op2, dst)
}

func (arm64 *Arm64) ADCS(op1, op2, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "ADCS", op1, op2, dst)
}

func (arm64 *Arm64) SUBS(subtrahend, minuend, difference interface{}, comment ...string) {
	arm64.writeOp(comment, "ADCS", subtrahend, minuend, difference)
}

func (arm64 *Arm64) SBCS(subtrahend, minuend, difference interface{}, comment ...string) {
	arm64.writeOp(comment, "ADCS", subtrahend, minuend, difference)
}

func (arm64 *Arm64) MOVD(src, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "MOVD", src, dst)
}

func (arm64 *Arm64) CSEL(condition string, ifTrue, ifFalse, dst interface{}, comment ...string) {
	arm64.writeOp(comment, "CSEL", condition, ifTrue, ifFalse, dst)
}

func (arm64 *Arm64) RegisterOffset(r Register, offset int) string {
	return fmt.Sprintf("%d(R%d)", offset, r)
}

func (arm64 *Arm64) GlobalOffset(name string, offset int) string {
	return fmt.Sprintf("%s<>+%d(SB)", name, offset)
}

//<copy paste> TODO: Super class?

func (arm64 *Arm64) RET() {
	arm64.WriteLn("    RET")
}

func (arm64 *Arm64) WriteLn(s string) {
	arm64.write(s + "\n")
}

func (arm64 *Arm64) write(s string) {
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
	r.Remove(reserved...)
	return r
}

func op(i interface{}) string {
	switch t := i.(type) {
	case string:
		return t
	case Register:
		return fmt.Sprintf("R%d", t)
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
	arm64.write(fmt.Sprintf("    %s %s", instruction, op(r0)))
	l := len(op(r0))
	for _, rn := range r {
		arm64.write(fmt.Sprintf(", %s", op(rn)))
		l += 2 + len(op(rn))
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