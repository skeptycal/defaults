package defaults

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const (
	defaultDebugState = true
	defaultTraceState = true
)

var Defaults DefaultMapper = NewDefaults()

type Setting interface {
	Booler
}

func NewSetting(key string, value Any) Setting { return AnyBooler(value) }

type (
	defaultMap map[string]Setting

	DefaultMapper interface {
		GetSetter
		Stringer
	}
)

func (d defaultMap) Set(key Any, value Any) error {
	switch v := key.(type) {
	case string:
		d[v] = AnyBooler(value)
		return nil
	default:
		return fmt.Errorf("key type not string")
	}
}

func (d defaultMap) Get(key Any) (Any, error) {
	if CheckType(key, "string") {
		return nil, fmt.Errorf("key type not string")
	}
	switch v := key.(type) {
	case string:
		if v, ok := d[v]; ok {
			return v, nil
		}
		return nil, errors.New("key not found: " + v)
		// d[v] = AnyBooler(value)
		// return nil
	default:
		return nil, fmt.Errorf("key type not string")
	}

}

func GetKind(any Any) string             { return reflect.ValueOf(any).Kind().String() }
func CheckType(any Any, typ string) bool { return GetType(any) == typ }

func GetType(any Any) string {
	if any == nil {
		return "nil"
	}
	return reflect.ValueOf(any).Type().String()
}

func (d defaultMap) String() string {

	const format = "%-20s = %-20v\n"
	sb := &strings.Builder{}
	defer sb.Reset()

	fmt.Fprintln(sb, "Default Settings Map:")
	fmt.Fprintf(sb, format, "Key", "Value")

	for key, value := range d {
		fmt.Fprintf(sb, format, key, value)
	}

	return sb.String()
}

func (d defaultMap) IsDebug() bool { return d["debugState"].AsBool() }
func (d defaultMap) IsTrace() bool { return d["traceState"].AsBool() }

func NewDefaults() DefaultMapper {
	m := make(defaultMap, 2)
	m.Set("debugState", defaultDebugState)
	m.Set("traceState", defaultTraceState)
	return &m
}

func IsTrue(v Any) bool  { return AnyBooler(v).AsBool() }
func IsFalse(v Any) bool { return !AnyBooler(v).AsBool() }

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
