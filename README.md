# go-color
![test](https://github.com/TwiN/go-color/workflows/test/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/TwiN/go-color)](https://goreportcard.com/report/github.com/TwiN/go-color)
[![Go version](https://img.shields.io/github/go-mod/go-version/TwiN/go-color.svg)](https://github.com/TwiN/go-color)
[![Go Reference](https://pkg.go.dev/badge/github.com/TwiN/go-color.svg)](https://pkg.go.dev/github.com/TwiN/go-color)
[![Follow TwiN](https://img.shields.io/github/followers/TwiN?label=Follow&style=social)](https://github.com/TwiN)

An extremely lightweight cross-platform package to colorize text in terminals.

This is not meant for maximal compatibility, nor is it meant to handle a plethora of scenarios.
It will simply wrap a message with the necessary characters, if the OS handles it.


## Usage
```console
go get github.com/TwiN/go-color
```


### Using Functions
You can use the `color.Colorize(color, str)`, the `color.Ize(color, str)`, or the `color.With(color, str)` function
in conjunction with a variable like so:
```go
package main

import "github.com/TwiN/go-color"

func main() {
    // These all have the same effect:
    println(color.With(color.Red, "This is red"))
    println(color.Ize(color.Red, "This is red"))
    println(color.Colorize(color.Red, "This is red"))
    println(color.InRed("This is red"))
}
```

Because I felt reading `color.With()`/`color.Ize()` to be more visually pleasant than `color.Colorize()`, 
I included `Ize()` and `With()` as aliases for `Colorize()`.

I'm not usually a big fan of having two methods doing the same thing, but since
this package doesn't have much room for growth (its only purpose is to colorize
terminal outputs, after all, and there aren't hundreds of ways to go about it),
I believe it's acceptable to have both.

Alternatively, you can use color-specific functions:
```go
package main

import "github.com/TwiN/go-color"

func main() {
    // Special
    println(color.InBold("This is bold"))
    println(color.InUnderline("This is underlined"))
    // Text colors
    println(color.InBlack("This is in black"))
    println(color.InRed("This is in red"))
    println(color.InGreen("This is in green"))
    println(color.InYellow("This is in yellow"))
    println(color.InBlue("This is in blue"))
    println(color.InPurple("This is in purple"))
    println(color.InCyan("This is in cyan"))
    println(color.InGray("This is in gray"))
    println(color.InWhite("This is in white"))
    // Background colors
    println(color.OverBlack("This is over a black background"))
    println(color.OverRed("This is over a red background"))
    println(color.OverGreen("This is over a green background"))
    println(color.OverYellow("This is over a yellow background"))
    println(color.OverBlue("This is over a blue background"))
    println(color.OverPurple("This is over a purple background"))
    println(color.OverCyan("This is over a cyan background"))
    println(color.OverGray("This is over a gray background"))
    println(color.OverWhite("This is over a white background"))
}
```

If you wish, you can also combine text and background colors by using the `In<color>Over<color>` functions, e.g.:
```go
color.InWhiteOverBlack("This is white text on a black background")
color.InRedOverYellow("This is red text on a yellow background")
color.InGreenOverBlue("This is green text on a blue background")
```
Note that the example above is not exhaustive.

#### Automatic coloring
> ⚠ **WARNING**: This is an experimental feature and may be removed depending on user feedback.
> If you have any feedback, please comment on [poll: Keep or Remove color.Autof?](https://github.com/TwiN/go-color/discussions/13).

You may use `color.Autof(string, args...)` as a replacement to `fmt.Sprintf(string, args...)` but with colors for
each argument:
```go
// 20 will be in red while "cookie jar" will be in green.
println(color.Autof("I have $%.02f in my %s", 20, "cookie jar"))
```


### Using Variables
> ⚠ **WARNING**: By using this approach, you will not be able to disable colors with `color.Toggle(false)`. 
> Consider using the function approach instead for maximal cross-library compatibility.

Unlike using the aforementioned approach, directly using the color variables will require you to manually
prepend `color.Reset` at the end of your string.

You can directly use the variables like so:
```go
package main

import "github.com/TwiN/go-color"

func main() {
    println(color.Bold + "This is bold" + color.Reset)
    println(color.Underline + "This is underlined" + color.Reset)
    println(color.Black + "This is black" + color.Reset)
    println(color.Red + "This is red" + color.Reset)
    println(color.Green + "This is green" + color.Reset)
    println(color.Yellow + "This is yellow" + color.Reset)
    println(color.Blue + "This is blue" + color.Reset)
    println(color.Purple + "This is purple" + color.Reset)
    println(color.Cyan + "This is cyan" + color.Reset)
    println(color.Gray + "This is gray" + color.Reset)
    println(color.White + "This is white" + color.Reset)
}
```

**NOTE**: If you're going to use this approach, don't forget to prepend your string with `color.Reset`, 
otherwise everything else in your terminal will be that color until the color is reset or overridden.


### Disabling Colors
You can disable colors by using `color.Toggle(false)`, which will cause all functions to not colorize the text.

As mentioned in [Using Variables](#using-variables), disabling colors will have no effect on the variables, so 
make sure you are using functions if you have a need to toggle colors.


### Examples
```go
println("My name is " + color.InGreen("John Doe") + " and I have " + color.InRed(32) + " apples.")
```
