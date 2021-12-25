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
	scenarios := []struct {
		Color          string
		ExpectedOutput string
	}{
		{
			Color:          Bold,
			ExpectedOutput: "\033[1mtest\033[0m",
		},
		{
			Color:          Red,
			ExpectedOutput: "\033[31mtest\033[0m",
		},
		{
			Color:          Green,
			ExpectedOutput: "\033[32mtest\033[0m",
		},
		{
			Color:          Yellow,
			ExpectedOutput: "\033[33mtest\033[0m",
		},
		{
			Color:          Blue,
			ExpectedOutput: "\033[34mtest\033[0m",
		},
		{
			Color:          Purple,
			ExpectedOutput: "\033[35mtest\033[0m",
		},
		{
			Color:          Cyan,
			ExpectedOutput: "\033[36mtest\033[0m",
		},
		{
			Color:          Gray,
			ExpectedOutput: "\033[37mtest\033[0m",
		},
		{
			Color:          White,
			ExpectedOutput: "\033[97mtest\033[0m",
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.Color, func(t *testing.T) {
			output := Colorize(scenario.Color, "test")
			if output != scenario.ExpectedOutput {
				t.Errorf("Expected %s, got %s", scenario.ExpectedOutput, output)
			}
		})
	}
}

func TestIn(t *testing.T) {
	scenarios := []struct {
		Func           func(string) string
		Color          string
		ExpectedOutput string
	}{
		{
			Func:           InBold,
			Color:          Bold,
			ExpectedOutput: "\033[1mtest\033[0m",
		},
		{
			Func:           InRed,
			Color:          Red,
			ExpectedOutput: "\033[31mtest\033[0m",
		},
		{
			Func:           InGreen,
			Color:          Green,
			ExpectedOutput: "\033[32mtest\033[0m",
		},
		{
			Func:           InYellow,
			Color:          Yellow,
			ExpectedOutput: "\033[33mtest\033[0m",
		},
		{
			Func:           InBlue,
			Color:          Blue,
			ExpectedOutput: "\033[34mtest\033[0m",
		},
		{
			Func:           InPurple,
			Color:          Purple,
			ExpectedOutput: "\033[35mtest\033[0m",
		},
		{
			Func:           InCyan,
			Color:          Cyan,
			ExpectedOutput: "\033[36mtest\033[0m",
		},
		{
			Func:           InGray,
			Color:          Gray,
			ExpectedOutput: "\033[37mtest\033[0m",
		},
		{
			Func:           InWhite,
			Color:          White,
			ExpectedOutput: "\033[97mtest\033[0m",
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.Color, func(t *testing.T) {
			output := scenario.Func("test")
			if output != scenario.ExpectedOutput {
				t.Errorf("Expected %s, got %s", scenario.ExpectedOutput, output)
			}
		})
	}
}
