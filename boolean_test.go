package defaults

import (
	"testing"

	"golang.org/x/tools/go/pointer"
)

func TestAnyBool(t *testing.T) {

	var fakePtr pointer.Pointer
	tests := []struct {
		name  string
		input Any
		want  string
	}{
		// TODO: Add test cases.
		{"the empty string", "", "false"},
		{"0", 0, "false"},
		{"0.0", 0.0, "false"},
		{"false", false, "false"},
		{"zero", nil, "false"},
		{"empty slice", []byte{}, "false"},
		{"empty map", make(map[string]string), "false"},
		{"empty chan", make(chan string), "false"},
		{"empty array", [4]byte{}, "false"},
		{"nil pointer", fakePtr, "false"},

		{"empty struct (thing{})", thing{}, "false"},

		{"true string", "fake", "true"},
		{"42", 42, "true"},
		{"42.0", 42.0, "true"},
		{"true", true, "true"},
		{"slice", []byte{'m', 'n'}, "true"},
		{"map", map[string]string{"m": "n"}, "true"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyBool(tt.input).String(); got != tt.want {
				t.Errorf("AnyBool(%v) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestIPAddr_String(t *testing.T) {
	tests := []struct {
		name string
		i    IPAddr
		want string
	}{
		// TODO: Add test cases.
		{"127.0.0.1", IPAddr{127, 0, 0, 1}, "127.0.0.1"},
		{"8.8.8.8", IPAddr{8, 8, 8, 8}, "8.8.8.8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("IPAddr.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
