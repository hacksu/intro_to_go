# intro_to_go
![Go](https://golang.org/lib/godoc/images/go-logo-blue.svg)

The aim of this lesson is to cover the basics of the Go programming language

Some important things to know about Go:
* It was designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson
* Syntactically similar to C
* It is not exactly an object oriented language
    * No classes, but there are structures
    * You can define methods for these structures
    * Still similar to OOP
* [How to write Go code](https://golang.org/doc/code.html)
    * The `go` tool requires you to organize your code in a specific way
    * We will cover the basics in this lesson, but use this for further reading or as a reference

## Getting started

First, head to [the official Go website](https://golang.org/) and download the latest version of Go

I would recommend using VS Code with the Go extension for development. That is what we'll be using in this lesson.
* We will compile and run Go in the command line
* There is a Go IDE, GoLand by JetBrains, as well if you are interested (but we won't be really covering that)

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

To run this, type `go run main.go` into the command line

Congrats, you've written your first Go program!

## Getting some input

An important part of this application is reading in user input, so we'll learn about that next

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("What is your name?")

	var name string
	fmt.Scanln(&name)

	fmt.Println("Hello,", name)
}
```

Again, type `go run main.go` into the command line to run this

## Adding some structures

Now that basic I/O has been covered, let's move on to structures

Here's a basic structure that we're going to be using:

```go
type product struct {
    name string
    cost int
}
```

Now we can create a new product:

```go
func main() {
	p := product{
		name: "Brick",
		cost: 10,
	}

	fmt.Println(p.name, "costs $", p.cost)
}
```

Using structs, we can begin to build a more complex application

## Creating the store

We're going to start building our store application

```go
func main() {
	brick := product{name: "brick", cost: 10}
	corn := product{name: "corn", cost: 6}
	gamecube := product{name: "gamecube", cost: 100}

	inventory := []product{brick, corn, gamecube}

	fmt.Println("Welcome to the store!")
	fmt.Println("Options:")
	fmt.Println("(i)nventory")
	fmt.Println("(a)dd item")
	fmt.Println("(q)uit")

	var choice string
	fmt.Scanln(&choice)

	for choice != "q" {
		switch choice {
		case "i":
			for _, p := range inventory {
				fmt.Println("name:", p.name)
				fmt.Println("cost:", p.cost)
				fmt.Println()
			}
			break
		case "a":
			var newName string
			fmt.Print("Enter a product name: ")
			fmt.Scanln(&newName)

			var newCost int
			fmt.Print("Enter a cost: ")
			fmt.Scanf("%d", &newCost)

			inventory = append(inventory, product{name: newName, cost: newCost})
			break
		}

		fmt.Println("Welcome to the store!")
		fmt.Println("Options:")
		fmt.Println("(i)nventory")
		fmt.Println("(a)dd item")
		fmt.Println("(q)uit")
		fmt.Scanln(&choice)
	}
}
```

The completed code can also be found in the `main.go` file in this repository

## Further reading

If you are interested in learning more about Go, I would recommend checking out these resources:
* [A Tour of Go](https://tour.golang.org)
* [Go by Example](https://gobyexample.com/)

Another great resource in the official docs for when you're more familiar with the language:
* [Effective Go](https://golang.org/doc/effective_go.html)
