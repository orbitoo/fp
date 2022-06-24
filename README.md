# FP -- a package designed for functional programming

## boxed --  a collection of useful functors and monads

### maybe

Maybe is designed for calculations which might cause error.

```go
type Maybe[T any] struct {
    value     T
    isNothing bool
}
```

* Just

```go
x := Just(1)
fmt.Println(x)
// Just 1
```

* Nothing

```go
x := Nothing[int]()
fmt.Println(x)
// Nothing
```

* Map

```go
x := Just(1)
add1 := func (x int) int { return x + 1 }
x = x.Map(add1)
fmt.Println(x)
// Just 2
```

* FlatMap

When a mapping returns a `Maybe`, you should use `FlatMap`.

```go
x := Just(-1.0)
sqrt := func (x float64) Maybe[float64] {
    if x < 0 {
        return Nothing[float64]()
    }
    return Just(math.Sqrt(x))
}
x = x.FlatMap(x)
fmt.Println(x)
// Nothing
```

* Get

``` go
x := Just(1)
y := Nothing[int]()
xv, xok := x.Get()
if xok {
    fmt.Println("x =", xv)
} else {
    fmt.Println("x is empty")
}
// x = 1
yv, yok := y.Get()
if yok {
    fmt.Println("y =", yv)
} else {
    fmt.Println("y is empty")
}
// y is empty
```

* IsNothing

```go
x := Just(1)
y := Nothing[int]()
fmt.Println(x.IsNothing())
// false
fmt.Println(y.IsNothing())
// true
```

* LiftMaybe2

`LiftMaybe` converts a normal function to a function which handles `Maybe`s.
`LiftMayben` means the original function has n parameters.

```go
add := func (x, y int) int { return x + y }
addMaybe := LiftMaybe2(add)
x := Just(1)
y := Just(2)
z := Nothing[int]();
fmt.Println(addMaybe(x, y))
// Just 3
fmt.Println(addMaybe(x, z))
// Nothing
```

`LiftMaybe1` to `LiftMaybe5` are provided.

The function you get from`LiftMaybe` will always return `Nothing` if any parameter is `Nothing`.
If you have any special needs, consider rewrite manually.
`IsNothing()` will give you some help.

This is a substitution of `Apply` function of applicative. Currently we can't implement this due to the design of generic in Golang.

### Either

`Right` is the normal path, which is the same as Haskell.
You can maintain more information in the `Left` path when calculation failed.

```go
type Either[T1, T2 any] struct {
    left   T1
    right  T2
    isLeft bool
}
```

* Left

```go
x := Left[string, int]("square root of a negative number")
fmt.Println(x)
// Left square root of a negative number
```

* Right

```go
x := Right[string](1)
fmt.Println(x)
// Right 1
```

* Map

```go
add1 := func(x int) int { return x + 1 }
x := Left[string, int]("student does not exist")
y := Right[string](84)
x.Map(add1)
y.Map(add1)
fmt.Println(x)
// Left student does not exist
fmt.Println(y)
// Right 85
```

* FlatMap

```go
sqrt := func(x float64) Either[string, float64] {
    if x < 0 {
        return Left[string, float64]("square root of a negative number")
    }
    return Right[string](math.Sqrt(x))
}
x := Right[string](1.0)
y := Right[string](-1.0)
x = x.FlatMap(sqrt)
y = y.FlatMap(sqrt)
fmt.Println(x)
// Right 1
fmt.Println(y)
// Left square root of a negative number
```

* Get

```go
x := Left[string, int]("student does not exist")
y := Right[string](84)
xl, xr, xok := x.Get()
yl, yr, yok := y.Get()
fmt.Println(xl, xr, xok)
// student does not exist 0 false
fmt.Println(yl, yr, yok)
//  84 true
```

* IsLeft

```go
x := Left[string, int]("student does not exist")
y := Right[string](84)
xleft := x.IsLeft()
yleft := y.IsLeft()
fmt.Println(xleft)
// true
fmt.Println(yleft)
// false
```

* LiftEither2

`LiftEither` converts a normal function to a function which handles `Either`s.
`LiftEithern` means the original function has n parameters.

```go
add := func(x, y int) int { return x + y }
addEither := LiftEither2[string](add)
x := Left[string, int]("student does not exist")
y := Right[string](91)
z := Right[string](95)
fmt.Println(addEither(z, y))
// Right 186
fmt.Println(addEither(z, x))
// Left student does not exist
```

`LiftEither1` to `LiftEither5` are provided.

The function you get from`LiftEither` will always return `Left` if any parameter is `Left`.
If you have any special needs, consider rewrite manually.
`IsLeft()` will give you some help.

This is a substitution of `Apply` function of applicative. Currently we can't implement this due to the design of generic in Golang.

### Result

This is inspired by Rust.

```go
type Result[T any] struct {
    v       T
    e       error
    isError bool
}
```

* Ok

```go
x := Ok(1)
```

* Err

```go
x := Err(errors.New("file not found"))
```

* Expect

```go
openFile := func(path string) Result[*os.File] {
    v, e := os.Open(path)
    if e != nil {
        return Err[*os.File](e)
    }
    return Ok(v)
}
file := openFile("test.txt").Expect("error occurred when opening the file")
// panic if there is an error,
// and string in `Expect` will be printed as log
```

* Unwrap

```go
openFile := func(path string) Result[*os.File] {
    v, e := os.Open(path)
    if e != nil {
        return Err[*os.File](e)
    }
    return Ok(v)
}
file := openFile("test.txt").Unwrap()
// panic if there is an error
```

* Default

When error occurred, you may want to return a default value.

```go
sqrt := func(x float64) f.Result[float64] {
    if x < 0 {
        return f.Err[float64](errors.New("square root of a negative number"))
    }
    return f.Ok(math.Sqrt(x))
}
zero := func() float64 { return 0 }
x := sqrt(-1).Default(zero)
// x = 0
y := sqrt(1).Default(zero)
// y = 1
```

* Handle

`Handle` allows you to handle an error using a specific function.

```go
sqrt := func(x float64) f.Result[float64] {
    if x < 0 {
        return f.Err[float64](errors.New("square root of a negative number"))
    }
    return f.Ok(math.Sqrt(x))
}
handler := func() {
    log.Println("error occurred")
    panic(1)
}
sqrt(-1).Handle(handler)
// log: error occurred
// panic 1
```

## function -- a set of operations helping you to play with functions

### control

* Ternary

```go
x := Ternary(1 > 0, "1 is bigger", "0 is bigger")
// x := 1>0 ? "1 is bigger" : "0 is bigger"
fmt.Println(x)
// 1 is bigger
```

* If, ElseIf, Else

```go
x := 1
y := If(x == 0, "x is 0").ElseIf(x == 1, "x is 1").
    Else("x is not 0 or 1")
fmt.Println(y)
// x is 1
```

* For

```go
x, n := 0, 100
addIndex := func (i int) { x += i }
For(n, addIndex)
// 0 + 1 + ... + 99
```

* ForReverse

```go
x, n := 0, 100
addIndex := func (i int) { x += i }
ForReverse(n, addIndex)
// 99 + 98 + ... + 0
```

* While

```go
x, n := 2, 1
cond := func () bool { n++; return n <= 10 }
fn := func () { x += x }
While(cond, fn)
fmt.Println("x =", x)
// x = 1024
```

### convert

* Dot

```go
add1 := func (x int) int { return x + 1 }
add2 := func (x int) int { return x + x }
fn := Dot(add1, add2)
// fn(x) = add1(add2(x))
fmt.Println(fn(3))
// 7
```

* Flip

```go
Div := func (x, y int) Maybe[float64] {
    if y == 0 {
        return Nothing[float64]()
    } else {
        return float64(x)/float64(y)
    }
}
newDiv := Flip(Div)
fmt.Println(newDiv(3, 1))
// Just 0.333333333333333
```

### curry

* Curry2

`Curryn` means the original function has n parameters.

```go
add := func (x, y int) int { return x + y }
addCurry := Curry2(add)
add1 := addCurry(1)
fmt.Println(add1(2))
// 3
fmt.Println(addCurry(2)(3))
// 5
```

`Curry2` to `Curry5` are provided.

## stream -- a convenient tool for calculations on maps and slices

## Q&A

> What's the difference between `Either` and `Result`?

`Either` helps you to finish your calculation even if there is an error, and you can use a string or something else as a substitution of the error, when you don't need to handle it at very first time.

`Result` should be the return value of an operation which may cause an emergency error.
As a result, `Result` provides you many ways to handle a possible error, while `Either` does not do.
