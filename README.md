# learn-go

This repository contains Go example scripts that covers a wide range of Go concepts from basic to advanced. The goal is to provide a comprehensive guide for developers looking to understand the concepts in an easy way.

The code examples provided are straight forward and self explainatory, chill-out comments are provided 🥵.

## Concepts Covered:
1. **Variables**
2. **Arrays**
3. **Strings**
4. **Fuctions**
5. **Concurrency**


## Concurrency Patterns:
1. **Generator**: Generator functions are used to generate a sequence of values.
2. **Consumers**: Consumer consumes the value produced by Generator.
3. **Pipeline**: A series of stages where each stage processes data and passes it to the next stage. This pattern is useful when you have a chain of computations or transformations.
4. **Fan-Out and Fan-In Pattern**: 
   - **Fan-out**: Distribute tasks to multiple goroutines (workers) to process data concurrently.
   - **Fan-in**: Merge results from multiple goroutines into a single result.
5. **Worker Pool Pattern**: A common pattern where multiple worker goroutines process a queue of tasks. This pattern is helpful when dealing with resource-heavy or computationally expensive tasks.
6. **Mutexes and Synchronization**: To prevent race conditions, Go provides synchronization primitives like `sync.Mutex` and `sync.WaitGroup` to ensure proper coordination between goroutines.
7. **Wait Groups**: A synchronization primitive used to wait for a collection of goroutines to finish executing.

## How to Use:

### 1. Clone the repository:

```bash
git clone https://github.com/vedashruta/learn-go.git
cd learn-go
```
### 2. Uncomment the code in main() function
Uncomment the code in the main() and run main.go to see the code in action
```
func main() {
	// concurrency.Channels()
	// strings.StringMethods()

	// concurrency.Concurrency()
	// variables.Variables()

	// uncomment below line for Generator pattern
	concurrency.Generator(1, 10)
}
```