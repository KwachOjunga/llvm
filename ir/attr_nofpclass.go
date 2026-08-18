package ir

import (
	"fmt"
	"strings"
)

// FPClassMask is a bitmask of floating-point value classes for the nofpclass attribute.
// Added in LLVM 17.
type FPClassMask uint32

// FP class mask bits.
const (
	FPClassSNaN    FPClassMask = 1 << 0 // snan
	FPClassQNaN    FPClassMask = 1 << 1 // qnan
	FPClassNegInf  FPClassMask = 1 << 2 // ninf
	FPClassNegNorm FPClassMask = 1 << 3 // nnorm
	FPClassNegSub  FPClassMask = 1 << 4 // nsub
	FPClassNegZero FPClassMask = 1 << 5 // nzero
	FPClassPosZero FPClassMask = 1 << 6 // pzero
	FPClassPosSub  FPClassMask = 1 << 7 // psub
	FPClassPosNorm FPClassMask = 1 << 8 // pnorm
	FPClassPosInf  FPClassMask = 1 << 9 // pinf
	// Composite masks.
	FPClassNaN  = FPClassSNaN | FPClassQNaN       // nan
	FPClassInf  = FPClassNegInf | FPClassPosInf   // inf
	FPClassNorm = FPClassNegNorm | FPClassPosNorm // norm
	FPClassSub  = FPClassNegSub | FPClassPosSub   // sub
	FPClassZero = FPClassNegZero | FPClassPosZero // zero
)

// AttrNoFPClass is the nofpclass(...) parameter or return attribute.
// Indicates that the value is guaranteed NOT to be any of the specified FP classes.
// Added in LLVM 17.
type AttrNoFPClass struct {
	Mask FPClassMask
}

// NewAttrNoFPClass returns a nofpclass attribute with the given FP class mask.
func NewAttrNoFPClass(mask FPClassMask) *AttrNoFPClass {
	return &AttrNoFPClass{Mask: mask}
}

// String returns the LLVM syntax representation of the nofpclass attribute.
func (a *AttrNoFPClass) String() string {
	return fmt.Sprintf("nofpclass(%s)", fpClassMaskString(a.Mask))
}

func fpClassMaskString(m FPClassMask) string {
	// Check composite masks first for brevity.
	var parts []string
	remaining := m
	if remaining&FPClassNaN == FPClassNaN {
		parts = append(parts, "nan")
		remaining &^= FPClassNaN
	} else {
		if remaining&FPClassSNaN != 0 {
			parts = append(parts, "snan")
			remaining &^= FPClassSNaN
		}
		if remaining&FPClassQNaN != 0 {
			parts = append(parts, "qnan")
			remaining &^= FPClassQNaN
		}
	}
	if remaining&FPClassInf == FPClassInf {
		parts = append(parts, "inf")
		remaining &^= FPClassInf
	} else {
		if remaining&FPClassNegInf != 0 {
			parts = append(parts, "ninf")
			remaining &^= FPClassNegInf
		}
		if remaining&FPClassPosInf != 0 {
			parts = append(parts, "pinf")
			remaining &^= FPClassPosInf
		}
	}
	if remaining&FPClassNorm == FPClassNorm {
		parts = append(parts, "norm")
		remaining &^= FPClassNorm
	} else {
		if remaining&FPClassNegNorm != 0 {
			parts = append(parts, "nnorm")
			remaining &^= FPClassNegNorm
		}
		if remaining&FPClassPosNorm != 0 {
			parts = append(parts, "pnorm")
			remaining &^= FPClassPosNorm
		}
	}
	if remaining&FPClassSub == FPClassSub {
		parts = append(parts, "sub")
		remaining &^= FPClassSub
	} else {
		if remaining&FPClassNegSub != 0 {
			parts = append(parts, "nsub")
			remaining &^= FPClassNegSub
		}
		if remaining&FPClassPosSub != 0 {
			parts = append(parts, "psub")
			remaining &^= FPClassPosSub
		}
	}
	if remaining&FPClassZero == FPClassZero {
		parts = append(parts, "zero")
		remaining &^= FPClassZero
	} else {
		if remaining&FPClassNegZero != 0 {
			parts = append(parts, "nzero")
			remaining &^= FPClassNegZero
		}
		if remaining&FPClassPosZero != 0 {
			parts = append(parts, "pzero")
			remaining &^= FPClassPosZero
		}
	}
	return strings.Join(parts, " ")
}

// Ensure AttrNoFPClass implements ParamAttribute and ReturnAttribute.
var (
	_ ParamAttribute  = (*AttrNoFPClass)(nil)
	_ ReturnAttribute = (*AttrNoFPClass)(nil)
)
