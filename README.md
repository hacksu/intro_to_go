# intro_to_go
The aim of this lesson is to cover the basics of the Go programming language

Some important things to know about Go:
* It was designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson
* Syntactically similar to C
* It is not exactly an object oriented language
    * No classes, but there are structures (types)
    * You can define methods for types
    * Still similar to OOP
* [How to write Go code](https://golang.org/doc/code.html)
    * The `go` tool requires you to organize your code in a specific way
    * We will cover the basics in this lesson, but use this for further reading or as a reference

## Getting started

First, head to [the official Go website](https://golang.org/) and download the latest version of Go

I would recommend using VS Code with the Go extension for development. That is what we'll be using in this lesson.
* We will compile and run Go in the command line
* There is a Go IDE, GoLand, as well if you are interested (but we won't be really covering that)

## Your first Go program

Open up the command line, and navigate to your Go workspace:
* `C:\Users\YourName\go\src` on Windows
* `$HOME/go/src` on Unix based systems
* The location of your workspace can be changed. See "How to write Go code" above

Make a new directory (folder) called `shopping`

Navigate into `shopping` and create a file called `main.go`. We will write the following code

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Wow I did it")
}
```

Congrats, you've written your first go program!
