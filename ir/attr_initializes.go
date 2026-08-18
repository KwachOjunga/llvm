package ir

import (
	"fmt"
	"strings"
)

// ByteRange represents a [lo, hi) byte range for the initializes attribute.
type ByteRange struct {
	Lo int64
	Hi int64
}

// String returns the LLVM syntax representation of the byte range.
func (r ByteRange) String() string {
	return fmt.Sprintf("(%d, %d)", r.Lo, r.Hi)
}

// AttrInitializes is the initializes(...) parameter attribute (LLVM 20+).
// Indicates that a pointer parameter is initialized in the specified byte ranges.
// Syntax: initializes((0, 4), (8, 12))
//
// See: https://llvm.org/docs/LangRef.html#parameter-attributes
type AttrInitializes struct {
	Ranges []ByteRange
}

// NewAttrInitializes returns an initializes attribute with the given byte ranges.
func NewAttrInitializes(ranges ...ByteRange) *AttrInitializes {
	return &AttrInitializes{Ranges: ranges}
}

// String returns the LLVM syntax representation of the initializes attribute.
func (a *AttrInitializes) String() string {
	buf := &strings.Builder{}
	buf.WriteString("initializes(")
	for i, r := range a.Ranges {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(r.String())
	}
	buf.WriteString(")")
	return buf.String()
}

// Ensure AttrInitializes implements ParamAttribute.
var _ ParamAttribute = (*AttrInitializes)(nil)
