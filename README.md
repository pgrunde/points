# Learn Go fundamentals writing a little point tracker

`go run scrabble.go`

I made this to keep track of points (originally for Scrabble, but anything with points will work) while showcasing some Go fundamentals for a beginner. If you're just starting out with Go or programming in general, hopefully you'll find this useful. The final application can be viewed by forking the project, though the tutorial will build out each section as if you were writing it for the first time.

## Prerequisites

A working [installation of Go](https://golang.org/doc/install), and some [command line basics](http://cli.learncodethehardway.org/book/). You should at some point read [Effective Go](https://golang.org/doc/effective_go.html), including a read of the language spec, but it isn't necessary to complete this tutorial, it's just a good resource for your learning adventure.

# Setup

Make a folder in your [GOPATH](https://golang.org/doc/code.html#GOPATH)- if you're using Github, I recommend `github.com/$username$/scrabble` and create our main file, `touch scrabble.go`. Open that up in your favorite editor and let's begin!

# The main function

Every application which compiles and runs will have a `main()` function which starts off your program. Libraries, called packages in Go, do not have a main function, they just give you cool functions and tools to play with in your own app- we're going to only use packages that come from Go's [powerful standard library](https://golang.org/pkg/#stdlib) (and a package of our own design!).

Each Go file begins with a package declaration that says to your application "Hey App, all the contents of this here file belongs to this package." The beginning of our program, the file containing the `main()` function, is called `package main`. If you tested your Go installation, you should have seen something like this:
```
package main

import "fmt"

func main() {
  fmt.Printf("hello, world\n")
}
```
A similar default setup can be found at [The Go Playground](http://play.golang.org/), a nifty tool which allows you to quickly play with and share Go code snippets. 

After the package declaration we see the word `import`, a [keyword](https://golang.org/ref/spec#Keywords) which tells the app which imported packages will be used in this file.`"fmt"` means we're importing the [fmt package](https://golang.org/pkg/fmt/) from the standard library.
