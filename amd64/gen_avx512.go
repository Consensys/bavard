//go:build ignore

// This file generates assembly test functions using bavard.
// Run with: go run generate_test.go

package main

import (
	"os"

	"github.com/consensys/bavard/amd64"
)

func main() {
	f, err := os.Create("avx512_test_amd64.s")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("//go:build amd64 && !purego\n\n")
	f.WriteString("#include \"textflag.h\"\n\n")

	asm := amd64.NewAmd64(f)

	// Generate all test functions
	generateVALIGND(asm)
	generateVALIGNQ(asm)
	generateVPBLENDMQ(asm)
	generateVPBLENDMD(asm)
	generateVPERMQ(asm)
	generateVPERMD(asm)
	generateVPERMI2Q(asm)
	generateVPERMT2Q(asm)
	generateVSHUFI64X2(asm)
	generateVSHUFPD(asm)
	generateVPSHUFD(asm)
	generateVPUNPCKLDQ(asm)
	generateVPUNPCKHDQ(asm)
	generateVPUNPCKLQDQ(asm)
	generateVPUNPCKHQDQ(asm)
	generateVPMADD52LUQ(asm)
	generateVPMADD52HUQ(asm)
	generateVPTERNLOGD(asm)
}

// generateVALIGND generates test function for VALIGND instruction
func generateVALIGND(asm *amd64.Amd64) {
	asm.FnHeader("testVALIGND", 0, 32)

	asm.MOVQ("src1+0(FP)", amd64.AX)
	asm.MOVQ("src2+8(FP)", amd64.BX)
	asm.MOVQ("imm+16(FP)", amd64.CX)
	asm.MOVQ("dst+24(FP)", amd64.DX)

	asm.VMOVDQU32("(AX)", amd64.Z0)
	asm.VMOVDQU32("(BX)", amd64.Z1)

	asm.CMPQ(amd64.CX, 0)
	l0 := asm.NewLabel("imm0")
	asm.JEQ(l0)
	asm.CMPQ(amd64.CX, 1)
	l1 := asm.NewLabel("imm1")
	asm.JEQ(l1)
	asm.CMPQ(amd64.CX, 2)
	l2 := asm.NewLabel("imm2")
	asm.JEQ(l2)
	asm.CMPQ(amd64.CX, 4)
	l4 := asm.NewLabel("imm4")
	asm.JEQ(l4)
	asm.CMPQ(amd64.CX, 8)
	l8 := asm.NewLabel("imm8")
	asm.JEQ(l8)

	asm.LABEL(l0)
	asm.WriteLn("    VALIGND $0, Z1, Z0, Z2")
	done := asm.NewLabel("done")
	asm.JMP(done)

	asm.LABEL(l1)
	asm.WriteLn("    VALIGND $1, Z1, Z0, Z2")
	asm.JMP(done)

	asm.LABEL(l2)
	asm.WriteLn("    VALIGND $2, Z1, Z0, Z2")
	asm.JMP(done)

	asm.LABEL(l4)
	asm.WriteLn("    VALIGND $4, Z1, Z0, Z2")
	asm.JMP(done)

	asm.LABEL(l8)
	asm.WriteLn("    VALIGND $8, Z1, Z0, Z2")

	asm.LABEL(done)
	asm.VMOVDQU32(amd64.Z2, "(DX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVALIGNQ(asm *amd64.Amd64) {
	asm.FnHeader("testVALIGNQ", 0, 32)

	asm.MOVQ("src1+0(FP)", amd64.AX)
	asm.MOVQ("src2+8(FP)", amd64.BX)
	asm.MOVQ("imm+16(FP)", amd64.CX)
	asm.MOVQ("dst+24(FP)", amd64.DX)

	asm.VMOVDQU64("(AX)", amd64.Z0)
	asm.VMOVDQU64("(BX)", amd64.Z1)

	asm.CMPQ(amd64.CX, 0)
	l0 := asm.NewLabel("imm0")
	asm.JEQ(l0)
	asm.CMPQ(amd64.CX, 1)
	l1 := asm.NewLabel("imm1")
	asm.JEQ(l1)
	asm.CMPQ(amd64.CX, 2)
	l2 := asm.NewLabel("imm2")
	asm.JEQ(l2)
	asm.CMPQ(amd64.CX, 4)
	l4 := asm.NewLabel("imm4")
	asm.JEQ(l4)

	asm.LABEL(l0)
	asm.WriteLn("    VALIGNQ $0, Z1, Z0, Z2")
	done := asm.NewLabel("done")
	asm.JMP(done)

	asm.LABEL(l1)
	asm.WriteLn("    VALIGNQ $1, Z1, Z0, Z2")
	asm.JMP(done)

	asm.LABEL(l2)
	asm.WriteLn("    VALIGNQ $2, Z1, Z0, Z2")
	asm.JMP(done)

	asm.LABEL(l4)
	asm.WriteLn("    VALIGNQ $4, Z1, Z0, Z2")

	asm.LABEL(done)
	asm.VMOVDQU64(amd64.Z2, "(DX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPBLENDMQ(asm *amd64.Amd64) {
	asm.FnHeader("testVPBLENDMQ", 0, 32)

	asm.MOVQ("src1+0(FP)", amd64.AX)
	asm.MOVQ("src2+8(FP)", amd64.BX)
	asm.MOVQ("mask+16(FP)", amd64.CX)
	asm.MOVQ("dst+24(FP)", amd64.DX)

	asm.VMOVDQU64("(AX)", amd64.Z0)
	asm.VMOVDQU64("(BX)", amd64.Z1)
	asm.KMOVB(amd64.CX, amd64.K1)

	asm.WriteLn("    VPBLENDMQ Z1, Z0, K1, Z2")

	asm.VMOVDQU64(amd64.Z2, "(DX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPBLENDMD(asm *amd64.Amd64) {
	asm.FnHeader("testVPBLENDMD", 0, 32)

	asm.MOVQ("src1+0(FP)", amd64.AX)
	asm.MOVQ("src2+8(FP)", amd64.BX)
	asm.MOVQ("mask+16(FP)", amd64.CX)
	asm.MOVQ("dst+24(FP)", amd64.DX)

	asm.VMOVDQU32("(AX)", amd64.Z0)
	asm.VMOVDQU32("(BX)", amd64.Z1)
	asm.KMOVW(amd64.CX, amd64.K1)

	asm.WriteLn("    VPBLENDMD Z1, Z0, K1, Z2")

	asm.VMOVDQU32(amd64.Z2, "(DX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPERMQ(asm *amd64.Amd64) {
	asm.FnHeader("testVPERMQ", 0, 24)

	asm.MOVQ("src+0(FP)", amd64.AX)
	asm.MOVQ("imm+8(FP)", amd64.BX)
	asm.MOVQ("dst+16(FP)", amd64.CX)

	asm.VMOVDQU64("(AX)", amd64.Z0)

	asm.CMPQ(amd64.BX, 0x00)
	l00 := asm.NewLabel("imm00")
	asm.JEQ(l00)
	asm.CMPQ(amd64.BX, 0x55)
	l55 := asm.NewLabel("imm55")
	asm.JEQ(l55)
	asm.CMPQ(amd64.BX, 0xAA)
	lAA := asm.NewLabel("immAA")
	asm.JEQ(lAA)
	asm.CMPQ(amd64.BX, 0xD8)
	lD8 := asm.NewLabel("immD8")
	asm.JEQ(lD8)
	asm.CMPQ(amd64.BX, 0x1B)
	l1B := asm.NewLabel("imm1B")
	asm.JEQ(l1B)

	asm.LABEL(l00)
	asm.WriteLn("    VPERMQ $0x00, Z0, Z1")
	done := asm.NewLabel("done")
	asm.JMP(done)

	asm.LABEL(l55)
	asm.WriteLn("    VPERMQ $0x55, Z0, Z1")
	asm.JMP(done)

	asm.LABEL(lAA)
	asm.WriteLn("    VPERMQ $0xAA, Z0, Z1")
	asm.JMP(done)

	asm.LABEL(lD8)
	asm.WriteLn("    VPERMQ $0xD8, Z0, Z1")
	asm.JMP(done)

	asm.LABEL(l1B)
	asm.WriteLn("    VPERMQ $0x1B, Z0, Z1")

	asm.LABEL(done)
	asm.VMOVDQU64(amd64.Z1, "(CX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPERMD(asm *amd64.Amd64) {
	asm.FnHeader("testVPERMD", 0, 24)

	asm.MOVQ("idx+0(FP)", amd64.AX)
	asm.MOVQ("src+8(FP)", amd64.BX)
	asm.MOVQ("dst+16(FP)", amd64.CX)

	asm.VMOVDQU32("(AX)", amd64.Z0) // idx
	asm.VMOVDQU32("(BX)", amd64.Z1) // src

	// Go asm VPERMD is: VPERMD src, idx, dst
	asm.VPERMD(amd64.Z1, amd64.Z0, amd64.Z2)

	asm.VMOVDQU32(amd64.Z2, "(CX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPERMI2Q(asm *amd64.Amd64) {
	asm.FnHeader("testVPERMI2Q", 0, 32)

	asm.MOVQ("src1+0(FP)", amd64.AX)
	asm.MOVQ("idx+8(FP)", amd64.BX)
	asm.MOVQ("src2+16(FP)", amd64.CX)
	asm.MOVQ("dst+24(FP)", amd64.DX)

	asm.VMOVDQU64("(AX)", amd64.Z0)
	asm.VMOVDQU64("(BX)", amd64.Z1)
	asm.VMOVDQU64("(CX)", amd64.Z2)

	asm.VPERMI2Q(amd64.Z2, amd64.Z0, amd64.Z1)

	asm.VMOVDQU64(amd64.Z1, "(DX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPERMT2Q(asm *amd64.Amd64) {
	asm.FnHeader("testVPERMT2Q", 0, 32)

	asm.MOVQ("src1+0(FP)", amd64.AX)
	asm.MOVQ("idx+8(FP)", amd64.BX)
	asm.MOVQ("src2+16(FP)", amd64.CX)
	asm.MOVQ("dst+24(FP)", amd64.DX)

	asm.VMOVDQU64("(AX)", amd64.Z0)
	asm.VMOVDQU64("(BX)", amd64.Z1)
	asm.VMOVDQU64("(CX)", amd64.Z2)

	asm.VPERMT2Q(amd64.Z2, amd64.Z1, amd64.Z0)

	asm.VMOVDQU64(amd64.Z0, "(DX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVSHUFI64X2(asm *amd64.Amd64) {
	asm.FnHeader("testVSHUFI64X2", 0, 32)

	asm.MOVQ("src1+0(FP)", amd64.AX)
	asm.MOVQ("src2+8(FP)", amd64.BX)
	asm.MOVQ("imm+16(FP)", amd64.CX)
	asm.MOVQ("dst+24(FP)", amd64.DX)

	asm.VMOVDQU64("(AX)", amd64.Z0)
	asm.VMOVDQU64("(BX)", amd64.Z1)

	asm.CMPQ(amd64.CX, 0x00)
	l00 := asm.NewLabel("imm00")
	asm.JEQ(l00)
	asm.CMPQ(amd64.CX, 0x44)
	l44 := asm.NewLabel("imm44")
	asm.JEQ(l44)
	asm.CMPQ(amd64.CX, 0xEE)
	lEE := asm.NewLabel("immEE")
	asm.JEQ(lEE)

	asm.LABEL(l00)
	asm.WriteLn("    VSHUFI64X2 $0x00, Z1, Z0, Z2")
	done := asm.NewLabel("done")
	asm.JMP(done)

	asm.LABEL(l44)
	asm.WriteLn("    VSHUFI64X2 $0x44, Z1, Z0, Z2")
	asm.JMP(done)

	asm.LABEL(lEE)
	asm.WriteLn("    VSHUFI64X2 $0xEE, Z1, Z0, Z2")

	asm.LABEL(done)
	asm.VMOVDQU64(amd64.Z2, "(DX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVSHUFPD(asm *amd64.Amd64) {
	asm.FnHeader("testVSHUFPD", 0, 32)

	asm.MOVQ("src1+0(FP)", amd64.AX)
	asm.MOVQ("src2+8(FP)", amd64.BX)
	asm.MOVQ("imm+16(FP)", amd64.CX)
	asm.MOVQ("dst+24(FP)", amd64.DX)

	asm.VMOVDQU64("(AX)", amd64.Z0)
	asm.VMOVDQU64("(BX)", amd64.Z1)

	asm.CMPQ(amd64.CX, 0x00)
	l00 := asm.NewLabel("imm00")
	asm.JEQ(l00)
	asm.CMPQ(amd64.CX, 0x55)
	l55 := asm.NewLabel("imm55")
	asm.JEQ(l55)
	asm.CMPQ(amd64.CX, 0xAA)
	lAA := asm.NewLabel("immAA")
	asm.JEQ(lAA)
	asm.CMPQ(amd64.CX, 0xFF)
	lFF := asm.NewLabel("immFF")
	asm.JEQ(lFF)

	asm.LABEL(l00)
	asm.WriteLn("    VSHUFPD $0x00, Z1, Z0, Z2")
	done := asm.NewLabel("done")
	asm.JMP(done)

	asm.LABEL(l55)
	asm.WriteLn("    VSHUFPD $0x55, Z1, Z0, Z2")
	asm.JMP(done)

	asm.LABEL(lAA)
	asm.WriteLn("    VSHUFPD $0xAA, Z1, Z0, Z2")
	asm.JMP(done)

	asm.LABEL(lFF)
	asm.WriteLn("    VSHUFPD $0xFF, Z1, Z0, Z2")

	asm.LABEL(done)
	asm.VMOVDQU64(amd64.Z2, "(DX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPSHUFD(asm *amd64.Amd64) {
	asm.FnHeader("testVPSHUFD", 0, 24)

	asm.MOVQ("src+0(FP)", amd64.AX)
	asm.MOVQ("imm+8(FP)", amd64.BX)
	asm.MOVQ("dst+16(FP)", amd64.CX)

	asm.VMOVDQU32("(AX)", amd64.Z0)

	asm.CMPQ(amd64.BX, 0x00)
	l00 := asm.NewLabel("imm00")
	asm.JEQ(l00)
	asm.CMPQ(amd64.BX, 0x1B)
	l1B := asm.NewLabel("imm1B")
	asm.JEQ(l1B)
	asm.CMPQ(amd64.BX, 0xB1)
	lB1 := asm.NewLabel("immB1")
	asm.JEQ(lB1)
	asm.CMPQ(amd64.BX, 0xD8)
	lD8 := asm.NewLabel("immD8")
	asm.JEQ(lD8)

	asm.LABEL(l00)
	asm.WriteLn("    VPSHUFD $0x00, Z0, Z1")
	done := asm.NewLabel("done")
	asm.JMP(done)

	asm.LABEL(l1B)
	asm.WriteLn("    VPSHUFD $0x1B, Z0, Z1")
	asm.JMP(done)

	asm.LABEL(lB1)
	asm.WriteLn("    VPSHUFD $0xB1, Z0, Z1")
	asm.JMP(done)

	asm.LABEL(lD8)
	asm.WriteLn("    VPSHUFD $0xD8, Z0, Z1")

	asm.LABEL(done)
	asm.VMOVDQU32(amd64.Z1, "(CX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPUNPCKLDQ(asm *amd64.Amd64) {
	asm.FnHeader("testVPUNPCKLDQ", 0, 24)

	asm.MOVQ("src1+0(FP)", amd64.AX)
	asm.MOVQ("src2+8(FP)", amd64.BX)
	asm.MOVQ("dst+16(FP)", amd64.CX)

	asm.VMOVDQU32("(AX)", amd64.Z0)
	asm.VMOVDQU32("(BX)", amd64.Z1)

	asm.VPUNPCKLDQ(amd64.Z1, amd64.Z0, amd64.Z2)

	asm.VMOVDQU32(amd64.Z2, "(CX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPUNPCKHDQ(asm *amd64.Amd64) {
	asm.FnHeader("testVPUNPCKHDQ", 0, 24)

	asm.MOVQ("src1+0(FP)", amd64.AX)
	asm.MOVQ("src2+8(FP)", amd64.BX)
	asm.MOVQ("dst+16(FP)", amd64.CX)

	asm.VMOVDQU32("(AX)", amd64.Z0)
	asm.VMOVDQU32("(BX)", amd64.Z1)

	asm.VPUNPCKHDQ(amd64.Z1, amd64.Z0, amd64.Z2)

	asm.VMOVDQU32(amd64.Z2, "(CX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPUNPCKLQDQ(asm *amd64.Amd64) {
	asm.FnHeader("testVPUNPCKLQDQ", 0, 24)

	asm.MOVQ("src1+0(FP)", amd64.AX)
	asm.MOVQ("src2+8(FP)", amd64.BX)
	asm.MOVQ("dst+16(FP)", amd64.CX)

	asm.VMOVDQU64("(AX)", amd64.Z0)
	asm.VMOVDQU64("(BX)", amd64.Z1)

	asm.VPUNPCKLQDQ(amd64.Z1, amd64.Z0, amd64.Z2)

	asm.VMOVDQU64(amd64.Z2, "(CX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPUNPCKHQDQ(asm *amd64.Amd64) {
	asm.FnHeader("testVPUNPCKHQDQ", 0, 24)

	asm.MOVQ("src1+0(FP)", amd64.AX)
	asm.MOVQ("src2+8(FP)", amd64.BX)
	asm.MOVQ("dst+16(FP)", amd64.CX)

	asm.VMOVDQU64("(AX)", amd64.Z0)
	asm.VMOVDQU64("(BX)", amd64.Z1)

	asm.VPUNPCKHQDQ(amd64.Z1, amd64.Z0, amd64.Z2)

	asm.VMOVDQU64(amd64.Z2, "(CX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPMADD52LUQ(asm *amd64.Amd64) {
	asm.FnHeader("testVPMADD52LUQ", 0, 32)

	asm.MOVQ("a+0(FP)", amd64.AX)
	asm.MOVQ("b+8(FP)", amd64.BX)
	asm.MOVQ("c+16(FP)", amd64.CX)
	asm.MOVQ("dst+24(FP)", amd64.DX)

	asm.VMOVDQU64("(AX)", amd64.Z0)
	asm.VMOVDQU64("(BX)", amd64.Z1)
	asm.VMOVDQU64("(CX)", amd64.Z2)

	asm.VPMADD52LUQ(amd64.Z1, amd64.Z0, amd64.Z2)

	asm.VMOVDQU64(amd64.Z2, "(DX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPMADD52HUQ(asm *amd64.Amd64) {
	asm.FnHeader("testVPMADD52HUQ", 0, 32)

	asm.MOVQ("a+0(FP)", amd64.AX)
	asm.MOVQ("b+8(FP)", amd64.BX)
	asm.MOVQ("c+16(FP)", amd64.CX)
	asm.MOVQ("dst+24(FP)", amd64.DX)

	asm.VMOVDQU64("(AX)", amd64.Z0)
	asm.VMOVDQU64("(BX)", amd64.Z1)
	asm.VMOVDQU64("(CX)", amd64.Z2)

	asm.VPMADD52HUQ(amd64.Z1, amd64.Z0, amd64.Z2)

	asm.VMOVDQU64(amd64.Z2, "(DX)")
	asm.RET()
	asm.WriteLn("")
}

func generateVPTERNLOGD(asm *amd64.Amd64) {
	asm.FnHeader("testVPTERNLOGD", 0, 40)

	asm.MOVQ("a+0(FP)", amd64.AX)
	asm.MOVQ("b+8(FP)", amd64.BX)
	asm.MOVQ("c+16(FP)", amd64.CX)
	asm.MOVQ("imm+24(FP)", amd64.R8)
	asm.MOVQ("dst+32(FP)", amd64.DX)

	asm.VMOVDQU32("(AX)", amd64.Z0)
	asm.VMOVDQU32("(BX)", amd64.Z1)
	asm.VMOVDQU32("(CX)", amd64.Z2)

	asm.CMPQ(amd64.R8, 0x96)
	l96 := asm.NewLabel("imm96")
	asm.JEQ(l96)
	asm.CMPQ(amd64.R8, 0x80)
	l80 := asm.NewLabel("imm80")
	asm.JEQ(l80)
	asm.CMPQ(amd64.R8, 0xFE)
	lFE := asm.NewLabel("immFE")
	asm.JEQ(lFE)

	asm.LABEL(l96)
	asm.WriteLn("    VPTERNLOGD $0x96, Z2, Z1, Z0")
	done := asm.NewLabel("done")
	asm.JMP(done)

	asm.LABEL(l80)
	asm.WriteLn("    VPTERNLOGD $0x80, Z2, Z1, Z0")
	asm.JMP(done)

	asm.LABEL(lFE)
	asm.WriteLn("    VPTERNLOGD $0xFE, Z2, Z1, Z0")

	asm.LABEL(done)
	asm.VMOVDQU32(amd64.Z0, "(DX)")
	asm.RET()
	asm.WriteLn("")
}
