# Lithium Programming Language Tour

## Welcome

This tour aims to cover the most important parts of the Lithium Programming
Language. This tour borrows heavily on the
[Go Programming Language tour](https://tour.golang.org) as I have based quite a
significant part of this programming language on Go.


## Start

    Li 0
    
    import:
        fmt
        math
    
    export main fn():
        fmt.printLn("My favorite number is", math.rand.int(10))

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
        fmt
        math

    export main fn():
        fmt.printf("Now you have %g problems.\n", math.sqrt(7))

Here we import `fmt` and `math` from the standard library so that they can be
used in this file. We could also write:

    import fmt
    import math


## The main function

    export main fn():

The `fn` keyword is used to define functions (which are blocks of code that can
be executed. When a program is executed, the first thing that will run is the
"main" function.


## Exporting components

    Li 0
    
    private sqr2 fn(x int) int:
        // This function is only available within the current package
        return x * x
    
    export sqr fn(x int) int:
        // This function can be used by other packages
        return sqr2(x)

In Lithium, anything marked as `export` will be made available for use by other
packages. Anything that starts with `private` can only be used in the current
package, but note that this also includes any other files that are part of the
current package.


## Importing custom packages

    Li 0
    
    import:
        fmt
        myMath "./myMathFolder"
    
    export main fn():
        fmt.print(myMath.sqr(3))

You can import your own packages just as easily as packages from the standard
library by providing the file path to the folder where the source code is
located.


## Functions

    Li 0
    
    import:
        fmt
    
    private add fn(x int, y int) int:
        return x + y
    
    export main fn():
        fmt.printLn(add(42, 13))

A function can take zero or more arguments. In this example, `add` takes two
arguments of type `int`. Note that the type comes after the variable name.


## Multiple results

    Li 0
    
    import fmt
    
    private:
        split fn(sum int) (int, int):
            x, y int = sum * 4 / 9, sum - x
            return x, y
    
    export:
        main fn(): fmt.printLn(split(17))

A function can return any number of results.

    Li 0
    
    import:
        fmt
    
    private swap fn(x, y string) (string, string): return y, x
    
    export main fn():
        a, b string = swap("hello", "world")
        fmt.printLn(a, b)

When two or more parameters are of the same type, you can omit the
type from all but the last parameter.


## Named parameters

    Li 0
    
    import:
        fmt
    
    private swap fn(x = "hello", y string) (string, string): return y, x
    
    export main fn():
        a, b string = swap(y = "world")
        fmt.printLn(a, b)

When calling a function, the parameters may be named. A function definition can
specify the default value of any parameters. If these parameters are omitted,
they will be set to the default.


## Variables

    Li 0
    
    import fmt
    
    private:
        c, python, java bool
    
    export go bool
    
    export main fn():
        i int
        fmt.printLn(i, c, python, java, go)

Variables are declared by specifying the name of the variable followed by the
type. Variables can be declared at the package as well as function level. If a
variable is declared at the package level using the `export` keyword, it will
be available for use by other packages as well.


## Initializers

    Li 0
    
    import fmt, math
    
    private i, j int, p float = 1, 2, math.pi
    
    export main fn():
        c, python bool, java, go string = true, false, "no!", "yes!"
        fmt.printLn(i, j, p, c, python, java, go)

A variable declaration can include initializers. If an initializer is omitted,
the variable will be assigned the default value of the type.


## Basic types

    Li 0
    
    import fmt, math
    
    private:
        ToBe bool = false
        MaxInt int.UInt64 = 2 ** 64 - 1
        z complex.Complex128 = math.complex.sqrt(-5 + 0i12)
    
    export main fn():
        fmt.printF("Type: %T Value: %v\n", ToBe, ToBe)
        fmt.printF("Type: %T Value: %v\n", MaxInt, MaxInt)
        fmt.printF("Type: %T Value: %v\n", z, z)

Lithium basic types are:

    bool complex float int string

Each of these types are defined in a standard library package that is imported
implicitly, so you do not need to list them in an import statement.

These packages also provide more fine-grained type definitions such as the
following:

    int.Int8 int.Int16 int.Int32 int.Int64
    int.UInt int.UInt8 int.UInt16 int.UInt32 int.UInt64
    
    float.Float32 float.Float64
    
    complex.Complex64 complex.Complex64


## Default values

    Li 0
    
    import fmt
    
    export main fn():
        i int
        f float
        b bool
        s string
        fmt.printF("%v %v %v %q\n", i, f, b, s)

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
        fmt
        math
    
    export main fn():
        x, y int = 3, 4
        f float = math.sqrt((x * x + y * y).Float)
        z int.UInt = f.UInt
        fmt.printLn(x, y, z)

The built in types have properties to convert to different types. Custom defined
types should also include similar functionality.

Some numeric conversions:

    i int = 42
    f float.Float64 = i.Float64
    u int.UInt = f.UInt

Unlike in C, assignment in Lithium between items of different type requires an
explicit conversion.


## Constants

    Li 0
    
    import fmt
    
    private PI const = 3.14
    
    export main fn = fn():
        WORLD const = "Mars"
        fmt.printLn("Hello", WORLD)
        fmt.printLn("Happy", PI, "Day")
        
        TRUTH const = true
        fmt.printLn("Humans rock?", TRUTH)

Constants are declared like variables, but with the `const` keyword as type.
Constants can be string, boolean, or numeric values. Constants cannot be
declared with an expression that requires runtime evaluation such as calling a
function.


## For

    Li 0
    
    import fmt
    
    export main fn():
        sum int = 0
        for i int = 0; i < 10; i++:
            sum += i
        fmt.printLn(sum)

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
    
    import fmt
    
    export main fn():
        sum int = 1
        for ; sum < 1000; :
            sum += sum
        fmt.printLn(sum)

The init and post statements are optional. At that point you can drop the
semicolons:

    Li 0
    
    import fmt
    
    export main fn():
        sum int = 1
        for sum < 1000:
            sum += sum
        fmt.printLn(sum)


## Forever

    Li 0
    
    export main fn():
        for:
            // Whatever code goes here will run in an infinite loop

If you omit the loop condition it loops forever.


## If

    Li 0
    
    import:
        fmt
        math
    
    private sqrt fn(x float) string:
        if x < 0:
            return sqrt(-x).String + "i"
        return fmt.sPrint(math.sqrt(x))
    
    export main fn():
        fmt.printLn(sqrt(2), sqrt(-4))

Lithium's `if` statements are like its `for` loops


## If with an init statement

    Li 0
    
    import:
        fmt
    
    private pow fn(x, n, lim float) float:
        if v float = x ** n; v < lim:
            return v
        return lim
    
    export main fn():
        fmt.printLn(
            pow(3, 2, 10),
            pow(3, 3, 20),
            )

Like `for`, the `if` statement can start with a short statement to execute
before the condition. Variables declared by the statement are only in scope
until the end of the `if`.


## If and else

    Li 0
    
    import:
        fmt
        math
    
    private pow fn(x, n, lim float) float:
        if x**= n; x < lim:
            return x
        else:
            fmt.printF("%g >= %g\n", x, lim)
        return lim
    
    export main fn():
        fmt.printLn(pow(3, 2, 10), pow(3, 3, 20), )

Variables declared inside an `if` init statement are also available inside any
of the else blocks.


## Defer

    Li 0
    
    import fmt
    
    export main fn():
        defer fmt.printLn("world")
        
        fmt.printLn("hello")

A `defer` statement defers the execution of a function until the surrounding
function returns.

The `defer` statement's arguments are evaluated immediately, but the function
call is not executed until the surrounding function returns.


## Stacking defers

    Li 0
    
    import fmt
    
    export main fn():
        fmt.printLn("counting")
        
        for i int = 0; i < 10; i++:
            defer fmt.printLn(i)
        
        fmt.printLn("done")

Deferred function calls are pushed onto a stack. When a function returns, its
deferred calls are executed in last-in-first-out order.


## Defining new types

    Li 0
    
    import fmt
    
    private Vertex type:
        X int
        Y int
    
    export main fn():
        fmt.printLn(Vertex(1, 2))

Lithium allows simple object-oriented programming. A type is a collection of
properties and methods and is often called a class in other languages.


## Type properties

    Li 0
    
    import fmt
    
    private Vertex type:
        X int
        Y int
    
    export main fn():
        v Vertex = Vertex(1, 2)
        v.X = 4
        fmt.printLn(v.X)

Type properties are accessed using a dot.


## Shorthand for initializing types

    private Vertex type:
        X int
        Y int

Given the above type, the following statements are equivalent:

    v Vertex = Vertex(1, 2)
    v Vertex(1, 2)
    v Vertex(X = 1, Y = 2)


## Methods

    Li 0
    
    import
        fmt
        math
    
    private Vertex type:
        X float
        Y float
        Dist fn(self type) float:
            return math.sqrt(self.X ** 2 + self.Y ** 2)
    
    export main fn():
        v Vertex(1, 2)
        fmt.printLn(v.Dist())

Methods are functions defined within a type.

Note the first parameter refers to the object which the method is associated
with. When calling the method, this parameter is omitted. Within a type, the
`type` keyword refers to the current type, so `self type` is the same as
`self Vertex`.


## Default values and the initializer function

    Li 0
    
    import
        fmt
        math
    
    private Vertex type:
        X float = 1
        Y float = 1
        fn(self type, X, Y float) type:
            self.X, self.Y = X, Y
            return self
        Dist fn(self type) float:
            return math.sqrt(self.X ** 2 + self.Y ** 2)
    
    export main fn():
        v Vertex(1, 2)
        fmt.printLn(v.Dist())

Just like normal variables, properties in a type can also be set to an initial
value which will then be the default for any objects of that type.

You can create an unnamed function in a type which returns the an object of
the type. This function then becomes the initializer and will be used when
creating an object of that type such as with `v Vertex(1, 2)`.


## Private values

    Li 0
    
    import
        fmt
        math
    
    private Vertex type:
        
        private: // X and Y can now only be used by methods within the object
            X float = 1
            Y float = 1
        
        fn(self type, X, Y float) type:
            self.X, self.Y = X, Y
            return self
        
        Dist fn(self type) float:
            return math.sqrt(self.X ** 2 + self.Y ** 2)
    
    export main fn():
        v Vertex(1, 2)
        fmt.printLn(v.Dist())

You can use the `private` keyword to make properties or even methods only
available to the methods in the object itself. If you try to use `v.X` in the
main function in this case, you will get a compile error.


## Dynamic properties

    Li 0
    
    import
        fmt
        math
    
    private Vertex type:
        fn(self type, X, Y float) type:
            self.X, self.Y = X, Y
            return self
        
        private XValue float
        X property:
            get fn(self type) float:
                return self.XValue
            set fn(self type, X float):
                if X >= 0:
                    self.XValue = X
                else:
                    self.XValue = X * -1
        
        private YValue float
        Y property:
            get fn(self type) float:
                return self.YValue
            set fn(self type, Y float):
                if Y >= 0:
                    self.YValue = Y
                else:
                    self.YValue = Y * -1
        
        Dist fn(self type) float:
            return math.sqrt(self.X ** 2 + self.Y ** 2)
    
    export main fn():
        v Vertex(1, 2)
        fmt.printLn(v.Dist())

Properties can have a `get` and `set` method which allows for more control over
what happens when you retrieve or assign a value to that property.

If the `get` method is omitted, you will get a compile error if you try to
retrieve the property. If the `set` method is omitted, you will get a compile
error if you try to assign a value to the property.

A `get` method can also be created at the top of the type which will allow you
to retrieve a custom value by directly using the object name. Note that a
similar `set` method will not work the same way as assigning a new value at top
level will create a new object rather that just changing its value.


## Inheritance

    Li 0
    
    import fmt
    
    export main fn():
        myInt type extends int:
            double fn(self type) type:
                return self * 2
        
        i myInt = 3
        fmt.printLn(i.double().double()) // This will print 12

The `extends` keyword allows you to create new types by extending existing ones.
It is possible to inherit from more than one type by listing all the types after
the `extends` keyword.


## Interfaces

    Li 0
    
    import fmt
    
    private HasX interface:
        X float
    
    private Vertex type:
        X, Y float
    
    private printX fn(h HasX):
        fmt.print(h.X)
    
    export main fn():
        v Vertex(3, 4)
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


## Parametric polymorphism

    Li 0
    
    import fmt
    
    private Plusable interface:
        plus fn(type, type) type
    
    private double fn[P Plusable](num P) P:
        return num + num
    
    export main fn():
        myInt int = 3
        myInt = double[int](myInt)
        fmt.printLn(myInt) // This will print 6
        
        myFloat float = 3.1
        myFloat = double[float](myFloat)
        fmt.printLn(myFloat) // This will print 6.2

This is often called generics in other languages. What is happening here is that
we can pass any type of variable to `double()` as long as it implements the
`Plusable` interface.

The advantage over just using interfaces is that we can
make sure that the variables retain their original type, so in the first
case, `double()` actually returns an `int` and in the second, `double()`
actually returns a `float`. If we used only interfaces, the return value would
only allow properties and methods of the interface, so we would not be able to
print it as `Plusable` does not specify a `String` property.


## Parametric polymorphism for types

    Li 0
    
    imports fmt
    
    private Vertex type[N fmt.Stringer]:
        X N
        Y N
    
    export main fn():
        v Vertex(1,2)
        fmt.printLn(v.X)

Parametric polymorphism can also be used with type definitions to create more
generic types. In the above example, `X` and `Y` would be of type `int`.


## Lists

    Li 0
    
    import fmt
    
    export main fn():
        a [int]string
        a(0) = "Hello"
        a(1) = "World"
        fmt.printLn(a(0), a(1))
        fmt.printLn(a)
        
        primes [int]int(2, 3, 5, 7, 11, 13)
        fmt.printLn(primes)

The type `[K]V` is a list with keys of type `K` and values of type `T`.

The statement `a [int]string` declares a variable `a` as a list of strings.


## Slices

    Li 0
    
    import fmt
    
    export main fn():
        abcd [int]string("a", "b", "c", "d")
        bc [int]string = abcd(1, 3)
        fmt.printLn(bc)

A slice is formed by specifying two indices, a low and high bound, separated by
a comma: `a(low , high)`.

This selects a half-open range which includes the first element, but excludes
the last one.


## Slice defaults

    Li 0
    
    import fmt
    
    export main fn():
        p [int]int(2, 3, 5, 7, 11, 13)
        
        p = p(1, 4)
        fmt.printLn(p)
        
        p = p(, 2)
        fmt.printLn(p)
        
        p = p(1, )
        fmt.printLn(p)

When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the list for the high
bound.


## List length

    Li 0
    
    import fmt
    
    export main fn():
        p [int]int(2, 3, 5, 7, 11, 13)
        fmt.print(p.length)

Lists have a length property that shows how many items are currently in the
list.


## Lists of lists

    Li 0
    
    import fmt
    
    export main fn():
        // Create a tic-tac-toe board.
        board [int][int]string(
            [int]string("_", "_", "_"),
            [int]string("_", "_", "_"),
            [int]string("_", "_", "_"),
        )
        
        // The players take turns.
        board(0)(0) = "X"
        board(2)(2) = "O"
        board(1)(2) = "X"
        board(1)(0) = "O"
        board(0)(2) = "X"
        
        for i int = 0; i < board.length; i++:
            fmt.printF("%s\n", strings.Join(board(i), " "))

Lists can contain any type, including other lists.


## Appending to a list

    Li 0
    
    import fmt
    
    export main fn():
        l [int]int
        fmt.printLn(l)
        
        l = l.push(0)
        fmt.printLn(l)
        
        l = l.push(1)
        fmt.printLn(l)
        
        // We can add more than one element at a time.
        l = l.push(2, 3, 4)
        fmt.printLn(l)

Lithium provides a built-in `push` method to append items to a list.


## Iterating over a list

    Li 0
    
    import fmt
    
    private pow [int]int(1, 2, 4, 8, 16, 32, 64, 128)
    
    export main fn():
        for i int, v int in pow.range:
            fmt.printF("2**%d = %d\n", i, v)

The built-in `range` property provides an iterator than can be used to loop over
the items in a list.


## Range continued

    Li 0
    
    import fmt
    
    export main fn():
        pow [int]int(1, 1, 1, 1, 1, 1, 1, 1)
        
        for i int in pow.range:
            pow(i) = 1 << i.UInt
        
        for _, value int in pow.range:
            fmt.printF("%d\n", value)

You can skip the key or value by assigning to `_`.

    for key, _ in list.range
    for _, value in list.range

If you only want the key, you can omit the second variable.

    for key in list.range


## Maps shorthand

    Li 0
    
    import fmt
    
    private Vertex type:
        lat, long float
    
    private m [string]Vertex(
        "Bell Labs" = (40.68433, -74.39967),
        "Google" = (37.42202, -122.08408),
    )
    
    export main fn():
        fmt.printLn(m)

If the top-level type is just a type name, you can omit it from the elements of
the literal.


## Mutating lists

    Li 0
    
    import fmt
    
    export main fn():
        m [string]int
        
        m("Answer") = 42
        fmt.printLn("The value:", m("Answer"))
        
        m["Answer"] = 48
        fmt.printLn("The value:", m("Answer"))
        
        m.delete("Answer")
        fmt.printLn("The value:", m("Answer"))
        
        v int, ok bool = m("Answer")
        fmt.printLn("The value:", v, "Present?", ok)

Insert or update an element in list `m`:

    m(key) = elem

Retrieve an element:

    elem = m(key)

Delete an element:

    m.delete(key)

Test that a key is present with a two-value assignment:

    elem, ok = m(key)

If key is in `m`, `ok` is `true`. If not, `ok` is `false`.

If key is not in the list, then `elem` is the zero value for the list's element
type.


## Function values

    Li 0
    
    import:
        fmt
        math
    
    private compute fn(func fn(float, float) float) float:
        return func(3, 4)
    
    export main fn():
        hypot fn(x, y float) float:
            return math.sqrt(x*x + y*y)
        
        fmt.printLn(hypot(5, 12))
        
        fmt.printLn(compute(hypot))
        fmt.printLn(compute(math.pow))

Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.


## Function closures

    Li 0
    
    import fmt
    
    private adder fn() fn(int) int:
        sum int = 0
        return fn(x int) int:
            sum += x
            return sum
    
    export main fn():
        pos int, neg int = adder(), adder()
        for i int; i < 10; i++:
            fmt.printLn(
                pos(i),
                neg(-2*i),
            )

Lithium functions may be closures. A closure is a function value that references
variables from outside its body. The function may access and assign to the
referenced variables. In this sense the function is "bound" to the variables.

For example, the `adder` function returns a closure. Each closure is bound to
its own `sum` variable.

