package goxc

var (
	counter       = 0
	packagePrefix = ""
	version       = "TRUNK"
	rev           = "unknown"
	built         = "19990101-235959"
)

func Version() string {
	return version
}

func Rev() string {
	return rev
}

func Built() string {
	return built
}
