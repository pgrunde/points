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
If you save and run this program with `go run scrabble.go` the console will print '*Creating game with 0 players*'. Now wait just a minute, I never even typed a `0` this whole time, how did `%d` turn into that? Variables in Go have what are called [zero values](https://golang.org/ref/spec#The_zero_value), and the zero value of an integer is, predictably, `0`. When we say `var count int`, two things are actually going in our computer's little digital brain. First it is allocating an *address* for your variable, this acts like a library reference card detailing the place you'd find a book. In the computer's case, the address is a reference that points you to a value, not a book (if you are too young to remember what a library reference card is, go ask an old person). Because we did not set the value of our variable explicitly in our code, its address automatically got the zero value of an integer, 0. You can declare *and* set a variable with the fancy tool `:=` like this `count := 5`, which will set the variable `count` to have an address referencing the value of 5. Go attempts to figure out what type you meant, and numbers will default to `int`. Change your main function to use this [short variable declaration](https://golang.org/ref/spec#Short_variable_declarations):
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

The `Printf` code is performing what is called [string interpolation](https://en.wikipedia.org/wiki/String_interpolation), which basically means the the words (type string) are replacing the `%d` with the contents of the next argument's value, which in our case means the value of `count` is replacing `%d` in the string, and it is then getting printed to the console. You can see more formatting options in Go's [fmt package](https://golang.org/pkg/fmt/). You can see what passing multiple arguments to `Printf` looks like in [this playground](http://play.golang.org/p/iiZAiXi2hZ). 

# Making your own function

While there is plenty of fun to be had only using the standard library's functions, it is much more enjoyable to write your own. Right now we've 'hard coded' the number of players into our program to always be 5 because we set `count` to that value. Instead, let us create a function which [returns](https://golang.org/ref/spec#Return_statements) the number of players in our game after prompting the user for a number. 

Fundamentally [functions](https://golang.org/ref/spec#Function_types) take inputs, and produce outputs. Writing them begins with the keyword `func`, a name, a set of optional arguments (the inputs to our function), and a set of optional return values (the outputs of our function) followed by the body of the function. For now, we'll write our function to take no arguments and return 5, so our program will function the same:
```
package main

import "fmt"

func main() {
  count := PlayerCount()
  fmt.Printf("Creating game with %d players\n", count)
}

func PlayerCount() int {
  return 5
}
```
The `fmt.Printf` function took arguments which were declared inside the parenthesis, however `PlayerCount` takes no arguments. It does, however, return an integer, so we need to declare that return type after the parenthesis. The next part of our function is wrapped in curly bracket delimiters and is called the function [block](https://golang.org/ref/spec#Blocks). All the lines of code inside `{` and `}` run whenever the function is used. When the function code is running, the `return` keyword tells the function to exit and return whatever value comes after `return`, which in our case is the integer 5 (go ahead and try to return `"five"` and try to run to program- it won't even compile when the type you return doesn't match the return type you specified). By the way, Go can perform math with integers, so `return 2 + 3` and `return 7 - 1` would all return the same thing as `return 5`.

# Give me the Methods to implement my Interface

In order to get input from the user, we're probably going to need to *read* something they give us, hope it is a number, and then create our game with that many players. Thankfully the standard library has a useful package called [bufio](http://golang.org/pkg/bufio/), which stands for buffered input/output. 'Buffer' just means that your computer will put stuff into memory while it gets used. Since we're going to read input from the user, we're going to use the [bufio.NewReader](http://golang.org/pkg/bufio/#NewReader) function to create a reader for our user's input. Check out the input for that function:
```
func NewReader(rd io.Reader) *Reader
``` 
The input for this function `rd` must be of a type `io.Reader`. If you click on that input type in the library, it will take you to a description of [type Reader](http://golang.org/pkg/io/#Reader). This is an unfamiliar type called an [interface](https://golang.org/doc/effective_go.html#interfaces_and_types), and you can see it is written `type Reader interface` followed by another *block* of code (we know it is a block because it is wrapped in curly brackets). The block, however, is simply a list of methods. In this case, the list contains only a single method, `Read(p []byte) (n int, err error)`. Don't worry if you don't recognize the types here just yet, it is more important to know that it is a method.  *Oh shoot dang what is a method?* you might be asking, *it looks a whole lot like a function, what gives?* A [method](https://golang.org/doc/effective_go.html#methods) is a kind of function, except only *types* can perform, or *call*, that method *if they've been already given that method*. Giving a type a method looks an awful lot like how you write a function, only inbetween the keyword `func` and the function name, you'll see parenthesis which declare the type which receives this method. For example, in the [os package](https://golang.org/pkg/os/) there is a a type called [File](https://golang.org/pkg/os/#File), and this type has a method we're about to find really useful- [Read](https://golang.org/pkg/os/#File.Read). It looks like this:
```
func (f *File) Read(b []byte) (n int, err error)
```
You see `(f *File)` after the keyword `func` and before `Read`- this tells you that, in this case, the type `*File` has the method `Read` (ignore the asterisk for now and just think `File`, I'll explain why that doo-hickey is there in a bit). More importantly, our `type Reader interface` has just one method, and its name, input type, and output types are *identical* to the method that `*File` has- you don't even need to know what `[]byte] or `error` are to see that they're the same. Using the Reader interface type as the input to `io.NewReader` lets us know that any type can be used as an input so long as it *implements the Reader interface*, which is to say it has the method `Read` with the same input and output types. Let us rewrite our `PlayerCount()` function to create a new reader, and we'll pass it a special `os.File` type called [Stdin](https://golang.org/pkg/os/#pkg-variables), which is a File that represents input coming from the user's console (Terminal).
```
func PlayerCount() int {
  r := bufio.NewReader(os.Stdin)
}
```
