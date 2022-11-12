	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 12, 0	sdk_version 12, 0
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:
	sub	sp, sp, #32                     ; =32
	stp	x29, x30, [sp, #16]             ; 16-byte Folded Spill
	add	x29, sp, #16                    ; =16
	.cfi_def_cfa w29, 16
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	adrp	x0, l_.str@PAGE
	add	x0, x0, l_.str@PAGEOFF
	bl	_printf
	adrp	x0, l_.str.1@PAGE
	add	x0, x0, l_.str.1@PAGEOFF
	mov	x9, sp
	sub	x8, x29, #4                     ; =4
	str	x8, [x9]
	bl	_scanf
	str	wzr, [sp, #8]
LBB0_1:                                 ; =>This Inner Loop Header: Depth=1
	ldr	w8, [sp, #8]
	ldur	w9, [x29, #-4]
	subs	w8, w8, w9
	b.ge	LBB0_4
; %bb.2:                                ;   in Loop: Header=BB0_1 Depth=1
	ldr	w8, [sp, #8]
	add	w10, w8, #1                     ; =1
	adrp	x0, l_.str.2@PAGE
	add	x0, x0, l_.str.2@PAGEOFF
	mov	x9, sp
                                        ; implicit-def: $x8
	mov	x8, x10
	str	x8, [x9]
	bl	_printf
; %bb.3:                                ;   in Loop: Header=BB0_1 Depth=1
	ldr	w8, [sp, #8]
	add	w8, w8, #1                      ; =1
	str	w8, [sp, #8]
	b	LBB0_1
LBB0_4:
	ldp	x29, x30, [sp, #16]             ; 16-byte Folded Reload
	add	sp, sp, #32                     ; =32
	ret
	.cfi_endproc
                                        ; -- End function
	.section	__TEXT,__cstring,cstring_literals
l_.str:                                 ; @.str
	.asciz	"Quante volte? "

l_.str.1:                               ; @.str.1
	.asciz	"%d"

l_.str.2:                               ; @.str.2
	.asciz	"%d Ciao simo culo \n"

.subsections_via_symbols
