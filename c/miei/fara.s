	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 12, 0	sdk_version 12, 3
	.section	__TEXT,__literal8,8byte_literals
	.p2align	3                               ; -- Begin function main
lCPI0_0:
	.quad	0x3f847ae147ae147b              ; double 0.01
	.section	__TEXT,__text,regular,pure_instructions
	.globl	_main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:
	sub	sp, sp, #80
	stp	x29, x30, [sp, #64]             ; 16-byte Folded Spill
	add	x29, sp, #64
	.cfi_def_cfa w29, 16
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	adrp	x8, ___stack_chk_guard@GOTPAGE
	ldr	x8, [x8, ___stack_chk_guard@GOTPAGEOFF]
	ldr	x8, [x8]
	stur	x8, [x29, #-8]
	str	wzr, [sp, #28]
	str	wzr, [sp, #24]
	b	LBB0_1
LBB0_1:                                 ; =>This Inner Loop Header: Depth=1
	ldr	w8, [sp, #24]
	subs	w8, w8, #6
	b.ge	LBB0_4
	b	LBB0_2
LBB0_2:                                 ;   in Loop: Header=BB0_1 Depth=1
	adrp	x0, l_.str@PAGE
	add	x0, x0, l_.str@PAGEOFF
	mov	x9, sp
	add	x8, sp, #20
	str	x8, [x9]
	bl	_scanf
	ldr	s0, [sp, #20]
	ldrsw	x9, [sp, #24]
	add	x8, sp, #32
	str	s0, [x8, x9, lsl #2]
	b	LBB0_3
LBB0_3:                                 ;   in Loop: Header=BB0_1 Depth=1
	ldr	w8, [sp, #24]
	add	w8, w8, #1
	str	w8, [sp, #24]
	b	LBB0_1
LBB0_4:
	str	wzr, [sp, #16]
	b	LBB0_5
LBB0_5:                                 ; =>This Inner Loop Header: Depth=1
	ldr	w8, [sp, #16]
	subs	w8, w8, #6
	b.ge	LBB0_10
	b	LBB0_6
LBB0_6:                                 ;   in Loop: Header=BB0_5 Depth=1
	ldrsw	x9, [sp, #16]
	add	x8, sp, #32
	ldr	s0, [x8, x9, lsl #2]
	fcvt	d0, s0
	adrp	x8, lCPI0_0@PAGE
	ldr	d1, [x8, lCPI0_0@PAGEOFF]
	fcmp	d0, d1
	b.le	LBB0_8
	b	LBB0_7
LBB0_7:                                 ;   in Loop: Header=BB0_5 Depth=1
	ldrsw	x9, [sp, #16]
	add	x8, sp, #32
	ldr	s0, [x8, x9, lsl #2]
	fcvt	d0, s0
	adrp	x0, l_.str.1@PAGE
	add	x0, x0, l_.str.1@PAGEOFF
	mov	x8, sp
	str	d0, [x8]
	bl	_printf
	b	LBB0_8
LBB0_8:                                 ;   in Loop: Header=BB0_5 Depth=1
	b	LBB0_9
LBB0_9:                                 ;   in Loop: Header=BB0_5 Depth=1
	ldr	w8, [sp, #16]
	add	w8, w8, #1
	str	w8, [sp, #16]
	b	LBB0_5
LBB0_10:
	adrp	x0, l_.str.2@PAGE
	add	x0, x0, l_.str.2@PAGEOFF
	bl	_printf
	adrp	x8, ___stack_chk_guard@GOTPAGE
	ldr	x8, [x8, ___stack_chk_guard@GOTPAGEOFF]
	ldr	x8, [x8]
	ldur	x9, [x29, #-8]
	subs	x8, x8, x9
	b.ne	LBB0_12
	b	LBB0_11
LBB0_11:
	mov	w0, #0
	ldp	x29, x30, [sp, #64]             ; 16-byte Folded Reload
	add	sp, sp, #80
	ret
LBB0_12:
	bl	___stack_chk_fail
	.cfi_endproc
                                        ; -- End function
	.section	__TEXT,__cstring,cstring_literals
l_.str:                                 ; @.str
	.asciz	"%f"

l_.str.1:                               ; @.str.1
	.asciz	"%.1f "

l_.str.2:                               ; @.str.2
	.asciz	"\n"

.subsections_via_symbols
