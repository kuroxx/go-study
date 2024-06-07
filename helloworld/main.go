package main // package == project == workspace

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

// To run code: $ go run main.go

// Go CLI
// 	go build : Compiles a bunch of go source code files
// 	go run : Compiles and executes one or two files
// 	go fmt : Formats all the code in each file in the current dir
// 	go install : Compiles & "installs" a package
// 	go get : Downloads raw source code of someone else's package
// 	go test : Runs any tests associated with the current project

// Types of packages
// 	Executable - Generates a file that we can run
// 	Reusable - Code used as 'helpers'. Good place to put reusable logic. (Think of code dependencies or packages)

// What does 'package main' mean? 
// 	Defines a package than can be compiled & then *executed*.
// 	Must have a func called 'main'
// 	i.e. package other -  Defines package that can be used as a dependency

// What does "import fmt" mean? 
// 	fmt - short for format, is a standard Go library
// 	import is used to gain access to another package
// 	https://golang.org/pkg

// How is the main.go file organized? 
// 	1. Package declaration
// 	2. import other packages that we need
// 	3. Declare functions, tell Go to do things
