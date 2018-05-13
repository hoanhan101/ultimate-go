# Ultimate Go

This repo contains my notes on learning Go and computer systems. Different people have different
learning style. For me, I learn best by doing and walking through examples. Hence, I am trying to
take notes carefully and comment directly on the source code, rather than writing up Markdown
files. That way, I can understand every single line of code as I am reading and also be mindful of
the theories behind the scene.

**Resources:**
- [Ultimate Go
  Programming](https://www.safaribooksonline.com/library/view/ultimate-go-programming/9780134757476/)
- [ardanlabs/gotraining/topics/courses/go
  ](https://github.com/ardanlabs/gotraining/blob/master/topics/courses/go/README.md)
- [Computer Systems: A Programmer's
  Perspective](https://www.amazon.com/Computer-Systems-Programmers-Perspective-3rd/dp/013409266X)

## Project Status

It is an on-going project. 

Below are the outline of the all topics. Normally, a topic is covered when there is a link, 
or several links next to it.

## Table of Contents 

- **Design Philosophy**:
  [Guideline](https://github.com/ardanlabs/gotraining/blob/master/topics/go/README.md)
- **Language Mechanics**
  - **Syntax**
    - Variable: [Part 1-end](go/language/variable.go)
    - Struct: [Part 1-end](go/language/struct.go)
    - Pointer: [Part 1-end](go/language/pointer.go)
    - Constant: [Part 1-end](go/language/constant.go)
    - Function: [Part 1-end](go/language/function.go)
  - **Data Structures**
    - Array: [Part 1-end](go/language/array.go)
    - Slice: [Part 1-end](go/language/slice.go)
    - Map: [Part 1-end](go/language/map.go)
  - **Decoupling**
    - Method: [Part 1](go/language/method_1.go) | [Part 2](go/language/method_2.go) | 
      [Part 3-end](go/language/method_3.go)
    - Interface: [Part 1](go/language/interface_1.go) | [Part 2-end](go/language/interface_2.go)
    - Embedding: [Part 1](go/language/embedding_1.go) | [Part 2](go/language/embedding_2.go) |
      [Part 3](go/language/embedding_3.go) | [Part 4-end](go/language/embedding_4.go)
    - Exporting: [Guideline](go/language/exporting/README.md) | [Part 1](go/language/exporting/exporting_1) | 
      [Part 2](go/language/exporting/exporting_2) | [Part 3](go/language/exporting/exporting_3) | 
      [Part 4-end](go/language/exporting/exporting_4)
- **Software Design**
  - Composition:
    [Guideline](https://github.com/ardanlabs/gotraining/tree/master/topics/go#interface-and-composition-design)
    - Grouping types: [Part 1](go/design/grouping_types_1.go) | [Part 2-end](go/design/grouping_types_2.go)
    - Decoupling: [Part 1](go/design/decoupling_1.go) | [Part 2](go/design/decoupling_2.go) |
    [Part 3](go/design/decoupling_3.go) | [Part 4-end](go/design/decoupling_4.go)
    - Conversion: [Part 1](go/design/conversion_1.go) | [Part 2-end](go/design/conversion_2.go)
    - Interface Pollution: [Part 1](go/design/pollution_1.go) | [Part 2-end](go/design/pollution_2.go)
    - Mocking: [Part 1](go/design/mocking_1.go) | [Part 2-end](go/design/mocking_2.go)
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
