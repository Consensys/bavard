//go:build amd64 && !purego

#include "textflag.h"

TEXT ·testVALIGND(SB), NOSPLIT, $0-32
    MOVQ src1+0(FP), AX
    MOVQ src2+8(FP), BX
    MOVQ imm+16(FP), CX
    MOVQ dst+24(FP), DX
    VMOVDQU32 (AX), Z0
    VMOVDQU32 (BX), Z1
    CMPQ CX, $0
    JEQ imm0_1
    CMPQ CX, $1
    JEQ imm1_2
    CMPQ CX, $0x0000000000000002
    JEQ imm2_3
    CMPQ CX, $0x0000000000000004
    JEQ imm4_4
    CMPQ CX, $0x0000000000000008
    JEQ imm8_5
imm0_1:
    VALIGND $0, Z1, Z0, Z2
    JMP done_6
imm1_2:
    VALIGND $1, Z1, Z0, Z2
    JMP done_6
imm2_3:
    VALIGND $2, Z1, Z0, Z2
    JMP done_6
imm4_4:
    VALIGND $4, Z1, Z0, Z2
    JMP done_6
imm8_5:
    VALIGND $8, Z1, Z0, Z2
done_6:
    VMOVDQU32 Z2, (DX)
    RET

TEXT ·testVALIGNQ(SB), NOSPLIT, $0-32
    MOVQ src1+0(FP), AX
    MOVQ src2+8(FP), BX
    MOVQ imm+16(FP), CX
    MOVQ dst+24(FP), DX
    VMOVDQU64 (AX), Z0
    VMOVDQU64 (BX), Z1
    CMPQ CX, $0
    JEQ imm0_7
    CMPQ CX, $1
    JEQ imm1_8
    CMPQ CX, $0x0000000000000002
    JEQ imm2_9
    CMPQ CX, $0x0000000000000004
    JEQ imm4_10
imm0_7:
    VALIGNQ $0, Z1, Z0, Z2
    JMP done_11
imm1_8:
    VALIGNQ $1, Z1, Z0, Z2
    JMP done_11
imm2_9:
    VALIGNQ $2, Z1, Z0, Z2
    JMP done_11
imm4_10:
    VALIGNQ $4, Z1, Z0, Z2
done_11:
    VMOVDQU64 Z2, (DX)
    RET

TEXT ·testVPBLENDMQ(SB), NOSPLIT, $0-32
    MOVQ src1+0(FP), AX
    MOVQ src2+8(FP), BX
    MOVQ mask+16(FP), CX
    MOVQ dst+24(FP), DX
    VMOVDQU64 (AX), Z0
    VMOVDQU64 (BX), Z1
    KMOVB CX, K1
    VPBLENDMQ Z1, Z0, K1, Z2
    VMOVDQU64 Z2, (DX)
    RET

TEXT ·testVPBLENDMD(SB), NOSPLIT, $0-32
    MOVQ src1+0(FP), AX
    MOVQ src2+8(FP), BX
    MOVQ mask+16(FP), CX
    MOVQ dst+24(FP), DX
    VMOVDQU32 (AX), Z0
    VMOVDQU32 (BX), Z1
    KMOVW CX, K1
    VPBLENDMD Z1, Z0, K1, Z2
    VMOVDQU32 Z2, (DX)
    RET

TEXT ·testVPERMQ(SB), NOSPLIT, $0-24
    MOVQ src+0(FP), AX
    MOVQ imm+8(FP), BX
    MOVQ dst+16(FP), CX
    VMOVDQU64 (AX), Z0
    CMPQ BX, $0
    JEQ imm00_12
    CMPQ BX, $0x0000000000000055
    JEQ imm55_13
    CMPQ BX, $0x00000000000000aa
    JEQ immAA_14
    CMPQ BX, $0x00000000000000d8
    JEQ immD8_15
    CMPQ BX, $0x000000000000001b
    JEQ imm1B_16
imm00_12:
    VPERMQ $0x00, Z0, Z1
    JMP done_17
imm55_13:
    VPERMQ $0x55, Z0, Z1
    JMP done_17
immAA_14:
    VPERMQ $0xAA, Z0, Z1
    JMP done_17
immD8_15:
    VPERMQ $0xD8, Z0, Z1
    JMP done_17
imm1B_16:
    VPERMQ $0x1B, Z0, Z1
done_17:
    VMOVDQU64 Z1, (CX)
    RET

TEXT ·testVPERMD(SB), NOSPLIT, $0-24
    MOVQ idx+0(FP), AX
    MOVQ src+8(FP), BX
    MOVQ dst+16(FP), CX
    VMOVDQU32 (AX), Z0
    VMOVDQU32 (BX), Z1
    VPERMD Z1, Z0, Z2
    VMOVDQU32 Z2, (CX)
    RET

TEXT ·testVPERMI2Q(SB), NOSPLIT, $0-32
    MOVQ src1+0(FP), AX
    MOVQ idx+8(FP), BX
    MOVQ src2+16(FP), CX
    MOVQ dst+24(FP), DX
    VMOVDQU64 (AX), Z0
    VMOVDQU64 (BX), Z1
    VMOVDQU64 (CX), Z2
    VPERMI2Q Z2, Z0, Z1
    VMOVDQU64 Z1, (DX)
    RET

TEXT ·testVPERMT2Q(SB), NOSPLIT, $0-32
    MOVQ src1+0(FP), AX
    MOVQ idx+8(FP), BX
    MOVQ src2+16(FP), CX
    MOVQ dst+24(FP), DX
    VMOVDQU64 (AX), Z0
    VMOVDQU64 (BX), Z1
    VMOVDQU64 (CX), Z2
    VPERMT2Q Z2, Z1, Z0
    VMOVDQU64 Z0, (DX)
    RET

TEXT ·testVSHUFI64X2(SB), NOSPLIT, $0-32
    MOVQ src1+0(FP), AX
    MOVQ src2+8(FP), BX
    MOVQ imm+16(FP), CX
    MOVQ dst+24(FP), DX
    VMOVDQU64 (AX), Z0
    VMOVDQU64 (BX), Z1
    CMPQ CX, $0
    JEQ imm00_18
    CMPQ CX, $0x0000000000000044
    JEQ imm44_19
    CMPQ CX, $0x00000000000000ee
    JEQ immEE_20
imm00_18:
    VSHUFI64X2 $0x00, Z1, Z0, Z2
    JMP done_21
imm44_19:
    VSHUFI64X2 $0x44, Z1, Z0, Z2
    JMP done_21
immEE_20:
    VSHUFI64X2 $0xEE, Z1, Z0, Z2
done_21:
    VMOVDQU64 Z2, (DX)
    RET

TEXT ·testVSHUFPD(SB), NOSPLIT, $0-32
    MOVQ src1+0(FP), AX
    MOVQ src2+8(FP), BX
    MOVQ imm+16(FP), CX
    MOVQ dst+24(FP), DX
    VMOVDQU64 (AX), Z0
    VMOVDQU64 (BX), Z1
    CMPQ CX, $0
    JEQ imm00_22
    CMPQ CX, $0x0000000000000055
    JEQ imm55_23
    CMPQ CX, $0x00000000000000aa
    JEQ immAA_24
    CMPQ CX, $0x00000000000000ff
    JEQ immFF_25
imm00_22:
    VSHUFPD $0x00, Z1, Z0, Z2
    JMP done_26
imm55_23:
    VSHUFPD $0x55, Z1, Z0, Z2
    JMP done_26
immAA_24:
    VSHUFPD $0xAA, Z1, Z0, Z2
    JMP done_26
immFF_25:
    VSHUFPD $0xFF, Z1, Z0, Z2
done_26:
    VMOVDQU64 Z2, (DX)
    RET

TEXT ·testVPSHUFD(SB), NOSPLIT, $0-24
    MOVQ src+0(FP), AX
    MOVQ imm+8(FP), BX
    MOVQ dst+16(FP), CX
    VMOVDQU32 (AX), Z0
    CMPQ BX, $0
    JEQ imm00_27
    CMPQ BX, $0x000000000000001b
    JEQ imm1B_28
    CMPQ BX, $0x00000000000000b1
    JEQ immB1_29
    CMPQ BX, $0x00000000000000d8
    JEQ immD8_30
imm00_27:
    VPSHUFD $0x00, Z0, Z1
    JMP done_31
imm1B_28:
    VPSHUFD $0x1B, Z0, Z1
    JMP done_31
immB1_29:
    VPSHUFD $0xB1, Z0, Z1
    JMP done_31
immD8_30:
    VPSHUFD $0xD8, Z0, Z1
done_31:
    VMOVDQU32 Z1, (CX)
    RET

TEXT ·testVPUNPCKLDQ(SB), NOSPLIT, $0-24
    MOVQ src1+0(FP), AX
    MOVQ src2+8(FP), BX
    MOVQ dst+16(FP), CX
    VMOVDQU32 (AX), Z0
    VMOVDQU32 (BX), Z1
    VPUNPCKLDQ Z1, Z0, Z2
    VMOVDQU32 Z2, (CX)
    RET

TEXT ·testVPUNPCKHDQ(SB), NOSPLIT, $0-24
    MOVQ src1+0(FP), AX
    MOVQ src2+8(FP), BX
    MOVQ dst+16(FP), CX
    VMOVDQU32 (AX), Z0
    VMOVDQU32 (BX), Z1
    VPUNPCKHDQ Z1, Z0, Z2
    VMOVDQU32 Z2, (CX)
    RET

TEXT ·testVPUNPCKLQDQ(SB), NOSPLIT, $0-24
    MOVQ src1+0(FP), AX
    MOVQ src2+8(FP), BX
    MOVQ dst+16(FP), CX
    VMOVDQU64 (AX), Z0
    VMOVDQU64 (BX), Z1
    VPUNPCKLQDQ Z1, Z0, Z2
    VMOVDQU64 Z2, (CX)
    RET

TEXT ·testVPUNPCKHQDQ(SB), NOSPLIT, $0-24
    MOVQ src1+0(FP), AX
    MOVQ src2+8(FP), BX
    MOVQ dst+16(FP), CX
    VMOVDQU64 (AX), Z0
    VMOVDQU64 (BX), Z1
    VPUNPCKHQDQ Z1, Z0, Z2
    VMOVDQU64 Z2, (CX)
    RET

TEXT ·testVPMADD52LUQ(SB), NOSPLIT, $0-32
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    MOVQ c+16(FP), CX
    MOVQ dst+24(FP), DX
    VMOVDQU64 (AX), Z0
    VMOVDQU64 (BX), Z1
    VMOVDQU64 (CX), Z2
    VPMADD52LUQ Z1, Z0, Z2
    VMOVDQU64 Z2, (DX)
    RET

TEXT ·testVPMADD52HUQ(SB), NOSPLIT, $0-32
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    MOVQ c+16(FP), CX
    MOVQ dst+24(FP), DX
    VMOVDQU64 (AX), Z0
    VMOVDQU64 (BX), Z1
    VMOVDQU64 (CX), Z2
    VPMADD52HUQ Z1, Z0, Z2
    VMOVDQU64 Z2, (DX)
    RET

TEXT ·testVPTERNLOGD(SB), NOSPLIT, $0-40
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    MOVQ c+16(FP), CX
    MOVQ imm+24(FP), R8
    MOVQ dst+32(FP), DX
    VMOVDQU32 (AX), Z0
    VMOVDQU32 (BX), Z1
    VMOVDQU32 (CX), Z2
    CMPQ R8, $0x0000000000000096
    JEQ imm96_32
    CMPQ R8, $0x0000000000000080
    JEQ imm80_33
    CMPQ R8, $0x00000000000000fe
    JEQ immFE_34
imm96_32:
    VPTERNLOGD $0x96, Z2, Z1, Z0
    JMP done_35
imm80_33:
    VPTERNLOGD $0x80, Z2, Z1, Z0
    JMP done_35
immFE_34:
    VPTERNLOGD $0xFE, Z2, Z1, Z0
done_35:
    VMOVDQU32 Z0, (DX)
    RET

