// Copyright 2018 Massimiliano Ghilardi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gc

#include "go_asm.h"
#include "textflag.h"
#include "../../src/runtime/go_tls.h"

// +build gc

TEXT ·goroutinePtr(SB),NOSPLIT,$0-8
	get_tls(CX)
	MOVQ g(CX), AX
	MOVQ AX, goid+0(FP)
	RET
