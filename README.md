# sidGo
GoLang Repository

This GoLang repository consists of basic program at the moment. 

Below are notes captured by me when learning GoLang. Will update the notes as and when I progress. 

GoLang

GoLang is a robust system-level language used for programming across large-scale network servers and big distributed systems.  

Go in a Nutshell 
• Imperative language 
• Statically typed 
• Syntax similar to Java/C/C++, but less parentheses and no semicolons 
• Compiles to native code (no JVM) **
• No classes, but structs with methods **
• Interfaces 
• No implementation inheritance. There's type embedding, though. **
• Functions are first class citizens 
• Functions can return multiple values 
• Has closures 
• Pointers, but not pointer arithmetic 
• Built-in concurrency primitives: Goroutines and Channels 

GoLang Benefits:

⁃ Multithreading & Concurrency
⁃ Handles garbage collections
⁃ Easy to learn (took me 3 hours to CRUD)
⁃ Inbuilt testing framework **
⁃ Dynamic AND non dynamic libraries **
⁃ No Generics **
⁃ Single Executable
⁃ Easy to maintain dependencies
⁃ Easy to read and document
⁃ No more classes **

Setup:
Install Golang https://golang.org/doc/install
Note: By default, the go command downloads and authenticates modules using the Go module mirror and Go checksum database run by Google

In Terminal:
go version
go version go1.15.6 darwin/amd64

Go Installation: /usr/local/go

#Print GOPATH:
sidz@Siddeshs-MacBook-Air ~ % go env GOPATH
/Users/sidz/go

The GOPATH environment variable
The GOPATH environment variable specifies the location of your workspace. It defaults to a directory named go inside your home directory, so $HOME/go on Unix, $home/go on Plan 9, and %USERPROFILE%\go (usually C:\Users\YourName\go) on Windows.
If you would like to work in a different location, you will need to set GOPATH to the path to that directory. (Another common setup is to set GOPATH=$HOME.) Note that GOPATH must not be the same path as your Go installation.
The command go env GOPATH prints the effective current GOPATH; it prints the default location if the environment variable is unset.
For convenience, add the workspace's bin subdirectory to your PATH:



1. For local Mac used below, just to be sure:
go env -w GOPATH=$HOME/go

2. $ export PATH=$PATH:$(go env GOPATH)/bin
or
export PATH=$PATH:$GOPATH/bin

3. Below command didn’t work mostly due to Catalina restriction. Hence created individual directories. 
mkdir -p $GOPATH/src/github.com/sidGo

/Users/sidz/go/src/github.com/sidGo


4. Create package directory
$ mkdir $GOPATH/src/github.com/user/hello

5. Create hello.go 
package main

import "fmt"

func main() {
fmt.Println("Hello, world.")
}

6. Install
go install

7. Execute the program
$GOPATH/bin/hello

hello 

Difference between go run, go build and go install
https://levelup.gitconnected.com/go-run-vs-go-build-vs-go-install-c7c0fd135cf9

How to write Go Code:
https://golang.org/doc/code.html

Code Organization:
Go programs are organized into packages. A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package. - Sid: Like java package


A module is a collection of related Go packages that are released together. A Go repository typically contains only one module, located at the root of the repository. A file named go.mod there declares the module path: the import path prefix for all packages within the module. The module contains the packages in the directory containing its go.mod file as well as subdirectories of that directory, up to the next subdirectory containing another go.mod file (if any). - Like java project

A repository contains one or more modules.

Each module's path not only serves as an import path prefix for its packages, but also indicates where the go command should look to download it. For example, in order to download the module golang.org/x/tools, the go command would consult the repository indicated by https://golang.org/x/tools (described more here).

An import path is a string used to import a package. A package's import path is its module path joined with its subdirectory within the module. For example, the module github.com/google/go-cmp contains a package in the directory cmp/. That package's import path is github.com/google/go-cmp/cmp. Packages in the standard library do not have a module path prefix.


GO Structs

Structs are a bit like a class, but don’t contain methods
Variables are created with types (eg int, Boolean, string)        
Can contain another struct

https://www.golangprograms.com/go-language/struct.html


Short Variable Declaration in Golang
The := short variable assignment operator indicates that short variable declaration is being used. There is no need to use the var keyword or declare the variable type.

Naming conventions for GoLang Variables:
• If the name of a variable begins with a lower-case letter, it can only be accessed within the current package this is considered as unexported variables.
• If the name of a variable begins with a capital letter, it can be accessed from packages outside the current package one this is considered as exported variables.
• Variable names are case-sensitive (car, Car and CAR are three different variables).
• If you declare a variable without assigning it a value, Golang will automatically bind a default value (or a zero-value) to the variable.
•	We use the operator : = for declaring and initializing variables with short variable declaration. When you declare variables with this method, you can't specify the type because the type is determined by the initializer expression.

Conversions in Go using strconv and reflect packages: 
https://www.golangprograms.com/go-language/integer-float-string-boolean.html

GoLang Operators

https://www.golangprograms.com/go-language/operators.html

•	Arithmetic Operators
•	Assignment Operators
•	Comparison Operators
•	Logical Operators
•	Bitwise Operators

If else : pretty much same like java just no brackets to define if and else if statements. Can initialize variable within if statement.
https://www.golangprograms.com/golang-if-else-statements.html
Switch Case: 
Pretty similar to Java. Can have multiple case line statements. fallthrough key word to continue flow in the following case, conditional case statement, var initialization in switch statement
https://www.golangprograms.com/golang-switch-case-statements.html

For Loop:
A for loop is used for iterating over a sequence (that is either a slice, an array, a map, or a string.
Golang has no while loop because the for loop serves the same purpose when used with a single condition.
Can be used in different ways as below:
package main
 
import "fmt"
 
func main() {
 
	k := 1
	for ; k <= 10; k++ {
		fmt.Println(k)
	}
 
	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}
 
	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}
}

Uses range key word to iterate over an expression that evaluates to an array, slice, map, string, or channel.
Infinite loop with just using for and to break the loop condition is needed inside the if block. Works as a do while. 

Function
Simple function: No parameters, no return type
Function with parameters
Function with return type
Named return values * can name the return value and define a variable in function to return and use only return without specifying variable name
Return multiple value * 
Note: If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package.


Kafka – GoLang Integration (Clients for Kafka) https://kafka.apache.org/ Navigate to DOCS > CLIENTS which takes you to below link
https://cwiki.apache.org/confluence/display/KAFKA/Clients#Clients-Go(AKAgolang)

Multiple clients are available for Kafka integration with GoLang. Then one we have chosen is confluent-kafka-go provided by Confluent.
confluent-kafka-go: Confluent's Kafka client for Golang wraps the librdkafka C library, providing full Kafka protocol support with great performance and reliability.
The Golang bindings provides a high-level Producer and Consumer with support for the balanced consumer groups of Apache Kafka 0.9 and above.
Kafka Version: 0.8.x, 0.9.x,  0.10.x, 0.11.x
Maintainer: Confluent
License: Apache v2.0
https://github.com/confluentinc/confluent-kafka-go
Docs: http://docs.confluent.io/current/clients/index.html


GoLang system packages:
fmt - Println
strconv – string conversion (Atoi, ParseInt, ParseFloat, ParseBool, FormatInt, FormatFloat)
reflect - TypeOf
time – Now eg: today := time.Now() 
var day int = today.Day()

