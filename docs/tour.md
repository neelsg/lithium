# Lithium Programming Language Tour

## Welcome

This tour aims to cover the most important parts of the Lithium Programming
Language. This tour borrows heavily on the
[Go Programming Language tour](https://tour.golang.org) as I have based quite a
significant part of this programming language on Go.


## Start

    Li 0.0
    
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

The version can be written in full such as `0.0.1` or shortened such as `0`. If
any part of the version is omitted, the compiler will assume the most recent
sub-version should be used.


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

    Li 0.0
    
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

    Li 0.0
    
    import:
        fmt
        myMath "./myMathFolder"
    
    export main fn():
        fmt.print(myMath.sqr(3))

You can import your own packages just as easily as packages from the standard
library by providing the file path to the folder where the source code is
located.


## Functions

    Li 0.0
    
    import:
        fmt
    
    local add fn(x int, y int) int:
        return x + y
    
    export main fn():
        fmt.printLn(add(42, 13))

A function can take zero or more arguments. In this example, `add` takes two
arguments of type `int`. Note that the type comes after the variable name.


## Multiple results

    Li 0.0
    
    import fmt
    
    local:
        split fn(sum int) (int, int):
            x, y int = sum * 4 / 9, sum - x
            return x, y
    
    export:
        main fn(): fmt.printLn(split(17))

A function can return any number of results.

    Li 0.0
    
    import:
        fmt
    
    local swap fn(x, y string) (string, string): return y, x
    
    export main fn():
        a, b string = swap("hello", "world")
        fmt.printLn(a, b)

When two or more parameters are of the same type, you can omit the
type from all but the last parameter.


## Variables

    Li 0.0
    
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

    Li 0.0
    
    import fmt, math
    
    local i, j int, p float = 1, 2, math.pi
    
    export main fn():
        c, python bool, java, go string = true, false, "no!", "yes!"
        fmt.printLn(i, j, p, c, python, java, go)

A variable declaration can include initializers. If an initializer is omitted,
the variable will be assigned the default value of the type.


## Basic types

    Li 0.0
    
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

    Li 0.0
    
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

    Li 0.0
    
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

    Li 0.0
    
    import fmt
    
    local Pi const = 3.14
    
    export main fn = fn():
        World const = "Mars"
        fmt.printLn("Hello", World)
        fmt.printLn("Happy", Pi, "Day")
        
        Truth const = true
        fmt.printLn("Humans rock?", Truth)

Constants are declared like variables, but with the `const` keyword as type.
Constants can be string, boolean, or numeric values. Constants cannot be
declared with an expression that requires runtime evaluation such as calling a
function.
