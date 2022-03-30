package arm64

import "fmt"

type Register uint
type Registers map[Register]struct{}

func (r Register) At(wordOffset int) string {
	return fmt.Sprintf("%d(R%d)", wordOffset*8, uint(r))
}

func (r Registers) Available() int {
	return len(r)
}

func (r Registers) Pop() (Register) {
	for k := range map[Register]struct{}(r) {
		delete(r, k)
		return k	// TODO: cleaner way?
	}
	panic("no registers available")
}

func (r Registers) PopN(n int) ([]Register) {
	toReturn := make([]Register, n)
	for i := 0; i < n; i++ {
		toReturn[i] = r.Pop()
	}
	return toReturn
}

func (r Registers) Push(rIn ...Register) {
	// ensure register is in our original list, and no duplicate
	for _, register := range rIn {

		if uint(register) >= 30 {
			panic("warning: unknown register")
		}

		if _, found := r[register]; found {
			panic("duplicate register, already present.")
		}

		r[register] = struct{}{}
	}
}

func (r Registers) Remove(registers ...Register) {
	for _, register := range registers {
		delete(r, register)
	}
}

func NewRegisters() Registers {
	r := make(map[Register]struct{}, 0)
	for i := 0; i < 30; i++ {
		r[Register(i)] = struct{}{}
	}
	return r
}