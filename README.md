# Learn Go fundamentals writing a little point tracker

`go run scrabble.go`

I made this to keep track of points (originally for Scrabble, but anything with points will work) while showcasing some Go fundamentals for a beginner. If you're just starting out with Go or programming in general, hopefully you'll find this useful. The final application can be viewed by forking the project, though the tutorial will build out each section as if you were writing it for the first time.

## Prerequisites

A working [installation of Go](https://golang.org/doc/install), and some [command line basics](http://cli.learncodethehardway.org/book/). You should at some point read [Effective Go](https://golang.org/doc/effective_go.html), including a read of the language spec, but it isn't necessary to complete this tutorial, it's just a good resource for your learning adventure.

# Setup


Make a folder in your [GOPATH](https://golang.org/doc/code.html#GOPATH)- if you're using Github, I recommend `github.com/$username$/scrabble`. Next create our main file inside, `touch scrabble.go`. Open that up in your favorite editor and let's begin!

# The main function

Every application which compiles and runs will have a `main()` function whose job is to start off your program. Libraries, called packages in Go, do not have a main function, they just give you cool functions and tools to play with in your own app- we're going to only use packages that come from Go's [powerful standard library](https://golang.org/pkg/#stdlib) (and a package of our own design!).

Each Go file begins with a package declaration that says to your application "*Hey App, all the contents of this here file belong to this package.*" The beginning of our program, the file containing the `main()` function, lives in `package main`. If you [tested your Go installation](https://golang.org/doc/install#testing), you should have seen something like this:
```
package main

import "fmt"

func main() {
  fmt.Printf("hello, world\n")
}
```
(A similar default setup can be found at [The Go Playground](http://play.golang.org/), a nifty tool which allows you to quickly play with and share Go code snippets.)

After the package declaration we see the word `import`, a [keyword](https://golang.org/ref/spec#Keywords) which tells the app which imported packages will be used in this file. The `"fmt"` means we're importing the [fmt package](https://golang.org/pkg/fmt/) from the standard library because we need to print something to the command line. It should be noted that the main [function](http://www.cs.utah.edu/~germain/PPS/Topics/functions.html), like all Go functions, starts with the keyword `func`. We'll go into further detail regarding functions later on.

# Variables and Functions

Before we can ever keep track of points, we need to know how many players are actually in the game. Let us create a [variable](https://golang.org/ref/spec#Variables) and store an integer in it, then print out to the user that the game is starting with that many players.
```
package main

import "fmt"

func main() {
  var count int
  fmt.Printf("Creating game with %d players\n", count)
}
```
If you save and run this program with `go run scrabble.go` the console will print '*Creating game with 0 players*'. Now wait just a minute, I never even typed a `0` this whole time, how did `%d` turn into that? The trick is in [the Printf function](https://golang.org/pkg/fmt/#Printf), which you'll notice is utilized by first typing out what package it comes from (fmt), followed by a period and then the function name. Our `fmt.Printf` takes two [arguments](https://en.wikipedia.org/wiki/Argument_of_a_function)- arguments are the names describing the inputs for a function. The are always inside the parenthesis following the function name. First up in our arguments is a [string of characters](https://golang.org/ref/spec#String_types) which you see inside the quotes, followed by that `count` variable we declared earler (note that the two arguments are comma separated, as are all arguments inside a function, `fmt.Printf(format string, a ...interface{})`).
