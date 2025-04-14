# Go Concurrency Learning Guide

This project is designed to help you learn and understand Go's concurrency model. Below is the recommended sequence to explore the files and their concepts.

## Learning Path

1. **Goroutines Basics** (`goRoutines/goRoutines.go`)
    - Learn how to create and use goroutines.
    - Understand how Go handles concurrency using the `go` keyword.
    - Example: Running multiple functions concurrently.

2. **Unbuffered Channels** (`channels/channels.go`)
    - Learn how to use unbuffered channels for communication between goroutines.
    - Understand the `select` statement to handle multiple channels.

3. **Buffered Channels** (`BufferChannel/bufferChannel.go`)
    - Explore buffered channels and their differences from unbuffered channels.
    - Learn how to send and receive data with buffered channels.
    - Understand how to close channels and iterate over them.

4. **Graceful Goroutine Termination** (`infiniteGoRoutines/infiniteRoutines.go`)
    - Learn how to stop goroutines gracefully using a `done` channel.
    - Understand the use of `select` with a `default` case for non-blocking operations.

## Files Overview

- **`goRoutines/goRoutines.go`**: Demonstrates the basics of goroutines.
- **`channels/channels.go`**: Explains unbuffered channels and the `select` statement.
- **`BufferChannel/bufferChannel.go`**: Covers buffered channels and their usage.
- **`infiniteGoRoutines/infiniteRoutines.go`**: Shows how to manage long-running goroutines and terminate them gracefully.

## How to Run

1. Navigate to the directory of the file you want to run.
2. Use the following command to execute the file:
   ```bash
   go run <filename>.go