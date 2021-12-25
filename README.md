# go-color

![build](https://github.com/TwiN/go-color/workflows/build/badge.svg?branch=master)

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
    // Or if you prefer the longer version
    println(color.Colorize(color.Red, "This is red"))
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
    println(color.InBold("This is bold"))
    println(color.InRed("This is red"))
    println(color.InGreen("This is green"))
    println(color.InYellow("This is yellow"))
    println(color.InBlue("This is blue"))
    println(color.InPurple("This is purple"))
    println(color.InCyan("This is cyan"))
    println(color.InGray("This is gray"))
    println(color.InWhite("This is white"))
}
```


### Variables only

Unlike using the aforementioned approach, directly using the color variables will required you to manually
prepend `color.Reset` at the end of your string.

You can either directly use the variables like so:

```go
package main

import "github.com/TwiN/go-color"

func main() {
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
