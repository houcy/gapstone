package gapstone

const (
	// Shifters
	ARM64_SFT_INVALID = 0
	ARM64_SFT_LSL     = 1
	ARM64_SFT_MSL     = 2
	ARM64_SFT_LSR     = 3
	ARM64_SFT_ASR     = 4
	ARM64_SFT_ROR     = 1
)

const (
	// Extenders
	ARM64_EXT_INVALID = 0
	ARM64_EXT_UXTB    = 1
	ARM64_EXT_UXTH    = 2
	ARM64_EXT_UXTW    = 3
	ARM64_EXT_UXTX    = 4
	ARM64_EXT_SXTB    = 5
	ARM64_EXT_SXTH    = 6
	ARM64_EXT_SXTW    = 7
	ARM64_EXT_SXTX    = 8
)

const (
	// Condition Codes
	ARM64_CC_INVALID = 0
	ARM64_CC_EQ      = 1  // Equal
	ARM64_CC_NE      = 2  // Not equal:                 Not equal, or unordered
	ARM64_CC_HS      = 3  // Unsigned higher or same:   >, ==, or unordered
	ARM64_CC_LO      = 4  // Unsigned lower or same:    Less than
	ARM64_CC_MI      = 5  // Minus, negative:           Less than
	ARM64_CC_PL      = 6  // Plus, positive or zero:    >, ==, or unordered
	ARM64_CC_VS      = 7  // Overflow:                  Unordered
	ARM64_CC_VC      = 8  // No overflow:               Ordered
	ARM64_CC_HI      = 9  // Unsigned higher:           Greater than, or unordered
	ARM64_CC_LS      = 10 // Unsigned lower or same:    Less than or equal
	ARM64_CC_GE      = 11 // Greater than or equal:     Greater than or equal
	ARM64_CC_LT      = 12 // Less than:                 Less than, or unordered
	ARM64_CC_GT      = 13 // Signed greater than:       Greater than
	ARM64_CC_LE      = 14 // Signed less than or equal: <, ==, or unordered
	ARM64_CC_AL      = 15 // Always (unconditional):    Always (unconditional)
	ARM64_CC_NV      = 16 // Always (unconditional):   Always (unconditional)
// Note the NV exists purely to disassemble 0b1111. Execution
// is "always".
)

const (
	// Operands
	ARM64_OP_INVALID = iota // Uninitialized.
	ARM64_OP_REG            // uint operand.
	ARM64_OP_CIMM           // C-Immediate
	ARM64_OP_IMM            // Immediate operand.
	ARM64_OP_FP             // Floating-Point immediate operand.
	ARM64_OP_MEM            // Memory operand
)

const (
	// Registers
	ARM64_REG_INVALID = 0
	ARM64_REG_NZCV    = 1
	ARM64_REG_WSP     = 2
	ARM64_REG_WZR     = 3
	ARM64_REG_SP      = 4
	ARM64_REG_XZR     = 5
	ARM64_REG_B0      = 6
	ARM64_REG_B1      = 7
	ARM64_REG_B2      = 8
	ARM64_REG_B3      = 9
	ARM64_REG_B4      = 10
	ARM64_REG_B5      = 11
	ARM64_REG_B6      = 12
	ARM64_REG_B7      = 13
	ARM64_REG_B8      = 14
	ARM64_REG_B9      = 15
	ARM64_REG_B10     = 16
	ARM64_REG_B11     = 17
	ARM64_REG_B12     = 18
	ARM64_REG_B13     = 19
	ARM64_REG_B14     = 20
	ARM64_REG_B15     = 21
	ARM64_REG_B16     = 22
	ARM64_REG_B17     = 23
	ARM64_REG_B18     = 24
	ARM64_REG_B19     = 25
	ARM64_REG_B20     = 26
	ARM64_REG_B21     = 27
	ARM64_REG_B22     = 28
	ARM64_REG_B23     = 29
	ARM64_REG_B24     = 30
	ARM64_REG_B25     = 31
	ARM64_REG_B26     = 32
	ARM64_REG_B27     = 33
	ARM64_REG_B28     = 34
	ARM64_REG_B29     = 35
	ARM64_REG_B30     = 36
	ARM64_REG_B31     = 37
	ARM64_REG_D0      = 38
	ARM64_REG_D1      = 39
	ARM64_REG_D2      = 40
	ARM64_REG_D3      = 41
	ARM64_REG_D4      = 42
	ARM64_REG_D5      = 43
	ARM64_REG_D6      = 44
	ARM64_REG_D7      = 45
	ARM64_REG_D8      = 46
	ARM64_REG_D9      = 47
	ARM64_REG_D10     = 48
	ARM64_REG_D11     = 49
	ARM64_REG_D12     = 50
	ARM64_REG_D13     = 51
	ARM64_REG_D14     = 52
	ARM64_REG_D15     = 53
	ARM64_REG_D16     = 54
	ARM64_REG_D17     = 55
	ARM64_REG_D18     = 56
	ARM64_REG_D19     = 57
	ARM64_REG_D20     = 58
	ARM64_REG_D21     = 59
	ARM64_REG_D22     = 60
	ARM64_REG_D23     = 61
	ARM64_REG_D24     = 62
	ARM64_REG_D25     = 63
	ARM64_REG_D26     = 64
	ARM64_REG_D27     = 65
	ARM64_REG_D28     = 66
	ARM64_REG_D29     = 67
	ARM64_REG_D30     = 68
	ARM64_REG_D31     = 69
	ARM64_REG_H0      = 70
	ARM64_REG_H1      = 71
	ARM64_REG_H2      = 72
	ARM64_REG_H3      = 73
	ARM64_REG_H4      = 74
	ARM64_REG_H5      = 75
	ARM64_REG_H6      = 76
	ARM64_REG_H7      = 77
	ARM64_REG_H8      = 78
	ARM64_REG_H9      = 79
	ARM64_REG_H10     = 80
	ARM64_REG_H11     = 81
	ARM64_REG_H12     = 82
	ARM64_REG_H13     = 83
	ARM64_REG_H14     = 84
	ARM64_REG_H15     = 85
	ARM64_REG_H16     = 86
	ARM64_REG_H17     = 87
	ARM64_REG_H18     = 88
	ARM64_REG_H19     = 89
	ARM64_REG_H20     = 90
	ARM64_REG_H21     = 91
	ARM64_REG_H22     = 92
	ARM64_REG_H23     = 93
	ARM64_REG_H24     = 94
	ARM64_REG_H25     = 95
	ARM64_REG_H26     = 96
	ARM64_REG_H27     = 97
	ARM64_REG_H28     = 98
	ARM64_REG_H29     = 99
	ARM64_REG_H30     = 100
	ARM64_REG_H31     = 101
	ARM64_REG_Q0      = 102
	ARM64_REG_Q1      = 103
	ARM64_REG_Q2      = 104
	ARM64_REG_Q3      = 105
	ARM64_REG_Q4      = 106
	ARM64_REG_Q5      = 107
	ARM64_REG_Q6      = 108
	ARM64_REG_Q7      = 109
	ARM64_REG_Q8      = 110
	ARM64_REG_Q9      = 111
	ARM64_REG_Q10     = 112
	ARM64_REG_Q11     = 113
	ARM64_REG_Q12     = 114
	ARM64_REG_Q13     = 115
	ARM64_REG_Q14     = 116
	ARM64_REG_Q15     = 117
	ARM64_REG_Q16     = 118
	ARM64_REG_Q17     = 119
	ARM64_REG_Q18     = 120
	ARM64_REG_Q19     = 121
	ARM64_REG_Q20     = 122
	ARM64_REG_Q21     = 123
	ARM64_REG_Q22     = 124
	ARM64_REG_Q23     = 125
	ARM64_REG_Q24     = 126
	ARM64_REG_Q25     = 127
	ARM64_REG_Q26     = 128
	ARM64_REG_Q27     = 129
	ARM64_REG_Q28     = 130
	ARM64_REG_Q29     = 131
	ARM64_REG_Q30     = 132
	ARM64_REG_Q31     = 133
	ARM64_REG_S0      = 134
	ARM64_REG_S1      = 135
	ARM64_REG_S2      = 136
	ARM64_REG_S3      = 137
	ARM64_REG_S4      = 138
	ARM64_REG_S5      = 139
	ARM64_REG_S6      = 140
	ARM64_REG_S7      = 141
	ARM64_REG_S8      = 142
	ARM64_REG_S9      = 143
	ARM64_REG_S10     = 144
	ARM64_REG_S11     = 145
	ARM64_REG_S12     = 146
	ARM64_REG_S13     = 147
	ARM64_REG_S14     = 148
	ARM64_REG_S15     = 149
	ARM64_REG_S16     = 150
	ARM64_REG_S17     = 151
	ARM64_REG_S18     = 152
	ARM64_REG_S19     = 153
	ARM64_REG_S20     = 154
	ARM64_REG_S21     = 155
	ARM64_REG_S22     = 156
	ARM64_REG_S23     = 157
	ARM64_REG_S24     = 158
	ARM64_REG_S25     = 159
	ARM64_REG_S26     = 160
	ARM64_REG_S27     = 161
	ARM64_REG_S28     = 162
	ARM64_REG_S29     = 163
	ARM64_REG_S30     = 164
	ARM64_REG_S31     = 165
	ARM64_REG_W0      = 166
	ARM64_REG_W1      = 167
	ARM64_REG_W2      = 168
	ARM64_REG_W3      = 169
	ARM64_REG_W4      = 170
	ARM64_REG_W5      = 171
	ARM64_REG_W6      = 172
	ARM64_REG_W7      = 173
	ARM64_REG_W8      = 174
	ARM64_REG_W9      = 175
	ARM64_REG_W10     = 176
	ARM64_REG_W11     = 177
	ARM64_REG_W12     = 178
	ARM64_REG_W13     = 179
	ARM64_REG_W14     = 180
	ARM64_REG_W15     = 181
	ARM64_REG_W16     = 182
	ARM64_REG_W17     = 183
	ARM64_REG_W18     = 184
	ARM64_REG_W19     = 185
	ARM64_REG_W20     = 186
	ARM64_REG_W21     = 187
	ARM64_REG_W22     = 188
	ARM64_REG_W23     = 189
	ARM64_REG_W24     = 190
	ARM64_REG_W25     = 191
	ARM64_REG_W26     = 192
	ARM64_REG_W27     = 193
	ARM64_REG_W28     = 194
	ARM64_REG_W29     = 195
	ARM64_REG_W30     = 196
	ARM64_REG_X0      = 197
	ARM64_REG_X1      = 198
	ARM64_REG_X2      = 199
	ARM64_REG_X3      = 200
	ARM64_REG_X4      = 201
	ARM64_REG_X5      = 202
	ARM64_REG_X6      = 203
	ARM64_REG_X7      = 204
	ARM64_REG_X8      = 205
	ARM64_REG_X9      = 206
	ARM64_REG_X10     = 207
	ARM64_REG_X11     = 208
	ARM64_REG_X12     = 209
	ARM64_REG_X13     = 210
	ARM64_REG_X14     = 211
	ARM64_REG_X15     = 212
	ARM64_REG_X16     = 213
	ARM64_REG_X17     = 214
	ARM64_REG_X18     = 215
	ARM64_REG_X19     = 216
	ARM64_REG_X20     = 217
	ARM64_REG_X21     = 218
	ARM64_REG_X22     = 219
	ARM64_REG_X23     = 220
	ARM64_REG_X24     = 221
	ARM64_REG_X25     = 222
	ARM64_REG_X26     = 223
	ARM64_REG_X27     = 224
	ARM64_REG_X28     = 225
	ARM64_REG_X29     = 226
	ARM64_REG_X30     = 227
	ARM64_REG_MAX     = 228
)

const (
	// Instructions
	ARM64_INS_INVALID = iota
	ARM64_INS_ADC
	ARM64_INS_ADDHN2
	ARM64_INS_ADDHN
	ARM64_INS_ADDP
	ARM64_INS_ADD
	ARM64_INS_CMN
	ARM64_INS_ADRP
	ARM64_INS_ADR
	ARM64_INS_AND
	ARM64_INS_ASR
	ARM64_INS_AT
	ARM64_INS_BFI
	ARM64_INS_BFM
	ARM64_INS_BFXIL
	ARM64_INS_BIC
	ARM64_INS_BIF
	ARM64_INS_BIT
	ARM64_INS_BLR
	ARM64_INS_BL
	ARM64_INS_BRK
	ARM64_INS_BR
	ARM64_INS_BSL
	ARM64_INS_B
	ARM64_INS_CBNZ
	ARM64_INS_CBZ
	ARM64_INS_CCMN
	ARM64_INS_CCMP
	ARM64_INS_CLREX
	ARM64_INS_CLS
	ARM64_INS_CLZ
	ARM64_INS_CMEQ
	ARM64_INS_CMGE
	ARM64_INS_CMGT
	ARM64_INS_CMHI
	ARM64_INS_CMHS
	ARM64_INS_CMLE
	ARM64_INS_CMLT
	ARM64_INS_CMP
	ARM64_INS_CMTST
	ARM64_INS_CRC32B
	ARM64_INS_CRC32CB
	ARM64_INS_CRC32CH
	ARM64_INS_CRC32CW
	ARM64_INS_CRC32CX
	ARM64_INS_CRC32H
	ARM64_INS_CRC32W
	ARM64_INS_CRC32X
	ARM64_INS_CSEL
	ARM64_INS_CSINC
	ARM64_INS_CSINV
	ARM64_INS_CSNEG
	ARM64_INS_DCPS1
	ARM64_INS_DCPS2
	ARM64_INS_DCPS3
	ARM64_INS_DC
	ARM64_INS_DMB
	ARM64_INS_DRPS
	ARM64_INS_DSB
	ARM64_INS_EON
	ARM64_INS_EOR
	ARM64_INS_ERET
	ARM64_INS_EXTR
	ARM64_INS_FABD
	ARM64_INS_FABS
	ARM64_INS_FACGE
	ARM64_INS_FACGT
	ARM64_INS_FADDP
	ARM64_INS_FADD
	ARM64_INS_FCCMPE
	ARM64_INS_FCCMP
	ARM64_INS_FCMEQ
	ARM64_INS_FCMGE
	ARM64_INS_FCMGT
	ARM64_INS_FCMLE
	ARM64_INS_FCMLT
	ARM64_INS_FCMP
	ARM64_INS_FCMPE
	ARM64_INS_FCSEL
	ARM64_INS_FCVTAS
	ARM64_INS_FCVTAU
	ARM64_INS_FCVTMS
	ARM64_INS_FCVTMU
	ARM64_INS_FCVTNS
	ARM64_INS_FCVTNU
	ARM64_INS_FCVTPS
	ARM64_INS_FCVTPU
	ARM64_INS_FCVTZS
	ARM64_INS_FCVTZU
	ARM64_INS_FCVT
	ARM64_INS_FDIV
	ARM64_INS_FMADD
	ARM64_INS_FMAXNMP
	ARM64_INS_FMAXNM
	ARM64_INS_FMAXP
	ARM64_INS_FMAX
	ARM64_INS_FMINNMP
	ARM64_INS_FMINNM
	ARM64_INS_FMINP
	ARM64_INS_FMIN
	ARM64_INS_FMLA
	ARM64_INS_FMLS
	ARM64_INS_FMOV
	ARM64_INS_FMSUB
	ARM64_INS_FMULX
	ARM64_INS_FMUL
	ARM64_INS_FNEG
	ARM64_INS_FNMADD
	ARM64_INS_FNMSUB
	ARM64_INS_FNMUL
	ARM64_INS_FRECPS
	ARM64_INS_FRINTA
	ARM64_INS_FRINTI
	ARM64_INS_FRINTM
	ARM64_INS_FRINTN
	ARM64_INS_FRINTP
	ARM64_INS_FRINTX
	ARM64_INS_FRINTZ
	ARM64_INS_FRSQRTS
	ARM64_INS_FSQRT
	ARM64_INS_FSUB
	ARM64_INS_HINT
	ARM64_INS_HLT
	ARM64_INS_HVC
	ARM64_INS_IC
	ARM64_INS_INS
	ARM64_INS_ISB
	ARM64_INS_LDARB
	ARM64_INS_LDAR
	ARM64_INS_LDARH
	ARM64_INS_LDAXP
	ARM64_INS_LDAXRB
	ARM64_INS_LDAXR
	ARM64_INS_LDAXRH
	ARM64_INS_LDPSW
	ARM64_INS_LDRSB
	ARM64_INS_LDURSB
	ARM64_INS_LDRSH
	ARM64_INS_LDURSH
	ARM64_INS_LDRSW
	ARM64_INS_LDR
	ARM64_INS_LDTRSB
	ARM64_INS_LDTRSH
	ARM64_INS_LDTRSW
	ARM64_INS_LDURSW
	ARM64_INS_LDXP
	ARM64_INS_LDXRB
	ARM64_INS_LDXR
	ARM64_INS_LDXRH
	ARM64_INS_LDRH
	ARM64_INS_LDURH
	ARM64_INS_STRH
	ARM64_INS_STURH
	ARM64_INS_LDTRH
	ARM64_INS_STTRH
	ARM64_INS_LDUR
	ARM64_INS_STR
	ARM64_INS_STUR
	ARM64_INS_LDTR
	ARM64_INS_STTR
	ARM64_INS_LDRB
	ARM64_INS_LDURB
	ARM64_INS_STRB
	ARM64_INS_STURB
	ARM64_INS_LDTRB
	ARM64_INS_STTRB
	ARM64_INS_LDP
	ARM64_INS_LDNP
	ARM64_INS_STNP
	ARM64_INS_STP
	ARM64_INS_LSL
	ARM64_INS_LSR
	ARM64_INS_MADD
	ARM64_INS_MLA
	ARM64_INS_MLS
	ARM64_INS_MOVI
	ARM64_INS_MOVK
	ARM64_INS_MOVN
	ARM64_INS_MOVZ
	ARM64_INS_MRS
	ARM64_INS_MSR
	ARM64_INS_MSUB
	ARM64_INS_MUL
	ARM64_INS_MVNI
	ARM64_INS_MVN
	ARM64_INS_ORN
	ARM64_INS_ORR
	ARM64_INS_PMULL2
	ARM64_INS_PMULL
	ARM64_INS_PMUL
	ARM64_INS_PRFM
	ARM64_INS_PRFUM
	ARM64_INS_SQRSHRUN2
	ARM64_INS_SQRSHRUN
	ARM64_INS_SQSHRUN2
	ARM64_INS_SQSHRUN
	ARM64_INS_RADDHN2
	ARM64_INS_RADDHN
	ARM64_INS_RBIT
	ARM64_INS_RET
	ARM64_INS_REV16
	ARM64_INS_REV32
	ARM64_INS_REV
	ARM64_INS_ROR
	ARM64_INS_RSHRN2
	ARM64_INS_RSHRN
	ARM64_INS_RSUBHN2
	ARM64_INS_RSUBHN
	ARM64_INS_SABAL2
	ARM64_INS_SABAL
	ARM64_INS_SABA
	ARM64_INS_SABDL2
	ARM64_INS_SABDL
	ARM64_INS_SABD
	ARM64_INS_SADDL2
	ARM64_INS_SADDL
	ARM64_INS_SADDW2
	ARM64_INS_SADDW
	ARM64_INS_SBC
	ARM64_INS_SBFIZ
	ARM64_INS_SBFM
	ARM64_INS_SBFX
	ARM64_INS_SCVTF
	ARM64_INS_SDIV
	ARM64_INS_SHADD
	ARM64_INS_SHL
	ARM64_INS_SHRN2
	ARM64_INS_SHRN
	ARM64_INS_SHSUB
	ARM64_INS_SLI
	ARM64_INS_SMADDL
	ARM64_INS_SMAXP
	ARM64_INS_SMAX
	ARM64_INS_SMC
	ARM64_INS_SMINP
	ARM64_INS_SMIN
	ARM64_INS_SMLAL2
	ARM64_INS_SMLAL
	ARM64_INS_SMLSL2
	ARM64_INS_SMLSL
	ARM64_INS_SMOV
	ARM64_INS_SMSUBL
	ARM64_INS_SMULH
	ARM64_INS_SMULL2
	ARM64_INS_SMULL
	ARM64_INS_SQADD
	ARM64_INS_SQDMLAL2
	ARM64_INS_SQDMLAL
	ARM64_INS_SQDMLSL2
	ARM64_INS_SQDMLSL
	ARM64_INS_SQDMULH
	ARM64_INS_SQDMULL2
	ARM64_INS_SQDMULL
	ARM64_INS_SQRDMULH
	ARM64_INS_SQRSHL
	ARM64_INS_SQRSHRN2
	ARM64_INS_SQRSHRN
	ARM64_INS_SQSHLU
	ARM64_INS_SQSHL
	ARM64_INS_SQSHRN2
	ARM64_INS_SQSHRN
	ARM64_INS_SQSUB
	ARM64_INS_SRHADD
	ARM64_INS_SRI
	ARM64_INS_SRSHL
	ARM64_INS_SRSHR
	ARM64_INS_SRSRA
	ARM64_INS_SSHLL2
	ARM64_INS_SSHLL
	ARM64_INS_SSHL
	ARM64_INS_SSHR
	ARM64_INS_SSRA
	ARM64_INS_SSUBL2
	ARM64_INS_SSUBL
	ARM64_INS_SSUBW2
	ARM64_INS_SSUBW
	ARM64_INS_STLRB
	ARM64_INS_STLR
	ARM64_INS_STLRH
	ARM64_INS_STLXP
	ARM64_INS_STLXRB
	ARM64_INS_STLXR
	ARM64_INS_STLXRH
	ARM64_INS_STXP
	ARM64_INS_STXRB
	ARM64_INS_STXR
	ARM64_INS_STXRH
	ARM64_INS_SUBHN2
	ARM64_INS_SUBHN
	ARM64_INS_SUB
	ARM64_INS_SVC
	ARM64_INS_SXTB
	ARM64_INS_SXTH
	ARM64_INS_SXTW
	ARM64_INS_SYSL
	ARM64_INS_SYS
	ARM64_INS_TBNZ
	ARM64_INS_TBZ
	ARM64_INS_TLBI
	ARM64_INS_TST
	ARM64_INS_UABAL2
	ARM64_INS_UABAL
	ARM64_INS_UABA
	ARM64_INS_UABDL2
	ARM64_INS_UABDL
	ARM64_INS_UABD
	ARM64_INS_UADDL2
	ARM64_INS_UADDL
	ARM64_INS_UADDW2
	ARM64_INS_UADDW
	ARM64_INS_UBFIZ
	ARM64_INS_UBFM
	ARM64_INS_UBFX
	ARM64_INS_UCVTF
	ARM64_INS_UDIV
	ARM64_INS_UHADD
	ARM64_INS_UHSUB
	ARM64_INS_UMADDL
	ARM64_INS_UMAXP
	ARM64_INS_UMAX
	ARM64_INS_UMINP
	ARM64_INS_UMIN
	ARM64_INS_UMLAL2
	ARM64_INS_UMLAL
	ARM64_INS_UMLSL2
	ARM64_INS_UMLSL
	ARM64_INS_UMOV
	ARM64_INS_UMSUBL
	ARM64_INS_UMULH
	ARM64_INS_UMULL2
	ARM64_INS_UMULL
	ARM64_INS_UQADD
	ARM64_INS_UQRSHL
	ARM64_INS_UQRSHRN2
	ARM64_INS_UQRSHRN
	ARM64_INS_UQSHL
	ARM64_INS_UQSHRN2
	ARM64_INS_UQSHRN
	ARM64_INS_UQSUB
	ARM64_INS_URHADD
	ARM64_INS_URSHL
	ARM64_INS_URSHR
	ARM64_INS_URSRA
	ARM64_INS_USHLL2
	ARM64_INS_USHLL
	ARM64_INS_USHL
	ARM64_INS_USHR
	ARM64_INS_USRA
	ARM64_INS_USUBL2
	ARM64_INS_USUBL
	ARM64_INS_USUBW2
	ARM64_INS_USUBW
	ARM64_INS_UXTB
	ARM64_INS_UXTH

	// alias Instructions
	ARM64_INS_MNEG
	ARM64_INS_UMNEGL
	ARM64_INS_SMNEGL
	ARM64_INS_MOV
	ARM64_INS_NOP
	ARM64_INS_YIELD
	ARM64_INS_WFE
	ARM64_INS_WFI
	ARM64_INS_SEV
	ARM64_INS_SEVL
	ARM64_INS_NGC
	ARM64_INS_NEG

	ARM64_INS_MAX
)

const (
	// Groups
	ARM64_GRP_INVALID = iota
	ARM64_GRP_NEON
	ARM64_GRP_JUMP
	ARM64_GRP_MAX
)
