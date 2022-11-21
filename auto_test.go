package color

import "testing"

func TestAutof(t *testing.T) {
	scenarios := []struct {
		InputFormat    string
		InputArguments []any
		ExpectedOutput string
	}{
		{
			InputFormat:    "no args, no colors.",
			InputArguments: nil,
			ExpectedOutput: "no args, no colors.",
		},
		{
			InputFormat:    "hello, %s.",
			InputArguments: []any{"world"},
			ExpectedOutput: "hello, " + Red + "world" + Reset + ".",
		},
		{
			InputFormat:    "%s got %d%% in math.",
			InputArguments: []any{"John Doe", 87},
			ExpectedOutput: InRed("John Doe") + " got " + InGreen(87) + "% in math.",
		},
		{
			InputFormat:    "%s sent %s$%.02f to %s",
			InputArguments: []any{"John", "USD", 123.4, "Jane"},
			ExpectedOutput: InRed("John") + " sent " + InGreen("USD") + "$" + InYellow("123.40") + " to " + InBlue("Jane"),
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.InputFormat, func(t *testing.T) {
			output := Autof(scenario.InputFormat, scenario.InputArguments...)
			if output != scenario.ExpectedOutput {
				t.Errorf("expected %s, got %s", scenario.ExpectedOutput, output)
			}
			println(output)
		})
	}
}
