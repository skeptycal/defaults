package defaults

import "reflect"

// NewBool returns a new boolean value that implements Booler
func NewBool(defaultValue bool) Booler { return &boolean{defaultValue} }

// AnyBool returns a new anyBool value that implements Booler
func AnyBool(defaultValue Any) Booler { return &anyBool{defaultValue} }

// Booler represents a boolean value that implements Enabler
// and Stringer interfaces.
type Booler interface {
	Enabler
	Stringer
}

// boolean represents a boolean value that implements Booler
type boolean struct{ bool }

func (b *boolean) Enable()  { b.bool = true }
func (b *boolean) Disable() { b.bool = false }
func (b *boolean) String() string {
	if b.bool {
		return "true"
	}
	return "false"
}

// anyBool represents any type of value that can be converted
// to a boolean to implement the Booler interfaces.
type anyBool struct{ any Any }

func (b *anyBool) Enable() {
	switch b.any.(type) {
	case bool:
		b.any = true
	case int, uint:
		b.any = 1
	case float32, float64:
		b.any = 1.0
	case string:
		b.any = "true"
	case []byte:
		b.any = []byte("true")
	default:
		b.any = true
	}
}

func (b *anyBool) Disable() {
	switch b.any.(type) {
	case bool:
		b.any = false
	case int, uint:
		b.any = 0
	case float32, float64:
		b.any = 0.0
	case string:
		b.any = "false"
	case []byte:
		b.any = []byte("false")
	default:
		b.any = nil
	}
}

func boolString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

/*
// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)
*/

func (b *anyBool) String() (s string) {
	return boolString(b.Bool())
}

func (b *anyBool) Bool() bool {
	if b.any == nil {
		return false
	}

	v := reflect.ValueOf(b.any)
	k := v.Kind()

	if k == reflect.UnsafePointer {
		return false
	}

	if k == reflect.Invalid {
		return false
	}

	if v.IsZero() {
		return false
	}

	if k == reflect.Bool {
		return v.Bool()
	}

	if k > 1 && k < 7 {
		return v.Int() != 0
	}

	if k > 6 && k < 12 {
		return v.Uint() != 0
	}

	if k == reflect.Float32 || k == reflect.Float64 {
		return v.Float() != 0.0
	}

	if k == reflect.Complex64 || k == reflect.Complex128 {
		return v.Complex() != 0.0
	}

	if k == reflect.Array || k == reflect.Map || k == reflect.Chan || k == reflect.Slice || k == reflect.String {
		return v.Len() != 0
	}

	if k == reflect.Ptr || k == reflect.Interface || k == reflect.Func {
		return !v.IsNil()
	}

	if k == reflect.Struct {
		return v.NumField() != 0
	}

	return false
}
