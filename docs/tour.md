# Lithium Programming Language Tour

## Welcome

This tour aims to cover the most important parts of the Lithium Programming
Language. This tour borrows heavily on the
[Go Programming Language tour](https://tour.golang.org) as I have based quite a
significant part of this programming language on Go.


## Start

    Li 0
    
    import:
        console
    
    main fn():
        console.log("My favorite number is", int.rand(10))

There is quite a bit happening here. This code will print the text
"My favorite number is" and a random number from 0 to 10.


## File format and doctype

    Li 0.0.1

Lithium source code is stored in plain text files that have the `.li` file
extension. All files that are located in the same folder are considered to be
part of the same package. Each file needs to start with the chars `Li ` followed
by the language version used in the file. Note that the "L" is uppercase and the
"i" is lowercase. This will allow newer compilers to compile files written using
older specs of the language and for older packages to be incorporated into newer
projects.

The version can be written in full such as `0.0.1` or shortened such as `0.0` or
`0`. If any part of the version is omitted, the compiler will assume the most
recent sub-version should be used.


## Importing from other packages

    Li 0.0

    import:
        console
        math

    main fn():
        console.log("Now you have % problems.".in(math.sqrt(7.0)))

Here we import `console` and `math` from the standard library so that they can
be used in this file. We could also write:

    import console
    import math


## The main function

    main fn():

The `fn` keyword is used to define functions (which are blocks of code that can
be executed. When a program is executed, the first thing that will run is the
"main" function.


## Exporting components

    Li 0
    
    private sqr2 fn(x int) int:
        // This function is only available within the current package
        return x * x
    
    sqr fn(x int) int:
        // This function can be used by other packages
        return sqr2(x)

In Lithium, anything that starts with `private` can only be used in the current
package, but note that this also includes any other files that are part of the
current package. Anything without the `private` keyword is made available to any
other packages that import this one.


## Importing custom packages

    Li 0
    
    import:
        console
        myMath "./myMathFolder"
    
    main fn():
        console.log(myMath.sqr(3.0))

You can import your own packages just as easily as packages from the standard
library by providing the file path to the folder where the source code is
located.


## Functions

    Li 0
    
    import:
        console
    
    private add fn(x int, y int) int:
        return x + y
    
    main fn():
        console.log(add(42, 13))

A function can take zero or more arguments. In this example, `add` takes two
arguments of type `int`. Note that the type comes after the variable name.


## Multiple results

    Li 0
    
    import console
    
    private:
        split fn(sum int) (int, int):
            x, y int = sum * 4 / 9, sum - x
            return x, y
    
    main fn(): console.log(split(17))

A function can return any number of results.

    Li 0
    
    import:
        console
    
    private swap fn(x, y string) (string, string): return y, x
    
    main fn():
        a, b string = swap("hello", "world")
        console.log(a, b)

When two or more parameters are of the same type, you can omit the
type from all but the last parameter.


## Named parameters

    Li 0
    
    import:
        console
    
    private swap fn(x = "hello", y string) (string, string): return y, x
    
    main fn():
        a, b string = swap(y = "world")
        console.log(a, b)

When calling a function, the parameters may be named. A function definition can
specify the default value of any parameters. If these parameters are omitted,
they will be set to the default.


## Variables

    Li 0
    
    import console
    
    private:
        c, python, java bool
    
    go bool
    
    main fn():
        i int
        console.log(i, c, python, java, go)

Variables are declared by specifying the name of the variable followed by the
type.


## Initializers

    Li 0
    
    import console, math
    
    private i, j int, p float = 1, 2, math.pi
    
    main fn():
        c, python bool, java, go string = true, false, "no!", "yes!"
        console.log(i, j, p, c, python, java, go)

A variable declaration can include initializers. If an initializer is omitted,
the variable will be assigned the default value of the type.


## Basic types

    Li 0
    
    import console, math
    
    private:
        toBe bool = false
        maxInt int.u64 = 2 ** 64 - 1
        z complex.p128 = math.sqrt[complex.p128](-5. + 12i)
    
    main fn():
        console.log("Value: %".in(toBe))
        console.log("Value: %".in(maxInt))
        console.log("Value: %".in(z))

Lithium basic types are:

    bool byte complex float int string

Each of these types are defined in a standard library package that is imported
implicitly, so you do not need to list them in an import statement.

These packages also provide more fine-grained type definitions such as the
following:

    int.s8 int.s16 int.s32 int.s64
    int.u int.u8 int.u16 int.u32 int.u64
    
    float.p32 float.p64
    
    complex.p64 complex.p128


## Default values

    Li 0
    
    import console
    
    main fn():
        i int
        f float
        b bool
        s string
        console.log("% , % , % , %".in(i, f, b, s))

Variables declared without an explicit initial value are given their default
value.

The default values for built in types are:

    0 for numeric types
    false for the boolean type
    "" for strings.

Custom types can have their own defined default values.


## Type conversions

    Li 0
    
    import:
        console
        math
    
    main fn():
        x, y int = 3, 4
        f float = math.sqrt((x * x + y * y).float)
        z int.u = f.intU
        console.log(x, y, z)

The built in types have properties to convert to different types. Custom defined
types should also include similar functionality.

Some numeric conversions:

    i int = 42
    f float.p64 = i.float64
    u int.u = f.intU

Unlike in C, assignment in Lithium between items of different type requires an
explicit conversion.


## Constants

    Li 0
    
    import console
    
    private PI const = 3.14
    
    main fn = fn():
        WORLD const = "Mars"
        console.log("Hello", WORLD)
        console.log("Happy", PI, "Day")
        
        TRUTH const = true
        console.log("Humans rock?", TRUTH)

Constants are declared like variables, but with the `const` keyword as type.
Constants can be string, boolean, or numeric values. Constants cannot be
declared with an expression that requires runtime evaluation such as calling a
function.


## For

    Li 0
    
    import console
    
    main fn():
        sum int = 0
        for i int = 0; i < 10; i++:
            sum += i
        console.log(sum)

Lithium has only one looping construct, the `for` loop.

The basic for loop has three components separated by semicolons:

    The init statement: executed before the first iteration
    The condition expression: evaluated before every iteration
    The post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables
declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to `false`.


## For continued

    Li 0
    
    import console
    
    main fn():
        sum int = 1
        for ; sum < 1000; :
            sum += sum
        console.log(sum)

The init and post statements are optional. At that point you can drop the
semicolons:

    Li 0
    
    import console
    
    main fn():
        sum int = 1
        for sum < 1000:
            sum += sum
        console.log(sum)


## Forever

    Li 0
    
    main fn():
        for:
            // Whatever code goes here will run in an infinite loop

If you omit the loop condition it loops forever.


## If

    Li 0
    
    import:
        console
        math
    
    private sqrt fn(x float) string:
        if x < 0:
            return sqrt(-x).string + "i"
        return math.sqrt(x).string
    
    main fn():
        console.log(sqrt(2.), sqrt(-4.))

Lithium's `if` statements are like its `for` loops


## If with an init statement

    Li 0
    
    import:
        console
    
    private pow fn(x, n, lim float) float:
        if v float = x ** n; v < lim:
            return v
        return lim
    
    main fn():
        console.log(
            pow(3, 2, 10),
            pow(3, 3, 20),
            )

Like `for`, the `if` statement can start with a short statement to execute
before the condition. Variables declared by the statement are only in scope
until the end of the `if`.


## If and else

    Li 0
    
    import:
        console
        math
    
    private pow fn(x, n, lim float) float:
        if x**= n; x < lim:
            return x
        else:
            console.log("% >= %".in(x, lim))
        return lim
    
    main fn():
        console.log(pow(3., 2., 10.), pow(3., 3., 20.), )

Variables declared inside an `if` init statement are also available inside any
of the else blocks.


## Function values

    Li 0
    
    import:
        console
        math
    
    private compute fn(func fn(float, float) float) float:
        return func(3, 4)
    
    main fn():
        hypot fn(x, y float) float:
            return math.sqrt(x*x + y*y)
        ave fn(x, y float) float:
            return (x + y) / 2
        
        console.log(hypot(5., 12.))
        
        console.log(compute(hypot))
        console.log(compute(ave))

Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.


## Defer

    Li 0
    
    import console
    
    main fn():
        defer(console.log, "world!")
        
        console.log("Hello")

The `defer` function is a special built-in function that defers the execution of
a function until the surrounding function returns. The first parameter is the
function to be called and any further parameters are the parameters that will be
passed to this function on execution.

The `defer` statement's arguments are evaluated immediately, but the function
call is not executed until the surrounding function returns.


## Stacking defers

    Li 0
    
    import console
    
    main fn():
        console.log("counting")
        
        for i int = 0; i < 10; i++:
            defer(console.log, i)
        
        console.log("done")

Deferred function calls are pushed onto a stack. When a function returns, its
deferred calls are executed in last-in-first-out order.


## Function closures

    Li 0
    
    import console
    
    private adder fn() fn(int) int:
        sum int = 0
        return fn(x int) int:
            sum += x
            return sum
    
    main fn():
        pos int, neg int = adder(), adder()
        for i int; i < 10; i++:
            console.log(
                pos(i),
                neg(-2 * i),
            )

Lithium functions may be closures. A closure is a function value that references
variables from outside its body. The function may access and assign to the
referenced variables. In this sense the function is "bound" to the variables.

For example, the `adder` function returns a closure. Each closure is bound to
its own `sum` variable.


## Defining new types

    Li 0
    
    import console
    
    private vertex type:
        x int
        y int
    
    main fn():
        console.log(vertex(1, 2))

Lithium allows simple object-oriented programming. A type is a collection of
properties and methods and is often called a class in other languages.


## Type properties

    Li 0
    
    import console
    
    private vertex type:
        x int
        y int
    
    main fn():
        v vertex = vertex(1, 2)
        v.x = 4
        console.log(v.x)

Type properties are accessed using a dot.


## Shorthand for initializing types

    private vertex type:
        x int
        y int

Given the above type, the following statements are equivalent:

    v vertex = vertex(1, 2)
    v vertex(1, 2)
    v vertex(x = 1, y = 2)


## Methods

    Li 0
    
    import
        console
        math
    
    private vertex type:
        x float
        y float
        dist fn(self vertex) float:
            return math.sqrt(self.x ** 2 + self.y ** 2)
    
    main fn():
        v vertex(1, 2)
        console.log(v.dist())

Methods are functions defined within a type.

Note the first parameter refers to the object which the method is associated
with. When calling the method, this parameter is omitted.


## Default values and the initializer function

    Li 0
    
    import
        console
        math
    
    private vertex type:
        x float = 1
        y float = 1
        init fn(x, y float) vertex:
            self vertex
            self.x, self.y = x, y
            return self
        dist fn(self vertex) float:
            return math.sqrt(self.x ** 2 + self.y ** 2)
    
    main fn():
        v vertex(1, 2)
        console.log(v.dist())

Just like normal variables, properties in a type can also be set to an initial
value which will then be the default for any objects of that type.

You can create a function with the label `init` in a type which returns an
object of that type. This function then becomes the initializer and will be used
when creating an object of that type such as with `v vertex(1, 2)`.


## Private values

    Li 0
    
    import
        console
        math
    
    private vertex type:
        
        private: // X and Y can now only be used by methods within the object
            x float = 1
            y float = 1
        
        init fn(x, y float) vertex:
            self vertex
            self.x, self.y = x, y
            return self
        
        dist fn(self vertex) float:
            return math.sqrt(self.x ** 2 + self.y ** 2)
    
    main fn():
        v vertex(1, 2)
        console.log(v.dist())

You can use the `private` keyword to make properties or even methods only
available to the methods in the object itself. If you try to use `v.x` in the
main function in this case, you will get a compile error.


## Dynamic properties

    Li 0
    
    import
        console
        math
    
    private vertex type:
        init fn(x, y float) vertex:
            self vertex
            self.x, self.y = x, y // This will call the respective set functions
            return self
        
        private x float
        get x fn(self vertex) float:
            return self.x
        set x fn(self vertex, x float):
            if x >= 0:
                self.x = x
            else:
                self.x = x * -1
        
        private y float
        get y fn(self vertex) float:
            return self.y
        set y fn(self vertex, y float):
            if y >= 0:
                self.y = y
            else:
                self.y = y * -1
        
        dist fn(self vertex) float:
            return math.sqrt(self.x ** 2 + self.y ** 2)
    
    main fn():
        v vertex(1, 2)
        console.log(v.dist())

Properties can have a `get` and `set` method which allows for more control over
what happens when you retrieve or assign a value to that property.

If both a `get` method for a property as well as the actual property is defined,
the `get` method will take preference whenever that method is used outside of
the actual code run by the `get` method itself. Similarly, the `set` method also
overrules an actual property in the type if such actual property is defined.

A `get` method can also be created for the type itself. This will allow you to
retrieve a custom value by directly using the object name, so `vertex` would be
the same as `vertex.get()`. Note that a similar `set` method is not supported.


## Inheritance

    Li 0
    
    import console
    
    main fn():
        myInt type extends int:
            double fn(self myInt) myInt:
                return self * 2
        
        i myInt = 3
        console.log(i.double().double()) // This will print 12

The `extends` keyword allows you to create new types by extending existing ones.
It is possible to inherit from more than one type by listing all the types after
the `extends` keyword.


## Interfaces

    Li 0
    
    import console
    
    private hasX interface:
        x float
    
    private vertex type:
        x, y float
    
    private printX fn(h hasX):
        console.log(h.x)
    
    main fn():
        v vertex(3, 4)
        printX(v) // This will print 3

Interfaces specify which properties and methods a type must have to satisfy it.
You can use an interface as a parameter in a function and ensure that any object
that is passed to that function must satisfy that interface. If you try to pass
an object that does not satisfy the interface or try to use the parameter in the
function in a way that is not allowed by the interface, you will get a compile
error.

Note that a type does not need to specify that it implements the interface, all
it needs to do is to actually conform to it. This allows for an enormous amount
of flexibility in using interfaces.


## Stringers

    Li 0
    
    import console
    
    private person type:
        name string
        age int
        string property:
            get fn(self person) string:
                return "% (% years)".in(self.name, self.age)
    
    main fn():
        a person("Arthur Dent", 42)
        z person("Zaphod Beeblebrox", 9001)
        console.log(a, z)

One of the most ubiquitous interfaces is `stringer` defined by the `string`
package.

    stringer interface:
        get string() string

A `stringer` is a type that can describe itself as a string. The `string`
package (and many others) look for this interface to print values.


## Errors

    Li 0
    
    import:
        console
        time
    
    private myError type:
        when time.Time
        what string
        ok bool = false
        type string = "MyError"
        string property:
            get fn(self myError) string:
                return "at %, %".in(self.when, self.what)
    
    private run fn() error:
        return myError(time.now(), "it didn't work")
    
    main fn():
        if err error = run(); !err.ok:
            console.log(err)

Lithium programs express error state with `error` values.

The `error` type is a built-in interface similar to `string.stringer`:

    default interface:
        string string
        ok bool
        type string

The name `default` as used for this interface is special in that it indicates
that this is the default item in this package, so instead of typing
`error.default`, you can simply use `error`.

Functions often return an `error` value, and calling code should handle errors
by testing whether the `error.ok` is false.

    i int, err error = strconv.aToI("42")
    if !err.ok:
        console.log("couldn't convert number: %".in(err))
        return
    console.log("Converted integer:", i)


## Parametric polymorphism

    Li 0
    
    import console
    
    private plusable interface[n]:
        plus fn(n, n) n
    
    private double fn[p plusable[p]](num p) p:
        return num + num
    
    main fn():
        myInt int = 3
        myInt = double[int](myInt)
        console.log(myInt) // This will print 6
        
        myFloat float = 3.1
        myFloat = double[float](myFloat)
        console.log(myFloat) // This will print 6.2

This is often called generics in other languages. What is happening here is that
we can pass any type of variable to `double()` as long as it implements the
`plusable` interface.

The advantage over just using interfaces is that we can make sure that the
variables retain their original type, so in the first case, `double()` actually
returns an `int` and in the second, `double()` actually returns a `float`. If we
used only interfaces, the return value would only allow properties and methods
of the interface, so we would not be able to print it as `plusable` does not
specify a `string` property.


## Default types for polymorphism

    Li 0
    
    import console
    
    private plusable interface[n]:
        plus fn(n, n) n
    
    private double fn[p plusable[p] int](num p) p:
        return num + num
    
    main fn():
        myInt int = 3
        myInt = double(myInt)
        console.log(myInt) // This will print 6
        
        myFloat float = 3.1
        myFloat = double[float](myFloat)
        console.log(myFloat) // This will print 6.2

A default type can be set for any polymorphic function. If the type is then
omitted when calling the function, this default type is assumed.


## Parametric polymorphism for types

    Li 0
    
    imports console
    
    private vertex type[n string.stringer]:
        x n
        y n
    
    main fn():
        v vertex[int](1, 2)
        console.log(v.x)

Parametric polymorphism can also be used with type definitions to create more
generic types. In the above example, `x` and `y` are of type `int`.


## Arrays

    Li 0
    
    import console
    
    main fn():
        a array[string]
        a(0) = "Hello"
        a(1) = "World"
        console.log(a(0), a(1))
        console.log(a)
        
        primes array[int](2, 3, 5, 7, 11, 13)
        console.log(primes)

`array[t]` is a built-in type that stores an array of variables of type `t`.

The statement `a array[string]` declares a new array `a` of strings. The keys in
an array are always of type `int`.


## Array slices

    Li 0
    
    import console
    
    main fn():
        abcd array[string]("a", "b", "c", "d")
        bc array[string] = abcd(1, 3)
        console.log(bc)

A slice is formed by specifying two indexes, a low and high bound, separated by
a comma: `array(low, high int)`.

This selects a half-open range which includes the first element, but excludes
the last one.


## Array slice defaults

    Li 0
    
    import console
    
    main fn():
        p array[int](2, 3, 5, 7, 11, 13)
        
        p = p(1, 4)
        console.log(p)
        
        p = p(, 2)
        console.log(p)
        
        p = p(1, )
        console.log(p)

When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the array for the high
bound.


## Array length

    Li 0
    
    import console
    
    main fn():
        p array[int](2, 3, 5, 7, 11, 13)
        console.log(p.length)

Array have a length property that shows how many items are currently in the
array.


## Arrays of arrays

    Li 0
    
    import console
    
    main fn():
        // Create a tic-tac-toe board.
        board array[array[string]](
            array[string]("_", "_", "_"),
            array[string]("_", "_", "_"),
            array[string]("_", "_", "_"),
        )
        
        // The players take turns.
        board(0)(0) = "X"
        board(2)(2) = "O"
        board(1)(2) = "X"
        board(1)(0) = "O"
        board(0)(2) = "X"
        
        for i int = 0; i < board.length; i++:
            console.log(board(i).join(" "))

Arrays can contain any type, including other arrays.


## Appending to an array

    Li 0
    
    import console
    
    main fn():
        l array[int]
        console.log(l)
        
        l = l.push(3)
        console.log(l)
        
        l = l.push(4)
        console.log(l)
        
        // We can add more than one element at a time.
        l = l.push(5, 6, 7)
        console.log(l)

Lithium provides a built-in `push` method to append items to an array.


## Iterating over an array

    Li 0
    
    import console
    
    private pow array[int](1, 2, 4, 8, 16, 32, 64, 128)
    
    main fn():
        for i, v int in pow:
            console.log("2 ** % = %".in(i, v))

The built-in `array` type is iterable, which means that it can be used in a for
loop with the `in` keyword.


## Iterating continued

    Li 0
    
    import console
    
    main fn():
        pow array[int](1, 2, 4, 8, 16, 32, 64, 128)
        
        for _, value int in pow:
            console.log(value)

You can skip the key or value by assigning to `_`.

    for key, _ in array
    for _, value in array

If you only want the key, you can omit the second variable.

    for key in array


## Maps

    Li 0
    
    import console
    
    main fn():
        m map[int, string]
        m(0) = "Hello"
        m(1) = "World"
        console.log(m(0), m(1))
        console.log(m)
        
        primes map[int, int](2, 3, 5, 7, 11, 13)
        console.log(primes)

`map[k, v]` is a built-in type that stores a map with keys of type `k` that
maps to values of type `v`.

The statement `m map[int, string]` declares a variable `m` as a map of strings.


## Map initializer with custom keys

    Li 0
    
    import console
    
    private vertex type:
        lat, long float
    
    private m [string, vertex](
        "Bell Labs" = vertex(40.68433, -74.39967),
        "Google" = vertex(37.42202, -122.08408),
    )
    
    main fn():
        console.log(m)

If the keys are not specified in the initialization of a map, keys are
automatically created starting at zero. You can specify the keys just like you
would specify the names of parameter in a function call.


## Mutating maps

    Li 0
    
    import console
    
    main fn():
        m map[string,int]
        
        m("Answer") = 42
        console.log("The value: ", m("Answer"))
        
        m["Answer"] = 48
        console.log("The value: ", m("Answer"))
        
        m.delete("Answer")
        console.log("The value: ", m("Answer"))
        
        v int, ok bool = m("Answer")
        console.log("The value: ", v, "Present? ", ok)

Insert or update an element in map `m`:

    map(key) = value

Retrieve an element:

    var = map(key)

Delete an element:

    map.delete(key)

Test that a key is present with a two-value assignment:

    var, ok = map(key)

If `key` is in the map, `ok` is `true`. If not, `ok` is `false`.

If `key` is not in the map, then `var` is the zero value for the map's value
type.


## Variadic functions

    Li 0
    
    import:
        console
    
    sum fn(...num array[int]) int:
        s int
        for _, n int in num:
            s += n
        return s
    
    main fn():
        console.log(sum(1, 2, 3, 4, 5)) // This will print 15

Variadic functions are functions that can take an arbitrary number of parameters
instead of a fixed set. The parameters are rolled up into a list type such as an
`array` or a `map`.


## Spread operator

    Li 0
    
    import:
        console
    
    sum fn(...num array[int]) int:
        s int
        for _, n int in num:
            s += n
        return s
    
    main fn():
        nums array[int](1, 2, 3, 4, 5)
        console.log(sum(nums...))

It is possible to use the spread operator (`...`) to pass the values in a list
such as an `array` or `map` as distinct parameters in a variadic function.

