package defaults

import (
	"fmt"
	"sync"
)

const (
	NL    = "\n"
	TAB   = "\t"
	SPACE = " "
)

type (

	// thing is an empty stuct used for preallocating zero resource objects
	thing struct{}

	// Any represents a object that may contain any
	// valid type.
	Any interface{}

	// Stringer is implemented by any value that has a
	// String method, which defines the “native” format
	// for that value. The String method is used to print
	// values passed as an operand to any format that accepts
	// a string or to an unformatted printer such as Print.
	//	 type Stringer interface {
	//	 	String() string
	//	 }
	Stringer = fmt.Stringer

	// A Locker represents an object that can be locked and unlocked.
	//	 type Locker interface {
	//	 	Lock()
	//	 	Unlock()
	//	 }
	Locker = sync.Locker

	// An Enabler represents an object that can be enabled or disabled.
	Enabler interface {
		Enable()
		Disable()
	}

	// A GetSetter represents an object that can be accessed using
	// Get and Set methods to access an underlying map.
	GetSetter interface {
		Get(key Any) (Any, error)
		Set(key Any, value Any) error
	}

	// A Printer implements common printing functions.
	Printer interface {
		Print(args ...interface{}) (n int, err error)
		Println(args ...interface{}) (n int, err error)
		Printf(format string, args ...interface{}) (n int, err error)
	}
)
