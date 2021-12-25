package color

var (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

// Ize is an alias for the Colorize function
//
// Example:
//     println(color.Ize(color.Red, "This is red"))
func Ize(color, s string) string {
	return Colorize(color, s)
}

// Colorize wraps a given message in a given color.
//
// Example:
//     println(color.Colorize(color.Red, "This is red"))
func Colorize(color, s string) string {
	return color + s + Reset
}

// InBold wraps the given string s in Bold
//
// Example:
//     println(color.InBold("This is bold"))
func InBold(s string) string {
	return Colorize(Bold, s)
}

// InRed wraps the given string s in Red
//
// Example:
//     println(color.InRed("This is red"))
func InRed(s string) string {
	return Colorize(Red, s)
}

// InGreen wraps the given string s in Green
//
// Example:
//     println(color.InGreen("This is green"))
func InGreen(s string) string {
	return Colorize(Green, s)
}

// InYellow wraps the given string s in Yellow
//
// Example:
//     println(color.InYellow("This is yellow"))
func InYellow(s string) string {
	return Colorize(Yellow, s)
}

// InBlue wraps the given string s in Blue
//
// Example:
//     println(color.InBlue("This is blue"))
func InBlue(s string) string {
	return Colorize(Blue, s)
}

// InPurple wraps the given string s in Purple
//
// Example:
//     println(color.InPurple("This is purple"))
func InPurple(s string) string {
	return Colorize(Purple, s)
}

// InCyan wraps the given string s in Cyan
//
// Example:
//     println(color.InCyan("This is cyan"))
func InCyan(s string) string {
	return Colorize(Cyan, s)
}

// InGray wraps the given string s in Gray
//
// Example:
//     println(color.InGray("This is gray"))
func InGray(s string) string {
	return Colorize(Gray, s)
}

// InWhite wraps the given string s in White
//
// Example:
//     println(color.InWhite("This is white"))
func InWhite(s string) string {
	return Colorize(White, s)
}
