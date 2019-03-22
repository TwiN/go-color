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
