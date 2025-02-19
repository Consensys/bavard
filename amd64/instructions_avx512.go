package amd64

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
func (amd64 *Amd64) VPADDDk(r1, r2, k, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPADDD", r1, r2, k, r3)
}

// VPSUBD: Subtract Packed Doubleword Integers
func (amd64 *Amd64) VPSUBD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSUBD", r1, r2, r3)
}

// VPMINUD: Minimum of Packed Unsigned Doubleword Integers
func (amd64 *Amd64) VPMINUD(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMINUD", r1, r2, r3)
}

// VPMINUQ: Minimum of Packed Unsigned Quadword Integers
func (amd64 *Amd64) VPMINUQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMINUQ", r1, r2, r3)
}

// VPSLLQ: Shift Packed Quadword Data Left Logical
func (amd64 *Amd64) VPSLLQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSLLQ", r1, r2, r3)
}

// VPSUBQ: Subtract Packed Quadword Integers
func (amd64 *Amd64) VPSUBQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPSUBQ", r1, r2, r3)
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

// VPMULUDQ Multiply Packed Unsigned Doubleword Integers
func (amd64 *Amd64) VPMULUDQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPMULUDQ", r1, r2, r3)
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

// VPEXTRQ Extract Quadword
func (amd64 *Amd64) VPEXTRQ(r1, r2, r3 interface{}, comment ...string) {
	amd64.writeOp(comment, "VPEXTRQ", r1, r2, r3)
}

// VALIGND Align Doubleword Vectors
func (amd64 *Amd64) VALIGND(r1, r2, r3, k, r4 interface{}, comment ...string) {
	amd64.writeOp(comment, "VALIGND", r1, r2, r3, k, r4)
}

// VALIGND_Z Align Doubleword Vectors (Zeroing Masking).
func (amd64 *Amd64) VALIGND_Z(r1, r2, r3, k, r4 interface{}, comment ...string) {
	amd64.writeOp(comment, "VALIGND.Z", r1, r2, r3, k, r4)
}

// VALIGNQ Align Quadword Vectors
func (amd64 *Amd64) VALIGNQ(r1, r2, r3, r4 interface{}, comment ...string) {
	amd64.writeOp(comment, "VALIGNQ", r1, r2, r3, r4)
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
