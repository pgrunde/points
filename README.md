# Learn Go fundamentals writing a little point tracker

`go run scrabble.go`

I made this to keep track of points (originally for Scrabble, but anything with points will work) while showcasing some Go fundamentals for a beginner. If you're just starting out with Go or programming in general, hopefully you'll find this useful. The final application can be viewed by forking the project, though the tutorial will build out each section as if you were writing it for the first time.

## Prerequisites

A working [installation of Go](https://golang.org/doc/install), and some [command line basics](http://cli.learncodethehardway.org/book/). You should at some point read [Effective Go](https://golang.org/doc/effective_go.html), including a read of the language spec, but it isn't necessary to complete this tutorial, it's just a good resource for your learning adventure.

# Setup


Make a folder in your [GOPATH](https://golang.org/doc/code.html#GOPATH)- if you're using Github, I recommend `github.com/$username$/scrabble`. Next create our main file inside, `touch scrabble.go`. Open that up in your favorite editor and let's begin!

# The main function

Every application which compiles and runs will have a `main()` function whose job is to start off your program. Libraries, called packages in Go, do not have a main function, they just give you cool functions and tools to play with in your own app- we're going to only use packages that come from Go's [powerful standard library](https://golang.org/pkg/#stdlib), which is like a huge collectionof tools that comes with Go itself.

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

# Variables

Before we can ever keep track of points, we need to know how many players are actually in the game. Let us create a [variable](https://golang.org/ref/spec#Variables) and store an integer in it, then print out to the user that the game is starting with that many players. We'll declare it with `var` followed by the name we give it, followed by the [type](https://golang.org/ref/spec#Types). The type of a variable is very important in Go as it is a [strongly typed](https://en.wikipedia.org/wiki/Strong_and_weak_typing) language, which means Go really needs to know what the heck type of thing you're creating or using. Maybe you can store memories as smells, sounds, tastes, and emotion, but a computer only remembers things as ones and zeros. Ones and zeros are still powerful though, because we can put them in patterns that a computer can recognize as more complex stuff. A computer doesn't see your Google searches as words, but ones and zeros assembled to represent those words- Go, by the by, uses what is called [UTF-8 encoding](https://en.wikipedia.org/wiki/UTF-8) to represent its words- they have a type `string`, which comes from the idea that words are characters 'strung' together. Since we need to represent a number for our player count, we can use a type of integer `int`, one of Go's [numeric types](https://golang.org/ref/spec#Numeric_types):
```
package main

import "fmt"

func main() {
  var count int
  fmt.Printf("Creating game with %d players\n", count)
}
```
If you save and run this program with `go run scrabble.go` the console will print '*Creating game with 0 players*'. Now wait just a minute, I never even typed a `0` this whole time, how did `%d` turn into that? Variables in Go have what are called [zero values](https://golang.org/ref/spec#The_zero_value), and the zero value of an integer is, predictably, `0`. When we say `var count int`, two things are actually going in our computer's little digital brain. First it is allocating an *address* for your variable, this acts like a library reference card detailing the place you'd find a book. In the computer's case, the address is a reference that points you to a value, not a book (if you are too young to remember what a library reference card is, go ask an old person). Because we did not set the value of our variable explicitly in our code, its address automatically got the zero value of an integer, 0. You can declare *and* set a variable with a fancy tool (called an operator) `:=` like this `count := 5`, which will set the variable `count` to have an address referencing the value of 5. Go attempts to figure out what type you meant, and numbers will default to `int`. Change your main function to use this [short variable declaration](https://golang.org/ref/spec#Short_variable_declarations):
```
func main() {
  count := 5
  fmt.Printf("Creating game with %d players\n", count)
}
```

# Functions

You might have already guessed what the [the Printf function](https://golang.org/pkg/fmt/#Printf) is doing in our program. It is replacing the `%d` with the value of our `count` variable and printing it to the console. This function comes from the fmt package in the standard library, so we call it by first typing out `fmt` followed by a period and then the function name. Our `fmt.Printf` takes two [arguments](https://en.wikipedia.org/wiki/Argument_of_a_function)- arguments are the names describing the inputs for a function. They are always inside the parenthesis following the function name. First up in our arguments is a [string of characters](https://golang.org/ref/spec#String_types) which you see inside the quotes, it is a type `string`,  followed by that `count` variable we declared earler. Note that the two arguments are comma separated, as are all arguments passed to a function. If you [look at the function](https://golang.org/pkg/fmt/#Printf) in the standard library it looks like this: 
```
fmt.Printf(format string, a ...interface{})
```
Here `format` has the type string, which we passed `"Creating game with %d players"`. But why is the argument `a` a type with the elipsis interface thingy? What is going on there? Didn't we pass it `count`, which has an `int` type? Well, the last argument in any of Go's functions can have an elipsis prefixed to it to mean that it is *variadic*, and may be called with zero, one, or many arguments of that type. In our case, we're only passing one argument. The type we've been given `interface{}` is what is called an *empty interface*- we'll cover interfaces more later, but just know that an empty interface is like a catch-all for types that can accept absolutely any type you give it. Empty interfaces are useful, but your programs can suffer if you overuse them as the code gets less readable and more prone to *runtime errors*; these are errors that happen while the program is running. 

The `Printf` code is performing what is called [string interpolation](https://en.wikipedia.org/wiki/String_interpolation), which basically means the the words (type string) are replacing the `%d` with the contents of the next argument's value, which in our case means the value of `count` is replacing `%d` in the string, and it is then getting printed to the console. You can see more formatting options in Go's [fmt package](https://golang.org/pkg/fmt/). You can see what passing multiple arguments to `Printf` looks like in[this playground](http://play.golang.org/p/iiZAiXi2hZ). 
