# Ultimate Go

This repo contains my notes on learning Go and computer systems.

**Resources:**
- [Ardan Labs's Ultimate Go course
  ](https://github.com/ardanlabs/gotraining/blob/master/topics/courses/go/README.md).
- [Computer Systems: A Programmer's
  Perspective](https://www.amazon.com/Computer-Systems-Programmers-Perspective-3rd/dp/013409266X)

## Project Status

### Tasks

- [ ] Finish Ultimate Go's training videos and reading resources.
- [ ] Finish Computer System book

## Table of Contents 

- [Design Philosophy](#design-philosophy)
- [Language Mechanics](#language-mechanics)
  - [Syntax](#syntax)
    - Variable
    - Struct
    - Pointer
    - Constant
    - Function
- Software Design
- Concurrency
- Testing and Profiling
- Packages

## Design Philosophy

- Does your performance better? Is it your highest priority?
- Performance vs Productivity?
- Trade-off? Everything comes with a cost
- Optimize for correctness first, then think about performance
- Code Reviews
- Integrity:
  - Be serious about writing code that reliable
  - Less code means less bugs
  - Error handling must be a part of the main code
- Readability:
  - Not hiding the cost of the code or the decision making, the impact
- Simplicity:
  - Hide complexity
- Performance:
  - Compute less
  - Never guess
  - Measurements must be relevant
  - Profile
  - Test

## Language Mechanics

### Syntax

- [Variable](variable.go)
- [Struct](struct.go)
- [Pointer](pointer.go)
- [Constant](constant.go)
- [Function](function.go)
