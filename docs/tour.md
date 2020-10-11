# Lithium Programming Language Tour

## Welcome

This tour aims to cover the basics of the Lithium Programming Language.


## Start

    # This imports a library so that we can output to the console and get random numbers
    import console "github.com/neelsg/lib/console"
    import math "github.com/neelsg/lib/math"
    
    # This is the main function
    fn main
        console.log("My favorite number is", math.rand[int](10))

This code will print the text "My favorite number is" and a random number from 0 to 10.

Anything after a `#` on a line is a comment and is ignored by the compiler.

The `fn` keyword is used to define functions (which are blocks of code that can be executed. When a program is executed, the first thing that will run is the "main" function.


## Exporting components

    fn multiply(x, y int) int
        # This function is only available within the current package
        return x * y
    
    export fn sqr(x int) int
        # This function can be used by other packages
        return multiply(x, x)

Use the `export` keyword to make your components available to other packages


## Importing packages

    import console "github.com/neelsg/lib/console"
    import myMath "./lib/myMath"
    
    fn main
        console.log(myMath.sqr(3))

You can import your own packages or third-party libraries using the `import` keyword.


## Functions

    import console "github.com/neelsg/lib/console"
    
    fn add(x, y int) int
        return x + y
    
    fn main
        console.log(add(42, 13))

A function can take zero or more arguments. In this example, `add` takes two arguments of type `int`. Note that the type comes after the variable name. If two or more consecutive arguments are of the same type, you only need to specify the type for the last argument.

A function that takes no arguments can optionally be declared without brackets:


## Multiple results

    import console "github.com/neelsg/lib/console"

    fn split(sum int) (int, int)
        let x int = sum * 4 / 9
        let y int = sum - x
        return x, y

    fn swap(x, y string) (string, string)
        return y, x

    main func
        console.log(split(17))
        
        let a, b = swap("hello", "world")
        console.log(a, b)

A function can return any number of results. If they do, these must be wrapped in brackets `( )`. Even if a function returns only one result, the result type can optionally be wrapped in brackets.


## Variables

    import console "github.com/neelsg/lib/console"

    var i, j, k bool
    
    export var l bool
    
    fn main
        m int
        console.log(i, j, k, l, m)

Variables are declared by specifying the name of the variable followed by the type.

    import console "github.com/neelsg/lib/console"
    import math "github.com/neelsg/lib/math"

    var i, j, p = 1, 2, math.pi
    
    fn main
        var k, l, m, n = true, false, "no!", "yes!"
        console.log(i, j, p, k, l, m, n)

A variable declaration can include initializers. If an initializer is omitted, the variable will be assigned the default value of the type.


## Basic types

    import console "github.com/neelsg/lib/console"
    import math "github.com/neelsg/lib/math"
    
    var toBe bool = false
    var maxInt int.u64 = math.power[int.u64](2, 64) - 1
    var z math.complex.p128 = math.sqrt[math.complex.p128](-5 + 12i)
    
    fn main
        console.log("Value: \(toBe)")
        console.log("Value: \(maxInt)")
        console.log("Value: \(z)")

Some of the lithium basic types are:

    bool int string float

Each of these types are defined in a built-in library package.

There are also more specific data types such as:

    int.s8 int.s16 int.s32 int.s64
    int.u int.u8 int.u16 int.u32 int.u64
    
    float.p32 float.p64


## Default values

    import console "github.com/neelsg/lib/console"

    fn main
        var i int
        var f float
        var b bool
        var s string
        console.log("\(i) , \(f) , \(b) , \(s)")

Variables declared without an explicit initial value are given their default value.

The default values for built in types are:

    0 for numeric types
    false for the boolean type
    "" for strings.

Custom types can have their own defined default values.


## Type conversions

    import console "github.com/neelsg/lib/console"
    import math "github.com/neelsg/lib/math"

    fn main
        let x, y = 3, 4
        let f = math.sqrt((x * x + y * y).float())
        let z int.u = f.int().unsigned()
        console.log(x, y, z)

The built in types have properties to convert to different types. Custom defined types should also include similar functionality.

Some numeric conversions:

    i int = 42
    f float.p64 = i.float64()
    u int.u = f.int().unsigned()

Unlike in C, assignment in Lithium between items of different type requires an explicit conversion.


## Constants

    import console "github.com/neelsg/lib/console"
    
    let pi = 3.14
    
    fn main
        world const = "Mars"
        console.log("Hello", world)
        console.log("Happy", pi, "Day")
        
        let truth = true
        console.log("Humans rock?", truth)

Constants are declared like variables, but declared with `let` keyword instead of `var`. Constants can be string, boolean, or numeric values.


## For

    import console "github.com/neelsg/lib/console"
    
    fn main
        var sum int = 0
        for i in [1:10]
            sum += i
        console.log(sum)

Lithium has only one looping construct, the `for` loop.

    import console "github.com/neelsg/lib/console"
    
    fn main
        var sum int = 1
        for sum < 1000
            sum += sum
        console.log(sum)

The loop can also include a conditional expression. The loop will exit as soon as the condition is false.

    import console "github.com/neelsg/lib/console"
    
    fn main
        var sum int
        for i in [1:10] && sum < 20
            sum += i
        console.log(sum)

Iterators and conditional expressions can be combined.

    fn main
        for
            # Whatever code goes here will run in an infinite loop

If you omit the loop condition it loops forever.


## If

    import console "github.com/neelsg/lib/console"
    
    fn sqrt(x float) string
        if x < 0
            return (-x ^ 0.5).string() + "i"
        return (x ^ 0.5).string()
    
    fn main
        console.log(sqrt(2), sqrt(-4))

Lithium's `if` statements are similar to its `for` loops

    import console "github.com/neelsg/lib/console"
    
    fn pow(x, n, lim float) float
        if (let v float = x ^ n) && v < lim
            return v
        return lim
    
    fn main
        console.log(
            pow(3, 2, 10),
            pow(3, 3, 20),
        )

Like `for`, the `if` statement can start with a short statement to execute before the condition using `&&`. Variables declared by the statement are only in scope until the end of the `if`.

    import console "github.com/neelsg/lib/console"
    
    fn pow(x, n, lim float) float
        if (let y = x ^ n) && y < lim
            return y
        else
            console.log("\(y) >= \(lim)")
        return lim
    
    fn main
        console.log(pow(3, 2, 10), pow(3, 3, 20), )

Variables declared inside an `if` init statement are also available inside any of the else blocks.


## Function values

    import console "github.com/neelsg/lib/console"
    
    fn compute(f fn(float, float) float) float
        return f(3, 4)
    
    fn main
        let hypot = fn(x, y float) float
            return (x * x + y * y) ^ 0.5
        let ave = fn(x, y float) float
            return (x + y) / 2
        
        console.log(hypot(5, 12))
        
        console.log(compute(hypot))
        console.log(compute(ave))

Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.


## Defer

    import console "github.com/neelsg/lib/console"
    
    fn main
        defer console.log("world!")
        
        console.log("Hello")

The `defer` statement defers the execution of a function until the surrounding function returns.

The `defer` statement's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

    import console "github.com/neelsg/lib/console"
    
    fn main
        console.log("counting")
        
        for i in [1:10]
            defer console.log(i)
        
        console.log("done")

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.


## Function closures

    import console "github.com/neelsg/lib/console"
    
    fn adder() (fn(int) int)
        var sum = 0
        return fn(x int) int:
            sum += x
            return sum
    
    fn main
        var pos, neg = adder(), adder()
        for i in [1:10]
            console.log(
                pos(i),
                neg(-2 * i),
            )

Lithium functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables. In this sense the function is "bound" to the variables.

For example, the `adder` function returns a closure. Each closure is bound to its own `sum` variable.


## Defining new types

    import console "github.com/neelsg/lib/console"
    
    type vertex struct(
        x int,
        y int
    )
    
    fn main
        let v vertex
        v.x, v.y = 1, 2
        console.log(v)

Lithium allows simple object-oriented programming. A new type can be defined from any existing type, inheriting the properties and methods of the existing type.


## Methods

    import console "github.com/neelsg/lib/console"
    
    type vertex struct(x, y float)
    
    fn vertex(v).dist() float
        return (v.x ^ 2 + v.y ^ 2) ^ 0.5
    
    fn main
        let v vertex
        v.x, v.y = 1, 2
        console.log(v.dist())

Methods are functions defined with the type defined at the start.


## The initializer function

    import console "github.com/neelsg/lib/console"
    
    type vertex struct(x, y float)
    
    fn vertex(v).=(xInit, yInit float)
        v.x, v.y = xInit, yInit
        return
    
    fn vertex(v).dist() float
        return (v.x ^ 2 + v.y ^ 2) ^ 0.5
    
    fn main
        let v vertex = (1, 2)
        console.log(v.dist())

You can create a function with the label `=` for a type which will act as the initializer for the type. The initializer is activated either by calling the type as a function in an expression such as `vertex(1, 2)` or by assigning to a variable of the specified type as in `var v vertex = (1, 2)`.


## Interfaces

    import console "github.com/neelsg/lib/console"
    
    interface hasX
        x float
    
    type vertex struct(x, y float)
    
    fn vertex(v).=(xInit, yInit float)
        v.x, v.y = xInit, yInit
        return
    
    fn printX(o hasX)
        console.log(o.x)
    
    fn main
        let v vertex = (3, 4)
        printX(v) # This will print 3

Interfaces specify which properties and methods a type must have to satisfy it. You can use an interface as a parameter in a function and ensure that any object that is passed to that function must satisfy that interface. If you try to pass an object that does not satisfy the interface or try to use the parameter in the function in a way that is not allowed by the interface, you will get a compile error.

Note that a type does not need to specify that it implements the interface, all it needs to do is to actually conform to it. This allows for quite a bit of flexibility in using interfaces.


## Stringers

    import console "github.com/neelsg/lib/console"
    
    type person struct(
        name string,
        age int
    )
    
    fn person(p).=(nameInit string, ageInit int)
        p.name, p.age = nameInit, ageInit
        return
    
    fn person(p).string() string
        return "\(p.name) (\(p.age) years)"
    
    fn main
        let a = person("Arthur Dent", 42)
        let z = person("Zaphod Beeblebrox", 9001)
        console.log(a, z)

One of the most ubiquitous interfaces is `string.stringer`.

    export interface stringer
        string fn() string

Many packages look for this interface to print values.


## Errors

    import console "github.com/neelsg/lib/console"
    import time "github.com/neelsg/lib/time"
    
    type myError struct(
        at time,
        msg string,
        ok bool,
        class string,
    )
    
    fn myError(e).=(when time, what string)
        e.at, e.msg = when, what
        e.ok, e.class = false, "MyError"
        return
    
    fn myError(e).string() string
        return "at \(e.at), \(e.msg)"
    
    fn run() error:
        return myError(time.now(), "it didn't work")
    
    fn main
        if let err = run() && !err.ok
            console.log(err)

Lithium programs express error state with `error` values.

The `error` type is a built-in interface:

    export interface self
        string.stringer
        ok bool

The name `self` as used for this interface is special in that it indicates that this is the default item in this package, so you can access it by simply using the package name which in this case is `error`.

Functions often return an `error` value, and calling code should handle errors by testing whether the `error.ok` is false.

    import console "github.com/neelsg/lib/console"
    
    fn main
        let i int, err error
        if i, err = "42".int() && !err.ok
            console.log("Couldn't convert number: \(err)")
        console.log("Converted integer:", i)
        return


## Parametric polymorphism

    import console "github.com/neelsg/lib/console"
    
    interface plusable[p]
        + fn(p) p
    
    fn double[p plusable](num p) p:
        return num + num
    
    fn main
        var myInt = 3
        myInt = double[int](myInt)
        console.log(myInt) # This will print 6
        
        var myFloat = 3.1
        myFloat = double[float](myFloat)
        console.log(myFloat) # This will print 6.2

This is often called generics in other languages. What is happening here is that we can pass any type of variable to `double()` as long as it implements the `plusable` interface.

The advantage over just using interfaces is that we can make sure that the variables retain their original type, so in the first case, `double()` actually returns an `int` and in the second, `double()` actually returns a `float`. If we used only interfaces, the return value would only allow properties and methods of the interface, so we would not be able to print it as `plusable` does not specify a `string` method, but `float` and `int` does.


## Default types for polymorphism

    import console "github.com/neelsg/lib/console"
    
    interfce plusable[n]
        + fn(n) n
    
    fn double[p plusable int](num p) p
        return num + num
    
    fn main
        var myInt int = 3
        myInt = double(myInt)
        console.log(myInt) # This will print 6
        
        var myFloat float = 3.1
        myFloat = double[float](myFloat)
        console.log(myFloat) # This will print 6.2

A default type can be set for any polymorphic function. If the type is then omitted when calling the function, this default type is assumed.


## Parametric polymorphism for types

    import console "github.com/neelsg/lib/console"
    
    type vertex struct[n string.stringer](x, y n)
    
    fn vertex(v, n).=(xInit, yInit n)
        v.x, v.y = xInit, yInit
        return
    
    fn main
        let v vertex[int] = (1, 2)
        console.log(v.x)

Parametric polymorphism can also be used with type definitions to create more generic types. In the above example, `x` and `y` are of type `int`.


## Tuples

    import console "github.com/neelsg/lib/console"
    
    fn main
        let triplets (string, string, string) = "Anna", "Betty", "Carmen"
        console.log(triplets)

Tuples are ordered pairs of variables. You have actually already seen them used for multiple return values for functions as well as with assignments on multiple variables. Tuples defined as a set of types wrapped in brackets.

    import console "github.com/neelsg/lib/console"
    
    fn main
        let triplets (string, string, string) = "Anna", "Betty", "Carmen"
        console.log(triplets.0) # This prints Anna
        console.log(triplets.1) # This prints Betty
        console.log(triplets.2) # This prints Carmen

The individual values in a tuple can be accessed using the index number starting from zero. The index number must be a literal number, it cannot be another variable or calculated in some way.


## Arrays

    import console "github.com/neelsg/lib/console"
    
    fn main
        var l array[string]
        l(0) = "Hello"
        l(1) = "World"
        console.log(l(0), l(1))
        console.log(l)
        
        let primes = [int](2, 3, 5, 7, 11, 13)
        console.log(primes)

`array[t]` is a built-in type that stores a sequence of values of type `t`.

The statement `l array[string]` declares a new array `l` of strings. The keys in an array are always of type `int` and always start at 0.

    import console "github.com/neelsg/lib/console"
    
    fn main
        let abcd = [string]("a", "b", "c", "d")
        let bc = abcd.slice(1, 3)
        console.log(bc)

A slice is formed by specifying two indexes, a low and high bound, separated by a comma: `arr.slice(low, high int)`.

This selects a half-open range which includes the first element, but excludes the last one.

    import console "github.com/neelsg/lib/console"
    
    fn main
        var p array[int] = [int](2, 3, 5, 7, 11, 13)
        
        p = p.slice(1, 4)
        console.log(p)
        
        p = p.slice(, 2)
        console.log(p)
        
        p = p.slice(1, )
        console.log(p)

When slicing, you may omit the high or low bounds to use their defaults instead.
The default is zero for the low bound and the length of the list for the high bound.

    import console "github.com/neelsg/lib/console"

    fn main
        let p = [int](2, 3, 5, 7, 11, 13)
        console.log(p.length())

Arrays have a length method that shows how many items are currently in the array.

    import console "github.com/neelsg/lib/console"
    
    fn main
        # Create a tic-tac-toe board.
        var board array[array[string]] = [][string](
            ("_", "_", "_"),
            ("_", "_", "_"),
            ("_", "_", "_"),
        )
        
        # The players take turns.
        board(0)(0) = "X"
        board(2)(2) = "O"
        board(1)(2) = "X"
        board(1)(0) = "O"
        board(0)(2) = "X"
        
        for i in board
            console.log(board(i).join(" "))

Arrays can hold any type, including other arrays.

    import console "github.com/neelsg/lib/console"

    fn main
        var a array[int]
        console.log(a)
        
        a.push(3)
        console.log(a)
        
        a.push(4)
        console.log(a)
        
        # We can add more than one element at a time.
        a.push(5, 6, 7)
        console.log(a)

Arrays have a `push` method to append items to it.


## Iterating over an array

    import console "github.com/neelsg/lib/console"
    
    let pow = [int](1, 2, 4, 8, 16, 32, 64, 128)
    
    fn main
        for i, v in pow
            console.log("2 ** \(i) = \(v)")

Arrays are iterable, which means that they can be used in a for loop with the `in` keyword.

    import console "github.com/neelsg/lib/console"
    
    fn main
        let pow = [int](1, 2, 4, 8, 16, 32, 64, 128)
        
        for _, value in pow
            console.log(value)

You can skip the key or value by assigning to `_`.

    for key, _ in collection
    for _, value in collection

If only the value is required, this can be further shortened to:

    for value in collection


## Maps

    import console "github.com/neelsg/lib/console"
    
    fn main
        var m map[int=string]
        m(1) = "Hello"
        m(2) = "World"
        console.log(m(1), m(2))
        console.log(m)
        
        var primes = [string=int](
            "first"  = 2,
            "second" = 3,
            "third"  = 5,
            "fourth" = 7,
            "fifth"  = 11,
            "sixth"  = 13,
        )
        console.log(primes)

`map[k=v]` is a data type that stores a map with keys of type `k` that maps to values of type `v`.

The statement `var m map[int=string]` declares a variable `m` as a map with `int` keys and `string` values.

    import console "github.com/neelsg/lib/console"
    
    type vertex struct(lat, long float)
    
    let m  = [string=vertex](
        "Bell Labs" = (40.68433, -74.39967),
        "Google"    = (37.42202, -122.08408),
    )
    
    fn main
        console.log(m)

The map initializer accepts an arbitrary number of tuples that consist of the key=value pairs that need to be added.

    import console "github.com/neelsg/lib/console"
    
    fn main
        var m map[string=int]
        
        m("Answer") = 42
        console.log("The value: ", m("Answer"))
        
        m("Answer") = 48
        console.log("The value: ", m("Answer"))
        
        m.delete("Answer")
        console.log("The value: ", m("Answer"))
        
        let ok = m.exists("Answer")
        console.log("Present? ", ok)
        
        let v, ok = m.get("Answer")
        console.log("The value: ", v, " (Present? \(ok)")

Insert or update an element in map `m`:

    map(key) = value

Retrieve an element:

    val = map(key)

If `key` is not in the map, then `val` is the zero value for the map's value
type.

Delete an element:

    map.delete(key)

Test that a key is present:

    ok = map.exists(key)

If `key` is in the map, `ok` is `true`. If not, `ok` is `false`.

You can use `val, ok = map.get(key)` as a shorthand for `val, ok = map(key), map.exists(key)`.


## Variadic functions

    import console "github.com/neelsg/lib/console"
    
    fn sum(...nums int) int
        var s int
        for n in nums
            s += n
        return s
    
    fn main
        console.log(sum(1, 2, 3, 4, 5)) # This will print 15

Variadic functions are functions that can take an arbitrary number of parameters instead of a fixed set. The parameters are passed into the initializer method of a list.

    import console "github.com/neelsg/lib/console"
    
    interface number[n]
        + fn(n) n
        ^ fn(n) n
    
    fn sumPower[num number](...pairs (num, num)) num
        var sum num
        for (base, exp) in pairs
            sum += base ^ exp
        return sum
    
    fn main
        console.log(sumPower[int]((2, 1), (2, 2), (2, 3), (2, 4))) # Output 30

Variadic functions can take any type of parameter that can be as a value in a list including tuples.

    import console "github.com/neelsg/lib/console"
    
    fn sum(...nums int) int
        var s int
        for n in nums
            s += n
        return s
    
    fn main
        let numbers = [int](1, 2, 3, 4, 5)
        console.log(sum(numbers...))

It is possible to use the spread operator (`...`) to pass the values in a collection as distinct parameters in a variadic function.


## Coroutines

    import
        console "github.com/neelsg/lib/console"
        time    "github.com/neelsg/lib/time"
    
    fn say(s string)
        for i in [0:4]
            time.sleep(100 * time.milliSecond)
            console.log(s)
    
    fn main
        co say("world")
        say("hello")

A coroutine is a lightweight thread managed by the Lithium runtime.

`co f(x, y, z)` starts a new coroutine running `f(x, y, z)`.

The evaluation of the passed parameters happen in the current coroutine, but the execution of the function happens in the new coroutine.

Coroutines run in the same address space, so access to shared memory must be synchronized. The `sync` package provides useful primitives, although you won't need them much in Lithium as there are other ways to handle concurrency as you will see shortly.

    import
        console "github.com/neelsg/lib/console"
        time    "github.com/neelsg/lib/time"
    
    fn say(s string)
        for i in [0:4]
            time.sleep(100 * time.milliSecond)
            console.log(s)
    
    fn main
        co
            for i in [0:4)
                time.sleep(100 * time.milliSecond)
                console.log("world")
        say("hello")

It is also possible to use `co` with a block of code to be executed instead of calling a function.


## Channels

    import console "github.com/neelsg/lib/console"
    
    fn sum(s array[int], c chan[int])
        var sum = 0
        for v in s
            sum += v
        c.push(sum) # Send sum to c
    
    fn main
        let s = [int](7, 2, 8, -9, 4, 0)
        let c = chan[int]
        co sum(s.slice(, s.length / 2), c)
        co sum(s.slice(s.length / 2, ), c)
        let x, y = c.pop(), c.pop() # receive from c
        console.log(x, y, x + y)

Channels are a typed conduit through which you can send values with `push()` and receive values with `pop()`

    ch.push(v)     # Send v to channel ch.
    v = ch.pop()   # Receive from ch, and assign value to v.

By default, sends and receives block until the other side is ready. This allows coroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in an array, distributing the work between two coroutines. Once both coroutines have completed their computation, it calculates the final result.

    import console "github.com/neelsg/lib/console"
    
    fn main
        let ch = chan[int](2)
        ch.push(1)
        ch.push(2)
        console.log(ch.pop())
        console.log(ch.pop())

Channels can be buffered. Provide the buffer length to the channel initializer to make a buffered channel:

    ch = chan[int](bufferLength)

Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.


## Closing a channel

    import console "github.com/neelsg/lib/console"
    
    fn fibonacci(n int, c chan[int])
        var x, y = 0, 1
        for i in [1:n]
            c.push(x)
            x, y = y, x + y
            c.close()
    
    fn main
        let c = chan[int](10)
        co fibonacci(c.capacity, c)
        for i in c
            console.log(i)

A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by checking the `ok` property.

    if ch.ok
        v int = ch.pop()

`ok` is `false` if there are no more values to receive and the channel is closed.

The loop `for i in c` receives values from the channel repeatedly until it is closed.

Note: Only the sender should close a channel, never the receiver.

Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a loop.

    import console "github.com/neelsg/lib/console"
    
    fn fibonacci(vals, quit chan[int])
        var x, y = 0, 1
        for
            if vals.capacity > 0 # This checks if you can push new values to the vals channel
                vals.push(x)
                x, y = y, x+y
            else if quit.length > 0: # This checks if any values have been pushed to the quit channel
                console.log("quit")
                return
    
    fn main
        let vals, quit chan[int]
        co
            for i in [0:9]
                console.log(vals.pop())
            quit.push(0)
        fibonacci(vals, quit)

The `for if` pattern lets a coroutine wait on multiple communication operations. This will wait until one of its conditions are true, then it executes that block.

    import
        console "github.com/neelsg/lib/console"
        time    "github.com/neelsg/lib/time"
    
    fn main
        let tick = time.tick(100 * time.millisecond)
        let boom = time.after(500 * time.millisecond)
        for
            if tick.length > 0
                tick.pop()
                console.log("tick.")
            else if boom.length > 0
                boom.pop()
                console.log("BOOM!")
                return
            else
                console.log(".")
                time.sleep(25 * time.millisecond)

An `else` block can be used to run if no other condition is ready.


## Mutual exclusion

    import
        sync    "github.com/neelsg/lib/sync"
        console "github.com/neelsg/lib/console"
        time    "github.com/neelsg/lib/time"
    
    # SafeCounter is safe to use concurrently.
    type safeCounter struct(
        counts map[string=int],
        mux sync.mutex,
    )
        
    # inc increments the counter for the given key.
    fn safeCounter(c).inc(key string)
        c.mux.lock()
        # Lock so only one coroutine at a time can access the map.
        c.counts(key)++
        c.mux.unlock()
        return
        
    # value returns the current value of the counter for the given key.
    fn safeCounter(c).value(key string) int
        c.mux.lock()
        # Lock so only one coroutine at a time can access the map.
        defer c.mux.unlock()
        return c.counts(key)
    
    fn main
        let c = safeCounter()
        for i in [1:1000]
            co c.inc("somekey")
        
        time.sleep(time.second)
        console.log(c.value("somekey"))

We've seen how channels are great for communication among coroutines.

But what if we don't need communication? What if we just want to make sure only one coroutine can access a variable at a time to avoid conflicts?

This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex.

The mutual exclusion type `sync.mutex` has two methods:

    lock()
    unlock()

We can define a block of code to be executed in mutual exclusion by surrounding it with a call to `lock` and `unlock` as shown on the `inc` method.

We can also use `defer` to ensure the mutex will be unlocked as in the `value` method.
