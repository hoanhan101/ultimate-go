# Ultimate Go

[![Go Report Card](https://goreportcard.com/badge/github.com/hoanhan101/ultimate-go)
](https://goreportcard.com/report/github.com/hoanhan101/ultimate-go)

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

**Tasks**

- [ ] Phase 1: Finish Ultimate Go Programming's video lectures
- [ ] Phase 2: Fill in all the missing details using Ardan Labs's links and examples
- [ ] Phase 3: Study Computer Systems book to reinforce the theory. Build more programs if needed.

## Table of Contents 

- **Design Philosophy**:
  [Guideline](https://github.com/ardanlabs/gotraining/blob/master/topics/go/README.md)
- **Language Mechanics**
  - **Syntax**
    - Variable: [Built-in types | Zero value concept | Declaration and initialization | 
      Conversion vs Casting](go/language/variable.go)
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
  - Error Handling: [Part 1](go/design/error_1.go) | [Part 2](go/design/error_2.go) |
    [Part 3](go/design/error_3.go) | [Part 4](go/design/error_4.go) | [Part 5](go/design/error_5.go) |
    [Part 6-end](go/design/error_6.go)
  - Packaging:
    [Guideline](https://github.com/ardanlabs/gotraining/blob/master/topics/go/design/packaging/README.md)
- **Concurrency**
  - **Mechanics**
    - Goroutine: [Part 1](go/concurrency/goroutine_1.go) | [Part 2](go/concurrency/goroutine_2.go) |
      [Part 3](go/concurrency/goroutine_3.go) | [Part 4-end](go/concurrency/goroutine_4.go)
    - Data race: [Part 1](go/concurrency/data_race_1.go) | [Part 2](go/concurrency/data_race_2.go) | 
    [Part 3](go/concurrency/data_race_3.go) | [Part 4-end](go/concurrency/data_race_4.go)
    - Channel: [Guideline](https://github.com/ardanlabs/gotraining/tree/master/topics/go#concurrent-software-design) |
    [Part 1](go/concurrency/channel_1.go) | [Part 2](go/concurrency/channel_2.go) |
    [Part 3](go/concurrency/channel_3.go) | [Part 4](go/concurrency/channel_4.go) |
    [Part 5](go/concurrency/channel_5.go) | [Part 6-end](go/concurrency/channel_6.go)
  - **Patterns**
    - Context: [Part 1](go/concurrency/context_1.go) | [Part 2](go/concurrency/context_2.go) |
    [Part 3](go/concurrency/context_3.go) | [Part 4](go/concurrency/context_4.go) | [Part 5-end](go/concurrency/context_5.go)
    - Pattern
- **Testing and Profiling**
  - Testing: [Part 1](go/testing/basic_test.go) | [Part 2](go/testing/table_test.go)
  - Benchmarking
  - Fuzzing
  - Profiling
- **Packages**
  - Context
  - Encoding
  - IO
  - Logging
  - Reflection
