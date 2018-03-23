# golang training notes

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
