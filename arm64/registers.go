package arm64

import (
	"fmt"
)

// R18 is reserved

const (
	R0  = Register("R0")
	R1  = Register("R1")
	R2  = Register("R2")
	R3  = Register("R3")
	R4  = Register("R4")
	R5  = Register("R5")
	R6  = Register("R6")
	R7  = Register("R7")
	R8  = Register("R8")
	R9  = Register("R9")
	R10 = Register("R10")
	R11 = Register("R11")
	R12 = Register("R12")
	R13 = Register("R13")
	R14 = Register("R14")
	R15 = Register("R15")
	R16 = Register("R16")
	R17 = Register("R17")
	R19 = Register("R19")
	R20 = Register("R20")
	R21 = Register("R21")
	R22 = Register("R22")
	R23 = Register("R23")
	R24 = Register("R24")
	R25 = Register("R25")
	R26 = Register("R26")
	R29 = Register("R29")
)

// type Label string
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
	// ensure register is in our original list, and no duplicate
	for _, register := range rIn {
		if _, ok := registerSet[register]; !ok {
			panic("warning: unknown register")
		}
		found := false
		for _, existing := range r.registers {
			if register == existing {
				found = true
				break
			}
		}
		if found {
			panic("duplicate register, already present.")
		}
		r.registers = append(r.registers, register)
	}

}

func NewRegisters() Registers {
	r := Registers{
		registers: make([]Register, len(registers)),
	}
	copy(r.registers, registers)
	return r
}

// NbRegisters contains nb default available registers, without BP
const NbRegisters = 27

var registers = []Register{
	R0,
	R1,
	R2,
	R3,
	R4,
	R5,
	R6,
	R7,
	R8,
	R9,
	R10,
	R11,
	R12,
	R13,
	R14,
	R15,
	R16,
	R17,
	R19,
	R20,
	R21,
	R22,
	R23,
	R24,
	R25,
	R26,
	R29, // risky. (reserved for FP)
}

var registerSet map[Register]struct{}

func init() {
	registerSet = make(map[Register]struct{}, 0)
	for _, register := range registers {
		registerSet[register] = struct{}{}
	}
	if len(registers) != NbRegisters {
		panic("update nb available registers")
	}
}

func (arm64 *Arm64) NewLabel(prefix ...string) Label {
	arm64.labelCounter++
	if len(prefix) > 0 {
		return Label(fmt.Sprintf("%s%d", prefix[0], arm64.labelCounter))
	}
	return Label(fmt.Sprintf("l%d", arm64.labelCounter))
}
