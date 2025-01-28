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

You can serve the profiling data via a server with
```bash
go tool pprof -http=:port profile.pb.gz
```

Source: [Graphviz Documentation](https://pkg.go.dev/github.com/goccy/go-graphviz#section-readme)

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

## Integrate with Tests: Use go test to directly collect traces
```bash
go test -trace=trace.out
```