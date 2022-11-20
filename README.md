# go-color

![test](https://github.com/TwiN/go-color/workflows/test/badge.svg?branch=master)

An extremely lightweight cross-platform package to colorize text in terminals.

This is not meant for maximal compatibility, nor is it meant to handle a plethora of scenarios.
It will simply wrap a message with the necessary characters, if the OS handles it.

There are many cases in which this would not work, such as the output being redirected to something other 
than a terminal (such as a file, i.e. `executable >> file.txt`)


## Usage

```console
go get github.com/TwiN/go-color
```


### Function

You can use the `color.Colorize(color, message)` or `color.Ize(color, message)` function 
in conjunction with a variable like so:
```go
package main

import "github.com/TwiN/go-color"

func main() {
    println(color.Ize(color.Red, "This is red"))
    // Or if you prefer the longer version:
    println(color.Colorize(color.Red, "This is red"))
    // Or if you prefer the non-parameterized version:
    println(color.InRed("This is red"))
}
```

Because I felt reading `color.Ize()` to be more visually pleasant than `color.Colorize()`, 
I included `Ize()` as an alias for `Colorize()`.

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


### Variables only

Unlike using the aforementioned approach, directly using the color variables will require you to manually
prepend `color.Reset` at the end of your string.

You can either directly use the variables like so:

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
