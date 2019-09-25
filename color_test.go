package color

import "testing"

// Clearly not a real test - just a visual reference
func TestOutput(t *testing.T) {
	println(Red + "Hello" + Reset)
	println(Green + "Hello" + Reset)
	println(Yellow + "Hello" + Reset)
	println(Blue + "Hello" + Reset)
	println(Purple + "Hello" + Reset)
	println(Cyan + "Hello" + Reset)
	println(Gray + "Hello" + Reset)
	println(White + "Hello" + Reset)
	println(Red + "H" + Green + "el" + Purple + "lo" + Reset)
	println(Ize(Red, "test"))
}

func TestColorize(t *testing.T) {
	if Colorize(Red, "red") != "\033[31mred\033[0m" {
		t.Fail()
	}
	if Colorize(Green, "green") != "\033[32mgreen\033[0m" {
		t.Fail()
	}
	if Colorize(Blue, "blue") != "\033[34mblue\033[0m" {
		t.Fail()
	}
}
