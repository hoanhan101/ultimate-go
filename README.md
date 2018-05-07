# Ultimate Go

This repo contains my notes on learning Go and computer systems. Different people have different
learning style. For me, I learn best by doing and walking through examples. Hence, I am trying to
take notes carefully and comment directly on the source code, rather than writing up Markdown
files. That way, I can understand every single line of code as I am reading and also be mindful of
the theories behind the scene.

**Resources:**
- [Ardan Labs's Ultimate Go course
  ](https://github.com/ardanlabs/gotraining/blob/master/topics/courses/go/README.md)
- [Computer Systems: A Programmer's
  Perspective](https://www.amazon.com/Computer-Systems-Programmers-Perspective-3rd/dp/013409266X)

## Project Status

It is an on-going project. 

Below are the outline of the all topics. Normally, a topic is covered when there is a link, 
or several links next to it.

## Table of Contents 

- **Design Philosophy**: [DESIGN.md](DESIGN.md)
- **Language Mechanics**
  - **Syntax**
    - Variable: [variable.go](variable.go)
    - Struct: [struct.go](struct.go)
    - Pointer: [pointer.go](pointer.go)
    - Constant: [constant.go](constant.go)
    - Function: [function.go](function.go)
  - **Data Structures**
    - Array: [array.go](array.go)
    - Slice: [slice.go](slice.go)
    - Map: [map.go](map.go)
  - **Decoupling**
    - Method: [method_1.go](method_1.go) | [method_2.go](method_2.go) | [method_3.go](method_3.go)
    - Interface
    - Embedding
    - Exporting
- **Software Design**
  - Composition
  - Error Handling
  - Packaging
- **Concurrency**
  - **Mechanics**
    - Goroutine
    - Data race
    - Channel
  - **Patterns**
    - Context
    - Pattern
- **Testing and Profiling**
  - Testing
  - Benchmarking
  - Fuzzing
  - Profiling
- **Packages**
  - Context
  - Encoding
  - IO
  - Logging
  - Reflection
