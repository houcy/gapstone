****************
Platform: ARM
Code:0xed 0xff 0xff 0xeb 0x04 0xe0 0x2d 0xe5 0x00 0x00 0x00 0x00 0xe0 0x83 0x22 0xe5 0xf1 0x02 0x03 0x0e 0x00 0x00 0xa0 0xe3 0x02 0x30 0xc1 0xe7 0x00 0x00 0x53 0xe3 0x00 0x02 0x01 0xf1 0x05 0x40 0xd0 0xe8 0xf4 0x80 0x00 0x00 
Disasm:
0x1000:	bl	#0xfbc
	op_count: 1
		operands[0].type: IMM = 0xfbc

0x1004:	str	lr, [sp, #-4]!
	op_count: 2
		operands[0].type: REG = lr
		operands[1].type: MEM
			operands[1].mem.base: REG = sp
			operands[1].mem.disp: 0xfffffffc
	Write-back: True

0x1008:	andeq	r0, r0, r0
	op_count: 3
		operands[0].type: REG = r0
		operands[1].type: REG = r0
		operands[2].type: REG = r0
	Code condition: 1

0x100c:	str	r8, [r2, #-0x3e0]!
	op_count: 2
		operands[0].type: REG = r8
		operands[1].type: MEM
			operands[1].mem.base: REG = r2
			operands[1].mem.disp: 0xfffffc20
	Write-back: True

0x1010:	mcreq	p2, #0, r0, c3, c1, #7
	op_count: 6
		operands[0].type: P-IMM = 2
		operands[1].type: IMM = 0x0
		operands[2].type: REG = r0
		operands[3].type: C-IMM = 3
		operands[4].type: C-IMM = 1
		operands[5].type: IMM = 0x7
	Code condition: 1

0x1014:	mov	r0, #0
	op_count: 2
		operands[0].type: REG = r0
		operands[1].type: IMM = 0x0

0x1018:	strb	r3, [r1, r2]
	op_count: 2
		operands[0].type: REG = r3
		operands[1].type: MEM
			operands[1].mem.base: REG = r1
			operands[1].mem.index: REG = r2

0x101c:	cmp	r3, #0
	op_count: 2
		operands[0].type: REG = r3
		operands[1].type: IMM = 0x0
	Update-flags: True

0x1020:	setend	be
	op_count: 1
		operands[0].type: SETEND = be

0x1024:	ldm	r0, {r0, r2, lr} ^
	op_count: 4
		operands[0].type: REG = r0
		operands[1].type: REG = r0
		operands[2].type: REG = r2
		operands[3].type: REG = lr
	User-mode: True

0x1028:	strdeq	r8, sb, [r0], -r4
	op_count: 4
		operands[0].type: REG = r8
		operands[1].type: REG = sb
		operands[2].type: MEM
			operands[2].mem.base: REG = r0
		operands[3].type: REG = r4
		Subtracted: True
	Code condition: 1
	Write-back: True

0x102c:

****************
Platform: Thumb
Code:0x70 0x47 0xeb 0x46 0x83 0xb0 0xc9 0x68 0x1f 0xb1 0x30 0xbf 0xaf 0xf3 0x20 0x84 
Disasm:
0x1000:	bx	lr
	op_count: 1
		operands[0].type: REG = lr

0x1002:	mov	fp, sp
	op_count: 2
		operands[0].type: REG = fp
		operands[1].type: REG = sp

0x1004:	sub	sp, #0xc
	op_count: 2
		operands[0].type: REG = sp
		operands[1].type: IMM = 0xc

0x1006:	ldr	r1, [r1, #0xc]
	op_count: 2
		operands[0].type: REG = r1
		operands[1].type: MEM
			operands[1].mem.base: REG = r1
			operands[1].mem.disp: 0xc

0x1008:	cbz	r7, #0x1012
	op_count: 2
		operands[0].type: REG = r7
		operands[1].type: IMM = 0x1012

0x100a:	wfi	

0x100c:	cpsie.w	f
	CPSI-mode: 2
	CPSI-flag: 1

0x1010:

****************
Platform: Thumb-mixed
Code:0xd1 0xe8 0x00 0xf0 0xf0 0x24 0x04 0x07 0x1f 0x3c 0xf2 0xc0 0x00 0x00 0x4f 0xf0 0x00 0x01 0x46 0x6c 
Disasm:
0x1000:	tbb	[r1, r0]
	op_count: 1
		operands[0].type: MEM
			operands[0].mem.base: REG = r1
			operands[0].mem.index: REG = r0

0x1004:	movs	r4, #0xf0
	op_count: 2
		operands[0].type: REG = r4
		operands[1].type: IMM = 0xf0
	Update-flags: True

0x1006:	lsls	r4, r0, #0x1c
	op_count: 3
		operands[0].type: REG = r4
		operands[1].type: REG = r0
		operands[2].type: IMM = 0x1c
	Update-flags: True

0x1008:	subs	r4, #0x1f
	op_count: 2
		operands[0].type: REG = r4
		operands[1].type: IMM = 0x1f
	Update-flags: True

0x100a:	stm	r0!, {r1, r4, r5, r6, r7}
	op_count: 6
		operands[0].type: REG = r0
		operands[1].type: REG = r1
		operands[2].type: REG = r4
		operands[3].type: REG = r5
		operands[4].type: REG = r6
		operands[5].type: REG = r7
	Write-back: True

0x100c:	movs	r0, r0
	op_count: 2
		operands[0].type: REG = r0
		operands[1].type: REG = r0
	Update-flags: True

0x100e:	mov.w	r1, #0
	op_count: 2
		operands[0].type: REG = r1
		operands[1].type: IMM = 0x0

0x1012:	ldr	r6, [r0, #0x44]
	op_count: 2
		operands[0].type: REG = r6
		operands[1].type: MEM
			operands[1].mem.base: REG = r0
			operands[1].mem.disp: 0x44

0x1014:

****************
Platform: Thumb-2 & register named with numbers
Code:0x4f 0xf0 0x00 0x01 0xbd 0xe8 0x00 0x88 0xd1 0xe8 0x00 0xf0 0x18 0xbf 0xad 0xbf 0xf3 0xff 0x0b 0x0c 0x86 0xf3 0x00 0x89 0x80 0xf3 0x00 0x8c 0x4f 0xfa 0x99 0xf6 0xd0 0xff 0xa2 0x01 
Disasm:
0x1000:	mov.w	r1, #0
	op_count: 2
		operands[0].type: REG = r1
		operands[1].type: IMM = 0x0

0x1004:	pop.w	{r11, pc}
	op_count: 2
		operands[0].type: REG = r11
		operands[1].type: REG = pc

0x1008:	tbb	[r1, r0]
	op_count: 1
		operands[0].type: MEM
			operands[0].mem.base: REG = r1
			operands[0].mem.index: REG = r0

0x100c:	it	ne
	Code condition: 2

0x100e:	iteet	ge
	Code condition: 11

0x1010:	vdupne.8	d16, d11[1]
	op_count: 2
		operands[0].type: REG = d16
		operands[1].type: REG = d11
		operands[1].vector_index = 1
	Code condition: 2
	Vector-size: 8

0x1014:	msr	cpsr_fc, r6
	op_count: 2
		operands[0].type: SYSREG = 144
		operands[1].type: REG = r6

0x1018:	msr	apsr_nzcvqg, r0
	op_count: 2
		operands[0].type: SYSREG = 259
		operands[1].type: REG = r0

0x101c:	sxtb.w	r6, r9, ror #8
	op_count: 2
		operands[0].type: REG = r6
		operands[1].type: REG = r9
			Shift: 4 = 8

0x1020:	vaddw.u16	q8, q8, d18
	op_count: 3
		operands[0].type: REG = q8
		operands[1].type: REG = q8
		operands[2].type: REG = d18
	Vector-data: 10

0x1024:

****************
Platform: Thumb-MClass
Code:0xef 0xf3 0x02 0x80 
Disasm:
0x1000:	mrs	r0, eapsr
	op_count: 2
		operands[0].type: REG = r0
		operands[1].type: SYSREG = 263

0x1004:

****************
Platform: Arm-V8
Code:0xe0 0x3b 0xb2 0xee 0x42 0x00 0x01 0xe1 0x51 0xf0 0x7f 0xf5 
Disasm:
0x1000:	vcvtt.f64.f16	d3, s1
	op_count: 2
		operands[0].type: REG = d3
		operands[1].type: REG = s1
	Vector-data: 17

0x1004:	crc32b	r0, r1, r2
	op_count: 3
		operands[0].type: REG = r0
		operands[1].type: REG = r1
		operands[2].type: REG = r2

0x1008:	dmb	oshld
	Memory-barrier: 2

0x100c:

