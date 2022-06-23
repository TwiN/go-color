# go-color

![test](https://github.com/TwiN/go-color/workflows/test/badge.svg?branch=master)

An extremely lightweight cross-platform package to colorize text in terminals.

This is not meant for maximal compatibility, nor is it meant to handle a plethora of scenarios.
It will simply wrap a message with the necessary characters, if the OS handles it.

There are many cases in which this would not work, such as the output being redirected to something other 
than a terminal (such as a file, i.e. `executable >> file.txt`)


## Install

`go get github.com/TwiN/go-color`


### Usage

You can use the `color.Colorize(color, message)` or `color.Ize(color, message)` function 
in conjunction with a variable like so:
```go
package main

import (
    "github.com/TwiN/go-color"
    "fmt"
)

func main() {
    fmt.Println(color.Ize(color.Red, "This is red"))
    // Or if you prefer the longer version
    fmt.Println(color.Colorize(color.Red, "This is red"))
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

import (
   "github.com/TwiN/go-color"
    "fmt"
func main() {
    fmt.Println(color.InBold("This is bold"))
    fmt.Println(color.InRed("This is red"))
    fmt.Println(color.InGreen("This is green"))
    fmt.Println(color.InYellow("This is yellow"))
    fmt.Println(color.InBlue("This is blue"))
    fmt.Println(color.InPurple("This is purple"))
    fmt.Println(color.InCyan("This is cyan"))
    fmt.Println(color.InGray("This is gray"))
    fmt.Println(color.InWhite("This is white"))
}
```


### Variables only

Unlike using the aforementioned approach, directly using the color variables will required you to manually
prepend `color.Reset` at the end of your string.

You can either directly use the variables like so:

```go
package main

import (
    "github.com/TwiN/go-color"
    "fmt"
)

func main() {
    fmt.Println(color.Bold + "This is bold" + color.Reset)
    fmt.Println(color.Red + "This is red" + color.Reset)
    fmt.Println(color.Green + "This is green" + color.Reset)
    fmt.Println(color.Yellow + "This is yellow" + color.Reset)
    fmt.Println(color.Blue + "This is blue" + color.Reset)
    fmt.Println(color.Purple + "This is purple" + color.Reset)
    fmt.Println(color.Cyan + "This is cyan" + color.Reset)
    fmt.Println(color.Gray + "This is gray" + color.Reset)
    fmt.Println(color.White + "This is white" + color.Reset)
}
```

**NOTE**: If you're going to use this approach, don't forget to prepend your string with `color.Reset`, 
otherwise everything else in your terminal will be that color until the color is reset or overridden.
