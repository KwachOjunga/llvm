package constant

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/types"
)

// ConstantPtrAuth is an LLVM IR ptrauth constant (LLVM 20+).
// Used for pointer authentication (ARM PAC/AUTIA).
// Syntax: ptrauth (ptr @fn, i32 0) or ptrauth (ptr @fn, i32 0, i64 0, ptr @disc)
//
// refs:
//
//	https://llvm.org/docs/LangRef.html#pointer-authentication-constants
type ConstantPtrAuth struct {
	// Base pointer (typically a function pointer).
	Ptr Constant
	// Key integer (i32).
	Key Constant
	// (optional) Discriminator; nil if not present.
	Disc Constant
	// (optional) Address discriminator pointer; nil if not present.
	AddrDisc Constant

	// extra.

	// Type of result produced by the constant.
	Typ types.Type
}

// NewConstantPtrAuth returns a new ptrauth constant.
func NewConstantPtrAuth(ptr, key Constant) *ConstantPtrAuth {
	c := &ConstantPtrAuth{Ptr: ptr, Key: key}
	// Compute type.
	c.Type()
	return c
}

// String returns the LLVM syntax representation of the constant as a
// type-value pair.
func (c *ConstantPtrAuth) String() string {
	return fmt.Sprintf("%s %s", c.Type(), c.Ident())
}

// Type returns the type of the constant (always a pointer type).
func (c *ConstantPtrAuth) Type() types.Type {
	if c.Typ == nil {
		c.Typ = c.Ptr.Type()
	}
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *ConstantPtrAuth) Ident() string {
	// 'ptrauth' '(' Ptr=TypeConst ',' Key=TypeConst (',' Disc=TypeConst (',' AddrDisc=TypeConst)?)? ')'
	buf := &strings.Builder{}
	buf.WriteString("ptrauth (")
	fmt.Fprintf(buf, "%s, %s", c.Ptr, c.Key)
	if c.Disc != nil {
		fmt.Fprintf(buf, ", %s", c.Disc)
		if c.AddrDisc != nil {
			fmt.Fprintf(buf, ", %s", c.AddrDisc)
		}
	}
	buf.WriteString(")")
	return buf.String()
}
