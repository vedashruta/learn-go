<style>
    body {
        font-family: Arial, sans-serif;
        line-height: 1.6;
        background-color: #f4f4f4;
    }
    .container {
        margin: auto;
        background: white;
        padding: 20px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    }
    h2 {
        color: #333;
        border-bottom: 2px solid #0073e6;
        padding-bottom: 5px;
    }
    ul {
        list-style-type: circle;
        margin-left: 20px;
    }
</style>
<div class="container">

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
<h2>Array</h2>
<p>An array is a fixed-size sequence of elements of the same type. The size is defined at declaration and cannot be changed.</p>
<p>This means once you declare an array of size <strong>n</strong>, you cannot extend or shrink it.</p>
<h3>Properties</h3>
<ul>
<li>Arrays are stored contiguously in memory.</li>
<li>Arrays are pass-by-value - meaning when passed to a function, arrays are copied, and any modifications inside the function do not affect the original array.</li>
</ul>
<h3>Declaration</h3>
<p>An array belongs to type <code>[n]T</code>. <code>n</code> denotes the number of elements in an array and <code>T</code> represents the type of each element.</p>

```bash
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
---
- The `[...]` notation tells the compiler to determine the array size based on the highest index provided.  
- In this case, `4: 10` means the array should have at least 5 elements (index `4` is the last valid index).  
- The compiler infers the size as `5`, making `arr4` an array of type `[5]int`.

### Memory Allocation
---
- The array is stored **contiguously** in memory.
- Since only **index 4** is explicitly set to `10`, the remaining elements default to `0` (Go initializes arrays with zero values).

### Array Type
---
- The type of `arr4` is **`[5]int`**, which is different from a slice (`[]int`).
- You **cannot use `append`** on `arr4` because it is a **fixed-size** array.
---
The size of the array is a part of the type. Hence `[5]int` and `[25]int` are distinct types. Because of this, arrays cannot be resized.
```bash
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

```bash
package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println("length : ", len(arr)) // 5
	fmt.Println("capacity : ", cap(arr)) // 5
	fmt.Println(len(arr) == cap(arr)) // true
}
```

<h2>Slice</h2>


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

---

## Detecting Data Race

A datarace occurs whenever two go routines access the same shared resource concuurently and at least one
of the access is a write

```bash
go run -race main.go
```

Methods to prevent datarace

1. Initialize variables and never modify it again
2. Avoid accessing resources from multiple go routines, share variables using channels
3. Using Mutex locks

Simple Race Porgram

```bash
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

```bash
go run -race main.go
```

Race Detector records all access to the shared resources ocuurred during the execution along with the identity of the go routines that read or wrote the variable.

---

## Go Routine leak:

Go Routine leak is a situation when more than one go routine is running and synchronizing over an unbuffered channel, and one go routine is writing into a channel frm which no go routine would ever receive.

Leaked go routines are not automatically collected by GC

---

## Memory Synchronization

Many processors have its own local cache of main memory
For efficiency, writes to the memory are buffered within each processor and flushed out to main memory only when necessary
If multiple go routines operate, the modified value may not be visible to other go routines.
Synchronization primitives like channels and mutex operations cause the processor to flush and commit all writes so that the writes upto that point are visible to go routines running on other processors.

---

## Channel Behaviour

### Blocking Behavior

**Blocking On Send**

1. Sending to an unbuffered channel blocks until another goroutine is ready to receive.

   ```bash
   ch := make(chan int) // Unbuffered channel
   go func() {
      ch <- 42 // Blocks until the value is received
      fmt.Println("Sent")
   }()
   fmt.Println(<-ch)
   ```

2. Sending to a buffered channel blocks only if the buffer is full.
   ```bash
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
   ```bash
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

```bash
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

```bash
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

```bash
ch := make(chan string)
close(ch)
v, ok := <-ch
fmt.Println(v, ok) // Prints: “” false
```

**Iterating Over a Channel**

for range on a Channel

`Case 1:` Channel is closed and buffer is not-empty.
Iterates over all values sent to the channel.

```bash
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

```bash
ch := make(chan int, 2)
close(ch)
for v := range ch {
 	fmt.Println(v) //No output
}
fmt.Println("Execution will reach here") //Execution will reach this line
```

`Case 3:` Channel is not closed(open) and the buffer is not empty
**UnBuffered Channel**

```bash
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

```bash
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

```bash
ch1 := make(chan int)
for v := range ch1 {
   fmt.Println(v) //No output
}
fmt.Println("Execution will not reach here") //Execution will not reach this line
```

### Closed Channel

Attempting to close an already closed channel causes panic.

```bash
ch := make(chan int)
go func() {
   ch <- 10
   close(ch)
	}()
<-ch
close(ch)
```

Attempting to send data to a closed channel causes panic.

```bash
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
   ```bash
   go test -cpuprofile=cpu.out
   ```
2. **Heap Profile :**
   <span style="color:teal; ">Identifies the statements responsible for allocating the most memory.</span>
   ```bash
   go test -memprofile=memory.out
   ```
3. **Block Profile :**
   <span style="color:teal; ">Identifies the operations responsible for blocking goroutines the longest.</span>
   ```bash
   go test -blockprofile=block.out
   ```

---

## Executing pprof Samples

To analyze the CPU profile, use the following command:

```bash
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
   ```bash
   curl <server-endpoint>/debug/pprof/profile -o cpu.pprof
   ```
   <span style="color:orange;">alternatively,</span>
   ```bash
   go tool pprof <server-endpoint>/debug/pprof/profile
   ```
2. <span style="color:green; font-weight:bold;">Heap Profile:</span>
   ```bash
   curl <server-endpoint>/debug/pprof/heap -o heap.pprof
   ```
3. <span style="color:red; font-weight:bold;">Block Profile:</span>
   ```bash
   curl <server-endpoint>/debug/pprof/block -o block.pprof
   ```

---

## Executing Using `go tool`

Once you've obtained the `profiling` files, you can analyze them using `go tool pprof`:

```bash
go tool pprof cpu.pprof
```

<span style="color:red;">Profiling introduces runtime overhead</span>, so avoid enabling it in production environments unless absolutely necessary.

To generate a visual representation of the profiling data, you need to install Graphviz:

```bash
go install github.com/goccy/go-graphviz/cmd/dot@latest
```

### alternative

```bash
go tool pprof -web cpu.pprof
```

```bash
go tool pprof -text cpu.pprof
```

To see a list if available options use

```bash
go tool pprof --help
```

You can serve the profiling data via http server with

```bash
go tool pprof -http=:port profile.pb.gz
```

Source: [Graphviz Documentation](https://pkg.go.dev/github.com/goccy/go-graphviz#section-readme)

---

## Tracing

The `go tool trace` command is used in Go to analyze execution traces collected during a program's run

To collect tracing data, use the following `curl` command to retrieve traces

<span style="color:blue; font-weight:bold;">Trace</span>

```bash
curl <server-endpoint>/debug/pprof/trace -o trace.out
```

Unlike pprof we donot have different types of traces

### Inspect the Report

The `go tool trace` command starts a local web server and opens a web-based interface with several views. Here are the key sections you can analyze:

```bash
go tool trace trace.out
```

- **Goroutines**: Analyze running goroutines and their states.
- **Heap**: View heap memory usage.
- **Scheduler**: Check the Go scheduler's behavior.
- **Network/IO**: Inspect network and I/O events.
- **User-defined Regions**: Examine custom trace annotations for deeper insights.

---

## Integrate with Tests: Use go test to directly collect traces

```bash
go test -trace=trace.out
```

</div>
