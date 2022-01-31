package defaults

import (
	"fmt"
	"testing"
)

func TestCheckType(t *testing.T) {
	type args struct {
		any Any
		typ string
	}
	tests := []struct {
		name  string
		input Any
		typ   string
		knd   string
		want  bool
	}{
		{"string", "fake", "string", "string", true},
		{"int", 42, "int", "int", true},
		{"nil", nil, "nil", "invalid", true},
		{"0", 0, "int", "int", true},
		{"bool", true, "bool", "bool", true},
		{"empty []int32", []rune{}, "[]int32", "slice", true},
		{"[]uint8", []byte("fake"), "[]uint8", "slice", true},
	}
	for _, tt := range tests {
		AssertEquals(t, "CheckType", tt.name+", "+tt.typ, tt.want, CheckType(tt.input, tt.typ))
		AssertEquals(t, "GetType", tt.name, tt.typ, GetType(tt.input))
		AssertEquals(t, "GetKind", tt.name, tt.knd, GetKind(tt.input))
	}
}

func AssertEquals(t *testing.T, testname string, argname string, want Any, got Any) {
	name := testname + "(" + argname + ")"
	t.Run(name, func(t *testing.T) {
		if got != want {
			t.Errorf("%v = %v, want %v", name, got, want)
		}
	})
}

func ExampleDefaultMapString() {
	fmt.Println(Defaults.String())
	// output:
	// Default Settings Map:
	// Key                  = Value
	// debugState           = true
	// traceState           = true
	//

}
