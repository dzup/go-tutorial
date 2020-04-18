package object

import (
	"fmt"
	"sort"
	"strings"
)

// Integer wraps int64 and implements Object and Hashable interfaces.
type Integer struct {
	// Value holds the integer value this object wraps
	Value int64
}

// Inspect returns a string-representation of the given object.
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type returns the type of this object.
func (i *Integer) Type() Type {
	return INTEGER_OBJ
}

// HashKey returns a hash key for the given object.
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (i *Integer) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "chr" {
		return &String{Value: string(i.Value)}
	}
	if method == "methods" {
		static := []string{"chr", "methods"}
		dynamic := env.Names("integer.")

		var names []string
		for _, e := range static {
			names = append(names, e)
		}
		for _, e := range dynamic {
			bits := strings.Split(e, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)

		result := make([]Object, len(names), len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	return nil
}
