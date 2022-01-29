package defaults

import (
	"testing"
)

/* Benchmark results for Booler (bool value vs interface value)

Takeaways:
- bools are very, very fast
- pointers are very fast
- ints, uints, floats, struct{}, maps, slices(not []byte), channels are fast
- strings and []byte are half speed of above (even other slices???)
- arrays are noticeably slower than all others

BenchmarkBooler/NewBool(true)-8         	196670611	         5.918 ns/op	   0 B/op	       0 allocs/op
BenchmarkBooler/AnyBool(bool)-8         	 34330753	        47.42  ns/op	   0 B/op	       0 allocs/op

BenchmarkBooler/AnyBool(false)-8           165912350	         7.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/AnyBool(true)-8            163697052	         7.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/AnyBool(true)-8          	55380460	        22.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/the_empty_string-8       	26632684	        51.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/'false'-8                	19048916	        58.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/'0'-8                    	26050334	        59.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/true_string-8            	24274688	        55.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/empty_[]byte-8           	22117113	        61.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/non-empty_[]byte-8       	22498663	        62.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int_0-8                  	30002030	        48.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int8_0-8                 	28536974	        43.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int16_0-8                	32840272	        38.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int32_0-8                	42892950	        29.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int64_0-8                	40022901	        29.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int_42-8                 	37534357	        31.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int8_42-8                	42608340	        30.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int16_42-8               	40746220	        28.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int32_42-8               	42150926	        28.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int64_42-8               	37750736	        28.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint_0-8                 	49273642	        24.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint8_0-8                	51577038	        29.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint16_0-8               	41422041	        26.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint32_0-8               	48347780	        24.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint64_0-8               	46554181	        27.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint_42-8                	42202501	        27.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint8_42-8               	46550343	        26.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint16_42-8              	38536478	        31.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint32_42-8              	39620086	        30.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint64_42-8              	34818856	        37.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/float_0.0-8              	41392989	        26.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/float32_0.0-8            	41017986	        29.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/float64_0.0-8            	41653528	        25.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/float_42.0-8             	44531652	        27.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/float32_42.0-8           	46420838	        27.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/float64_42.0-8           	30215378	        33.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/empty_struct_(thing{})-8 	32121490	        36.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/false_struct-8           	24694210	        45.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/true_struct-8            	25428420	        48.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/empty_slice-8            	35387748	        35.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/slice-8                  	33835500	        35.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/empty_array-8            	11728260	       102.3  ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/array-8                  	14829880	        70.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/empty_map-8              	26005053	        41.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/map[int]byte-8           	32912179	        36.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/map[string]string-8      	33690870	        37.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/empty_chan-8             	32297919	        36.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/nil_pointer-8            	77658804	        15.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/true_pointer_(not_nil)-8 	76069932	        15.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/false_pointer_(not_nil)-8   81939463	        15.67 ns/op	       0 B/op	       0 allocs/op
*/

var true_bool bool = true
var truePtr *bool = &true_bool

var false_bool bool = false
var falsePtr *bool = &false_bool

var nilPtr *uintptr = nil
var retval = ""

var boolerTests = []struct {
	name  string
	input Booler
	want  string
}{
	// nil
	{"nil", AnyBooler(nil), "false"},

	// bools
	{"AnyBool(false)", AnyBooler(false), "false"},
	{"AnyBool(true)", AnyBooler(true), "true"},

	// strings
	{"the empty string", AnyBooler(""), "false"},
	{"'false'", AnyBooler("false"), "false"},
	{"'0'", AnyBooler("0"), "false"},
	{"true string", AnyBooler("fake"), "true"},

	// []byte
	{"empty []byte", AnyBooler([]byte("")), "false"},
	{"non-empty []byte", AnyBooler([]byte("fake")), "true"},

	// ints
	{"int 0", AnyBooler(int(0)), "false"},
	{"int8 0", AnyBooler(int8(0)), "false"},
	{"int16 0", AnyBooler(int16(0)), "false"},
	{"int32 0", AnyBooler(int32(0)), "false"},
	{"int64 0", AnyBooler(int64(0)), "false"},
	{"int 42", AnyBooler(int(42)), "true"},
	{"int8 42", AnyBooler(int8(42)), "true"},
	{"int16 42", AnyBooler(int16(42)), "true"},
	{"int32 42", AnyBooler(int32(42)), "true"},
	{"int64 42", AnyBooler(int64(42)), "true"},

	// uints
	{"uint 0", AnyBooler(uint(0)), "false"},
	{"uint8 0", AnyBooler(uint8(0)), "false"},
	{"uint16 0", AnyBooler(uint16(0)), "false"},
	{"uint32 0", AnyBooler(uint32(0)), "false"},
	{"uint64 0", AnyBooler(uint64(0)), "false"},
	{"uint 42", AnyBooler(uint(42)), "true"},
	{"uint8 42", AnyBooler(uint8(42)), "true"},
	{"uint16 42", AnyBooler(uint16(42)), "true"},
	{"uint32 42", AnyBooler(uint32(42)), "true"},
	{"uint64 42", AnyBooler(uint64(42)), "true"},

	// floats
	{"float 0.0", AnyBooler(0.0), "false"},
	{"float32 0.0", AnyBooler(float32(0.0)), "false"},
	{"float64 0.0", AnyBooler(float64(0.0)), "false"},
	{"float 42.0", AnyBooler(42.0), "true"},
	{"float32 42.0", AnyBooler(float32(42.0)), "true"},
	{"float64 42.0", AnyBooler(float64(42.0)), "true"},
	// structs
	{"empty struct (thing{})", AnyBooler(thing{}), "false"},
	{"false struct", AnyBooler(struct{ i int }{i: 0}), "false"},
	{"true struct", AnyBooler(struct{ i int }{i: 1}), "true"},

	// slices
	{"empty slice", AnyBooler([]rune{}), "false"},
	{"slice", AnyBooler([]rune{'4', '2'}), "true"},

	// arrays
	{"empty array", AnyBooler([4]byte{}), "false"},
	{"array", AnyBooler([2]byte{'4', '2'}), "true"},

	// maps
	{"empty map", AnyBooler(make(map[string]string)), "false"},
	{"map[int]byte", AnyBooler(map[int]byte{1: '4', 2: '2'}), "true"},
	{"map[string]string", AnyBooler(map[string]string{"4": "2"}), "true"},

	// channels
	{"empty chan", AnyBooler(make(chan string)), "false"},

	// pointers
	{"nil pointer", AnyBooler(nilPtr), "false"},
	{"true pointer (not nil)", AnyBooler(truePtr), "true"},
	{"false pointer (not nil)", AnyBooler(falsePtr), "true"},
}

var newBoolTests = boolerTests[1:2]

func BenchmarkBooler(b *testing.B) {
	for _, bb := range boolerTests {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				retval = bb.input.String()
			}
		})
	}
}

func Test_boolean_Enable(t *testing.T) {
	want := "true"
	for _, tt := range newBoolTests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.input
			b.Enable()
			if got := b.String(); got != want {
				t.Errorf("AnyBool(%v) = %v, want %v", tt.name, got, want)
			}
		})
	}
}

func Test_boolean_Disable(t *testing.T) {
	want := "false"
	for _, tt := range newBoolTests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.input
			b.Disable()
			if got := b.String(); got != want {
				t.Errorf("NewBool(%v) = %v, want %v", tt.name, got, want)
			}
		})
	}
}

func TestAnyBool(t *testing.T) {
	for _, tt := range boolerTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.String(); got != tt.want {
				// v := reflect.ValueOf(tt.input)
				// ty := v.Type()
				// k := v.Kind()
				// s := ty.String()
				// t.Errorf("NewBool(%v, %v, %v, %v) = %v, want %v", tt.name, v, k, s, got, tt.want)
				t.Errorf("NewBool(%v) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_anyBool_Enable(t *testing.T) {
	want := "true"

	for _, tt := range boolerTests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.input
			b.Enable()
			if got := b.String(); got != want {
				t.Errorf("NewBool(%v) = %v, want %v", tt.name, got, want)
			}
		})
	}
}

func Test_anyBool_Disable(t *testing.T) {
	want := "false"

	for _, tt := range boolerTests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.input
			b.Disable()
			if got := b.String(); got != want {
				t.Errorf("NewBool(%v) = %v, want %v", tt.name, got, want)
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
