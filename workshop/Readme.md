# Workshop

## Download Go

(Current Version: 1.22.2)[https://go.dev/doc/install]

## Install Raylib

Raylib:
https://github.com/gen2brain/raylib-go

Installation:
go get -v -u github.com/gen2brain/raylib-go/raylib

Raylib documentation - Examples:
https://www.raylib.com/examples.html

## Start a new project

go mod init github/username/workshop

- Creates a file called go.mod
  The gogame/workshop path is called a Go Module.
  Very similar to how namespaces work.

Best practice is to keep it to lowercases and many will reflect this to their github profile and repo.
If you host the code as a github package it is recommended to set up like:
github.com/username/goworkshop
It makes the default packages easier to find and fetch necessary dependencies. But this is not required when hosting the code in projects.

The go.mod file is the setup of go-config/manifest.
Here you will find all dependencies and version control.

Similar to package.json in Node.js or Gemfile in Ruby.

## main.go

This is the startup, initialize and main package.

in the official Go tutorial, the main.go file only needs the following code to handle the infamous hello world:

```
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

In our example we handle text input, string formatting and the time package.

### Packages

Find other Go packages here:
pkg.go.dev
