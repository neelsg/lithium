# Lithium Programming Language Tour

## Welcome

This tour aims to cover the most important parts of the Lithium Programming Language. This tour borrows heavily
on the [Go Programming Language tour](https://tour.golang.org) as I have based quite a significant part of this
programming language on Go


## 1. Start

    li 0.0.1
    
    import:
    	fmt
    	math
    
    export main fn():
    	fmt.printLn("My favorite number is", math.rand.int(10))

#### File format and doctype
Lithium source code is stored in plain text files that have the `.li` file extension. All files that are located
in the same folder are considered to be part of the same package. Each file needs to start with the chars `li `
followed by the language version used in the file. This will allow newer compilers to compile files written
using older specs of the language and for older packages to be incorporated into newer projects.

#### import
The `import` statement allows for importing of other packages so that they can be used in the current file. In
this case we are importing `fmt` and `math` from the standard library.

#### export main fn():
The `fn` keyword is used to define functions (which are blocks of code that can be executed. When a program
is executed, the first thing that will run is the `main()` function

#### fmt.Println
The `fmt.Println` method is used to print a line out to the console. In this case, it will print the text "My
favorite number is" and a random number from 0 to 10


## 2. Imports

    li 0.0.1

    import:
    	fmt
    	math

    export main fn():
    	fmt.printf("Now you have %g problems.\n", math.sqrt(7))

Here we import `fmt` and `math` from the standard library so that they can be used in this file. We could also write:

    import fmt
    import math


## 3. Exports

    li 0.0.1
    
    local sqr2 fn(x int) int:
    	// This function is only available within the current package
    	return x * x
    
    export sqr fn(x int) int:
    	// This function can be used by other packages
    	return sqr2(x)

In Lithium, anything marked as `export ` will be made available for use by other packages. Anything that starts with
`local ` can only be used in the current package, but note that this also includes any other files that are part of
the current package.


## 4. Importing custom packages

    li 0.0.1
    
    import:
    	fmt
    	myMath "./myMathFolder"
    
    export main fn():
    	fmt.print(myMath.sqr(3))

You can import your own packages just as easily as packages from the standard library by providing the file path to
the folder where the source code is located.


## 5. Functions

    li 0.0.1
    
    import:
    	fmt
    
    local add fn(x int, y int) int:
    	return x + y
    
    export main fn():
    	fmt.printLn(add(42, 13))

A function can take zero or more arguments.
In this example, `add` takes two arguments of type `int`.
Note that the type comes after the variable name.

### Multiple results

    li 0.0.1

    import:
    	fmt
    
    local swap fn(x string, y string) (string, string):
        return y, x
    
    export main fn():
    	a, b := swap("hello", "world")
    	fmt.printLn(a, b)

 A function can return any number of results.
 The swap function returns two `strings`.
 
 
 
