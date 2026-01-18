// Copyright 2020-2024 Consensys Software Inc.
// Licensed under the Apache License, Version 2.0. See the LICENSE file for details.

package amd64

import (
	"fmt"
)

const (
	AX  = Register("AX")
	DX  = Register("DX")
	CX  = Register("CX")
	BX  = Register("BX")
	BP  = Register("BP")
	SI  = Register("SI")
	DI  = Register("DI")
	R8  = Register("R8")
	R9  = Register("R9")
	R10 = Register("R10")
	R11 = Register("R11")
	R12 = Register("R12")
	R13 = Register("R13")
	R14 = Register("R14")

	R15 = Register("R15") // use with caution, see https://github.com/Consensys/gnark-crypto/issues/707
)

// Z0 .. Z31 AVX512 registers
const (
	Z0  = VectorRegister("Z0")
	Z1  = VectorRegister("Z1")
	Z2  = VectorRegister("Z2")
	Z3  = VectorRegister("Z3")
	Z4  = VectorRegister("Z4")
	Z5  = VectorRegister("Z5")
	Z6  = VectorRegister("Z6")
	Z7  = VectorRegister("Z7")
	Z8  = VectorRegister("Z8")
	Z9  = VectorRegister("Z9")
	Z10 = VectorRegister("Z10")
	Z11 = VectorRegister("Z11")
	Z12 = VectorRegister("Z12")
	Z13 = VectorRegister("Z13")
	Z14 = VectorRegister("Z14")
	Z15 = VectorRegister("Z15")
	Z16 = VectorRegister("Z16")
	Z17 = VectorRegister("Z17")
	Z18 = VectorRegister("Z18")
	Z19 = VectorRegister("Z19")
	Z20 = VectorRegister("Z20")
	Z21 = VectorRegister("Z21")
	Z22 = VectorRegister("Z22")
	Z23 = VectorRegister("Z23")
	Z24 = VectorRegister("Z24")
	Z25 = VectorRegister("Z25")
	Z26 = VectorRegister("Z26")
	Z27 = VectorRegister("Z27")
	Z28 = VectorRegister("Z28")
	Z29 = VectorRegister("Z29")
	Z30 = VectorRegister("Z30")
	Z31 = VectorRegister("Z31")
)

// Mask registers K0-K7 for AVX-512 predicated operations.
// K0 is special: when used as a write mask, it means "no masking" (all elements active).
// K1-K7 are general-purpose mask registers for predicated operations.
const (
	K0 = MaskRegister("K0") // special: no masking when used as write mask
	K1 = MaskRegister("K1")
	K2 = MaskRegister("K2")
	K3 = MaskRegister("K3")
	K4 = MaskRegister("K4")
	K5 = MaskRegister("K5")
	K6 = MaskRegister("K6")
	K7 = MaskRegister("K7")
)

type Label string
type Register string
type VectorRegister string
type MaskRegister string

func (vr VectorRegister) Y() VectorRegister {
	// replace first letter by Y
	return VectorRegister("Y" + string(vr[1:]))
}

func (vr VectorRegister) X() VectorRegister {
	// replace first letter by X
	return VectorRegister("X" + string(vr[1:]))
}

func (vr VectorRegister) Z() VectorRegister {
	// replace first letter by Z
	return VectorRegister("Z" + string(vr[1:]))
}

type Registers struct {
	registers  []Register
	vRegisters []VectorRegister
}

func (r *Register) At(wordOffset int) string {
	return fmt.Sprintf("%d(%s)", wordOffset*8, string(*r))
}

func (r *Register) AtD(wordOffset int) Register {
	return Register(fmt.Sprintf("%d(%s)", wordOffset*4, string(*r)))
}

func (r *Registers) Available() int {
	return len(r.registers)
}

func (r *Registers) AvailableV() int {
	return len(r.vRegisters)
}

func (r *Registers) Pop() Register {
	toReturn := r.registers[0]
	r.registers = r.registers[1:]
	return toReturn
}

func (r *Registers) PopV() VectorRegister {
	toReturn := r.vRegisters[0]
	r.vRegisters = r.vRegisters[1:]
	return toReturn
}

func (r *Registers) PopN(n int) []Register {
	toReturn := make([]Register, n)
	for i := 0; i < n; i++ {
		toReturn[i] = r.Pop()
	}
	return toReturn
}

func (r *Registers) PopVN(n int) []VectorRegister {
	toReturn := make([]VectorRegister, n)
	for i := 0; i < n; i++ {
		toReturn[i] = r.PopV()
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

// UnsafePush is used to push registers without checking if they are known registers.
func (r *Registers) UnsafePush(rIn ...Register) {
	// ensure register is in our original list, and no duplicate
	for _, register := range rIn {
		if _, ok := registerSet[register]; !ok {
			// fmt.Printf("warning: unknown register %s\n", register)
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

func (r *Registers) PushV(vIn ...VectorRegister) {
	// ensure register is in our original list, and no duplicate
	for _, register := range vIn {
		register = register.Z()
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

func NewRegisters() Registers {
	r := Registers{
		registers:  make([]Register, len(registers)),
		vRegisters: make([]VectorRegister, len(vRegisters)),
	}
	copy(r.registers, registers)
	copy(r.vRegisters, vRegisters)
	return r
}

// NbRegisters contains nb default available registers, without BP
const NbRegisters = 13

var registers = []Register{
	AX,
	DX,
	CX,
	BX,
	SI,
	DI,
	R8,
	R9,
	R10,
	R11,
	R12,
	R13,
	R14,
	// R15,
}

var vRegisters = []VectorRegister{
	Z0,
	Z1,
	Z2,
	Z3,
	Z4,
	Z5,
	Z6,
	Z7,
	Z8,
	Z9,
	Z10,
	Z11,
	Z12,
	Z13,
	Z14,
	Z15,
	Z16,
	Z17,
	Z18,
	Z19,
	Z20,
	Z21,
	Z22,
	Z23,
	Z24,
	Z25,
	Z26,
	Z27,
	Z28,
	Z29,
	Z30,
	Z31,
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

func (amd64 *Amd64) NewLabel(prefix ...string) Label {
	amd64.labelCounter++
	if len(prefix) > 0 {
		return Label(fmt.Sprintf("%s_%d", prefix[0], amd64.labelCounter))
	}
	return Label(fmt.Sprintf("l%d", amd64.labelCounter))
}
