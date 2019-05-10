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
    
    local sqr2 fn(x int) int:
        // This function is only available within the current package
        return x * x
    
    export sqr fn(x int) int:
        // This function can be used by other packages
        return sqr2(x)

In Lithium, anything marked as `export ` will be made available for use by other
packages. Anything that starts with `local ` can only be used in the current
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
    
    local add fn(x int, y int) int:
        return x + y
    
    export main fn():
        fmt.printLn(add(42, 13))

A function can take zero or more arguments. In this example, `add` takes two
arguments of type `int`. Note that the type comes after the variable name.


## Multiple results

    Li 0
    
    import fmt
    
    local:
        split fn(sum int) (int, int):
            x, y int = sum * 4 / 9, sum - x
            return x, y
    
    export:
        main fn(): fmt.printLn(split(17))

A function can return any number of results.

    Li 0
    
    import:
        fmt
    
    local swap fn(x, y string) (string, string): return y, x
    
    export main fn():
        a, b string = swap("hello", "world")
        fmt.printLn(a, b)

When two or more parameters are of the same type, you can omit the
type from all but the last parameter.


## Named parameters

    Li 0
    
    import:
        fmt
    
    local swap fn(x = "hello", y string) (string, string): return y, x
    
    export main fn():
        a, b string = swap(y = "world")
        fmt.printLn(a, b)

When calling a function, the parameters may be named. A function definition can
specify the default value of any parameters. If these parameters are omitted,
they will be set to the default.


## Variables

    Li 0
    
    import fmt
    
    local:
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
    
    local i, j int, p float = 1, 2, math.pi
    
    export main fn():
        c, python bool, java, go string = true, false, "no!", "yes!"
        fmt.printLn(i, j, p, c, python, java, go)

A variable declaration can include initializers. If an initializer is omitted,
the variable will be assigned the default value of the type.


## Basic types

    Li 0
    
    import fmt, math
    
    local:
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
    
    local PI const = 3.14
    
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
    
    local sqrt fn(x float) string:
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
    
    local pow fn(x, n, lim float) float:
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
    
    local pow fn(x, n, lim float) float:
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


## Classes

    Li 0
    
    import fmt
    
    local Vertex class:
        X int
        Y int
    
    export main fn():
        fmt.printLn(Vertex(1, 2))

Lithium allows simple object-oriented programming. A class is a collection of
properties and methods.


## Class properties

    Li 0
    
    import fmt
    
    local Vertex class:
        X int
        Y int
    
    export main fn():
        v Vertex = Vertex(1, 2)
        v.X = 4
        fmt.printLn(v.X)

Class properties are accessed using a dot.


## Shorthand for initializing classes

    local Vertex class:
        X int
        Y int

Given the above class, the following statements are equivalent:

    v Vertex = Vertex(1, 2)
    v Vertex(1, 2)
    v Vertex(X = 1, Y = 2)


## Methods

    Li 0
    
    import
        fmt
        math
    
    local Vertex class:
        X float
        Y float
        Dist (V) fn() float:
            return math.sqrt(V.X ** 2 + V.Y ** 2)
    
    export main fn():
        v Vertex(1, 2)
        fmt.printLn(v.Dist())

Methods are functions defined within a class. Note the `(V)` before `fn()`, this
is the name with which the object can be accessed within the function itself.


## Default values and the initializer function

    Li 0
    
    import
        fmt
        math
    
    local Vertex class:
        X float = 1
        Y float = 1
        (V) fn(X, Y float) Vertex:
            V.X, V.Y = X, Y
            return V
        Dist (V) fn() float:
            return math.sqrt(V.X ** 2 + V.Y ** 2)
    
    export main fn():
        v Vertex(1, 2)
        fmt.printLn(v.Dist())

Just like normal variables, properties in a class can also be set to an initial
value which will then be the default for any objects of that class.

You can create an unnamed function in a class which returns the an object of
the class type. This function then becomes the initializer and will be used when
creating an object of that class such as with `v Vertex(1, 2)`.


## Private values

    Li 0
    
    import
        fmt
        math
    
    local Vertex class:
        
        local: // X and Y can now only be accessed by methods within the object
            X float = 1
            Y float = 1
        
        (V) fn(X, Y float) Vertex:
            V.X, V.Y = X, Y
            return V
        
        Dist (V) fn() float:
            return math.sqrt(V.X ** 2 + V.Y ** 2)
    
    export main fn():
        v Vertex(1, 2)
        fmt.printLn(v.Dist())

You can use the `local` keyword to make properties or even methods only
available to the methods in the object itself. If you try to use `v.X` in the
main function in this case, you will get a compile error.


## Dynamic properties

    Li 0
    
    import
        fmt
        math
    
    local Vertex class:
        
        local XValue float
        X property:
            get (V) fn() float:
                return V.XValue
            set (V) fn(X float):
                if X >= 0:
                    V.XValue = X
                else:
                    V.XValue = X * -1
        
        local YValue float
        Y property:
            get (V) fn() float:
                return V.YValue
            set (V) fn(Y float):
                if Y >= 0:
                    V.YValue = Y
                else:
                    V.YValue = Y * -1
        
        (V) fn(X, Y float) Vertex:
            V.X, V.Y = X, Y
            return V
        
        Dist (V) fn() float:
            return math.sqrt(V.X ** 2 + V.Y ** 2)
    
    export main fn():
        v Vertex(1, 2)
        fmt.printLn(v.Dist())

Properties can have a `get` and `set` method which allows for more control over
what happens when you retrieve or assign a value to that property.

If the `get` method is omitted, you will get a compile error if you try to
retrieve the property. If the `set` method is omitted, you will get a compile
error if you try to assign a value to the property.

A `get` and `set` method can also be assigned at the top of the class which will
allow you to assign or retrieve a custom value by directly using the object name


## Inheritance

    Li 0
    
    import fmt
    
    export main fn():
        myInt class extends int:
            double (I) fn() myInt:
                return I * 2
        
        i myInt = 3
        fmt.printLn(i.double().double()) // This will print 12

The `extends` keyword allows you to create new classes by extending existing
ones. It is possible to inherit from more than one class by listing all the
classes after the `extends` keyword.

