package defaults

type (
	defaultMap struct {
		Debug   bool
		TraceOn bool
	}
)

var Defaults defaultMap = defaultMap{
	Debug:   true,
	TraceOn: true,
}

func IsDebug() bool { return Defaults.Debug }
func IsTrace() bool { return Defaults.TraceOn }
