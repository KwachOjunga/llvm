package ir

import (
	"fmt"
	"strings"
)

// MemoryEffect specifies the memory effect (none, read, write, readwrite).
// Used in the memory(...) function attribute (LLVM 16+).
type MemoryEffect uint8

const (
	MemoryEffectNone      MemoryEffect = 0 // none
	MemoryEffectRead      MemoryEffect = 1 // read
	MemoryEffectWrite     MemoryEffect = 2 // write
	MemoryEffectReadWrite MemoryEffect = 3 // readwrite
)

func (e MemoryEffect) String() string {
	switch e {
	case MemoryEffectNone:
		return "none"
	case MemoryEffectRead:
		return "read"
	case MemoryEffectWrite:
		return "write"
	case MemoryEffectReadWrite:
		return "readwrite"
	default:
		return fmt.Sprintf("MemoryEffect(%d)", uint8(e))
	}
}

// AttrMemory is the memory(...) function/call attribute (LLVM 16+).
// Replaces the old readnone, readonly, writeonly, argmemonly, inaccessiblememonly,
// and inaccessiblemem_or_argmemonly attributes.
// Syntax: memory(read) or memory(argmem: write, inaccessiblemem: read)
type AttrMemory struct {
	// Default effect for unspecified locations.
	Default MemoryEffect
	// Effect for argument memory (argmem).
	ArgMem MemoryEffect
	// Effect for inaccessible memory (inaccessiblemem).
	InaccessibleMem MemoryEffect
	// Whether ArgMem is explicitly set.
	ArgMemSet bool
	// Whether InaccessibleMem is explicitly set.
	InaccessibleMemSet bool
}

// NewAttrMemory returns a memory attribute with the given default effect.
func NewAttrMemory(defaultEffect MemoryEffect) *AttrMemory {
	return &AttrMemory{Default: defaultEffect}
}

// String returns the LLVM syntax representation of the memory attribute.
func (a *AttrMemory) String() string {
	// If no per-location overrides, emit simple form.
	if !a.ArgMemSet && !a.InaccessibleMemSet {
		return fmt.Sprintf("memory(%s)", a.Default)
	}
	buf := &strings.Builder{}
	buf.WriteString("memory(")
	first := true
	if a.ArgMemSet {
		if !first {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "argmem: %s", a.ArgMem)
		first = false
	}
	if a.InaccessibleMemSet {
		if !first {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "inaccessiblemem: %s", a.InaccessibleMem)
		first = false
	}
	if a.Default != MemoryEffectNone {
		if !first {
			buf.WriteString(", ")
		}
		buf.WriteString(a.Default.String())
	}
	buf.WriteString(")")
	return buf.String()
}

// Ensure AttrMemory implements FuncAttribute.
var _ FuncAttribute = (*AttrMemory)(nil)
