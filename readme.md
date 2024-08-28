# Advanced Go

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
| What | When |
|------|------|
| Commence | 9:00 AM |
| Tea Break | 10:30 AM (20 mins) |
| Lunch Break | 12:30 PM (1 hr) |
| Tea Break | 3:00 PM (20 mins) |
| Wind up | 4:30 PM |

## Software Requirements
- Go Tools
- VS Code (or any editor)

## Methodology
- No powerpoints
- Code & Discuss
- No dedicated Q & A time

## Repository
- https://github.com/tkmagesh/Cisco-AdvGo-Aug-2024

## Prerequisites
- Data Types, Variables, Constants, iota
- Programming Constructs (if else, for, switch case)
- Functions
    - Higher Order Functions
    - Deferred Functions
- Errors
- Panic & Recovery
- Structs & Methods
    - Struct Composition
- Interfaces
- Modules & Packages

## Agenda
- Recap
- Concurrency
- Adv Concurreny Patterns
- Context
- Database programming choices
- HTTP Services
- GRPC Services
- Testing
- Micro benchmarking
- Profiling

## Recap
- Higher Order Functions
- iota
- interfaces

## Concurrency

## Concurrency Programming
- Designing the application with more than one execution path
- Achieved using OS Threads
- Scheduling Strategies
    - Cooperative Multitasking
    - Pre-emptive Multitasking
- OS Threads are costly
    - ~2MB of memory
    - thread Context switch are costly
    - creating & destroying

## Go Concurrency Model
- Built-in scheduler 
- Concurrent operations are represented as Goroutines
- Goroutines are cheap (~4KB)
- Support for concurrency is built in the language
    - "go" keyword, "channel" data type, channel "<-" operator, range & select-case constructs
- Standard Library Packages
    - "sync" package
    - "sync/atomic" package

### sync.WaitGroup
- semaphore based counter
- has the ability to block the execution of a function until the counter becomes 0