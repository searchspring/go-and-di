# go-and-di

Basics of Go and Dependency Injection

# Installation

## Mac

- Directly from golang.org: https://golang.org/doc/install
- Using brew: `brew install golang`

## Windows

- Directly from golang.org: https://golang.org/doc/install

## Ubuntu

- Probably want to use package manager: https://github.com/golang/go/wiki/Ubuntu

## Arch

- `yay -S go go-tools`

# Learning the Syntax

Strongly recommend you go through as much as possible of the "Tour of Go": https://tour.golang.org

It's pretty simple and fast, and provides a good starting point for working with the language.

## Notable differences from other languages

- [Multiple return values](https://tour.golang.org/basics/6)
- [Named return values](https://tour.golang.org/basics/7)
- [Short assignment](https://tour.golang.org/basics/10) 
- [Zero Values](https://tour.golang.org/basics/12)
- [All loops are for-loops](https://tour.golang.org/flowcontrol/1)
- [Switches are a bit different](https://tour.golang.org/flowcontrol/9)
- [Defer is like finally for functions](https://tour.golang.org/flowcontrol/12)
- [Pointers](https://tour.golang.org/moretypes/1)
- [Arrays vs Slices](https://tour.golang.org/moretypes/7)
- [Methods and receivers (class-like stuff)](https://tour.golang.org/methods/1)
- [Implicit Interfaces](https://tour.golang.org/methods/9)
- [Empty Interface (untyped variables)](https://tour.golang.org/methods/14)
- [Errors](https://tour.golang.org/methods/19)
- [Defer, Panic, Recover](https://blog.golang.org/defer-panic-and-recover)
- [Goroutines and Concurrency](https://tour.golang.org/concurrency/1)
- [Go Modules vs Workspace](https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16)
- [Go Module Versioning](https://blog.golang.org/v2-go-modules)

# Dependency Injection

[Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection) (DI) is a form of Inversion of Control
intended to improve code modularity and testability. Instead a class directly finding and accessing it's dependencies,
those dependencies are passed in by the caller.

The form of DI that we use is called constructor injection, where we pass in any complex or large dependencies at
construction time for use at runtime. This has the benefit of ensuring that all components are properly created at the
moment the app is started, rather than finding that out later when a new piece of code is first executed.

## Why?

DI allows us to easily mock every major dependency during testing, making unit testing much easier.

For example, without DI, in order to test a handler (aka controller), you would have the following call chain:

```
HTTP request -> controller -> service -> DAL -> client -> DB -> client -> service -> controller -> HTTP response
```

The larger the application becomes, the more elements are involved in the call chain and the more complicated the test
becomes. Additionally, some individual branches become actually impossible to test, and tests become unmaintainable.

Testing at a lower level is simpler, but still has parts of the same deep call tree.

```
service -> DAL -> client -> DB -> client -> service
```

But with DI, all complex dependencies are mocked, so regardless of which level you are testing, the call tree is only one level deep.

```
service -> Mock DAL -> service
```

Or:

```
HTTP request -> controller -> Mock service -> controller -> HTTP response
```

## Consistent Patterns

Another strong benefit of using DI consistently is that pretty much every component of the application can be templated in a similar way. 

Define your dependencies, config (if needed), and interface. Then provide a constructor that satisfies that interface. Repeat for every piece of logic at every level.

Within a short amount of time, all the code ends up looking very similar, and all the tests follow the same pattern. Code reviews are much easier because deviations from the pattern are immediately obvious, and you can focus on reviewing the business logic. 

## How it's done

```go
package campaigns

// Struct containing all complex external dependencies
type Deps struct {
    SSCore sscore.SSCore
}

// Configurable parameters needed at construction
type Config struct {
	Port int
}

// Definition of the interface this package will provide externally
type Campaigns interface {
    Get(id string) (*Campaign, error)
}

// Private struct that provides the based for the implementation of that interface
// Internal properties are retained as deps, config, and then any other private properties that may be constructed
type impl struct {
    deps *Deps
    config *Config
}

// Constructor that returns an implementation matching the interface
func New(deps *Deps, config *Config) Campaigns {
    return &impl{deps: deps, config: config}
}

// Method definition attached to implementation as described in the interface
func (impl *impl) Get(id string) (*Campaign, error) {
	return &Campaign{}, nil
}
```

# Testing

- Tests are in files next to the src, but suffixed with `_test.go`
- Optionally, you can put your tests in an `_test` package to force the use of only public functions/properties

## Mocking

- Generally found the simplest method is to use `gomock` for automatically generating
  mocks: https://github.com/golang/mock
- Alternatively, you can also create them manually with `testify/mock` though this is more work for the same
  thing: https://github.com/stretchr/testify
