package color

type Name string

var (
	Reset  Name = "\033[0m"
	Red    Name = "\033[31m"
	Green  Name = "\033[32m"
	Yellow Name = "\033[33m"
	Blue   Name = "\033[34m"
	Purple Name = "\033[35m"
	Cyan   Name = "\033[36m"
	Gray   Name = "\033[37m"
	White  Name = "\033[97m"
)

func Ize(color Name, message string) string {
	return Colorize(color, message)
}

func Colorize(color Name, message string) string {
	return string(color) + message + string(Reset)
}
