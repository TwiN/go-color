package color

import "fmt"

var (
	Reset = "\033[0m"

	/////////////
	// Special //
	/////////////

	Bold      = "\033[1m"
	Underline = "\033[4m"

	/////////////////
	// Text colors //
	/////////////////

	Black  = "\033[30m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"

	///////////////////////
	// Background colors //
	///////////////////////

	BlackBackground  = "\033[40m"
	RedBackground    = "\033[41m"
	GreenBackground  = "\033[42m"
	YellowBackground = "\033[43m"
	BlueBackground   = "\033[44m"
	PurpleBackground = "\033[45m"
	CyanBackground   = "\033[46m"
	GrayBackground   = "\033[47m"
	WhiteBackground  = "\033[107m"
)

var (
	enabled = true
)

// Toggle enables or disables color output
//
// Note that this does not apply if you are coloring by concatenating strings with color variables directly (e.g. color.Red+"Hello"+color.Reset).
// Enabling/disabling coloring only applies to all functions (Colorize, Ize, With, InRed, OverRed, etc.)
func Toggle(enable bool) {
	enabled = enable
}

// Colorize wraps a given message in a given color.
//
// Example:
//
//	println(color.Colorize(color.Red, "This is red"))
func Colorize(color string, s any) string {
	switch s := s.(type) {
	case string:
		if !enabled {
			return s
		}
		return color + s + Reset
	default:
		if !enabled {
			return fmt.Sprint(s)
		}
		return color + fmt.Sprint(s) + Reset
	}
}

// Ize is an alias for the Colorize function
//
// Example:
//
//	println(color.Ize(color.Red, "This is red"))
func Ize(color string, s any) string {
	return Colorize(color, s)
}

// With is an alias for the Colorize function
//
// Example:
//
//	println(color.With(color.Red, "This is red"))
func With(color string, s any) string {
	return Colorize(color, s)
}

// InBold wraps the given string s in Bold
//
// Example:
//
//	println(color.InBold("This is bold"))
func InBold(s any) string {
	return Colorize(Bold, s)
}

// InUnderline wraps the given string s in Underline
//
// Example:
//
//	println(color.InUnderline("This is underlined"))
func InUnderline(s any) string {
	return Colorize(Underline, s)
}

// InBlack wraps the given string s in Black
//
// Example:
//
//	println(color.InBlack("This is black"))
func InBlack(s any) string {
	return Colorize(Black, s)
}

// InRed wraps the given string s in Red
//
// Example:
//
//	println(color.InRed("This is red"))
func InRed(s any) string {
	return Colorize(Red, s)
}

// InGreen wraps the given string s in Green
//
// Example:
//
//	println(color.InGreen("This is green"))
func InGreen(s any) string {
	return Colorize(Green, s)
}

// InYellow wraps the given string s in Yellow
//
// Example:
//
//	println(color.InYellow("This is yellow"))
func InYellow(s any) string {
	return Colorize(Yellow, s)
}

// InBlue wraps the given string s in Blue
//
// Example:
//
//	println(color.InBlue("This is blue"))
func InBlue(s any) string {
	return Colorize(Blue, s)
}

// InPurple wraps the given string s in Purple
//
// Example:
//
//	println(color.InPurple("This is purple"))
func InPurple(s any) string {
	return Colorize(Purple, s)
}

// InCyan wraps the given string s in Cyan
//
// Example:
//
//	println(color.InCyan("This is cyan"))
func InCyan(s any) string {
	return Colorize(Cyan, s)
}

// InGray wraps the given string s in Gray
//
// Example:
//
//	println(color.InGray("This is gray"))
func InGray(s any) string {
	return Colorize(Gray, s)
}

// InWhite wraps the given string s in White
//
// Example:
//
//	println(color.InWhite("This is white"))
func InWhite(s any) string {
	return Colorize(White, s)
}

// OverBlack wraps the given string s in BlackBackground
//
// Example:
//
//	println(color.OverBlack("This is over black"))
func OverBlack(s any) string {
	return Colorize(BlackBackground, s)
}

// OverRed wraps the given string s in RedBackground
//
// Example:
//
//	println(color.OverRed("This is over red"))
func OverRed(s any) string {
	return Colorize(RedBackground, s)
}

// OverGreen wraps the given string s in GreenBackground
//
// Example:
//
//	println(color.OverGreen("This is over green"))
func OverGreen(s any) string {
	return Colorize(GreenBackground, s)
}

// OverYellow wraps the given string s in YellowBackground
//
// Example:
//
//	println(color.OverYellow("This is over yellow"))
func OverYellow(s any) string {
	return Colorize(YellowBackground, s)
}

// OverBlue wraps the given string s in BlueBackground
//
// Example:
//
//	println(color.OverBlue("This is over blue"))
func OverBlue(s any) string {
	return Colorize(BlueBackground, s)
}

// OverPurple wraps the given string s in PurpleBackground
//
// Example:
//
//	println(color.OverPurple("This is over purple"))
func OverPurple(s any) string {
	return Colorize(PurpleBackground, s)
}

// OverCyan wraps the given string s in CyanBackground
//
// Example:
//
//	println(color.OverCyan("This is over cyan"))
func OverCyan(s any) string {
	return Colorize(CyanBackground, s)
}

// OverGray wraps the given string s in GrayBackground
//
// Example:
//
//	println(color.OverGray("This is over gray"))
func OverGray(s any) string {
	return Colorize(GrayBackground, s)
}

// OverWhite wraps the given string s in WhiteBackground
//
// Example:
//
//	println(color.OverWhite("This is over white"))
func OverWhite(s any) string {
	return Colorize(WhiteBackground, s)
}

func InBlackOverBlack(s any) string   { return Colorize(Black+BlackBackground, s) }
func InBlackOverRed(s any) string     { return Colorize(Black+RedBackground, s) }
func InBlackOverGreen(s any) string   { return Colorize(Black+GreenBackground, s) }
func InBlackOverYellow(s any) string  { return Colorize(Black+YellowBackground, s) }
func InBlackOverBlue(s any) string    { return Colorize(Black+BlueBackground, s) }
func InBlackOverPurple(s any) string  { return Colorize(Black+PurpleBackground, s) }
func InBlackOverCyan(s any) string    { return Colorize(Black+CyanBackground, s) }
func InBlackOverGray(s any) string    { return Colorize(Black+GrayBackground, s) }
func InBlackOverWhite(s any) string   { return Colorize(Black+WhiteBackground, s) }
func InRedOverBlack(s any) string     { return Colorize(Red+BlackBackground, s) }
func InRedOverRed(s any) string       { return Colorize(Red+RedBackground, s) }
func InRedOverGreen(s any) string     { return Colorize(Red+GreenBackground, s) }
func InRedOverYellow(s any) string    { return Colorize(Red+YellowBackground, s) }
func InRedOverBlue(s any) string      { return Colorize(Red+BlueBackground, s) }
func InRedOverPurple(s any) string    { return Colorize(Red+PurpleBackground, s) }
func InRedOverCyan(s any) string      { return Colorize(Red+CyanBackground, s) }
func InRedOverGray(s any) string      { return Colorize(Red+GrayBackground, s) }
func InRedOverWhite(s any) string     { return Colorize(Red+WhiteBackground, s) }
func InGreenOverBlack(s any) string   { return Colorize(Green+BlackBackground, s) }
func InGreenOverRed(s any) string     { return Colorize(Green+RedBackground, s) }
func InGreenOverGreen(s any) string   { return Colorize(Green+GreenBackground, s) }
func InGreenOverYellow(s any) string  { return Colorize(Green+YellowBackground, s) }
func InGreenOverBlue(s any) string    { return Colorize(Green+BlueBackground, s) }
func InGreenOverPurple(s any) string  { return Colorize(Green+PurpleBackground, s) }
func InGreenOverCyan(s any) string    { return Colorize(Green+CyanBackground, s) }
func InGreenOverGray(s any) string    { return Colorize(Green+GrayBackground, s) }
func InGreenOverWhite(s any) string   { return Colorize(Green+WhiteBackground, s) }
func InYellowOverBlack(s any) string  { return Colorize(Yellow+BlackBackground, s) }
func InYellowOverRed(s any) string    { return Colorize(Yellow+RedBackground, s) }
func InYellowOverGreen(s any) string  { return Colorize(Yellow+GreenBackground, s) }
func InYellowOverYellow(s any) string { return Colorize(Yellow+YellowBackground, s) }
func InYellowOverBlue(s any) string   { return Colorize(Yellow+BlueBackground, s) }
func InYellowOverPurple(s any) string { return Colorize(Yellow+PurpleBackground, s) }
func InYellowOverCyan(s any) string   { return Colorize(Yellow+CyanBackground, s) }
func InYellowOverGray(s any) string   { return Colorize(Yellow+GrayBackground, s) }
func InYellowOverWhite(s any) string  { return Colorize(Yellow+WhiteBackground, s) }
func InBlueOverBlack(s any) string    { return Colorize(Blue+BlackBackground, s) }
func InBlueOverRed(s any) string      { return Colorize(Blue+RedBackground, s) }
func InBlueOverGreen(s any) string    { return Colorize(Blue+GreenBackground, s) }
func InBlueOverYellow(s any) string   { return Colorize(Blue+YellowBackground, s) }
func InBlueOverBlue(s any) string     { return Colorize(Blue+BlueBackground, s) }
func InBlueOverPurple(s any) string   { return Colorize(Blue+PurpleBackground, s) }
func InBlueOverCyan(s any) string     { return Colorize(Blue+CyanBackground, s) }
func InBlueOverGray(s any) string     { return Colorize(Blue+GrayBackground, s) }
func InBlueOverWhite(s any) string    { return Colorize(Blue+WhiteBackground, s) }
func InPurpleOverBlack(s any) string  { return Colorize(Purple+BlackBackground, s) }
func InPurpleOverRed(s any) string    { return Colorize(Purple+RedBackground, s) }
func InPurpleOverGreen(s any) string  { return Colorize(Purple+GreenBackground, s) }
func InPurpleOverYellow(s any) string { return Colorize(Purple+YellowBackground, s) }
func InPurpleOverBlue(s any) string   { return Colorize(Purple+BlueBackground, s) }
func InPurpleOverPurple(s any) string { return Colorize(Purple+PurpleBackground, s) }
func InPurpleOverCyan(s any) string   { return Colorize(Purple+CyanBackground, s) }
func InPurpleOverGray(s any) string   { return Colorize(Purple+GrayBackground, s) }
func InPurpleOverWhite(s any) string  { return Colorize(Purple+WhiteBackground, s) }
func InCyanOverBlack(s any) string    { return Colorize(Cyan+BlackBackground, s) }
func InCyanOverRed(s any) string      { return Colorize(Cyan+RedBackground, s) }
func InCyanOverGreen(s any) string    { return Colorize(Cyan+GreenBackground, s) }
func InCyanOverYellow(s any) string   { return Colorize(Cyan+YellowBackground, s) }
func InCyanOverBlue(s any) string     { return Colorize(Cyan+BlueBackground, s) }
func InCyanOverPurple(s any) string   { return Colorize(Cyan+PurpleBackground, s) }
func InCyanOverCyan(s any) string     { return Colorize(Cyan+CyanBackground, s) }
func InCyanOverGray(s any) string     { return Colorize(Cyan+GrayBackground, s) }
func InCyanOverWhite(s any) string    { return Colorize(Cyan+WhiteBackground, s) }
func InGrayOverBlack(s any) string    { return Colorize(Gray+BlackBackground, s) }
func InGrayOverRed(s any) string      { return Colorize(Gray+RedBackground, s) }
func InGrayOverGreen(s any) string    { return Colorize(Gray+GreenBackground, s) }
func InGrayOverYellow(s any) string   { return Colorize(Gray+YellowBackground, s) }
func InGrayOverBlue(s any) string     { return Colorize(Gray+BlueBackground, s) }
func InGrayOverPurple(s any) string   { return Colorize(Gray+PurpleBackground, s) }
func InGrayOverCyan(s any) string     { return Colorize(Gray+CyanBackground, s) }
func InGrayOverGray(s any) string     { return Colorize(Gray+GrayBackground, s) }
func InGrayOverWhite(s any) string    { return Colorize(Gray+WhiteBackground, s) }
func InWhiteOverBlack(s any) string   { return Colorize(White+BlackBackground, s) }
func InWhiteOverRed(s any) string     { return Colorize(White+RedBackground, s) }
func InWhiteOverGreen(s any) string   { return Colorize(White+GreenBackground, s) }
func InWhiteOverYellow(s any) string  { return Colorize(White+YellowBackground, s) }
func InWhiteOverBlue(s any) string    { return Colorize(White+BlueBackground, s) }
func InWhiteOverPurple(s any) string  { return Colorize(White+PurpleBackground, s) }
func InWhiteOverCyan(s any) string    { return Colorize(White+CyanBackground, s) }
func InWhiteOverGray(s any) string    { return Colorize(White+GrayBackground, s) }
func InWhiteOverWhite(s any) string   { return Colorize(White+WhiteBackground, s) }
