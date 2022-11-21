package color

import "testing"

// Clearly not a real test - just a visual reference
func TestOutput(t *testing.T) {
	println(Bold + "Hello" + Reset)
	println(Underline + "Hello" + Reset)
	println(Black + "Hello" + Reset)
	println(Red + "Hello" + Reset)
	println(Green + "Hello" + Reset)
	println(Yellow + "Hello" + Reset)
	println(Blue + "Hello" + Reset)
	println(Purple + "Hello" + Reset)
	println(Cyan + "Hello" + Reset)
	println(Gray + "Hello" + Reset)
	println(White + "Hello" + Reset)
	println(BlackBackground + "Hello" + Reset)
	println(Red + "H" + Green + "el" + Purple + "lo" + Reset)
	println(Ize(Red, "test"))
}

func TestColorize(t *testing.T) {
	scenarios := []struct {
		Color          string
		ExpectedOutput string
	}{
		{Color: Bold, ExpectedOutput: "\033[1mtest\033[0m"},
		{Color: Underline, ExpectedOutput: "\033[4mtest\033[0m"},
		{Color: Black, ExpectedOutput: "\033[30mtest\033[0m"},
		{Color: Red, ExpectedOutput: "\033[31mtest\033[0m"},
		{Color: Green, ExpectedOutput: "\033[32mtest\033[0m"},
		{Color: Yellow, ExpectedOutput: "\033[33mtest\033[0m"},
		{Color: Blue, ExpectedOutput: "\033[34mtest\033[0m"},
		{Color: Purple, ExpectedOutput: "\033[35mtest\033[0m"},
		{Color: Cyan, ExpectedOutput: "\033[36mtest\033[0m"},
		{Color: Gray, ExpectedOutput: "\033[37mtest\033[0m"},
		{Color: White, ExpectedOutput: "\033[97mtest\033[0m"},
		{Color: BlackBackground, ExpectedOutput: "\033[40mtest\033[0m"},
		{Color: RedBackground, ExpectedOutput: "\033[41mtest\033[0m"},
		{Color: GreenBackground, ExpectedOutput: "\033[42mtest\033[0m"},
		{Color: YellowBackground, ExpectedOutput: "\033[43mtest\033[0m"},
		{Color: BlueBackground, ExpectedOutput: "\033[44mtest\033[0m"},
		{Color: PurpleBackground, ExpectedOutput: "\033[45mtest\033[0m"},
		{Color: CyanBackground, ExpectedOutput: "\033[46mtest\033[0m"},
		{Color: GrayBackground, ExpectedOutput: "\033[47mtest\033[0m"},
		{Color: WhiteBackground, ExpectedOutput: "\033[107mtest\033[0m"},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.Color, func(t *testing.T) {
			output := Colorize(scenario.Color, "test")
			if output != scenario.ExpectedOutput {
				t.Errorf("expected %s, got %s", scenario.ExpectedOutput, output)
			}
			// Ize is an alias of Colorize, therefore the result should be the same
			usingIze := Ize(scenario.Color, "test")
			if usingIze != scenario.ExpectedOutput {
				t.Errorf("expected %s, got %s", scenario.ExpectedOutput, usingIze)
			}
			// With is an alias of Colorize, therefore the result should be the same
			usingWith := With(scenario.Color, "test")
			if usingWith != scenario.ExpectedOutput {
				t.Errorf("expected %s, got %s", scenario.ExpectedOutput, usingWith)
			}
		})
	}
}

func TestIn(t *testing.T) {
	scenarios := []struct {
		Func           func(any) string
		Name           string
		ExpectedOutput string
	}{
		{Name: "bold", Func: InBold, ExpectedOutput: "\033[1mtest\033[0m"},
		{Name: "underline", Func: InUnderline, ExpectedOutput: "\033[4mtest\033[0m"},
		{Name: "black", Func: InBlack, ExpectedOutput: "\033[30mtest\033[0m"},
		{Name: "red", Func: InRed, ExpectedOutput: "\033[31mtest\033[0m"},
		{Name: "green", Func: InGreen, ExpectedOutput: "\033[32mtest\033[0m"},
		{Name: "yellow", Func: InYellow, ExpectedOutput: "\033[33mtest\033[0m"},
		{Name: "blue", Func: InBlue, ExpectedOutput: "\033[34mtest\033[0m"},
		{Name: "purple", Func: InPurple, ExpectedOutput: "\033[35mtest\033[0m"},
		{Name: "cyan", Func: InCyan, ExpectedOutput: "\033[36mtest\033[0m"},
		{Name: "gray", Func: InGray, ExpectedOutput: "\033[37mtest\033[0m"},
		{Name: "white", Func: InWhite, ExpectedOutput: "\033[97mtest\033[0m"},
		{Name: "bg-black", Func: OverBlack, ExpectedOutput: "\033[40mtest\033[0m"},
		{Name: "bg-red", Func: OverRed, ExpectedOutput: "\033[41mtest\033[0m"},
		{Name: "bg-green", Func: OverGreen, ExpectedOutput: "\033[42mtest\033[0m"},
		{Name: "bg-yellow", Func: OverYellow, ExpectedOutput: "\033[43mtest\033[0m"},
		{Name: "bg-blue", Func: OverBlue, ExpectedOutput: "\033[44mtest\033[0m"},
		{Name: "bg-purple", Func: OverPurple, ExpectedOutput: "\033[45mtest\033[0m"},
		{Name: "bg-cyan", Func: OverCyan, ExpectedOutput: "\033[46mtest\033[0m"},
		{Name: "bg-gray", Func: OverGray, ExpectedOutput: "\033[47mtest\033[0m"},
		{Name: "bg-white", Func: OverWhite, ExpectedOutput: "\033[107mtest\033[0m"},
		{Name: "black-over-black", Func: InBlackOverBlack, ExpectedOutput: "\033[30m\033[40mtest\033[0m"},
		{Name: "black-over-red", Func: InBlackOverRed, ExpectedOutput: "\033[30m\033[41mtest\033[0m"},
		{Name: "black-over-green", Func: InBlackOverGreen, ExpectedOutput: "\033[30m\033[42mtest\033[0m"},
		{Name: "black-over-yellow", Func: InBlackOverYellow, ExpectedOutput: "\033[30m\033[43mtest\033[0m"},
		{Name: "black-over-blue", Func: InBlackOverBlue, ExpectedOutput: "\033[30m\033[44mtest\033[0m"},
		{Name: "black-over-purple", Func: InBlackOverPurple, ExpectedOutput: "\033[30m\033[45mtest\033[0m"},
		{Name: "black-over-cyan", Func: InBlackOverCyan, ExpectedOutput: "\033[30m\033[46mtest\033[0m"},
		{Name: "black-over-gray", Func: InBlackOverGray, ExpectedOutput: "\033[30m\033[47mtest\033[0m"},
		{Name: "black-over-white", Func: InBlackOverWhite, ExpectedOutput: "\033[30m\033[107mtest\033[0m"},
		{Name: "red-over-black", Func: InRedOverBlack, ExpectedOutput: "\033[31m\033[40mtest\033[0m"},
		{Name: "red-over-red", Func: InRedOverRed, ExpectedOutput: "\033[31m\033[41mtest\033[0m"},
		{Name: "red-over-green", Func: InRedOverGreen, ExpectedOutput: "\033[31m\033[42mtest\033[0m"},
		{Name: "red-over-yellow", Func: InRedOverYellow, ExpectedOutput: "\033[31m\033[43mtest\033[0m"},
		{Name: "red-over-blue", Func: InRedOverBlue, ExpectedOutput: "\033[31m\033[44mtest\033[0m"},
		{Name: "red-over-purple", Func: InRedOverPurple, ExpectedOutput: "\033[31m\033[45mtest\033[0m"},
		{Name: "red-over-cyan", Func: InRedOverCyan, ExpectedOutput: "\033[31m\033[46mtest\033[0m"},
		{Name: "red-over-gray", Func: InRedOverGray, ExpectedOutput: "\033[31m\033[47mtest\033[0m"},
		{Name: "red-over-white", Func: InRedOverWhite, ExpectedOutput: "\033[31m\033[107mtest\033[0m"},
		{Name: "green-over-black", Func: InGreenOverBlack, ExpectedOutput: "\033[32m\033[40mtest\033[0m"},
		{Name: "green-over-red", Func: InGreenOverRed, ExpectedOutput: "\033[32m\033[41mtest\033[0m"},
		{Name: "green-over-green", Func: InGreenOverGreen, ExpectedOutput: "\033[32m\033[42mtest\033[0m"},
		{Name: "green-over-yellow", Func: InGreenOverYellow, ExpectedOutput: "\033[32m\033[43mtest\033[0m"},
		{Name: "green-over-blue", Func: InGreenOverBlue, ExpectedOutput: "\033[32m\033[44mtest\033[0m"},
		{Name: "green-over-purple", Func: InGreenOverPurple, ExpectedOutput: "\033[32m\033[45mtest\033[0m"},
		{Name: "green-over-cyan", Func: InGreenOverCyan, ExpectedOutput: "\033[32m\033[46mtest\033[0m"},
		{Name: "green-over-gray", Func: InGreenOverGray, ExpectedOutput: "\033[32m\033[47mtest\033[0m"},
		{Name: "green-over-white", Func: InGreenOverWhite, ExpectedOutput: "\033[32m\033[107mtest\033[0m"},
		{Name: "yellow-over-black", Func: InYellowOverBlack, ExpectedOutput: "\033[33m\033[40mtest\033[0m"},
		{Name: "yellow-over-red", Func: InYellowOverRed, ExpectedOutput: "\033[33m\033[41mtest\033[0m"},
		{Name: "yellow-over-green", Func: InYellowOverGreen, ExpectedOutput: "\033[33m\033[42mtest\033[0m"},
		{Name: "yellow-over-yellow", Func: InYellowOverYellow, ExpectedOutput: "\033[33m\033[43mtest\033[0m"},
		{Name: "yellow-over-blue", Func: InYellowOverBlue, ExpectedOutput: "\033[33m\033[44mtest\033[0m"},
		{Name: "yellow-over-purple", Func: InYellowOverPurple, ExpectedOutput: "\033[33m\033[45mtest\033[0m"},
		{Name: "yellow-over-cyan", Func: InYellowOverCyan, ExpectedOutput: "\033[33m\033[46mtest\033[0m"},
		{Name: "yellow-over-gray", Func: InYellowOverGray, ExpectedOutput: "\033[33m\033[47mtest\033[0m"},
		{Name: "yellow-over-white", Func: InYellowOverWhite, ExpectedOutput: "\033[33m\033[107mtest\033[0m"},
		{Name: "blue-over-black", Func: InBlueOverBlack, ExpectedOutput: "\033[34m\033[40mtest\033[0m"},
		{Name: "blue-over-red", Func: InBlueOverRed, ExpectedOutput: "\033[34m\033[41mtest\033[0m"},
		{Name: "blue-over-green", Func: InBlueOverGreen, ExpectedOutput: "\033[34m\033[42mtest\033[0m"},
		{Name: "blue-over-yellow", Func: InBlueOverYellow, ExpectedOutput: "\033[34m\033[43mtest\033[0m"},
		{Name: "blue-over-blue", Func: InBlueOverBlue, ExpectedOutput: "\033[34m\033[44mtest\033[0m"},
		{Name: "blue-over-purple", Func: InBlueOverPurple, ExpectedOutput: "\033[34m\033[45mtest\033[0m"},
		{Name: "blue-over-cyan", Func: InBlueOverCyan, ExpectedOutput: "\033[34m\033[46mtest\033[0m"},
		{Name: "blue-over-gray", Func: InBlueOverGray, ExpectedOutput: "\033[34m\033[47mtest\033[0m"},
		{Name: "blue-over-white", Func: InBlueOverWhite, ExpectedOutput: "\033[34m\033[107mtest\033[0m"},
		{Name: "purple-over-black", Func: InPurpleOverBlack, ExpectedOutput: "\033[35m\033[40mtest\033[0m"},
		{Name: "purple-over-red", Func: InPurpleOverRed, ExpectedOutput: "\033[35m\033[41mtest\033[0m"},
		{Name: "purple-over-green", Func: InPurpleOverGreen, ExpectedOutput: "\033[35m\033[42mtest\033[0m"},
		{Name: "purple-over-yellow", Func: InPurpleOverYellow, ExpectedOutput: "\033[35m\033[43mtest\033[0m"},
		{Name: "purple-over-blue", Func: InPurpleOverBlue, ExpectedOutput: "\033[35m\033[44mtest\033[0m"},
		{Name: "purple-over-purple", Func: InPurpleOverPurple, ExpectedOutput: "\033[35m\033[45mtest\033[0m"},
		{Name: "purple-over-cyan", Func: InPurpleOverCyan, ExpectedOutput: "\033[35m\033[46mtest\033[0m"},
		{Name: "purple-over-gray", Func: InPurpleOverGray, ExpectedOutput: "\033[35m\033[47mtest\033[0m"},
		{Name: "purple-over-white", Func: InPurpleOverWhite, ExpectedOutput: "\033[35m\033[107mtest\033[0m"},
		{Name: "cyan-over-black", Func: InCyanOverBlack, ExpectedOutput: "\033[36m\033[40mtest\033[0m"},
		{Name: "cyan-over-red", Func: InCyanOverRed, ExpectedOutput: "\033[36m\033[41mtest\033[0m"},
		{Name: "cyan-over-green", Func: InCyanOverGreen, ExpectedOutput: "\033[36m\033[42mtest\033[0m"},
		{Name: "cyan-over-yellow", Func: InCyanOverYellow, ExpectedOutput: "\033[36m\033[43mtest\033[0m"},
		{Name: "cyan-over-blue", Func: InCyanOverBlue, ExpectedOutput: "\033[36m\033[44mtest\033[0m"},
		{Name: "cyan-over-purple", Func: InCyanOverPurple, ExpectedOutput: "\033[36m\033[45mtest\033[0m"},
		{Name: "cyan-over-cyan", Func: InCyanOverCyan, ExpectedOutput: "\033[36m\033[46mtest\033[0m"},
		{Name: "cyan-over-gray", Func: InCyanOverGray, ExpectedOutput: "\033[36m\033[47mtest\033[0m"},
		{Name: "cyan-over-white", Func: InCyanOverWhite, ExpectedOutput: "\033[36m\033[107mtest\033[0m"},
		{Name: "gray-over-black", Func: InGrayOverBlack, ExpectedOutput: "\033[37m\033[40mtest\033[0m"},
		{Name: "gray-over-red", Func: InGrayOverRed, ExpectedOutput: "\033[37m\033[41mtest\033[0m"},
		{Name: "gray-over-green", Func: InGrayOverGreen, ExpectedOutput: "\033[37m\033[42mtest\033[0m"},
		{Name: "gray-over-yellow", Func: InGrayOverYellow, ExpectedOutput: "\033[37m\033[43mtest\033[0m"},
		{Name: "gray-over-blue", Func: InGrayOverBlue, ExpectedOutput: "\033[37m\033[44mtest\033[0m"},
		{Name: "gray-over-purple", Func: InGrayOverPurple, ExpectedOutput: "\033[37m\033[45mtest\033[0m"},
		{Name: "gray-over-cyan", Func: InGrayOverCyan, ExpectedOutput: "\033[37m\033[46mtest\033[0m"},
		{Name: "gray-over-gray", Func: InGrayOverGray, ExpectedOutput: "\033[37m\033[47mtest\033[0m"},
		{Name: "gray-over-white", Func: InGrayOverWhite, ExpectedOutput: "\033[37m\033[107mtest\033[0m"},
		{Name: "white-over-black", Func: InWhiteOverBlack, ExpectedOutput: "\033[97m\033[40mtest\033[0m"},
		{Name: "white-over-red", Func: InWhiteOverRed, ExpectedOutput: "\033[97m\033[41mtest\033[0m"},
		{Name: "white-over-green", Func: InWhiteOverGreen, ExpectedOutput: "\033[97m\033[42mtest\033[0m"},
		{Name: "white-over-yellow", Func: InWhiteOverYellow, ExpectedOutput: "\033[97m\033[43mtest\033[0m"},
		{Name: "white-over-blue", Func: InWhiteOverBlue, ExpectedOutput: "\033[97m\033[44mtest\033[0m"},
		{Name: "white-over-purple", Func: InWhiteOverPurple, ExpectedOutput: "\033[97m\033[45mtest\033[0m"},
		{Name: "white-over-cyan", Func: InWhiteOverCyan, ExpectedOutput: "\033[97m\033[46mtest\033[0m"},
		{Name: "white-over-gray", Func: InWhiteOverGray, ExpectedOutput: "\033[97m\033[47mtest\033[0m"},
		{Name: "white-over-white", Func: InWhiteOverWhite, ExpectedOutput: "\033[97m\033[107mtest\033[0m"},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.Name, func(t *testing.T) {
			output := scenario.Func("test")
			if output != scenario.ExpectedOutput {
				t.Errorf("Expected %s, got %s", scenario.ExpectedOutput, output)
			}
			println(output)
		})
	}
}

func TestToggle(t *testing.T) {
	Toggle(false)
	if output := InRed("test"); output != "test" {
		t.Errorf("Expected %s, got %s", "test", output)
	}
	if output := InRed(123); output != "123" {
		t.Errorf("Expected %s, got %s", "test", output)
	}
	Toggle(true)
	if output := InRed("test"); output != "\033[31mtest\033[0m" {
		t.Errorf("Expected %s, got %s", "\033[31mtest\033[0m", output)
	}
	if output := InRed(123); output != "\033[31m123\033[0m" {
		t.Errorf("Expected %s, got %s", "\033[31mtest\033[0m", output)
	}
}
