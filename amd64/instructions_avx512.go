package amd64

import "fmt"

// AVX 512 instructions
// some of the documentation (in particular for AVX512 ops) is taken from https://github.com/mmcloughlin/avo

// VPMULLD: Multiply Packed Signed Doubleword Integers and Store Low Result.
func (amd64 *Amd64) VPMULLD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMULLD", r1, r2, r3)
}

// VPMULLD_BCST: Multiply Packed Signed Doubleword Integers and Store Low Result (Broadcast).
func (amd64 *Amd64) VPMULLD_BCST(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMULLD.BCST", r1, r2, r3)
}

// VPMULLQ: Multiply Packed Signed Quadword Integers and Store Low Result.
func (amd64 *Amd64) VPMULLQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMULLQ", r1, r2, r3)
}

// VPMULLQ_BCST: Multiply Packed Signed Quadword Integers and Store Low Result (Broadcast).
func (amd64 *Amd64) VPMULLQ_BCST(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMULLQ.BCST", r1, r2, r3)
}

// VMOVSHDUP: Move Packed Single-FP High and Duplicate
func (amd64 *Amd64) VMOVSHDUP(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVSHDUP", r1, r2)
}

// VMOVSHDUPk: Move Packed Single-FP High and Duplicate
func (amd64 *Amd64) VMOVSHDUPk(r1, k, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVSHDUP", r1, k, r2)
}

// VPCMPUD: Compare Packed Unsigned Doubleword Values.
//
// Forms:
//
//	VPCMPUD imm8 m128 xmm k k
//	VPCMPUD imm8 m128 xmm k
//	VPCMPUD imm8 m256 ymm k k
//	VPCMPUD imm8 m256 ymm k
//	VPCMPUD imm8 xmm  xmm k k
//	VPCMPUD imm8 xmm  xmm k
//	VPCMPUD imm8 ymm  ymm k k
//	VPCMPUD imm8 ymm  ymm k
//	VPCMPUD imm8 m512 zmm k k
//	VPCMPUD imm8 m512 zmm k
//	VPCMPUD imm8 zmm  zmm k k
//	VPCMPUD imm8 zmm  zmm k
func (amd64 *Amd64) VPCMPUD(imm8, r1, r2, k interface{}, comment ...string) {
	amd64.writeOp(comment, "VPCMPUD", imm8, r1, r2, k)
}

// VPCMPLTUD: Compare Packed Unsigned Doubleword Values for Less Than.
// This is a shorthand for VPCMPUD with imm8=1 (VPCMPULT).
// Compares r1 < r2 and stores mask result in k.
//
// Forms:
//
//	VPCMPLTUD r1 r2 k
func (amd64 *Amd64) VPCMPLTUD(r1, r2, k interface{}, comment ...string) {
	amd64.writeOp(comment, "VPCMPUD", "$1", r1, r2, k)
}

// VPMADD52HUQ: Packed Multiply of Unsigned 52-bit Unsigned Integers and Add High 52-bit Products to Quadword Accumulators.
//
// Forms:
//
//	VPMADD52HUQ m128 xmm k xmm
//	VPMADD52HUQ m128 xmm xmm
//	VPMADD52HUQ m256 ymm k ymm
//	VPMADD52HUQ m256 ymm ymm
//	VPMADD52HUQ xmm  xmm k xmm
//	VPMADD52HUQ xmm  xmm xmm
//	VPMADD52HUQ ymm  ymm k ymm
//	VPMADD52HUQ ymm  ymm ymm
//	VPMADD52HUQ m512 zmm k zmm
//	VPMADD52HUQ m512 zmm zmm
//	VPMADD52HUQ zmm  zmm k zmm
//	VPMADD52HUQ zmm  zmm zmm
func (amd64 *Amd64) VPMADD52HUQ(r1, r2, r3 any, comment ...string) {
	amd64.writeOp(comment, "VPMADD52HUQ", r1, r2, r3)
}

// VPMADD52LUQ: Packed Multiply of Unsigned 52-bit Integers and Add the Low 52-bit Products to Quadword Accumulators.
func (amd64 *Amd64) VPMADD52LUQ(r1, r2, r3 any, comment ...string) {
	amd64.writeOp(comment, "VPMADD52LUQ", r1, r2, r3)
}

// VPMADD52LUQ_BCST: Packed Multiply of Unsigned 52-bit Integers and Add Low 52-bit Products (Broadcast).
func (amd64 *Amd64) VPMADD52LUQ_BCST(r1, r2, r3 any, comment ...string) {
	amd64.writeOp(comment, "VPMADD52LUQ.BCST", r1, r2, r3)
}

// VPMADD52HUQ_BCST: Packed Multiply of Unsigned 52-bit Integers and Add High 52-bit Products (Broadcast).
func (amd64 *Amd64) VPMADD52HUQ_BCST(r1, r2, r3 any, comment ...string) {
	amd64.writeOp(comment, "VPMADD52HUQ.BCST", r1, r2, r3)
}

// VPBROADCASTQ: Broadcast Quadword Integer
func (amd64 *Amd64) VPBROADCASTQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPBROADCASTQ", r1, r2)
}

// VPBROADCASTD: Broadcast Doubleword Integer
func (amd64 *Amd64) VPBROADCASTD(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPBROADCASTD", r1, r2)
}

// VPADDD: Add Packed Doubleword Integers
func (amd64 *Amd64) VPADDD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPADDD", r1, r2, r3)
}

// VPADDDk: Add Packed Doubleword Integers
func (amd64 *Amd64) VPADDDk(r1, r2, r3, k interface{}, comment ...string) {
	amd64.writeOp(comment, "VPADDD", r1, r2, k, r3)
}

// VPTESTMD zmm  zmm k
func (amd64 *Amd64) VPTESTMD(r1, r2, k interface{}, comment ...string) {
	amd64.writeOp(comment, "VPTESTMD", r1, r2, k)
}

// VPMADDWD zmm  zmm zmm
func (amd64 *Amd64) VPMADDWD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMADDWD", r1, r2, r3)
}

// VPSUBD: Subtract Packed Doubleword Integers
func (amd64 *Amd64) VPSUBD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSUBD", r1, r2, r3)
}

func (amd64 *Amd64) VPSUBDk(r1, r2, r3, k interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSUBD", r1, r2, k, r3)
}

// VPMINUD: Minimum of Packed Unsigned Doubleword Integers
func (amd64 *Amd64) VPMINUD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMINUD", r1, r2, r3)
}

func (amd64 *Amd64) VPMINUDk(r1, r2, r3, k interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMINUD", r1, r2, k, r3)
}

// VPMINUQ: Minimum of Packed Unsigned Quadword Integers
func (amd64 *Amd64) VPMINUQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMINUQ", r1, r2, r3)
}

// VPSLLQ: Shift Packed Quadword Data Left Logical
func (amd64 *Amd64) VPSLLQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSLLQ", r1, r2, r3)
}

func (amd64 *Amd64) VPSLLD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSLLD", r1, r2, r3)
}

func (amd64 *Amd64) VPSLLDk(r1, r2, k, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSLLD", r1, r2, k, r3)
}

// VPSUBQ: Subtract Packed Quadword Integers
func (amd64 *Amd64) VPSUBQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSUBQ", r1, r2, r3)
}

// VPSRAQ: Shift Packed Quadword Integers Right Arithmetic
// Performs arithmetic (signed) right shift on packed 64-bit integers.
//
// Forms:
//
//	VPSRAQ imm8, zmm, zmm
//	VPSRAQ zmm, zmm, zmm
func (amd64 *Amd64) VPSRAQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSRAQ", r1, r2, r3)
}

// VPANDNQ: Bitwise AND NOT of Packed Quadword Integers
// Performs bitwise AND NOT: dst = NOT(src1) AND src2
//
// Forms:
//
//	VPANDNQ zmm, zmm, zmm
func (amd64 *Amd64) VPANDNQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPANDNQ", r1, r2, r3)
}

// KNOTB: NOT 8-bit Mask Register
func (amd64 *Amd64) KNOTB(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KNOTB", r1, r2)
}

// VSHUFI64X2: Shuffle 128-Bit Packed Quadword Integer Values
func (amd64 *Amd64) VSHUFI64X2(r1, r2, r3, r4 interface{}, comment ...string) {
	amd64.writeOp(comment, "VSHUFI64X2", r1, r2, r3, r4)
}

// VPERMQ: Permute Quadword Integers
func (amd64 *Amd64) VPERMQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPERMQ", r1, r2, r3)
}

// VEXTRACTI32X8: Extract 256 Bits of Packed Doubleword Integer Values.
//
// Forms:
//
//	VEXTRACTI32X8 imm8 zmm k m256
//	VEXTRACTI32X8 imm8 zmm k ymm
//	VEXTRACTI32X8 imm8 zmm m256
//	VEXTRACTI32X8 imm8 zmm ymm
func (amd64 *Amd64) VEXTRACTI32X8(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VEXTRACTI32X8", imm8, r1, r2)
}

// VPBLENDMQ: Blend Quadword Vectors Using an OpMask Control.
//
// Forms:
//
//	VPBLENDMQ m128 xmm k xmm
//	VPBLENDMQ m128 xmm xmm
//	VPBLENDMQ m256 ymm k ymm
//	VPBLENDMQ m256 ymm ymm
//	VPBLENDMQ xmm  xmm k xmm
//	VPBLENDMQ xmm  xmm xmm
//	VPBLENDMQ ymm  ymm k ymm
//	VPBLENDMQ ymm  ymm ymm
//	VPBLENDMQ m512 zmm k zmm
//	VPBLENDMQ m512 zmm zmm
//	VPBLENDMQ zmm  zmm k zmm
//	VPBLENDMQ zmm  zmm zmm
func (amd64 *Amd64) VPBLENDMQ(r1, r2, r3, k interface{}, comment ...string) {
	amd64.writeOp(comment, "VPBLENDMQ", r1, r2, k, r3)
}

// VEXTRACTI64X4: Extract 256 Bits of Packed Quadword Integer Values.
//
// Forms:
//
//	VEXTRACTI64X4 imm8 zmm k m256
//	VEXTRACTI64X4 imm8 zmm k ymm
//	VEXTRACTI64X4 imm8 zmm m256
//	VEXTRACTI64X4 imm8 zmm ymm
func (amd64 *Amd64) VEXTRACTI64X4(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VEXTRACTI64X4", imm8, r1, r2)
}

// VPSHRDQ: Concatenate Quadwords and Shift Packed Data Right Logical.
//
// Forms:
//
//	VPSHRDQ imm8 m128 xmm k xmm
//	VPSHRDQ imm8 m128 xmm xmm
//	VPSHRDQ imm8 m256 ymm k ymm
//	VPSHRDQ imm8 m256 ymm ymm
//	VPSHRDQ imm8 xmm  xmm k xmm
//	VPSHRDQ imm8 xmm  xmm xmm
//	VPSHRDQ imm8 ymm  ymm k ymm
//	VPSHRDQ imm8 ymm  ymm ymm
//	VPSHRDQ imm8 m512 zmm k zmm
//	VPSHRDQ imm8 m512 zmm zmm
//	VPSHRDQ imm8 zmm  zmm k zmm
//	VPSHRDQ imm8 zmm  zmm zmm
func (amd64 *Amd64) VPSHRDQ(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSHRDQ", imm8, r1, r2, r3)
}

// VPBLENDMD: Blend Doubleword Vectors Using an OpMask Control.
//
// Forms:
//
//	VPBLENDMD m128 xmm k xmm
//	VPBLENDMD m128 xmm xmm
//	VPBLENDMD m256 ymm k ymm
//	VPBLENDMD m256 ymm ymm
//	VPBLENDMD xmm  xmm k xmm
//	VPBLENDMD xmm  xmm xmm
//	VPBLENDMD ymm  ymm k ymm
//	VPBLENDMD ymm  ymm ymm
//	VPBLENDMD m512 zmm k zmm
//	VPBLENDMD m512 zmm zmm
//	VPBLENDMD zmm  zmm k zmm
//	VPBLENDMD zmm  zmm zmm
func (amd64 *Amd64) VPBLENDMD(r1, r2, r3, k interface{}, comment ...string) {
	amd64.writeOp(comment, "VPBLENDMD", r1, r2, k, r3)
}

// VPBLENDD
func (amd64 *Amd64) VPBLENDD(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPBLENDD", imm8, r1, r2, r3)
}

// VEXTRACTI64X2
func (amd64 *Amd64) VEXTRACTI64X2(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VEXTRACTI64X2", imm8, r1, r2)
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

// KMOVQ Move 64-bit Mask
func (amd64 *Amd64) KMOVQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KMOVQ", r1, r2)
}

// VPINSRQ: Insert Quadword.
//
// Forms:
//
//	VPINSRQ imm8 m64 xmm xmm
//	VPINSRQ imm8 r64 xmm xmm
func (amd64 *Amd64) VPINSRQ(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPINSRQ", imm8, r1, r2, r3)
}

// VPINSRD: Insert Doubleword.
func (amd64 *Amd64) VPINSRD(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPINSRD", imm8, r1, r2, r3)
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

// VPORQ: Bitwise Logical OR of Packed Quadword Integers
func (amd64 *Amd64) VPORQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPORQ", r1, r2, r3)
}

// VPERMT2Q: Full Permute of Quadwords From Two Tables Overwriting a Table
func (amd64 *Amd64) VPERMT2Q(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPERMT2Q", r1, r2, r3)
}

// VPMOVQ2M: Move Signs of Packed Quadword Integers to Mask Register
func (amd64 *Amd64) VPMOVQ2M(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMOVQ2M", r1, r2)
}

// VMOVDQA32: Move Aligned Doubleword Values
func (amd64 *Amd64) VMOVDQA32(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVDQA32", r1, r2)
}

// VPSRLD: Shift Packed Doubleword Data Right Logical.
func (amd64 *Amd64) VPSRLD(imm8, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSRLD", imm8, r2, r3)
}

// VPSHUFLW: Shuffle Packed Low Words.
func (amd64 *Amd64) VPSHUFLW(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSHUFLW", imm8, r1, r2)
}

// VPSHUFHW: Shuffle Packed High Words.
func (amd64 *Amd64) VPSHUFHW(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSHUFHW", imm8, r1, r2)
}

// VPMOVDW: Down Convert Packed Doubleword Values to Word Values with Truncation.
//
// Forms:
//
//	VPMOVDW xmm k m64
//	VPMOVDW xmm k xmm
//	VPMOVDW xmm m64
//	VPMOVDW xmm xmm
//	VPMOVDW ymm k m128
//	VPMOVDW ymm k xmm
//	VPMOVDW ymm m128
//	VPMOVDW ymm xmm
//	VPMOVDW zmm k m256
//	VPMOVDW zmm k ymm
//	VPMOVDW zmm m256
//	VPMOVDW zmm ymm
func (amd64 *Amd64) VPMOVDW(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMOVDW", r1, r2)
}

// VMOVDQA64 Move Aligned Quadword Values
func (amd64 *Amd64) VMOVDQA64(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVDQA64", r1, r2)
}

// VMOVDQA64_Z Move Aligned Quadword Values  (Zeroing Masking).
func (amd64 *Amd64) VMOVDQA64_Z(r1, k, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVDQA64.Z", r1, k, r2)
}

// VMOVDQA32_Z Move Aligned Quadword Values  (Zeroing Masking).
func (amd64 *Amd64) VMOVDQA32_Z(r1, k, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVDQA32.Z", r1, k, r2)
}

// VPMOVQD: Down Convert Packed Quadword Values to Doubleword Values with Truncation.
func (amd64 *Amd64) VPMOVQD(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMOVQD", r1, r2)
}

// VPMOVZXDQ Move Packed Doubleword Integers to Quadword Integers
func (amd64 *Amd64) VPMOVZXDQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMOVZXDQ", r1, r2)
}

// VPMOVZXWD: Move Packed Word Integers to Doubleword Integers with Zero Extension.
//
// Forms:
//
//	VPMOVZXWD m128 ymm
//	VPMOVZXWD xmm  ymm
//	VPMOVZXWD m64  xmm
//	VPMOVZXWD xmm  xmm
//	VPMOVZXWD m128 k ymm
//	VPMOVZXWD m64  k xmm
//	VPMOVZXWD xmm  k xmm
//	VPMOVZXWD xmm  k ymm
//	VPMOVZXWD m256 k zmm
//	VPMOVZXWD m256 zmm
//	VPMOVZXWD ymm  k zmm
//	VPMOVZXWD ymm  zmm
func (amd64 *Amd64) VPMOVZXWD(r1, r2 VectorRegister, comment ...string) {
	amd64.writeOp(comment, "VPMOVZXWD", r1, r2)
}

// VPSHUFD: Shuffle Packed Doublewords.
//
// Forms:
//
//	VPSHUFD imm8 m256 ymm
//	VPSHUFD imm8 ymm  ymm
//	VPSHUFD imm8 m128 xmm
//	VPSHUFD imm8 xmm  xmm
//	VPSHUFD imm8 m128 k xmm
//	VPSHUFD imm8 m256 k ymm
//	VPSHUFD imm8 xmm  k xmm
//	VPSHUFD imm8 ymm  k ymm
//	VPSHUFD imm8 m512 k zmm
//	VPSHUFD imm8 m512 zmm
//	VPSHUFD imm8 zmm  k zmm
//	VPSHUFD imm8 zmm  zmm
func (amd64 *Amd64) VPSHUFD(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSHUFD", imm8, r1, r2)
}

// VSHUFPD: Shuffle Packed Double-Precision Floating-Point Values.
//
// Forms:
//
//	VSHUFPD imm8 m128 xmm xmm
//	VSHUFPD imm8 m256 ymm ymm
//	VSHUFPD imm8 xmm  xmm xmm
//	VSHUFPD imm8 ymm  ymm ymm
//	VSHUFPD imm8 m128 xmm k xmm
//	VSHUFPD imm8 m256 ymm k ymm
//	VSHUFPD imm8 xmm  xmm k xmm
//	VSHUFPD imm8 ymm  ymm k ymm
//	VSHUFPD imm8 m512 zmm k zmm
//	VSHUFPD imm8 m512 zmm zmm
//	VSHUFPD imm8 zmm  zmm k zmm
//	VSHUFPD imm8 zmm  zmm zmm
func (amd64 *Amd64) VSHUFPD(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VSHUFPD", imm8, r1, r2, r3)
}

// VPUNPCKLDQ: Unpack and Interleave Low-Order Doublewords into Quadwords.
//
// Forms:
//
//	VPUNPCKLDQ m256 ymm ymm
//	VPUNPCKLDQ ymm  ymm ymm
//	VPUNPCKLDQ m128 xmm xmm
//	VPUNPCKLDQ xmm  xmm xmm
//	VPUNPCKLDQ m128 xmm k xmm
//	VPUNPCKLDQ m256 ymm k ymm
//	VPUNPCKLDQ xmm  xmm k xmm
//	VPUNPCKLDQ ymm  ymm k ymm
//	VPUNPCKLDQ m512 zmm k zmm
//	VPUNPCKLDQ m512 zmm zmm
//	VPUNPCKLDQ zmm  zmm k zmm
//	VPUNPCKLDQ zmm  zmm zmm
func (amd64 *Amd64) VPUNPCKLDQ(r1, r2, r3 VectorRegister, comment ...string) {
	amd64.writeOp(comment, "VPUNPCKLDQ", r1, r2, r3)
}

// VPUNPCKHDQ: Unpack and Interleave High-Order Doublewords into Quadwords.
//
// Forms:
//
//	VPUNPCKHDQ m256 ymm ymm
//	VPUNPCKHDQ ymm  ymm ymm
//	VPUNPCKHDQ m128 xmm xmm
//	VPUNPCKHDQ xmm  xmm xmm
//	VPUNPCKHDQ m128 xmm k xmm
//	VPUNPCKHDQ m256 ymm k ymm
//	VPUNPCKHDQ xmm  xmm k xmm
//	VPUNPCKHDQ ymm  ymm k ymm
//	VPUNPCKHDQ m512 zmm k zmm
//	VPUNPCKHDQ m512 zmm zmm
//	VPUNPCKHDQ zmm  zmm k zmm
//	VPUNPCKHDQ zmm  zmm zmm
func (amd64 *Amd64) VPUNPCKHDQ(r1, r2, r3 VectorRegister, comment ...string) {
	amd64.writeOp(comment, "VPUNPCKHDQ", r1, r2, r3)
}

// VSHUFF64X2: Shuffle 128-Bit Packed Double-Precision Floating-Point Values.
//
// Forms:
//
//	VSHUFF64X2 imm8 m256 ymm k ymm
//	VSHUFF64X2 imm8 m256 ymm ymm
//	VSHUFF64X2 imm8 ymm  ymm k ymm
//	VSHUFF64X2 imm8 ymm  ymm ymm
//	VSHUFF64X2 imm8 m512 zmm k zmm
//	VSHUFF64X2 imm8 m512 zmm zmm
//	VSHUFF64X2 imm8 zmm  zmm k zmm
//	VSHUFF64X2 imm8 zmm  zmm zmm
func (amd64 *Amd64) VSHUFF64X2(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VSHUFF64X2", imm8, r1, r2, r3)
}

// VSHUFF32X4: Shuffle 128-Bit Packed Single-Precision Floating-Point Values.
//
// Forms:
//
//	VSHUFF32X4 imm8 m256 ymm k ymm
//	VSHUFF32X4 imm8 m256 ymm ymm
//	VSHUFF32X4 imm8 ymm  ymm k ymm
//	VSHUFF32X4 imm8 ymm  ymm ymm
//	VSHUFF32X4 imm8 m512 zmm k zmm
//	VSHUFF32X4 imm8 m512 zmm zmm
//	VSHUFF32X4 imm8 zmm  zmm k zmm
//	VSHUFF32X4 imm8 zmm  zmm zmm
func (amd64 *Amd64) VSHUFF32X4(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VSHUFF32X4", imm8, r1, r2, r3)
}

// VINSERTI64X4: Insert 256 Bits of Packed Quadword Integer Values.
//
// Forms:
//
//	VINSERTI64X4 imm8 m256 zmm k zmm
//	VINSERTI64X4 imm8 m256 zmm zmm
//	VINSERTI64X4 imm8 ymm  zmm k zmm
//	VINSERTI64X4 imm8 ymm  zmm zmm
func (amd64 *Amd64) VINSERTI64X4(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VINSERTI64X4", imm8, r1, r2, r3)
}

// VINSERTI64X2: Insert 128 Bits of Packed Quadword Integer Values.
//
// Forms:
//
//	VINSERTI64X2 imm8 m128 ymm k ymm
//	VINSERTI64X2 imm8 m128 ymm ymm
//	VINSERTI64X2 imm8 xmm  ymm k ymm
//	VINSERTI64X2 imm8 xmm  ymm ymm
//	VINSERTI64X2 imm8 m128 zmm k zmm
//	VINSERTI64X2 imm8 m128 zmm zmm
//	VINSERTI64X2 imm8 xmm  zmm k zmm
//	VINSERTI64X2 imm8 xmm  zmm zmm
func (amd64 *Amd64) VINSERTI64X2(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VINSERTI64X2", imm8, r1, r2, r3)
}

// VMOVDQU32: Move Unaligned Doubleword Values
func (amd64 *Amd64) VMOVDQU32(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVDQU32", r1, r2)
}

// VMOVDQU64 Move Unaligned Quadword Values
func (amd64 *Amd64) VMOVDQU64(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVDQU64", r1, r2)
}

// VPGATHERDD: Gather Packed Doubleword Values Using Signed Doubleword Indices.
// example: VPGATHERDD  8(AX)(Z18*4), K7, Z6
func (amd64 *Amd64) VPGATHERDD(baseAddrOffset int, baseAddr Register, indices VectorRegister, scale int, mask MaskRegister, res VectorRegister, comment ...string) {
	amd64.writeOp(comment, "VPGATHERDD", fmt.Sprintf("%d(%s)(%s*%d)", baseAddrOffset, baseAddr, indices, scale), mask, res)
}

// VPSCATTERDD: Scatter Packed Doubleword Values with Signed Doubleword Indices.
func (amd64 *Amd64) VPSCATTERDD(baseAddrOffset int, baseAddr Register, indices VectorRegister, scale int, mask MaskRegister, src VectorRegister, comment ...string) {
	amd64.writeOp(comment, "VPSCATTERDD", src, mask, fmt.Sprintf("%d(%s)(%s*%d)", baseAddrOffset, baseAddr, indices, scale))
}

// VPERMI2Q: Full Permute of Quadwords From Two Tables Overwriting the Index.
//
// Forms:
//
//	VPERMI2Q m128 xmm k xmm
//	VPERMI2Q m128 xmm xmm
//	VPERMI2Q m256 ymm k ymm
//	VPERMI2Q m256 ymm ymm
//	VPERMI2Q xmm  xmm k xmm
//	VPERMI2Q xmm  xmm xmm
//	VPERMI2Q ymm  ymm k ymm
//	VPERMI2Q ymm  ymm ymm
//	VPERMI2Q m512 zmm k zmm
//	VPERMI2Q m512 zmm zmm
//	VPERMI2Q zmm  zmm k zmm
//	VPERMI2Q zmm  zmm zmm
func (amd64 *Amd64) VPERMI2Q(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPERMI2Q", r1, r2, r3)
}

// VPERMI2D: Full Permute of Doublewords From Two Tables Overwriting the Index.
func (amd64 *Amd64) VPERMI2D(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPERMI2D", r1, r2, r3)
}

// VMOVDQU64 Move Unaligned Quadword Values
func (amd64 *Amd64) VMOVDQU64k(r1, k, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VMOVDQU64", r1, k, r2)
}

// VPTERNLOGD: Bitwise Ternary Logical Operation on Doubleword Values.
//
// Forms:
//
//	VPTERNLOGD imm8 m128 xmm k xmm
//	VPTERNLOGD imm8 m128 xmm xmm
//	VPTERNLOGD imm8 m256 ymm k ymm
//	VPTERNLOGD imm8 m256 ymm ymm
//	VPTERNLOGD imm8 xmm  xmm k xmm
//	VPTERNLOGD imm8 xmm  xmm xmm
//	VPTERNLOGD imm8 ymm  ymm k ymm
//	VPTERNLOGD imm8 ymm  ymm ymm
//	VPTERNLOGD imm8 m512 zmm k zmm
//	VPTERNLOGD imm8 m512 zmm zmm
//	VPTERNLOGD imm8 zmm  zmm k zmm
//	VPTERNLOGD imm8 zmm  zmm zmm
func (amd64 *Amd64) VPTERNLOGD(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPTERNLOGD", imm8, r1, r2, r3)
}

// VPADDQ Add Packed Quadword Integers
func (amd64 *Amd64) VPADDQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPADDQ", r1, r2, r3)
}

func (amd64 *Amd64) VPADDQk(r1, r2, r3, k interface{}, comment ...string) {
	amd64.writeOp(comment, "VPADDQ", r1, r2, k, r3)
}

// VPMULUDQ Multiply Packed Unsigned Doubleword Integers
func (amd64 *Amd64) VPMULUDQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMULUDQ", r1, r2, r3)
}

func (amd64 *Amd64) VPMULUDQk(r1, r2, r3, k interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMULUDQ", r1, r2, k, r3)
}

// VPMULUDQ_BCST Multiply Packed Unsigned Doubleword Integers (Broadcast).
func (amd64 *Amd64) VPMULUDQ_BCST(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMULUDQ.BCST", r1, r2, r3)
}

// VPANDQ Bitwise Logical AND of Packed Quadword Integers
func (amd64 *Amd64) VPANDQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPANDQ", r1, r2, r3)
}

func (amd64 *Amd64) VPANDD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPANDD", r1, r2, r3)
}

func (amd64 *Amd64) VPANDDk(r1, r2, k, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPANDD", r1, r2, k, r3)
}

func (amd64 *Amd64) VPANDDkz(r1, r2, k, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPANDD.Z", r1, r2, k, r3)
}

// VPSRLQ Shift Packed Quadword Data Right Logical
func (amd64 *Amd64) VPSRLQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSRLQ", r1, r2, r3)
}

func (amd64 *Amd64) VPSRLQk(r1, r2, r3, k interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSRLQ", r1, r2, k, r3)
}

// VPEXTRQ Extract Quadword
func (amd64 *Amd64) VPEXTRQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPEXTRQ", r1, r2, r3)
}

// VALIGND concatenates and shifts doubleword values from two sources.
// Concatenates src2:src1, shifts right by imm8 doublewords, and stores the low part in dst.
// This is useful for byte-level data alignment and rotation operations.
//
// Forms:
//
//	VALIGND imm8, zmm, zmm, zmm           // 4-operand form (no mask)
//	VALIGND imm8, zmm, zmm, k, zmm        // 5-operand form (with mask)
//
// Operation (4-operand form, 512-bit):
//
//	temp[1023:0] = src2[511:0]:src1[511:0]
//	dst[511:0] = temp[imm8*32+511:imm8*32]
//
// Example:
//
//	VALIGND $0, Z15, Z11, Z11  // effectively copies Z11 elements shifted by Z15 position
func (amd64 *Amd64) VALIGND(r1, r2, r3, r4 interface{}, r5 ...interface{}) {
	if len(r5) == 0 {
		// 4-operand form: imm8, src2, src1, dst
		amd64.writeOp(nil, "VALIGND", r1, r2, r3, r4)
	} else {
		// 5-operand form: imm8, src2, src1, k, dst
		comment, _ := r5[0].(string)
		var comments []string
		if comment != "" {
			comments = []string{comment}
		}
		amd64.writeOp(comments, "VALIGND", r1, r2, r3, r4, r5[0])
	}
}

// VALIGNDk concatenates and shifts doubleword values with mask.
// 5-operand form with explicit mask parameter.
//
// Forms:
//
//	VALIGND imm8, zmm, zmm, k, zmm
func (amd64 *Amd64) VALIGNDk(imm8, src2, src1, k, dst interface{}, comment ...string) {
	amd64.writeOp(comment, "VALIGND", imm8, src2, src1, k, dst)
}

// VALIGND_Z concatenates and shifts with zeroing masking.
// Elements not selected by the mask are zeroed.
//
// Forms:
//
//	VALIGND.Z imm8, zmm, zmm, k, zmm
func (amd64 *Amd64) VALIGND_Z(imm8, src2, src1, k, dst interface{}, comment ...string) {
	amd64.writeOp(comment, "VALIGND.Z", imm8, src2, src1, k, dst)
}

// VALIGNQ concatenates and shifts quadword values from two sources.
// Concatenates src2:src1, shifts right by imm8 quadwords, and stores the low part in dst.
//
// Forms:
//
//	VALIGNQ imm8, zmm, zmm, zmm
//
// Operation (512-bit):
//
//	temp[1023:0] = src2[511:0]:src1[511:0]
//	dst[511:0] = temp[imm8*64+511:imm8*64]
func (amd64 *Amd64) VALIGNQ(imm8, src2, src1, dst interface{}, comment ...string) {
	amd64.writeOp(comment, "VALIGNQ", imm8, src2, src1, dst)
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

// PEXTRQ: Extract Quadword.
//
// Forms:
//
//	PEXTRQ imm8 xmm m64
//	PEXTRQ imm8 xmm r64
func (amd64 *Amd64) PEXTRQ(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "PEXTRQ", imm8, r1, r2)
}

// PEXTRD: Extract Doubleword.
func (amd64 *Amd64) PEXTRD(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "PEXTRD", imm8, r1, r2)
}

// KORTESTQ: OR 64-bit Masks and Set Flags
func (amd64 *Amd64) KORTESTQ(r1, r2 interface{}, comment ...string) {
        amd64.writeOp(comment, "KORTESTQ", r1, r2)
}

// VPCMPUQ: Compare Packed Unsigned Quadword Values.
func (amd64 *Amd64) VPCMPUQ(imm8, r1, r2, k interface{}, comment ...string) {
        amd64.writeOp(comment, "VPCMPUQ", imm8, r1, r2, k)
}

// VPTESTMQ: Packed Quadword Test Mask
func (amd64 *Amd64) VPTESTMQ(r1, r2, k interface{}, comment ...string) {
        amd64.writeOp(comment, "VPTESTMQ", r1, r2, k)
}

// VPUNPCKLQDQ: Unpack and Interleave Low-Order Quadwords
func (amd64 *Amd64) VPUNPCKLQDQ(r1, r2, r3 interface{}, comment ...string) {
        amd64.writeOp(comment, "VPUNPCKLQDQ", r1, r2, r3)
}

// VPUNPCKHQDQ: Unpack and Interleave High-Order Quadwords
func (amd64 *Amd64) VPUNPCKHQDQ(r1, r2, r3 interface{}, comment ...string) {
        amd64.writeOp(comment, "VPUNPCKHQDQ", r1, r2, r3)
}

// KMOVB Move 8-bit Mask
func (amd64 *Amd64) KMOVB(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KMOVB", r1, r2)
}

// -----------------------------------------------------------------------------
// Additional Mask Register Operations
// -----------------------------------------------------------------------------

// KNOTW performs bitwise NOT on 16-bit mask register.
//
// Forms:
//
//	KNOTW k, k
func (amd64 *Amd64) KNOTW(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KNOTW", r1, r2)
}

// KNOTD performs bitwise NOT on 32-bit mask register.
//
// Forms:
//
//	KNOTD k, k
func (amd64 *Amd64) KNOTD(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KNOTD", r1, r2)
}

// KNOTQ performs bitwise NOT on 64-bit mask register.
//
// Forms:
//
//	KNOTQ k, k
func (amd64 *Amd64) KNOTQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KNOTQ", r1, r2)
}

// KANDW performs bitwise AND on 16-bit mask registers.
//
// Forms:
//
//	KANDW k, k, k
func (amd64 *Amd64) KANDW(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KANDW", r1, r2, r3)
}

// KANDD performs bitwise AND on 32-bit mask registers.
//
// Forms:
//
//	KANDD k, k, k
func (amd64 *Amd64) KANDD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KANDD", r1, r2, r3)
}

// KANDQ performs bitwise AND on 64-bit mask registers.
//
// Forms:
//
//	KANDQ k, k, k
func (amd64 *Amd64) KANDQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KANDQ", r1, r2, r3)
}

// KORW performs bitwise OR on 16-bit mask registers.
//
// Forms:
//
//	KORW k, k, k
func (amd64 *Amd64) KORW(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KORW", r1, r2, r3)
}

// KORD performs bitwise OR on 32-bit mask registers.
//
// Forms:
//
//	KORD k, k, k
func (amd64 *Amd64) KORD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KORD", r1, r2, r3)
}

// KORQ performs bitwise OR on 64-bit mask registers.
//
// Forms:
//
//	KORQ k, k, k
func (amd64 *Amd64) KORQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KORQ", r1, r2, r3)
}

// KXORW performs bitwise XOR on 16-bit mask registers.
//
// Forms:
//
//	KXORW k, k, k
func (amd64 *Amd64) KXORW(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KXORW", r1, r2, r3)
}

// KXORD performs bitwise XOR on 32-bit mask registers.
//
// Forms:
//
//	KXORD k, k, k
func (amd64 *Amd64) KXORD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KXORD", r1, r2, r3)
}

// KXORQ performs bitwise XOR on 64-bit mask registers.
//
// Forms:
//
//	KXORQ k, k, k
func (amd64 *Amd64) KXORQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KXORQ", r1, r2, r3)
}

// KSHIFTRW shifts 16-bit mask right by immediate.
//
// Forms:
//
//	KSHIFTRW imm8, k, k
func (amd64 *Amd64) KSHIFTRW(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KSHIFTRW", imm8, r1, r2)
}

// KSHIFTRD shifts 32-bit mask right by immediate.
//
// Forms:
//
//	KSHIFTRD imm8, k, k
func (amd64 *Amd64) KSHIFTRD(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KSHIFTRD", imm8, r1, r2)
}

// KSHIFTRQ shifts 64-bit mask right by immediate.
//
// Forms:
//
//	KSHIFTRQ imm8, k, k
func (amd64 *Amd64) KSHIFTRQ(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KSHIFTRQ", imm8, r1, r2)
}

// KSHIFTLD shifts 32-bit mask left by immediate.
//
// Forms:
//
//	KSHIFTLD imm8, k, k
func (amd64 *Amd64) KSHIFTLD(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KSHIFTLD", imm8, r1, r2)
}

// KSHIFTLQ shifts 64-bit mask left by immediate.
//
// Forms:
//
//	KSHIFTLQ imm8, k, k
func (amd64 *Amd64) KSHIFTLQ(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KSHIFTLQ", imm8, r1, r2)
}

// KADDB adds 8-bit masks.
//
// Forms:
//
//	KADDB k, k, k
func (amd64 *Amd64) KADDB(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KADDB", r1, r2, r3)
}

// KADDD adds 32-bit masks.
//
// Forms:
//
//	KADDD k, k, k
func (amd64 *Amd64) KADDD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KADDD", r1, r2, r3)
}

// KADDQ adds 64-bit masks.
//
// Forms:
//
//	KADDQ k, k, k
func (amd64 *Amd64) KADDQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "KADDQ", r1, r2, r3)
}

// KTESTB tests 8-bit masks and sets flags.
//
// Forms:
//
//	KTESTB k, k
func (amd64 *Amd64) KTESTB(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KTESTB", r1, r2)
}

// KTESTW tests 16-bit masks and sets flags.
//
// Forms:
//
//	KTESTW k, k
func (amd64 *Amd64) KTESTW(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KTESTW", r1, r2)
}

// KTESTD tests 32-bit masks and sets flags.
//
// Forms:
//
//	KTESTD k, k
func (amd64 *Amd64) KTESTD(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KTESTD", r1, r2)
}

// KTESTQ tests 64-bit masks and sets flags.
//
// Forms:
//
//	KTESTQ k, k
func (amd64 *Amd64) KTESTQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KTESTQ", r1, r2)
}

// KORTESTW: OR 16-bit Masks and Set Flags
func (amd64 *Amd64) KORTESTW(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KORTESTW", r1, r2)
}

// KORTESTD: OR 32-bit Masks and Set Flags
func (amd64 *Amd64) KORTESTD(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "KORTESTD", r1, r2)
}

// -----------------------------------------------------------------------------
// Additional AVX-512 Vector Operations
// -----------------------------------------------------------------------------

// VPSLLDQ shifts 128-bit lanes left by bytes (byte granularity).
// Shifts the destination operand left by the number of bytes specified in the count operand.
//
// Forms:
//
//	VPSLLDQ imm8, zmm, zmm
func (amd64 *Amd64) VPSLLDQ(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSLLDQ", imm8, r1, r2)
}

// VPSRLDQ shifts 128-bit lanes right by bytes (byte granularity).
//
// Forms:
//
//	VPSRLDQ imm8, zmm, zmm
func (amd64 *Amd64) VPSRLDQ(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSRLDQ", imm8, r1, r2)
}

// VPSLLVD shifts packed doublewords left by variable amounts.
// Each element is shifted by the corresponding shift count in the shift vector.
//
// Forms:
//
//	VPSLLVD zmm, zmm, zmm
func (amd64 *Amd64) VPSLLVD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSLLVD", r1, r2, r3)
}

// VPSLLVQ shifts packed quadwords left by variable amounts.
//
// Forms:
//
//	VPSLLVQ zmm, zmm, zmm
func (amd64 *Amd64) VPSLLVQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSLLVQ", r1, r2, r3)
}

// VPSRLVD shifts packed doublewords right logically by variable amounts.
//
// Forms:
//
//	VPSRLVD zmm, zmm, zmm
func (amd64 *Amd64) VPSRLVD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSRLVD", r1, r2, r3)
}

// VPSRLVQ shifts packed quadwords right logically by variable amounts.
//
// Forms:
//
//	VPSRLVQ zmm, zmm, zmm
func (amd64 *Amd64) VPSRLVQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSRLVQ", r1, r2, r3)
}

// VPSRAVD shifts packed doublewords right arithmetically by variable amounts.
//
// Forms:
//
//	VPSRAVD zmm, zmm, zmm
func (amd64 *Amd64) VPSRAVD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSRAVD", r1, r2, r3)
}

// VPSRAVQ shifts packed quadwords right arithmetically by variable amounts.
//
// Forms:
//
//	VPSRAVQ zmm, zmm, zmm
func (amd64 *Amd64) VPSRAVQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSRAVQ", r1, r2, r3)
}

// VPROLD rotates packed doublewords left by immediate.
//
// Forms:
//
//	VPROLD imm8, zmm, zmm
func (amd64 *Amd64) VPROLD(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPROLD", imm8, r1, r2)
}

// VPROLQ rotates packed quadwords left by immediate.
//
// Forms:
//
//	VPROLQ imm8, zmm, zmm
func (amd64 *Amd64) VPROLQ(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPROLQ", imm8, r1, r2)
}

// VPRORD rotates packed doublewords right by immediate.
//
// Forms:
//
//	VPRORD imm8, zmm, zmm
func (amd64 *Amd64) VPRORD(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPRORD", imm8, r1, r2)
}

// VPRORQ rotates packed quadwords right by immediate.
//
// Forms:
//
//	VPRORQ imm8, zmm, zmm
func (amd64 *Amd64) VPRORQ(imm8, r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPRORQ", imm8, r1, r2)
}

// VPROLVD rotates packed doublewords left by variable amounts.
//
// Forms:
//
//	VPROLVD zmm, zmm, zmm
func (amd64 *Amd64) VPROLVD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPROLVD", r1, r2, r3)
}

// VPROLVQ rotates packed quadwords left by variable amounts.
//
// Forms:
//
//	VPROLVQ zmm, zmm, zmm
func (amd64 *Amd64) VPROLVQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPROLVQ", r1, r2, r3)
}

// VPRORVD rotates packed doublewords right by variable amounts.
//
// Forms:
//
//	VPRORVD zmm, zmm, zmm
func (amd64 *Amd64) VPRORVD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPRORVD", r1, r2, r3)
}

// VPRORVQ rotates packed quadwords right by variable amounts.
//
// Forms:
//
//	VPRORVQ zmm, zmm, zmm
func (amd64 *Amd64) VPRORVQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPRORVQ", r1, r2, r3)
}

// -----------------------------------------------------------------------------
// Additional AVX-512 Comparison and Selection
// -----------------------------------------------------------------------------

// VPMAXUD computes maximum of packed unsigned doublewords.
//
// Forms:
//
//	VPMAXUD zmm, zmm, zmm
func (amd64 *Amd64) VPMAXUD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMAXUD", r1, r2, r3)
}

// VPMAXUQ computes maximum of packed unsigned quadwords.
//
// Forms:
//
//	VPMAXUQ zmm, zmm, zmm
func (amd64 *Amd64) VPMAXUQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMAXUQ", r1, r2, r3)
}

// VPMAXSD computes maximum of packed signed doublewords.
//
// Forms:
//
//	VPMAXSD zmm, zmm, zmm
func (amd64 *Amd64) VPMAXSD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMAXSD", r1, r2, r3)
}

// VPMAXSQ computes maximum of packed signed quadwords.
//
// Forms:
//
//	VPMAXSQ zmm, zmm, zmm
func (amd64 *Amd64) VPMAXSQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMAXSQ", r1, r2, r3)
}

// VPMINSD computes minimum of packed signed doublewords.
//
// Forms:
//
//	VPMINSD zmm, zmm, zmm
func (amd64 *Amd64) VPMINSD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMINSD", r1, r2, r3)
}

// VPMINSQ computes minimum of packed signed quadwords.
//
// Forms:
//
//	VPMINSQ zmm, zmm, zmm
func (amd64 *Amd64) VPMINSQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMINSQ", r1, r2, r3)
}

// VPABSD computes absolute value of packed signed doublewords.
//
// Forms:
//
//	VPABSD zmm, zmm
func (amd64 *Amd64) VPABSD(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPABSD", r1, r2)
}

// VPABSQ computes absolute value of packed signed quadwords.
//
// Forms:
//
//	VPABSQ zmm, zmm
func (amd64 *Amd64) VPABSQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPABSQ", r1, r2)
}

// -----------------------------------------------------------------------------
// Additional Shuffle/Permutation Instructions
// -----------------------------------------------------------------------------

// VSHUFI32X4 shuffles 128-bit groups of packed doubleword integers.
//
// Forms:
//
//	VSHUFI32X4 imm8, zmm, zmm, zmm
func (amd64 *Amd64) VSHUFI32X4(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VSHUFI32X4", imm8, r1, r2, r3)
}

// VPERMT2D performs full permute of doublewords from two tables.
//
// Forms:
//
//	VPERMT2D zmm, zmm, zmm
func (amd64 *Amd64) VPERMT2D(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPERMT2D", r1, r2, r3)
}

// VPERMW permutes packed words.
//
// Forms:
//
//	VPERMW zmm, zmm, zmm
func (amd64 *Amd64) VPERMW(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPERMW", r1, r2, r3)
}

// VPCOMPRESSD stores sparse packed doublewords into dense memory/register.
//
// Forms:
//
//	VPCOMPRESSD zmm, k, zmm
//	VPCOMPRESSD zmm, k, m512
func (amd64 *Amd64) VPCOMPRESSD(r1, k, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPCOMPRESSD", r1, k, r2)
}

// VPCOMPRESSQ stores sparse packed quadwords into dense memory/register.
//
// Forms:
//
//	VPCOMPRESSQ zmm, k, zmm
//	VPCOMPRESSQ zmm, k, m512
func (amd64 *Amd64) VPCOMPRESSQ(r1, k, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPCOMPRESSQ", r1, k, r2)
}

// VPEXPANDD loads sparse packed doublewords from dense memory/register.
//
// Forms:
//
//	VPEXPANDD zmm, k, zmm
//	VPEXPANDD m512, k, zmm
func (amd64 *Amd64) VPEXPANDD(r1, k, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPEXPANDD", r1, k, r2)
}

// VPEXPANDQ loads sparse packed quadwords from dense memory/register.
//
// Forms:
//
//	VPEXPANDQ zmm, k, zmm
//	VPEXPANDQ m512, k, zmm
func (amd64 *Amd64) VPEXPANDQ(r1, k, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPEXPANDQ", r1, k, r2)
}

// -----------------------------------------------------------------------------
// Conflict Detection Instructions (AVX-512CD)
// -----------------------------------------------------------------------------

// VPCONFLICTD detects conflicts within packed doublewords.
// Returns for each element a bitmask of elements in the same vector
// that have the same value and come before it.
//
// Forms:
//
//	VPCONFLICTD zmm, zmm
func (amd64 *Amd64) VPCONFLICTD(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPCONFLICTD", r1, r2)
}

// VPCONFLICTQ detects conflicts within packed quadwords.
//
// Forms:
//
//	VPCONFLICTQ zmm, zmm
func (amd64 *Amd64) VPCONFLICTQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPCONFLICTQ", r1, r2)
}

// VPLZCNTD counts leading zero bits of packed doublewords.
//
// Forms:
//
//	VPLZCNTD zmm, zmm
func (amd64 *Amd64) VPLZCNTD(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPLZCNTD", r1, r2)
}

// VPLZCNTQ counts leading zero bits of packed quadwords.
//
// Forms:
//
//	VPLZCNTQ zmm, zmm
func (amd64 *Amd64) VPLZCNTQ(r1, r2 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPLZCNTQ", r1, r2)
}

// -----------------------------------------------------------------------------
// AVX-512 VBMI2 Instructions (Bit Manipulation)
// -----------------------------------------------------------------------------

// VPSHLDQ concatenates and shifts packed quadwords left.
// Each quadword result is formed by concatenating the corresponding
// elements from src1 and src2, then shifting left by the count.
//
// Forms:
//
//	VPSHLDQ imm8, zmm, zmm, zmm
//
// Operation:
//
//	for each quadword element i:
//	    temp = (src1[i] << 64) | src2[i]  // 128-bit concatenation
//	    dst[i] = (temp << count)[127:64]   // extract upper 64 bits after shift
func (amd64 *Amd64) VPSHLDQ(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSHLDQ", imm8, r1, r2, r3)
}

// VPSHLDD concatenates and shifts packed doublewords left.
//
// Forms:
//
//	VPSHLDD imm8, zmm, zmm, zmm
func (amd64 *Amd64) VPSHLDD(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSHLDD", imm8, r1, r2, r3)
}

// VPSHRDD concatenates and shifts packed doublewords right.
//
// Forms:
//
//	VPSHRDD imm8, zmm, zmm, zmm
func (amd64 *Amd64) VPSHRDD(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSHRDD", imm8, r1, r2, r3)
}

// VPSHLDVQ concatenates and shifts quadwords left by variable amounts.
//
// Forms:
//
//	VPSHLDVQ zmm, zmm, zmm
func (amd64 *Amd64) VPSHLDVQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSHLDVQ", r1, r2, r3)
}

// VPSHLDVD concatenates and shifts doublewords left by variable amounts.
//
// Forms:
//
//	VPSHLDVD zmm, zmm, zmm
func (amd64 *Amd64) VPSHLDVD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSHLDVD", r1, r2, r3)
}

// VPSHRDVQ concatenates and shifts quadwords right by variable amounts.
//
// Forms:
//
//	VPSHRDVQ zmm, zmm, zmm
func (amd64 *Amd64) VPSHRDVQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSHRDVQ", r1, r2, r3)
}

// VPSHRDVD concatenates and shifts doublewords right by variable amounts.
//
// Forms:
//
//	VPSHRDVD zmm, zmm, zmm
func (amd64 *Amd64) VPSHRDVD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSHRDVD", r1, r2, r3)
}

// -----------------------------------------------------------------------------
// Gather/Scatter Instructions (additional forms)
// -----------------------------------------------------------------------------

// VPGATHERDQ gathers packed quadwords using signed dword indices.
//
// Forms:
//
//	VPGATHERDQ baseOffset(base)(index*scale), k, dst
func (amd64 *Amd64) VPGATHERDQ(baseAddrOffset int, baseAddr Register, indices VectorRegister, scale int, mask MaskRegister, res VectorRegister, comment ...string) {
	amd64.writeOp(comment, "VPGATHERDQ", fmt.Sprintf("%d(%s)(%s*%d)", baseAddrOffset, baseAddr, indices, scale), mask, res)
}

// VPGATHERQD gathers packed doublewords using signed qword indices.
//
// Forms:
//
//	VPGATHERQD baseOffset(base)(index*scale), k, dst
func (amd64 *Amd64) VPGATHERQD(baseAddrOffset int, baseAddr Register, indices VectorRegister, scale int, mask MaskRegister, res VectorRegister, comment ...string) {
	amd64.writeOp(comment, "VPGATHERQD", fmt.Sprintf("%d(%s)(%s*%d)", baseAddrOffset, baseAddr, indices, scale), mask, res)
}

// VPGATHERQQ gathers packed quadwords using signed qword indices.
//
// Forms:
//
//	VPGATHERQQ baseOffset(base)(index*scale), k, dst
func (amd64 *Amd64) VPGATHERQQ(baseAddrOffset int, baseAddr Register, indices VectorRegister, scale int, mask MaskRegister, res VectorRegister, comment ...string) {
	amd64.writeOp(comment, "VPGATHERQQ", fmt.Sprintf("%d(%s)(%s*%d)", baseAddrOffset, baseAddr, indices, scale), mask, res)
}

// VPSCATTERDQ scatters packed quadwords using signed dword indices.
//
// Forms:
//
//	VPSCATTERDQ src, k, baseOffset(base)(index*scale)
func (amd64 *Amd64) VPSCATTERDQ(baseAddrOffset int, baseAddr Register, indices VectorRegister, scale int, mask MaskRegister, src VectorRegister, comment ...string) {
	amd64.writeOp(comment, "VPSCATTERDQ", src, mask, fmt.Sprintf("%d(%s)(%s*%d)", baseAddrOffset, baseAddr, indices, scale))
}

// VPSCATTERQD scatters packed doublewords using signed qword indices.
//
// Forms:
//
//	VPSCATTERQD src, k, baseOffset(base)(index*scale)
func (amd64 *Amd64) VPSCATTERQD(baseAddrOffset int, baseAddr Register, indices VectorRegister, scale int, mask MaskRegister, src VectorRegister, comment ...string) {
	amd64.writeOp(comment, "VPSCATTERQD", src, mask, fmt.Sprintf("%d(%s)(%s*%d)", baseAddrOffset, baseAddr, indices, scale))
}

// VPSCATTERQQ scatters packed quadwords using signed qword indices.
//
// Forms:
//
//	VPSCATTERQQ src, k, baseOffset(base)(index*scale)
func (amd64 *Amd64) VPSCATTERQQ(baseAddrOffset int, baseAddr Register, indices VectorRegister, scale int, mask MaskRegister, src VectorRegister, comment ...string) {
	amd64.writeOp(comment, "VPSCATTERQQ", src, mask, fmt.Sprintf("%d(%s)(%s*%d)", baseAddrOffset, baseAddr, indices, scale))
}

// -----------------------------------------------------------------------------
// Reduction and Horizontal Operations
// -----------------------------------------------------------------------------

// VPTERNLOGQ performs bitwise ternary logical operation on quadwords.
// The operation is specified by the 8-bit immediate which encodes
// a 256-entry truth table.
//
// Forms:
//
//	VPTERNLOGQ imm8, zmm, zmm, zmm
//
// Common immediate values:
//
//	0x00: result = 0
//	0xFF: result = all 1s
//	0xF0: result = A
//	0xCC: result = B
//	0xAA: result = C
//	0x96: result = A XOR B XOR C
//	0xCA: result = (A AND B) OR (NOT(A) AND C)
func (amd64 *Amd64) VPTERNLOGQ(imm8, r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPTERNLOGQ", imm8, r1, r2, r3)
}
