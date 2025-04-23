# learn-go

This repository contains Go example scripts that covers a wide range of Go concepts from basic to advanced. The goal is to provide a comprehensive guide for developers looking to understand the concepts in an easy way.

The code examples provided are straightforward and self-explanatory. Chill out <span style="color:orange; ">comments</span> are provided 🥵.

## Concepts Covered:

1. **Variables**
2. **Arrays**
3. **Strings**
4. **Functions**
5. **Concurrency**
6. **Profiling**
7. **Tracing**

---

## Arrays & Slices

Both arrays and slices store sequences of elements, but they differ in flexibility, memory management, and usage.

## Arrays

An array is a fixed-size sequence of elements of the same type. The size is defined at declaration and cannot be changed.

<p>This means once you declare an array of size <strong>n</strong>, you cannot extend or shrink it.</p>
<h3>Properties</h3>
<ul>
<li>Arrays are stored contiguously in memory.</li>
<li>Arrays are pass-by-value - meaning when passed to a function, arrays are copied, and any modifications inside the function do not affect the original array.</li>
</ul>
<h3>Declaration</h3>
<p>An array belongs to type <code>[n]T</code>. <code>n</code> denotes the number of elements in an array and <code>T</code> represents the type of each element.</p>

```go
package main

import "fmt"

func main() {
	var arr1 [5]int                  // declares an array of type int whose size is 5
	var arr2 = [5]int{1, 2, 3, 4, 5} // declares an array of type int whose size is 5 and initialize values to it.

	fmt.Println(arr1) // output : [0,0,0,0,0]
	fmt.Println(arr2) // output : [1,2,3,4,5]

	// Assigning values to arr1
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5
	fmt.Println(arr1) // output : [1,2,3,4,5]

   // Short Hand initialization
	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr3) // output : [1,2,3,4,5]

	// Declares and initialize an array with default value of the type T except the last index
	arr4 := [...]int{4: 10}
	fmt.Println(arr4) // output : [0,0,0,0,10]
}

```

### Use of `[...]`

- The `[...]` notation tells the compiler to determine the array size based on the highest index provided.
- In this case, `4: 10` means the array should have at least 5 elements (index `4` is the last valid index).
- The compiler infers the size as `5`, making `arr4` an array of type `[5]int`.
- Elipsis determine the length of the array based on the number of initializers

### Memory Allocation

- The array is stored **contiguously** in memory.
- Since only **index 4** is explicitly set to `10`, the remaining elements default to `0` (Go initializes arrays with zero values).

### Array Type

- The type of `arr4` is **`[5]int`**, which is different from a slice (`[]int`).
- You **cannot use `append`** on `arr4` because it is a **fixed-size** array.

The size of the array is a part of the type. Hence `[5]int` and `[25]int` are distinct types. Because of this, arrays cannot be resized.

```go
package main

func main() {
	var arr1 [5]int
	var arr2 = [10]int
	arr2 = arr1 // Not Possible
}
```

<span style="color:red">error : cannot use a (type [5]int) as type [10]int in assignment</span>

### Length and Capacity of an Array

The `len` function returns the number of elements in the array.

Since arrays in Go have a fixed size, `len(arr)` will always return the declared size of the array

The `cap` function returns the number of elements array can hold.

Unlike slices, where `cap` can be different from len, for arrays, `len(arr) == cap(arr)` always holds `true`.

```go
package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println("length : ", len(arr)) // 5
	fmt.Println("capacity : ", cap(arr)) // 5
	fmt.Println(len(arr) == cap(arr)) // true
}
```

## Slices

A slice is a dynamically sized, flexible view into an underlying array. Unlike arrays, slices do not have a fixed size and can grow or shrink as needed.

### Are slices contiguously stored?
Yes, only when they are first created from an array.
No, if append triggers a reallocation.

### Slice Declaration
A slice is declared using []T, where T is the type of elements.
Unlike arrays, a slice does not specify a fixed length.

```go
package main

import (
	"fmt"
)

func main() {
	var slice []int                   // declare a slice
	var slice2 = []int{1, 2, 3, 4, 5} // declare and initialize a slice

	fmt.Println(slice)        // [], slice is declared but not yet initialized
	fmt.Println(slice2)       // [1,2,3,4,5]
	fmt.Println(slice == nil) // true

	slice3 := make([]int, 5)     // Creates a slice of length 5 and initialized default value 0
	slice4 := make([]int, 3, 10) // Length 3, Capacity 10

	fmt.Println(slice3) // [0,0,0,0,0] slice is initialized
	fmt.Println(slice4) // [0,0,0] slice is initialized
}
```

<span style="color:red">The main difference between the array and slice is that, for array you specify the size in advance (static allocation), but for slices you need not know the size in prior and the size can change dynamically
</span>

### Slicing an Array

Slices are created by taking a portion of an array.

```go
package main

import (
	"fmt"
)

func main() {
	var slice = []int{1, 2, 3, 4, 5}
	slice = slice[:3]

	fmt.Println(slice)      // [1,2,3]
	fmt.Println(len(slice)) // 3
	fmt.Println(cap(slice)) // 5
}
```
***Note: The capacity of the slice is determined from the starting index to the end of the underlying array.***

### Growing a Slice with `append`

Slices are dynamic and new elements can be appended to the slice using `append` function

```go
package main

import (
	"fmt"
)

func main() {
	var slice = make([]int, 2, 3)

	fmt.Println(slice)      // [0,0]
	fmt.Println(len(slice)) // 2
	fmt.Println(cap(slice)) // 3

	slice = append(slice, 1, 1, 2, 2)
	fmt.Println(slice)      // [0,0,1,1,2,2]
	fmt.Println(len(slice)) // 6
	fmt.Println(cap(slice)) // 6
}
```

### Slices are mutable

A slice in Go is mutable because it shares the same underlying array. Modifying elements in a slice affects the original array.

```go
package main

import (
	"fmt"
)

func main() {
	var x = []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	fmt.Println(x) // [1,2,3,4,5,6]

	y := x
	x = x[:3] // x is sliced after y has a refernce
	fmt.Println(x) // [1,2,3]
	fmt.Println(y) // [1,2,3,4,5,6]

	y = x // y is updated with new x's reference
	fmt.Println(y) // [1,2,3]
}
```
---
## Interfaces

Interfaces provide a way to define behavior. An interface is a type that specifies a set of method signatures, and any type that implements those methods implicitly satisfies the interface. Interfaces are a key part of Go’s approach to polymorphism and abstraction

```go

package main
import "fmt"

// Define an interface
type Speaker interface {
   Speak() string
}

// Define a type that implements the interface
type Dog struct{}

func (d Dog) Speak() string {
   return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
   return "Meow!"
}

func main() {
   var s Speaker

   s = Dog{}
   fmt.Println(s.Speak()) // Output: Woof!

   s = Cat{}
   fmt.Println(s.Speak()) // Output: Meow!
}
```


---
## Concurrency Patterns:

1. **Generator :**
   Generator functions are used to <span style="color:blue; ">generate</span> a sequence of values.</span>
2. **Consumers :**
   Consumer <span style="color:blue; ">consumes</span> the value produced by Generator.</span>
3. **Pipeline :**
   A series of stages where each stage processes data and passes it to the next stage. This pattern is useful when you have a <span style="color:blue; ">chain</span> of computations or transformations.
4. **Fan-Out and Fan-In Pattern :**
   - **Fan-out**:
     Distribute tasks to multiple goroutines (workers) to process data concurrently.
   - **Fan-in**: Merge results from multiple goroutines into a single result.
5. **Worker Pool Pattern :**
   A common pattern where multiple worker goroutines process a queue of tasks. This pattern is helpful when dealing with <span style="color:red; ">resource-heavy or computationally expensive</span> tasks.
6. **Mutexes and Synchronization :**
   To prevent race conditions, Go provides synchronization primitives like `sync.Mutex` and `sync.WaitGroup`.
7. **Wait Groups**: Used to wait for a collection of goroutines to finish executing.

## Achieving Parallelism in Go

## Concurrency vs Parallelism

Although often used interchangeably, **concurrency** and **parallelism** are distinct concepts in computer science

---

### Concurrency

- **Definition**: Concurrency is about **dealing with many tasks at once**, by allowing tasks to make progress without necessarily finishing simultaneously.
- **How it works**: Multiple tasks are executed by **time-sharing** on a single CPU core.
- **Use case**: Useful when tasks are **I/O-bound** or when you need to manage many tasks efficiently.

> Think of concurrency as a single chef preparing multiple dishes by switching between them based on readiness of ingredients.

---

### Parallelism

- **Definition**: Parallelism is about **doing many tasks at exactly the same time**.
- **How it works**: Multiple tasks are executed **simultaneously on multiple CPU cores**.
- **Use case**: Ideal when tasks are **CPU-bound** and you want to reduce total execution time.

> Think of parallelism as having multiple chefs, each cooking their own dish at the same time.

---

### Summary Table

| Aspect       | Concurrency                          | Parallelism                          |
|--------------|--------------------------------------|--------------------------------------|
| Definition   | Structuring tasks to run independently | Executing tasks simultaneously       |
| CPU Usage    | Can happen on a single core           | Requires multiple cores              |
| Main Benefit | Responsiveness, managing complexity   | Speed and performance                |
| In Go        | Achieved using goroutines             | Achieved using goroutines + multiple cores via `runtime.GOMAXPROCS()` |


You can achieve **parallelism** in Go by leveraging **goroutines**, the **Go scheduler**, and proper configuration of **CPU core usage**.

---

### Goroutines

- Goroutines are **lightweight threads** managed by the Go runtime.
- They are more efficient than traditional threads and are ideal for concurrent and parallel programming.

---

### CPU Core Utilization

- By default, Go can utilize **multiple CPU cores**.
- Use the following to set the number of CPU cores the program can use:

```go
import "runtime"

runtime.GOMAXPROCS(runtime.NumCPU())
```
---
This lets the Go scheduler distribute goroutines across multiple OS threads and cores, enabling true parallelism.CPU, then multiple goroutines would be run on same cpu core, making the execution concurrent, if there is more than one cpu core available we can run goroutines on each available core achieving parallelism

## Detecting Data Race

A datarace occurs whenever two go routines access the same shared resource concuurently and at least one
of the access is a write

```go
go run -race main.go
```

Methods to prevent datarace

1. Initialize variables and never modify it again
2. Avoid accessing resources from multiple go routines, share variables using channels
3. Using Mutex locks

Simple Race Porgram

```go
package main

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go Increment()
	go Increment()
	wg.Wait()
	fmt.Println("Counter:", counter)
}

func Increment() {
	for i := 0; i < 1000; i++ {
		counter++ // Race condition: multiple goroutines modifying counter
	}
	wg.Done()
}
```

```go
go run -race main.go
```

Race Detector records all access to the shared resources ocuurred during the execution along with the identity of the go routines that read or wrote the variable.

---

## Go Routine leak

Go Routine leak is a situation when more than one go routine is running and synchronizing over an unbuffered channel, and one go routine is writing into a channel frm which no go routine would ever receive.

Leaked go routines are not automatically collected by GC

---

## Memory Synchronization

Many processors have its own local cache of main memory
For efficiency, writes to the memory are buffered within each processor and flushed out to main memory only when necessary
If multiple go routines operate, the modified value may not be visible to other go routines.
Synchronization primitives like channels and mutex operations cause the processor to flush and commit all writes so that the writes upto that point are visible to go routines running on other processors.

Mutex and Semaphores are kernel resources that provide synchronization services (also known as synchronization primitives).

### Mutex Lock
Mutex is mainly used to provide mutual exclusion to a specific portion of the code (critical section) and is used in concurrent programming to manage access to shared resources, ensuring that only one thread can access the shared resource at a time.

### Mutex Locks in Golang
Package `sync` provides basic synchronization primitives such as mutual exclusion locks.
A mutex (mutual exclusion lock) in Golang is used to prevent race conditions by ensuring that only one goroutine accesses a shared resource at a time.

### Types of Mutexes 
Golang provides `sync.Mutex` and `sync.RWMutex` for managing concurrent access
1. `sync.Mutex` (Exclusive Lock)</br>
Only one go routine can read/write at a time by acquiring the lock. When one goroutine acquires a sync.Mutex lock, no other goroutine can read or write until the lock is released.
```go
package main
import (
   "fmt"
   "sync"
)
var (
   counter int
   mutex   sync.Mutex
)
func increment(wg *sync.WaitGroup) {
   defer wg.Done()
   mutex.Lock()   // Acquire lock
   counter++      // Critical section
   mutex.Unlock() // Release lock
}

func main() {
   var wg sync.WaitGroup
   for i := 0; i < 10; i++ {
       wg.Add(1)
       go increment(&wg)
   }
   wg.Wait()
   fmt.Println("Final Counter:", counter)
}
```

2. `sync.RWMutex` (Read-Write Mutex)</br>
Multiple goroutines can read (read only, no write is allowed in this method, to write use exclusive lock) at a time by acquiring the lock.
Multiple goroutines can acquire `RLock()` simultaneously.
```go
package main
import (
   "fmt"
   "sync"
)

var (
   data  int
   mutex sync.RWMutex
)

func readData(id int, wg *sync.WaitGroup) {
   defer wg.Done()
   mutex.RLock() // Multiple readers allowed
   defer mutex.RUnlock()
   fmt.Println("Goroutine", id, "reading data:", data)
}

func writeData(wg *sync.WaitGroup) {
   defer wg.Done()
   mutex.Lock() // Exclusive write lock
   defer mutex.Unlock()
   data++
   fmt.Println("Writing data:", data)
}

func main() {
   var wg sync.WaitGroup
   for i := 0; i < 5; i++ {
       wg.Add(1)
       go readData(i, &wg)
   }
   wg.Add(1)
   go writeData(&wg)
   wg.Wait()
}
```

## Scenarios

### Case 1: One Goroutine Holds `Lock()`

#### Case 1.1: Another Goroutine Wants to Read or Write
- The second goroutine must wait until the first goroutine releases the acquired lock.

---

### Case 2: One Goroutine Holds `RLock()`

#### Case 2.1: Another Goroutine Wants to Read
- Multiple goroutines can acquire `RLock()` simultaneously.

#### Case 2.2: Another Goroutine Wants to Write
- A writer (`Lock()`) must wait until all readers release `RLock()`.

---

## Race Condition

A **data race** occurs whenever two goroutines access the same shared resource concurrently and **at least one of the accesses is a write**.

---

## Detecting Data Races

Use the Go race detector:

```go
go run -race main.go
```

## Methods to Prevent Data Races

1. **Initialize Variables and Never Modify Again**
2. **Avoid Accessing Shared Resources from Multiple Goroutines**, share data using **channels** instead.
3. **Use Mutex Locks**
   - Use `sync.Mutex` or `sync.RWMutex` to lock access to shared resources.

---

## How the Race Detector Works

- Records **all accesses** to shared resources during execution.
- Tracks the **identity of goroutines** that read or wrote to the variables.

---

## Example: Race Condition

```go
package main

import (
    "fmt"
    "sync"
)

var (
    counter int
    wg      sync.WaitGroup
)

func main() {
    wg.Add(2)
    go Increment()
    go Increment()
    wg.Wait()
    fmt.Println("Counter:", counter)
}

func Increment() {
    for i := 0; i < 1000; i++ {
        counter++ // Race condition: multiple goroutines modifying counter
    }
    wg.Done()
}
```

---

## Channel Behaviour

### Blocking Behavior

**Blocking On Send**

1. Sending to an unbuffered channel blocks until another goroutine is ready to receive.

   ```go
   ch := make(chan int) // Unbuffered channel
   go func() {
      ch <- 42 // Blocks until the value is received
      fmt.Println("Sent")
   }()
   fmt.Println(<-ch)
   ```

2. Sending to a buffered channel blocks only if the buffer is full.
   ```go
   ch := make(chan int, 1) // Buffered channel
   go func() {
      ch <- 42 // will not be blocked
      ch <- 43 // Blocks until the value is received
      fmt.Println("Sent")
   }()
   fmt.Println(<-ch)
   ```
   Once the buffer is full, buffered channel will behave like an unbuffered channel

**Blocking On Receive**

1. Receiving from an unbuffered channel blocks until another goroutine sends a value.
   Receiving from a buffered channel blocks only if the buffer is empty.
   ```go
   ch := make(chan int, 1) // Buffered channel
   go func() {
      fmt.Println("Sent")
      ch <- 42
   }()
   fmt.Println("Receiving")
   fmt.Println(<-ch) // Blocks until a value is sent
   ```
   Once the buffer is empty, buffered channel will behave like an unbuffered channel

### Non-Blocking Behavior

Use a select statement with a default case to make non-blocking operations.

```go
ch := make(chan int)
select {
   case v := <-ch:
      fmt.Println("Received:", v) // Won't execute (no value to receive)
   default:
      fmt.Println("No data") // Executes immediately will not be blocked
}
```

### Reading from a closed channel

**`case 1:` channel buffer is not empty**

If the channel is closed but still has buffered values, those values are read first before getting the zero value.

```go
ch := make(chan int, 2)
ch <- 10
ch <- 20
close(ch)
fmt.Println(<-ch) // Prints: 10
fmt.Println(<-ch) // Prints: 20
fmt.Println(<-ch) // Prints: 0 (channel is closed and empty),this will not be blocked
v, ok := <-ch
fmt.Println(v, ok) // Prints: 0 false

```

**`case 2:` channel buffer is empty (no data in channel)**

Reading from a closed channel returns the zero value of the channel's type and false for `ok`

```go
ch := make(chan string)
close(ch)
v, ok := <-ch
fmt.Println(v, ok) // Prints: “” false
```

**Iterating Over a Channel**

for range on a Channel

`Case 1:` Channel is closed and buffer is not-empty.
Iterates over all values sent to the channel.

```go
// Buffered channel is used,if not it will be blocked on second send and will not reach the loop
ch := make(chan int, 2)
ch <- 10
ch <- 20
close(ch)
for v := range ch {
   fmt.Println(v)
}

// Same Example with UnBuffered channel
ch := make(chan int, 2)
go func() {
    ch <- 10
    ch <- 20
    close(ch)
}()
for v := range ch {
    fmt.Println(v)
}
```

`Case 2:` channel is closed and buffer is empty.

No Output ,as cannot iterate over a channel which is empty and which is closed

```go
ch := make(chan int, 2)
close(ch)
for v := range ch {
 	fmt.Println(v) //No output
}
fmt.Println("Execution will reach here") //Execution will reach this line
```

`Case 3:` Channel is not closed(open) and the buffer is not empty
**UnBuffered Channel**

```go
// Use a goroutine with an unbuffered channel,since the receiver is not yet ready, send would block forever.
ch := make(chan int)
go func() {
 	ch <- 10
	}()
for v := range ch {
 	fmt.Println(v) //Prints 10 & then Deadlock
}
```

**Buffered Channel**

```go
ch := make(chan int,1)
ch <- 10
for v := range ch {
   fmt.Println(v) //Prints 10 & then Deadlock
}
fmt.Println("Execution will not reach here") //Execution will not reach this line
```

`Reason for Deadlock in both cases: The channel is not closed, and the for range loop waits indefinitely for more values.`

`Case 4:` Channel is not closed and the buffer is empty

In case of both buffered and unbuffered channel this would lead to deadlock

```go
ch1 := make(chan int)
for v := range ch1 {
   fmt.Println(v) //No output
}
fmt.Println("Execution will not reach here") //Execution will not reach this line
```

### Closed Channel

Attempting to close an already closed channel causes panic.

```go
ch := make(chan int)
go func() {
   ch <- 10
   close(ch)
	}()
<-ch
close(ch)
```

Attempting to send data to a closed channel causes panic.

```go
ch := make(chan int, 2)
close(ch1)
ch <- 10
```

### Summary of Channel Behaviour

<table>
  <thead>
    <tr>
      <th>Scenario</th>
      <th>Blocking</th>
      <th>Non-Blocking</th>
      <th>Remarks</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Send to unbuffered channel</td>
      <td>YES</td>
      <td>No (if receiver)</td>
      <td>Blocks until received</td>
    </tr>
    <tr>
      <td>Receive from unbuffered channel</td>
      <td>YES</td>
      <td>No (if sender)</td>
      <td>Blocks until value is sent</td>
    </tr>
    <tr>
      <td>Send to buffered channel</td>
      <td>YES (If buffer is full)</td>
      <td>No (if buffer is not full)</td>
      <td>Blocks if buffer is full</td>
    </tr>
    <tr>
      <td>Receive from buffered channel</td>
      <td>Yes (if buffer is empty)</td>
      <td>No (if buffer has some data)</td>
      <td>Blocks if buffer is empty</td>
    </tr>
    <tr>
      <td>Read from closed channel</td>
      <td>NO</td>
      <td>YES</td>
      <td>Returns zero value</td>
    </tr>
    <tr>
      <td>for range on open channel</td>
      <td>YES</td>
      <td>NO</td>
      <td>Blocks if no senders or not closed</td>
    </tr>
    <tr>
      <td>for range on closed channel</td>
      <td>NO</td>
      <td>YES</td>
      <td>Terminates automatically</td>
    </tr>
  </tbody>
</table>

---

## Profiling Go Applications

Profiling is an automated approach to measure performance based on sampling a number of profile events during execution, then extrapolating during a post-processing step. The resulting statistical summary is called a <span style="color:red; ">profile</span>.

Complete reference: [https://go.dev/blog/pprof](https://go.dev/blog/pprof)

### Different Types of Profiles

1. **CPU Profile :**
   <span style="color:teal; ">Identifies the functions that require the most CPU time.</span>
   ```go
   go test -cpuprofile=cpu.out
   ```
2. **Heap Profile :**
   <span style="color:teal; ">Identifies the statements responsible for allocating the most memory.</span>
   ```go
   go test -memprofile=memory.out
   ```
3. **Block Profile :**
   <span style="color:teal; ">Identifies the operations responsible for blocking goroutines the longest.</span>
   ```go
   go test -blockprofile=block.out
   ```

---

## Executing pprof Samples

To analyze the CPU profile, use the following command:

```go
go tool pprof cpu.out
```

## Options

1. <span style="color:blue; font-weight:bold;">web</span>: Generates a visual representation (requires Graphviz).
2. <span style="color:olive; font-weight:bold;">text</span>: Displays the profiling data in text format.
3. <span style="color:green; font-weight:bold;">top</span>: Displays the most CPU-intensive functions.
4. <span style="color:olive; font-weight:bold;">topn</span>: Displays the top n CPU-intensive functions.(n=1,2...)
5. <span style="color:grey; font-weight:bold;">list <function_name></span>: The listing shows the source code for the provided <function_name> (really, for every function matching the regular expression function_name)

```go
app.Use(pprof.New())
```

---

## Profiling Endpoints

To collect profiling data, use the following `curl` command to retrieve different types of profiles:

1. <span style="color:blue; font-weight:bold;">CPU Profile:</span>
   ```go
   curl <server-endpoint>/debug/pprof/profile -o cpu.pprof
   ```
   <span style="color:orange;">alternatively,</span>
   ```go
   go tool pprof <server-endpoint>/debug/pprof/profile
   ```
2. <span style="color:green; font-weight:bold;">Heap Profile:</span>
   ```go
   curl <server-endpoint>/debug/pprof/heap -o heap.pprof
   ```
3. <span style="color:red; font-weight:bold;">Block Profile:</span>
   ```go
   curl <server-endpoint>/debug/pprof/block -o block.pprof
   ```

---

## Executing Using `go tool`

Once you've obtained the `profiling` files, you can analyze them using `go tool pprof`:

```go
go tool pprof cpu.pprof
```

<span style="color:red;">Profiling introduces runtime overhead</span>, so avoid enabling it in production environments unless absolutely necessary.

To generate a visual representation of the profiling data, you need to install Graphviz:

```go
go install github.com/goccy/go-graphviz/cmd/dot@latest
```

### alternative

```go
go tool pprof -web cpu.pprof
```

```go
go tool pprof -text cpu.pprof
```

To see a list if available options use

```go
go tool pprof --help
```

You can serve the profiling data via http server with

```go
go tool pprof -http=:port profile.pb.gz
```

Source: [Graphviz Documentation](https://pkg.go.dev/github.com/goccy/go-graphviz#section-readme)

---

## Tracing

The `go tool trace` command is used in Go to analyze execution traces collected during a program's run

To collect tracing data, use the following `curl` command to retrieve traces

<span style="color:blue; font-weight:bold;">Trace</span>

```go
curl <server-endpoint>/debug/pprof/trace -o trace.out
```

Unlike pprof we donot have different types of traces

### Inspect the Report

The `go tool trace` command starts a local web server and opens a web-based interface with several views. Here are the key sections you can analyze:

```go
go tool trace trace.out
```

- **Goroutines**: Analyze running goroutines and their states.
- **Heap**: View heap memory usage.
- **Scheduler**: Check the Go scheduler's behavior.
- **Network/IO**: Inspect network and I/O events.
- **User-defined Regions**: Examine custom trace annotations for deeper insights.

---

## Integrate with Tests: Use go test to directly collect traces

```go
go test -trace=trace.out
```
