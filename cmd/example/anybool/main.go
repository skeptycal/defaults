package main

import (
	"fmt"

	"github.com/skeptycal/defaults"
)

type Booler = defaults.Booler

var (
	pf = fmt.Printf
	ab = defaults.AnyBooler

	samples = []struct {
		name string
		b    Booler
	}{
		{"boolean true", ab(true)},
		{"boolean false", ab(false)},
		{"integer 0", ab(int(0))},
		{"integer 42", ab(int(42))},
		{"empty string", ab("")},
		{"'false' string", ab("false")},
		{"'true' string", ab("true")},
		{"'0' string", ab("0")},
		{"nil", ab(nil)},
		{"empty slice", ab([]rune{})},
		{"non-empty slice", ab([]rune{'4', '2'})},
	}
)

func main() {
	for _, ss := range samples {
		pf("%v: %v\n", ss.name, ss.b)
	}
}
