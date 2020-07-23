package interfaces

type ExitCode int

const (
	ConfigParseError ExitCode = 1
	ConfigNotFound
	UnavailableTags
	UnparseableTags
)

func (d ExitCode) String() string {
	return [...]string{"Config can't be parsed", "Config not found", "Can't fetch tags", "Can't parse tags"}[d]
}
