# go-color

An extremely lightweight cross-platform package to colorize text in terminals.

This is not meant for maximal compatibility, nor is it meant to handle a plethora of scenarios.
It will simply wrap a message with the necessary characters, if the OS handles it.

There are many cases in which this would not work, such as the output being redirected to something other 
than a terminal (such as a file, i.e. `executable >> file.txt`)


## Usage


### Variables only

You can either directly use the variables like so:

```go
package main

import "github.com/TwinProduction/go-color"

func main() {
    println(color.Red + "Something something something" + color.Reset)
    // or
    println(color.Ize(color.Red, "Something something something"))
}
```

**NOTE**: If you're going to use this approach, don't forget to prepend your string with `color.Reset`, otherwise everything else in your terminal will be that color until the color is reset or overridden.


### Function and variable

You can use the `Colorize(color, message)`/`Ize(color, message)` function in conjunction with a variable:

```go
package main

import "github.com/TwinProduction/go-color"

func main() {
    println(color.Ize(color.Red, "Something something something"))
    // Or if you prefer the longer version
    println(color.Colorize(color.Red, "Something something something"))
}
```

Unlike the method above, using this function will automatically prepend `color.Reset` at the 
end of your message.

Because I felt reading `color.Ize()` to be more visually pleasant than `color.Colorize()`, I included `Ize()` as an alias for `Colorize()`.

I'm not usually a big fan of having two methods doing the same thing, but since
this package doesn't have much room for growth (its only purpose is to colorize
terminal outputs, after all, and there isn't hundreds of ways to go about it),
I believe it's acceptable to have both.


## Installation

```
go get github.com/TwinProduction/go-color
```
