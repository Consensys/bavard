//go:build amd64 && !purego

package amd64

import (
	"math/bits"
	"testing"
)

// Reference implementations for verification

// refVALIGND computes expected result for VALIGND
// Intel: TMP = (SRC1 << 512) | SRC2, then shift right by imm elements
// SRC2 is in the low bits (elements 0-15), SRC1 is in high bits (elements 16-31)
func refVALIGND(src1, src2 [16]uint32, imm int) [16]uint32 {
	var concat [32]uint32
	copy(concat[:16], src2[:])  // src2 in low positions (as per Go asm operand order)
	copy(concat[16:], src1[:]) // src1 in high positions
	var result [16]uint32
	for i := 0; i < 16; i++ {
		result[i] = concat[i+imm]
	}
	return result
}

// refVALIGNQ computes expected result for VALIGNQ
// Intel: TMP = (SRC1 << 512) | SRC2, then shift right by imm elements
// SRC2 is in the low bits (elements 0-7), SRC1 is in high bits (elements 8-15)
func refVALIGNQ(src1, src2 [8]uint64, imm int) [8]uint64 {
	var concat [16]uint64
	copy(concat[:8], src2[:])  // src2 in low positions (as per Go asm operand order)
	copy(concat[8:], src1[:]) // src1 in high positions
	var result [8]uint64
	for i := 0; i < 8; i++ {
		result[i] = concat[i+imm]
	}
	return result
}

// refVPBLENDMQ computes expected result for VPBLENDMQ
func refVPBLENDMQ(src1, src2 [8]uint64, mask uint8) [8]uint64 {
	var result [8]uint64
	for i := 0; i < 8; i++ {
		if (mask>>i)&1 == 1 {
			result[i] = src2[i]
		} else {
			result[i] = src1[i]
		}
	}
	return result
}

// refVPBLENDMD computes expected result for VPBLENDMD
func refVPBLENDMD(src1, src2 [16]uint32, mask uint16) [16]uint32 {
	var result [16]uint32
	for i := 0; i < 16; i++ {
		if (mask>>i)&1 == 1 {
			result[i] = src2[i]
		} else {
			result[i] = src1[i]
		}
	}
	return result
}

// refVPERMQ computes expected result for VPERMQ with immediate
func refVPERMQ(src [8]uint64, imm uint8) [8]uint64 {
	var result [8]uint64
	for lane := 0; lane < 2; lane++ {
		base := lane * 4
		for i := 0; i < 4; i++ {
			idx := (imm >> (i * 2)) & 0x3
			result[base+i] = src[base+int(idx)]
		}
	}
	return result
}

// refVPERMD computes expected result for VPERMD
func refVPERMD(idx, src [16]uint32) [16]uint32 {
	var result [16]uint32
	for i := 0; i < 16; i++ {
		result[i] = src[idx[i]&0xF]
	}
	return result
}

// refVPERMI2Q computes expected result for VPERMI2Q
func refVPERMI2Q(src1 [8]uint64, idx [8]uint64, src2 [8]uint64) [8]uint64 {
	var result [8]uint64
	for i := 0; i < 8; i++ {
		index := idx[i] & 0xF
		if index < 8 {
			result[i] = src1[index]
		} else {
			result[i] = src2[index-8]
		}
	}
	return result
}

// refVPERMT2Q computes expected result for VPERMT2Q
func refVPERMT2Q(src1 [8]uint64, idx [8]uint64, src2 [8]uint64) [8]uint64 {
	return refVPERMI2Q(src1, idx, src2)
}

// refVSHUFI64X2 computes expected result for VSHUFI64X2
func refVSHUFI64X2(src1, src2 [8]uint64, imm uint8) [8]uint64 {
	var result [8]uint64
	sel0 := (imm >> 0) & 0x3
	sel1 := (imm >> 2) & 0x3
	sel2 := (imm >> 4) & 0x3
	sel3 := (imm >> 6) & 0x3

	result[0] = src1[sel0*2]
	result[1] = src1[sel0*2+1]
	result[2] = src1[sel1*2]
	result[3] = src1[sel1*2+1]
	result[4] = src2[sel2*2]
	result[5] = src2[sel2*2+1]
	result[6] = src2[sel3*2]
	result[7] = src2[sel3*2+1]

	return result
}

// refVSHUFPD computes expected result for VSHUFPD
func refVSHUFPD(src1, src2 [8]uint64, imm uint8) [8]uint64 {
	var result [8]uint64
	for i := 0; i < 8; i += 2 {
		bit := (imm >> (i / 2 * 2)) & 0x1
		if bit == 0 {
			result[i] = src1[i]
		} else {
			result[i] = src1[i+1]
		}
		bit = (imm >> (i/2*2 + 1)) & 0x1
		if bit == 0 {
			result[i+1] = src2[i]
		} else {
			result[i+1] = src2[i+1]
		}
	}
	return result
}

// refVPSHUFD computes expected result for VPSHUFD
func refVPSHUFD(src [16]uint32, imm uint8) [16]uint32 {
	var result [16]uint32
	for lane := 0; lane < 4; lane++ {
		base := lane * 4
		for i := 0; i < 4; i++ {
			idx := (imm >> (i * 2)) & 0x3
			result[base+i] = src[base+int(idx)]
		}
	}
	return result
}

// refVPUNPCKLDQ computes expected result for VPUNPCKLDQ
func refVPUNPCKLDQ(src1, src2 [16]uint32) [16]uint32 {
	var result [16]uint32
	for lane := 0; lane < 4; lane++ {
		base := lane * 4
		result[base+0] = src1[base+0]
		result[base+1] = src2[base+0]
		result[base+2] = src1[base+1]
		result[base+3] = src2[base+1]
	}
	return result
}

// refVPUNPCKHDQ computes expected result for VPUNPCKHDQ
func refVPUNPCKHDQ(src1, src2 [16]uint32) [16]uint32 {
	var result [16]uint32
	for lane := 0; lane < 4; lane++ {
		base := lane * 4
		result[base+0] = src1[base+2]
		result[base+1] = src2[base+2]
		result[base+2] = src1[base+3]
		result[base+3] = src2[base+3]
	}
	return result
}

// refVPUNPCKLQDQ computes expected result for VPUNPCKLQDQ
func refVPUNPCKLQDQ(src1, src2 [8]uint64) [8]uint64 {
	var result [8]uint64
	for lane := 0; lane < 4; lane++ {
		base := lane * 2
		result[base+0] = src1[base+0]
		result[base+1] = src2[base+0]
	}
	return result
}

// refVPUNPCKHQDQ computes expected result for VPUNPCKHQDQ
func refVPUNPCKHQDQ(src1, src2 [8]uint64) [8]uint64 {
	var result [8]uint64
	for lane := 0; lane < 4; lane++ {
		base := lane * 2
		result[base+0] = src1[base+1]
		result[base+1] = src2[base+1]
	}
	return result
}

// refVPMADD52LUQ computes expected result for VPMADD52LUQ (IFMA)
func refVPMADD52LUQ(a, b, c [8]uint64) [8]uint64 {
	const mask52 = (uint64(1) << 52) - 1
	var result [8]uint64
	for i := 0; i < 8; i++ {
		aLo := a[i] & mask52
		bLo := b[i] & mask52
		_, lo := bits.Mul64(aLo, bLo)
		prodLo52 := lo & mask52
		result[i] = (c[i] + prodLo52)
	}
	return result
}

// refVPMADD52HUQ computes expected result for VPMADD52HUQ (IFMA)
func refVPMADD52HUQ(a, b, c [8]uint64) [8]uint64 {
	const mask52 = (uint64(1) << 52) - 1
	var result [8]uint64
	for i := 0; i < 8; i++ {
		aLo := a[i] & mask52
		bLo := b[i] & mask52
		hi, lo := bits.Mul64(aLo, bLo)
		prodHi52 := (hi << 12) | (lo >> 52)
		result[i] = (c[i] + prodHi52)
	}
	return result
}

// refVPTERNLOGD computes expected result for VPTERNLOGD
func refVPTERNLOGD(a, b, c [16]uint32, imm uint8) [16]uint32 {
	var result [16]uint32
	for i := 0; i < 16; i++ {
		var r uint32
		for bit := 0; bit < 32; bit++ {
			aBit := (a[i] >> bit) & 1
			bBit := (b[i] >> bit) & 1
			cBit := (c[i] >> bit) & 1
			idx := (aBit << 2) | (bBit << 1) | cBit
			resBit := (uint32(imm) >> idx) & 1
			r |= resBit << bit
		}
		result[i] = r
	}
	return result
}

// Test functions

func TestVALIGND(t *testing.T) {
	src1 := [16]uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	src2 := [16]uint32{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}

	for _, imm := range []int{0, 1, 2, 4, 8} {
		var got [16]uint32
		testVALIGND(&src1, &src2, uint64(imm), &got)
		want := refVALIGND(src1, src2, imm)

		for i := 0; i < 16; i++ {
			if got[i] != want[i] {
				t.Errorf("VALIGND imm=%d: got[%d]=%d, want[%d]=%d", imm, i, got[i], i, want[i])
			}
		}
	}
}

func TestVALIGNQ(t *testing.T) {
	src1 := [8]uint64{0, 1, 2, 3, 4, 5, 6, 7}
	src2 := [8]uint64{8, 9, 10, 11, 12, 13, 14, 15}

	for _, imm := range []int{0, 1, 2, 4} {
		var got [8]uint64
		testVALIGNQ(&src1, &src2, uint64(imm), &got)
		want := refVALIGNQ(src1, src2, imm)

		for i := 0; i < 8; i++ {
			if got[i] != want[i] {
				t.Errorf("VALIGNQ imm=%d: got[%d]=%d, want[%d]=%d", imm, i, got[i], i, want[i])
			}
		}
	}
}

func TestVPBLENDMQ(t *testing.T) {
	src1 := [8]uint64{0, 1, 2, 3, 4, 5, 6, 7}
	src2 := [8]uint64{100, 101, 102, 103, 104, 105, 106, 107}

	for _, mask := range []uint8{0x00, 0xFF, 0xAA, 0x55, 0x0F, 0xF0} {
		var got [8]uint64
		testVPBLENDMQ(&src1, &src2, uint64(mask), &got)
		want := refVPBLENDMQ(src1, src2, mask)

		for i := 0; i < 8; i++ {
			if got[i] != want[i] {
				t.Errorf("VPBLENDMQ mask=0x%02X: got[%d]=%d, want[%d]=%d", mask, i, got[i], i, want[i])
			}
		}
	}
}

func TestVPBLENDMD(t *testing.T) {
	var src1, src2 [16]uint32
	for i := 0; i < 16; i++ {
		src1[i] = uint32(i)
		src2[i] = uint32(100 + i)
	}

	for _, mask := range []uint16{0x0000, 0xFFFF, 0xAAAA, 0x5555, 0x00FF, 0xFF00} {
		var got [16]uint32
		testVPBLENDMD(&src1, &src2, uint64(mask), &got)
		want := refVPBLENDMD(src1, src2, mask)

		for i := 0; i < 16; i++ {
			if got[i] != want[i] {
				t.Errorf("VPBLENDMD mask=0x%04X: got[%d]=%d, want[%d]=%d", mask, i, got[i], i, want[i])
			}
		}
	}
}

func TestVPERMQ(t *testing.T) {
	src := [8]uint64{10, 20, 30, 40, 50, 60, 70, 80}

	for _, imm := range []uint8{0x00, 0x55, 0xAA, 0xD8, 0x1B} {
		var got [8]uint64
		testVPERMQ(&src, uint64(imm), &got)
		want := refVPERMQ(src, imm)

		for i := 0; i < 8; i++ {
			if got[i] != want[i] {
				t.Errorf("VPERMQ imm=0x%02X: got[%d]=%d, want[%d]=%d", imm, i, got[i], i, want[i])
			}
		}
	}
}

func TestVPERMD(t *testing.T) {
	var src [16]uint32
	for i := 0; i < 16; i++ {
		src[i] = uint32(i * 10)
	}

	// Identity permutation
	idx1 := [16]uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var got1 [16]uint32
	testVPERMD(&idx1, &src, &got1)
	want1 := refVPERMD(idx1, src)
	for i := 0; i < 16; i++ {
		if got1[i] != want1[i] {
			t.Errorf("VPERMD identity: got[%d]=%d, want[%d]=%d", i, got1[i], i, want1[i])
		}
	}

	// Reverse permutation
	idx2 := [16]uint32{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	var got2 [16]uint32
	testVPERMD(&idx2, &src, &got2)
	want2 := refVPERMD(idx2, src)
	for i := 0; i < 16; i++ {
		if got2[i] != want2[i] {
			t.Errorf("VPERMD reverse: got[%d]=%d, want[%d]=%d", i, got2[i], i, want2[i])
		}
	}

	// Broadcast element 0
	idx3 := [16]uint32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var got3 [16]uint32
	testVPERMD(&idx3, &src, &got3)
	want3 := refVPERMD(idx3, src)
	for i := 0; i < 16; i++ {
		if got3[i] != want3[i] {
			t.Errorf("VPERMD broadcast: got[%d]=%d, want[%d]=%d", i, got3[i], i, want3[i])
		}
	}
}

func TestVPERMI2Q(t *testing.T) {
	src1 := [8]uint64{10, 20, 30, 40, 50, 60, 70, 80}
	src2 := [8]uint64{110, 120, 130, 140, 150, 160, 170, 180}

	// Select all from src1
	idx1 := [8]uint64{0, 1, 2, 3, 4, 5, 6, 7}
	var got1 [8]uint64
	testVPERMI2Q(&src1, &idx1, &src2, &got1)
	want1 := refVPERMI2Q(src1, idx1, src2)
	for i := 0; i < 8; i++ {
		if got1[i] != want1[i] {
			t.Errorf("VPERMI2Q from src1: got[%d]=%d, want[%d]=%d", i, got1[i], i, want1[i])
		}
	}

	// Select all from src2
	idx2 := [8]uint64{8, 9, 10, 11, 12, 13, 14, 15}
	var got2 [8]uint64
	testVPERMI2Q(&src1, &idx2, &src2, &got2)
	want2 := refVPERMI2Q(src1, idx2, src2)
	for i := 0; i < 8; i++ {
		if got2[i] != want2[i] {
			t.Errorf("VPERMI2Q from src2: got[%d]=%d, want[%d]=%d", i, got2[i], i, want2[i])
		}
	}

	// Interleave
	idx3 := [8]uint64{0, 8, 1, 9, 2, 10, 3, 11}
	var got3 [8]uint64
	testVPERMI2Q(&src1, &idx3, &src2, &got3)
	want3 := refVPERMI2Q(src1, idx3, src2)
	for i := 0; i < 8; i++ {
		if got3[i] != want3[i] {
			t.Errorf("VPERMI2Q interleave: got[%d]=%d, want[%d]=%d", i, got3[i], i, want3[i])
		}
	}
}

func TestVPERMT2Q(t *testing.T) {
	src1 := [8]uint64{10, 20, 30, 40, 50, 60, 70, 80}
	src2 := [8]uint64{110, 120, 130, 140, 150, 160, 170, 180}
	idx := [8]uint64{0, 8, 1, 9, 2, 10, 3, 11}

	var got [8]uint64
	testVPERMT2Q(&src1, &idx, &src2, &got)
	want := refVPERMT2Q(src1, idx, src2)

	for i := 0; i < 8; i++ {
		if got[i] != want[i] {
			t.Errorf("VPERMT2Q: got[%d]=%d, want[%d]=%d", i, got[i], i, want[i])
		}
	}
}

func TestVSHUFI64X2(t *testing.T) {
	src1 := [8]uint64{0, 1, 2, 3, 4, 5, 6, 7}
	src2 := [8]uint64{10, 11, 12, 13, 14, 15, 16, 17}

	for _, imm := range []uint8{0x00, 0x44, 0xEE} {
		var got [8]uint64
		testVSHUFI64X2(&src1, &src2, uint64(imm), &got)
		want := refVSHUFI64X2(src1, src2, imm)

		for i := 0; i < 8; i++ {
			if got[i] != want[i] {
				t.Errorf("VSHUFI64X2 imm=0x%02X: got[%d]=%d, want[%d]=%d", imm, i, got[i], i, want[i])
			}
		}
	}
}

func TestVSHUFPD(t *testing.T) {
	src1 := [8]uint64{0, 1, 2, 3, 4, 5, 6, 7}
	src2 := [8]uint64{10, 11, 12, 13, 14, 15, 16, 17}

	for _, imm := range []uint8{0x00, 0x55, 0xAA, 0xFF} {
		var got [8]uint64
		testVSHUFPD(&src1, &src2, uint64(imm), &got)
		want := refVSHUFPD(src1, src2, imm)

		for i := 0; i < 8; i++ {
			if got[i] != want[i] {
				t.Errorf("VSHUFPD imm=0x%02X: got[%d]=%d, want[%d]=%d", imm, i, got[i], i, want[i])
			}
		}
	}
}

func TestVPSHUFD(t *testing.T) {
	var src [16]uint32
	for i := 0; i < 16; i++ {
		src[i] = uint32(i)
	}

	for _, imm := range []uint8{0x00, 0x1B, 0xB1, 0xD8} {
		var got [16]uint32
		testVPSHUFD(&src, uint64(imm), &got)
		want := refVPSHUFD(src, imm)

		for i := 0; i < 16; i++ {
			if got[i] != want[i] {
				t.Errorf("VPSHUFD imm=0x%02X: got[%d]=%d, want[%d]=%d", imm, i, got[i], i, want[i])
			}
		}
	}
}

func TestVPUNPCKLDQ(t *testing.T) {
	var src1, src2 [16]uint32
	for i := 0; i < 16; i++ {
		src1[i] = uint32(i)
		src2[i] = uint32(100 + i)
	}

	var got [16]uint32
	testVPUNPCKLDQ(&src1, &src2, &got)
	want := refVPUNPCKLDQ(src1, src2)

	for i := 0; i < 16; i++ {
		if got[i] != want[i] {
			t.Errorf("VPUNPCKLDQ: got[%d]=%d, want[%d]=%d", i, got[i], i, want[i])
		}
	}
}

func TestVPUNPCKHDQ(t *testing.T) {
	var src1, src2 [16]uint32
	for i := 0; i < 16; i++ {
		src1[i] = uint32(i)
		src2[i] = uint32(100 + i)
	}

	var got [16]uint32
	testVPUNPCKHDQ(&src1, &src2, &got)
	want := refVPUNPCKHDQ(src1, src2)

	for i := 0; i < 16; i++ {
		if got[i] != want[i] {
			t.Errorf("VPUNPCKHDQ: got[%d]=%d, want[%d]=%d", i, got[i], i, want[i])
		}
	}
}

func TestVPUNPCKLQDQ(t *testing.T) {
	src1 := [8]uint64{0, 1, 2, 3, 4, 5, 6, 7}
	src2 := [8]uint64{100, 101, 102, 103, 104, 105, 106, 107}

	var got [8]uint64
	testVPUNPCKLQDQ(&src1, &src2, &got)
	want := refVPUNPCKLQDQ(src1, src2)

	for i := 0; i < 8; i++ {
		if got[i] != want[i] {
			t.Errorf("VPUNPCKLQDQ: got[%d]=%d, want[%d]=%d", i, got[i], i, want[i])
		}
	}
}

func TestVPUNPCKHQDQ(t *testing.T) {
	src1 := [8]uint64{0, 1, 2, 3, 4, 5, 6, 7}
	src2 := [8]uint64{100, 101, 102, 103, 104, 105, 106, 107}

	var got [8]uint64
	testVPUNPCKHQDQ(&src1, &src2, &got)
	want := refVPUNPCKHQDQ(src1, src2)

	for i := 0; i < 8; i++ {
		if got[i] != want[i] {
			t.Errorf("VPUNPCKHQDQ: got[%d]=%d, want[%d]=%d", i, got[i], i, want[i])
		}
	}
}

func TestVPMADD52LUQ(t *testing.T) {
	a := [8]uint64{1, 2, 3, 4, 5, 6, 7, 8}
	b := [8]uint64{10, 20, 30, 40, 50, 60, 70, 80}
	c := [8]uint64{100, 200, 300, 400, 500, 600, 700, 800}

	var got [8]uint64
	testVPMADD52LUQ(&a, &b, &c, &got)
	want := refVPMADD52LUQ(a, b, c)

	for i := 0; i < 8; i++ {
		if got[i] != want[i] {
			t.Errorf("VPMADD52LUQ: got[%d]=%d, want[%d]=%d", i, got[i], i, want[i])
		}
	}
}

func TestVPMADD52HUQ(t *testing.T) {
	const mask52 = (uint64(1) << 52) - 1
	a := [8]uint64{
		mask52, mask52 >> 1, mask52 >> 2, mask52 >> 3,
		1 << 40, 1 << 41, 1 << 42, 1 << 43,
	}
	b := [8]uint64{
		mask52, mask52 >> 1, mask52 >> 2, mask52 >> 3,
		1 << 40, 1 << 41, 1 << 42, 1 << 43,
	}
	c := [8]uint64{0, 0, 0, 0, 0, 0, 0, 0}

	var got [8]uint64
	testVPMADD52HUQ(&a, &b, &c, &got)
	want := refVPMADD52HUQ(a, b, c)

	for i := 0; i < 8; i++ {
		if got[i] != want[i] {
			t.Errorf("VPMADD52HUQ: got[%d]=0x%x, want[%d]=0x%x", i, got[i], i, want[i])
		}
	}
}

func TestVPTERNLOGD(t *testing.T) {
	var a, b, c [16]uint32
	for i := 0; i < 16; i++ {
		a[i] = uint32(0xAAAAAAAA)
		b[i] = uint32(0xCCCCCCCC)
		c[i] = uint32(0xF0F0F0F0)
	}

	tests := []struct {
		imm  uint8
		name string
	}{
		{0x96, "XOR"},
		{0x80, "AND"},
		{0xFE, "OR"},
	}

	for _, tc := range tests {
		var got [16]uint32
		testVPTERNLOGD(&a, &b, &c, uint64(tc.imm), &got)
		want := refVPTERNLOGD(a, b, c, tc.imm)

		for i := 0; i < 16; i++ {
			if got[i] != want[i] {
				t.Errorf("VPTERNLOGD %s (imm=0x%02X): got[%d]=0x%08X, want[%d]=0x%08X",
					tc.name, tc.imm, i, got[i], i, want[i])
			}
		}
	}
}
