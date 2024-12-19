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

const (
	V0  = VectorRegister("V0")
	V1  = VectorRegister("V1")
	V2  = VectorRegister("V2")
	V3  = VectorRegister("V3")
	V4  = VectorRegister("V4")
	V5  = VectorRegister("V5")
	V6  = VectorRegister("V6")
	V7  = VectorRegister("V7")
	V8  = VectorRegister("V8")
	V9  = VectorRegister("V9")
	V10 = VectorRegister("V10")
	V11 = VectorRegister("V11")
	V12 = VectorRegister("V12")
	V13 = VectorRegister("V13")
	V14 = VectorRegister("V14")
	V15 = VectorRegister("V15")
	V16 = VectorRegister("V16")
	V17 = VectorRegister("V17")
	V18 = VectorRegister("V18")
	V19 = VectorRegister("V19")
	V20 = VectorRegister("V20")
	V21 = VectorRegister("V21")
	V22 = VectorRegister("V22")
	V23 = VectorRegister("V23")
	V24 = VectorRegister("V24")
	V25 = VectorRegister("V25")
	V26 = VectorRegister("V26")
	V27 = VectorRegister("V27")
	V28 = VectorRegister("V28")
	V29 = VectorRegister("V29")
	V30 = VectorRegister("V30")
	V31 = VectorRegister("V31")
)

// type Label string
type Register string
type VectorRegister string

func (vr VectorRegister) MemString() string {
	return "[" + string(vr) + "]"
}

func (vr VectorRegister) SAt(i int) string {
	return fmt.Sprintf("%s.S[%d]", string(vr), i)
}

// DAt
func (vr VectorRegister) DAt(i int) string {
	return fmt.Sprintf("%s.D[%d]", string(vr), i)
}

func (vr VectorRegister) S4() VectorRegister {
	return vr.withSuffix(".S4")
}

// B8, B16
func (vr VectorRegister) B8() VectorRegister {
	return vr.withSuffix(".B8")
}

func (vr VectorRegister) B16() VectorRegister {
	return vr.withSuffix(".B16")
}

// H8
func (vr VectorRegister) H8() VectorRegister {
	return vr.withSuffix(".H8")
}

func (vr VectorRegister) Q1() VectorRegister {
	return vr.withSuffix(".Q1")
}

func (vr VectorRegister) S2() VectorRegister {
	return vr.withSuffix(".S2")
}

func (vr VectorRegister) D1() VectorRegister {
	return vr.withSuffix(".D1")
}

func (vr VectorRegister) D2() VectorRegister {
	return vr.withSuffix(".D2")
}

func (vr VectorRegister) withSuffix(suffix string) VectorRegister {
	return VectorRegister(string(vr) + suffix)
}

type Registers struct {
	registers  []Register
	vRegisters []VectorRegister
	vAliases   map[string]VectorRegister
	f          *Arm64
}

func (r *Register) At(wordOffset int) string {
	return fmt.Sprintf("%d(%s)", wordOffset*8, string(*r))
}

func (r *Register) At2(wordOffset int) string {
	return fmt.Sprintf("%d(%s)", wordOffset*4, string(*r))
}

func (r *Registers) Available() int {
	return len(r.registers)
}

func (r *Registers) Pop() Register {
	toReturn := r.registers[0]
	r.registers = r.registers[1:]
	return toReturn
}

func (r *Registers) PopV(alias ...string) VectorRegister {
	toReturn := r.vRegisters[0]
	r.vRegisters = r.vRegisters[1:]

	if len(alias) > 0 {
		// check if alias is already used
		if _, ok := r.vAliases[alias[0]]; ok {
			panic("alias already used")
		}
		r.vAliases[alias[0]] = toReturn
		// write a #define
		r.f.WriteLn(fmt.Sprintf("#define %s %s", alias[0], string(toReturn)))
		return VectorRegister(alias[0])
	}

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

func (r *Registers) PushV(rIn ...VectorRegister) {
	// ensure register is in our original list, and no duplicate
	for _, register := range rIn {
		if _, ok := vRegisterSet[register]; !ok {
			// check if it's an alias
			realRegister, ok := r.vAliases[string(register)]
			if !ok {
				panic("warning: unknown register")
			}
			// remove the alias
			delete(r.vAliases, string(register))
			// undef
			r.f.WriteLn("#undef " + string(register))
			register = realRegister
		}
		found := false
		for _, existing := range r.vRegisters {
			if register == existing {
				found = true
				break
			}
		}
		if found {
			panic("duplicate register, already present.")
		}
		r.vRegisters = append(r.vRegisters, register)
	}

}

func NewRegisters(arm64 *Arm64) Registers {
	r := Registers{
		registers:  make([]Register, len(registers)),
		vRegisters: make([]VectorRegister, len(vRegisters)),
		vAliases:   make(map[string]VectorRegister),
		f:          arm64,
	}
	copy(r.registers, registers)
	copy(r.vRegisters, vRegisters)
	return r
}

func (r *Registers) AssertCleanState() {
	if len(r.vRegisters) != len(vRegisters) {
		// find the ones that are missing for a clear error message
		for _, vr := range vRegisters {
			found := false
			for _, vr2 := range r.vRegisters {
				if vr == vr2 {
					found = true
					break
				}
			}
			if !found {
				panic(fmt.Sprintf("missing push vector register %s", vr))
			}
		}
	}
	if len(r.registers) != len(registers) {
		// find the ones that are missing for a clear error message
		for _, vr := range registers {
			found := false
			for _, vr2 := range r.registers {
				if vr == vr2 {
					found = true
					break
				}
			}
			if !found {
				panic(fmt.Sprintf("missing push register %s", vr))
			}
		}
	}
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

var vRegisters = []VectorRegister{
	V0,
	V1,
	V2,
	V3,
	V4,
	V5,
	V6,
	V7,
	V8,
	V9,
	V10,
	V11,
	V12,
	V13,
	V14,
	V15,
	V16,
	V17,
	V18,
	V19,
	V20,
	V21,
	V22,
	V23,
	V24,
	V25,
	V26,
	V27,
	V28,
	V29,
	V30,
	V31,
}

var (
	registerSet  map[Register]struct{}
	vRegisterSet map[VectorRegister]struct{}
)

func init() {
	registerSet = make(map[Register]struct{}, 0)
	for _, register := range registers {
		registerSet[register] = struct{}{}
	}
	if len(registers) != NbRegisters {
		panic("update nb available registers")
	}

	vRegisterSet = make(map[VectorRegister]struct{}, 0)
	for _, register := range vRegisters {
		vRegisterSet[register] = struct{}{}
	}

}

func (arm64 *Arm64) NewLabel(prefix ...string) Label {
	arm64.labelCounter++
	if len(prefix) > 0 {
		return Label(fmt.Sprintf("%s%d", prefix[0], arm64.labelCounter))
	}
	return Label(fmt.Sprintf("l%d", arm64.labelCounter))
}
