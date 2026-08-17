package ir

import (
	"strings"
)

// CaptureComponent specifies a capture component for the captures(...) attribute.
// Added in LLVM 21 to replace nocapture.
type CaptureComponent uint8

// Capture components.
const (
	CaptureComponentNone           CaptureComponent = 0
	CaptureComponentAddress        CaptureComponent = 1 << 0 // address
	CaptureComponentAddressIsNull  CaptureComponent = 1 << 1 // address_is_null
	CaptureComponentReadProvenance CaptureComponent = 1 << 2 // read_provenance
)

// AttrCaptures is the captures(...) parameter attribute (LLVM 21+).
// Replaces the nocapture attribute.
// Syntax: captures(none) or captures(address, address_is_null)
type AttrCaptures struct {
	// Capture mask. Zero means captures(none).
	Mask CaptureComponent
	// Return capture mask (for ret: prefix).
	RetMask CaptureComponent
}

// NewAttrCaptures returns a captures attribute with the given mask.
func NewAttrCaptures(mask CaptureComponent) *AttrCaptures {
	return &AttrCaptures{Mask: mask}
}

// NoCapture returns a captures(none) attribute, equivalent to the old nocapture.
func NoCapture() *AttrCaptures {
	return &AttrCaptures{Mask: CaptureComponentNone}
}

// String returns the LLVM syntax representation of the captures attribute.
func (a *AttrCaptures) String() string {
	if a.Mask == CaptureComponentNone && a.RetMask == CaptureComponentNone {
		return "captures(none)"
	}
	buf := &strings.Builder{}
	buf.WriteString("captures(")
	first := true
	if a.RetMask != 0 {
		buf.WriteString("ret: ")
		buf.WriteString(captureComponentString(a.RetMask))
		first = false
	}
	if a.Mask != 0 {
		if !first {
			buf.WriteString(", ")
		}
		buf.WriteString(captureComponentString(a.Mask))
	}
	buf.WriteString(")")
	return buf.String()
}

func captureComponentString(m CaptureComponent) string {
	var parts []string
	if m&CaptureComponentAddress != 0 {
		parts = append(parts, "address")
	}
	if m&CaptureComponentAddressIsNull != 0 {
		parts = append(parts, "address_is_null")
	}
	if m&CaptureComponentReadProvenance != 0 {
		parts = append(parts, "read_provenance")
	}
	return strings.Join(parts, ", ")
}

// IsParamAttribute ensures that only parameter attributes can be assigned to
// the ir.ParamAttribute interface.
func (*AttrCaptures) IsParamAttribute() {}

// Ensure AttrCaptures implements ParamAttribute.
var _ ParamAttribute = (*AttrCaptures)(nil)
