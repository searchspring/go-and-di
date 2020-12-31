# go-basics
Basics of Go

# Terminology

- Module is basically a repo containing Go code and a go.mod file
- Package is any directory within a module, and packages are defined by their file path

# gopath and workspaces

- The GOPATH env var defines the location of your Go workspace in your system
- Prior to 1.13 all development would be done directly in your Go workspace
    - Code would usually be written in a path made up of your github repo ($GOPATH/src/github.com/searchspring/foo)
    - Installed dependency source would also be placed here
    - Installed binaries go in $GOPATH/bin
    - $GOPATH/pkg is where compiled objects go before being built into executables that end up in $GOPATH/bin

# GOROOT

$GOROOT holds the standard libraries

# Third Party Packages

- No central repo service, nothing like NPM or Maven
- Libraries are installed directly from URLs, usually the Github repo. 
    - It usually looks for a VCS root of some kind at that location
    - Can also install other types of stuff, but not important 
- Prior to 1.13 there was no way to install specific versions, or even different versions. 
  Everything on the machine used the same $GOPATH/src and all pointed to whatever the latest installed version was.
  
# Go Modules

1. Must be VCS (fine, obviously that's the intuitive thing anyways)
1. Initialize with `go mod init <github repo>`, like `go mod init github.com/searchspring/go-basics`
  - Creates a `go.mod` file that just names the module and ties it to a go version.
  - Later this will store dependency versions 

# Resources

[Go Modules vs Workspace](https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16)
[Go V2 Modules](https://blog.golang.org/v2-go-modules)
