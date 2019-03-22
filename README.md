# go-color

A lightweight, simple and cross-platform package to colorize text in terminals.


## Usage


### Variables only

You can either directly use the variables like so:

```go
package main

import "github.com/TwinProduction/go-color"

func main() {
    println(color.Red + "Something something something" + color.Reset)
}
```

**NOTE**: If you're going to use this approach, don't forget to prepend your string with `color.Reset`, otherwise everything else in your terminal will be that color until the color is reset or overridden.


### Function and variable

You can use the `Ize(color, message)` function in conjunction with a variable:

```go
package main

import "github.com/TwinProduction/go-color"

func main() {
    println(color.Ize(color.Red, "Something something something"))
}
```

Unlike the method above, using this function will automatically prepend `color.Reset` at the 
end of your message.


### Why is it called Ize

Because I felt reading `color.Ize(...)` to be more visually pleasant than `color.Colorize(...)`


## Installation

```
go get github.com/TwinProduction/go-color
```
