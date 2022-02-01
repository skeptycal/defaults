package defaults

import (
	"errors"
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

func AssertTypeEquals(arg Any, typ string) error {
	if a := GetType(arg); a != typ {
		return errors.New(fmt.Sprintf("incorrect type: %v want %v", a, typ))
	}
	return nil
}

func IsString(arg Any) error {
	return AssertTypeEquals(arg, "string")
}

func AssertStringEquals(t *testing.T, testname string, argname string, want Any, got Any) {

	if err := IsString(want); err != nil {
		t.Errorf("%v(%v): %v", testname, argname, err)
		t.FailNow()
	}

	if err := IsString(got); err != nil {
		t.Errorf("%v(%v): %v", testname, argname, err)
		t.FailNow()
	}

	w := want.(string)
	g := got.(string)

	length := len(w)
	if len(g) < len(w) {
		length = len(g)
	}

	for i := 0; i < length; i++ {
		if w[i] != g[i] {
			t.Errorf("%v(%v): first string mismatch at position %d - want: %q  got: %q", testname, argname, i, w[i], g[i])
			break
		}
	}
}

func TestDefaultMapString(t *testing.T) {
	const defaultsStringSample = "Default Settings Map:\nKey                  = Value\ndebugState           = true\ntraceState           = true\n"
	var b = []byte{68, 101, 102, 97, 117, 108, 116, 32, 83, 101, 116, 116, 105, 110, 103, 115, 32, 77, 97, 112, 58, 10, 75, 101, 121, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 61, 32, 86, 97, 108, 117, 101, 10, 100, 101, 98, 117, 103, 83, 116, 97, 116, 101, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 61, 32, 116, 114, 117, 101, 10, 116, 114, 97, 99, 101, 83, 116, 97, 116, 101, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 61, 32, 116, 114, 117, 101, 10}

	AssertEquals(t, "Defaults_String", "", string(b)[:49], Defaults.String()[:49])
}

func ExampleDefaultsString() {
	fmt.Println(Defaults.String())
	/*
				Default Settings Map:
		Key                  = Value
		debugState           = true
		traceState           = true
	*/
}
