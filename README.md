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

- **Design Philosophy**: [Guideline](go/README.md)
- **Language Mechanics**
  - **Syntax**
    - Variable: [variable.go](go/language/variable.go)
    - Struct: [struct.go](go/language/struct.go)
    - Pointer: [pointer.go](go/language/pointer.go)
    - Constant: [constant.go](go/language/constant.go)
    - Function: [function.go](go/language/function.go)
  - **Data Structures**
    - Array: [array.go](go/language/array.go)
    - Slice: [slice.go](go/language/slice.go)
    - Map: [map.go](go/language/map.go)
  - **Decoupling**
    - Method: [method_1.go](go/language/method_1.go) | [method_2.go](go/language/method_2.go) | 
      [method_3.go](go/language/method_3.go)
    - Interface: [interface_1.go](go/language/interface_1.go) | [interface_2.go](go/language/interface_2.go)
    - Embedding: [embedding_1.go](go/language/embedding_1.go) | [embedding_2.go](go/language/embedding_2.go) |
      [embedding_3.go](go/language/embedding_3.go) | [embedding_4.go](go/language/embedding_4.go)
    - Exporting: [Guideline](go/language/exporting/README.md) | [exporting_1](go/language/exporting/exporting_1) | 
      [exporting_2](go/language/exporting/exporting_2) | [exporting_3](go/language/exporting/exporting_3) | 
      [exporting_4](go/language/exporting/exporting_4)
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
