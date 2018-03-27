# Golang training notes

<!-- TOC -->

- [Golang training notes](#golang-training-notes)
    - [Resources](#resources)
    - [Basics](#basics)
        - [Variables](#variables)
        - [Constants](#constants)
        - [Types](#types)
        - [Arrays and slices](#arrays-and-slices)
        - [Maps](#maps)
        - [Flow control and loops](#flow-control-and-loops)
        - [Structs](#structs)
        - [Pointers](#pointers)
    - [Methods and interfaces](#methods-and-interfaces)
        - [Methods](#methods)
        - [Interfaces](#interfaces)
        - [Embedding](#embedding)
    - [Goroutines](#goroutines)
        - [Mutex](#mutex)
        - [Channels](#channels)
        - [Context](#context)
    - [Tests](#tests)
        - [Running tests](#running-tests)
        - [Code coverage](#code-coverage)
        - [Example](#example)
    - [Exercises / examples](#exercises--examples)
        - [Tour exercises](#tour-exercises)
        - [Context exercise](#context-exercise)

<!-- /TOC -->

---

## Resources

* [A tour of go](https://tour.golang.org)
* [The little go book](http://openmymind.net/The-Little-Go-Book/)
* [How to write go code](https://golang.org/doc/code.html)
* [Writting web applications](https://golang.org/doc/articles/wiki/)
* [Effective Go](https://golang.org/doc/effective_go.html#introduction)
* [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## Basics

### Variables

```go
// Declaration
var a int
a = 42
// Shortcut with type infering
b := 42
```

### Constants

To be placed after import

```go
const A = 3
const (
    timeout = 0.75
    url     = "https://github.com/"
    flag    = true
)
```

### Types

```
// false is Zero value for bools
bool

// "" is Zero value for string
string

// 0 is the Zero value for numbers
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr


byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

Type conversions:

```go
i := 42
f := float64(i)
u := uint(f)
```

### Arrays and slices

```go
// Arrays (fixed size)
var a [2]string
a[0] = "Hello"
a[1] = "World"
primes := [6]int{2, 3, 5, 7, 11, 13}

// slices (dynamic size), refereces of an underlying array which can be explicit or implicit
s := primes[1:4]
s[0] = 99
fmt.Println(primes)

s := []struct {
    i int
    b bool
}{
    {1, true},
    {2, true},
    {3, false},
}
// len of the slice
len(s) // 3
// capacity of the underlying array
cap(s) // 3
s = s[:1]
s[3] // would throw 'index out of range'
s = s[1:5] // We can resize the array

s := make([]int, 5) // Creates an slice of Zero elements with len 5
s := make([]int, 0, 5) // Creates an slice of 0 len an 5 cap

// Append will add elements to the slice even if excees the unerlying array capacity
var a []int
a = append(a, 3)
```

### Maps

```go
var m0 = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -74.39967,
    },
    "Google": Vertex{
        37.42202, -122.08408,
    },
}
// Type name can be omitted
m0 = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}
fmt.Println(m0)

m := make(map[string]int)

// map operations
m["Answer"] = 42
fmt.Println("The value:", m["Answer"])

m["Answer"] = 48
fmt.Println("The value:", m["Answer"])

delete(m, "Answer")
fmt.Println("The value:", m["Answer"])

v, ok := m["Answer"]
fmt.Println("The value:", v, "Present?", ok)
```

### Flow control and loops

Conditionals:

```go
if _, err := someFunct(); err != nil {
    // handle error
}

if someCondition {
    // ...
} else if someOtherCondition {
    // ...
} else {
    // ...
}
```

Loops:

```go
for i := 0; i < N; i++ {  // any of the three separated by ; is optional
    // "regular" loop
}

for {
    // infinite loop
}

for condition {
    // like while in other languages
}

// Iter slices
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
for i, v := range pow {
    fmt.Printf("%v^2 = %v", i, v)
}
for i := range pow {
    // Only iter indexes
}
for _, v := range pow {
    // Skip indexes
}

// Iter maps
m = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}
for k, v := range(m) {
    fmt.Printf("%v, coordinates: %v", k, v)
}
for k := range(m) {
    // Only iter keys
}
for _, v := range(m) {
    // Only values
}
```

### Structs

```go
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

// Copy a struct (or any other variable)
copy := v1
```

### Pointers

```go
i, j := 42, 2701

p := &i         // point to i
fmt.Println(*p) // read i through the pointer
*p = 21         // set i through the pointer
fmt.Println(i)  // see the new value of i

p = &j         // point to j
*p = *p / 37   // divide j through the pointer
fmt.Println(j) // see the new value of j

v := Vertex{1, 2}
q := &v  // point to v
q.X = 10 // set a struct field directly (no need to do (*q).X)
fmt.Println(v)
```

## Methods and interfaces

### Methods
```go
type MyFloat float64
// Abs for custom float
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
f := MyFloat(-math.Sqrt2)
fmt.Println(f.Abs())

// Usually methods are defined with pointers receivers as then, it is possible to modify the value
type Vertex struct {
    X, Y float64
}
func (v *Verteex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}
// The equivalent function
func ScaleFunc(v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

v := Vertex(3, 4)
v.Scale(2) // This is syntaxis-sugar for (&v).Scale(2)
ScaleFunc(&v, 2) // this would be the equivalent

v := &Vertex{3, 4}
v.Scale(2) // This work the same

// Method using valure receiver
func (v Vertex) Abs() float64 {
    return mat.Sqrt(v.X*v.X + v.Y*v.Y)
}

v := Vertex{3, 4}
v.Abs()

v := &Vertex{3, 4}
v.Abs() // syntaxic-sugar for (*v).Abs()
```

### Interfaces

An interface is a definition of signature (without implementation), the interface
implementation is implicit (a type implements an interface if it defines all its methods)

```go
type I interface {
    M()
}

type T struct {
    S string
}

type F float64

func (t T) M() {
    fmt.Println(t.s)
}

func (f F) M() {
    fmt.Println(f)
}

// Both types implement the interface I

func use(i I) {
    i.M()
}

use(F(42))
use(I{"Meaning"})
```

Methods can handle nil values

```go
func (t *T) Show() {
    if t == nil {
        fmt.Println("Nope")
    } else {

    }
}
```

The empty interface describes any type `interface{}`

```go
func describe(any interface{}) {
    fmt.Printf("(%v, %T\n)", any, any)
}
```

We can go from the interface to its type using type assertions

```go
var nothing interface{}
nothing = "hello"
s := nothing.(string)
nothing = 42
i := nothing.(int)
x := nothing.(string) // panic (conversion)
nothing = 1.4
// Type switch
switch v := empty.(type) {
case int:
    fmt.Println("Is int!")
case string:
    fmt.Println("Is string")
default:
    fmt.Println("Is other thing...") // will be here
}
```

Stringers implement the string method

```go
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

me := Person{"Christian", 32}
fmt.Println(me) // Will print "Christian (32 years)"
```

Errors interface is similar to Stringers

```go
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

if err := run(); err != nil {
    fmt.Println(err)
}
```

### Embedding

```go
// ReadWriter implementations must satisfy both Reader and Writer
type ReadWriter interface {
    Reader
    Writer
}

// Server exposes all the methods that Logger has
type Server struct {
    Host string
    Port int
    *log.Logger
}
```

## Goroutines

Goroutines are light threads (not OS threads, but routines managed by Go). To start a new goroutine it is just needed to include `go` before a function call. Example: `go f(a)`.

Also an anonymous function may be used:

```go
go func (a int) {
    // do stuff
}(42)
```

### Mutex

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```

The `sync` package also includes [RWMutex](https://golang.org/pkg/sync/#RWMutex) which allows to only lock read and/or write operations.

### Channels

Channels are concurrency-safe communication objects. See more info in the [channels section](https://tour.golang.org/concurrency/2) of the tutorial.

```go
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel
}

func main() {
    s := []int{7, 2, 8, -9, 4, 8}
    c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from channel
}
```

Buffered chanells have length (length can be tested with `cap` function)

```go
c := make(chan int, 2)
c <- 1
c <- 2
// Will get an error if overfill the channel
// c <- 3 // fatal error: all goroutines are asleep - deadlock
```

Sender can close channels to tell no more values are comming. Receivers can test if a channel is closed using the second parameter.

Receiving from a closed channels returns the zero value inmediately.

```go
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
    c := make(chan int, 10)
    go fibonnaci(cap(c), c)
    for i := range c {
        fmt.Println(i)
    }
}
```

`select` lets a goroutine wait on multiple communication operations. It blocks until one of the cases can run and choses ramdonly if multiple can run.

```go
func fibonacci(c, quit chan it) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <- quit:
            fmt.Println("I'm done!")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println( <- c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}
```

A default case in slect let us try a send/receive without blocking

```go
func main() {
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisencond)
    for {
        select {
        case <- tick:
            fmt.Println("tick.")
        case <- boom:
            fmt.Printl("BOOM!")
            return
        default:
            fmt.Println(".")
            time.Sleep(50 * time.Millisecond)
        }
    }
}
```

### Context

The `Context` type carries deadlines, cancelation signals and other request-scoped values across API boundaries and between processes. See:

* [context package docs](https://golang.org/pkg/context/)
* [Context as concurrency patter in go](https://blog.golang.org/context)

See [Context exercise](#context-exercise).


## Tests

Test files live next to the files containing the code that are testing:

```
foo.go
foo_test.go
```

The naming is a required pattern in Go.

Test (always included in `*_test.go`files) must be named `Test<name>(* testing.T)`.

```go
package main
import "testing"
func TestSimple(t *testing.T) {
    if true {
        t.Error("expected false, got true")
    }
}
```

`*testing.T` type has method available to control the tests flow. See [testing docs](https://golang.org/pkg/testing/).


To avoid circular dependency imports go let's us create a "magic" testing package.

```go
// foo.go
package foo

// foo_test.go
package foo_test
```

It is optional and is an exception to "one package per folder" requirement. It helps with circular dependencies but, as it is a new package, all rules about exporting names apply.

### Running tests

```bash
go test <package-path> # run tests of package in path
go test ./... # run all test in packages recursively
go test -v . # run test in current package in verbose mode
go test -run "Call" -v ./... # run all test maching the "Call" regular expression
go test -race race_test.go # run race_test.go checkin race conditions
```

### Code coverage

```bash
# For all tests
go test -coverprofile cover.out
go tool cover -html=cover.out
# For specific test
go test -coverprofile cover.out -run TestUser_Validate
go tool cover -html=cover.out
```

### Example

See [test exercise](./src/calculate/).

## Exercises / examples

### Tour exercises

To be able to run some exercises it is needed to install the tour in the GOPATH

```bash
go get golang.org/x/tour/gotour
````

* [basics](./src/basics)
* [methods](./src/methods)
* [goroutines](./src/goroutines)

### Context exercise

> At <http://hamlet.gopherguides.com> you will find a text copy of the play Hamlet.
>
> Write a program using contexts that searches each line of text from the play to search for a particular word. Examples are Hamlet, Mark, King etc...
>
> The program must:
>
> print out the line number the word was found, and the text of that line: `3545: QUEEN GERTRUDE O Hamlet, speak no more`
> stop when it finds 50 occurrences of the word
> stop after 5 seconds if it hasn't yet found 50 occurrences

* [Context exercise solution](./src/context)

