package color

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
	Bold   = "\033[1m"
	Orange = "\033[93m"
)

// Ize is an alias for the Colorize function
func Ize(color, message string) string {
	return Colorize(color, message)
}

// Colorize wraps a given message in a given color.
func Colorize(color, message string) string {
	return color + message + Reset
}
