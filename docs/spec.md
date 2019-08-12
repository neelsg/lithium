# Lithium Programming Language Specification

This document defines the Lithium programming language.

Note: Where curly braces (`{ }`) are used in this documentation, these indicate customizable parts of the source code and does not literally include curly braces.


## Statements

Statements are discrete steps that are executed by the program. Statements can be terminated with a `;`, but will also end at the end of the current line unless the next line is indented. If the next line is indented, this line and any further lines at the deeper indentation level will still be considered to be part of the same statement.


## Expressions

An expression can be part of a statement or stand as a statement by itself. Expressions evaluate to a value. For instance, in the statement `a = b + c`, the part `b + c` is an expression. The part `b` is also an expression within the `b + c` expression.


## Indents and dedents

Indentation is created by adding spaces or tabs to the front of a line. A line is indented when it has more tabs or spaces at the start of the line than there was for the previous non-blank line. A line is dedented if it contains fewer tabs or spaces than the previous non-blank line. Blank lines are ignored completely, including for indentation purposes.

A block is a set of lines that have the same level of indentation. Blocks may be nested within other blocks if they have a higher level of indentation.

For indentation purposes, either tabs or spaces may be used, but they may not be mixed.


## Blocks

A block is an indented set of statements or other text that is preceded by a block statement followed by a `:`. The indentation of all the lines in a block must be deeper than the block statement and must be the same for all the lines within the block.

A block ends when the next line has the same or shallower indentation as the block statement.


## Comments

Comments start with `#` and end at the end of the line. There is no special support for multi-line comments as you might find in other languages such as C


## Files

All the files under a single folder are considered part of the same package. Files must have the extension ".li".

The underscore ("\_") is used in file names to indicate various special affixes used by Lithium and should only be used for this purpose.

A file can have multiple affixes separated by underscores.

Affix for test files:

    test

Affixes for platforms, operating systems or architectures:

    amd64 android arduino arm jvm linux none riscv riscv64 webasm webjs windows


## Keywords

The following keywords are reserved in Lithium:

    co const defer else enum false func for get if import in interface is mut pub return set true type


## Import

An `import` statement is used to indicate packages that need to be imported into the current package. This statement must be placed at the top level of each file where the imported packages are used.

The package name is a variable name used within the current file to refer to the imported package.

The package path is a string literal that specifies the path where the imported package's source files are located.

The `import` statement can take two forms:

### Inline form:

    import {package name} "{package path}"{, ...}

### Block form:

    import:
        {package name} "{package path}"
        {...}


## Built-in library packages

The built-in library packages provide basic functionality. The availability as well as contents of any built-in packages may vary depending on the architecture and operating system which is targeted.

Built-in packages do not need to be imported with the `import` keyword and are available implicitly.


## Operators

The following operators are available:

    !  ||  &&  ==  !=  <  >  <=  >=  +  -  *  /  %  =  +=  -=  ++  --

### Boolean operators

Boolean operators only work on boolean values. If they are used by any other types, it will first implicitly convert to `bool` by trying the `x.bool` property. If no such property is present or that property does not return a boolean, a compile time error will result. The boolean operators are:

- `!x` Boolean not. Will return the opposite of `x`.
- `x || y` Boolean or. Will return true if either `x` or `y` is true.
- `x && y` Boolean and. Will return true if both `x` and `y` is true.

For any boolean expressions, the order will first be `!` (Not) and then `||` and `&&` from left to right. Brackets can be used to change the order.

Short circuiting is used, so in the following cases, `y` will not even be evaluated:

- `x || y` if `x` is `true`.
- `x && y` if `x` is `false`.

Note that we do not include an XOR operation as that can be trivially written as `x != y` where both `x` and `y` are boolean or more verbosely as `(x || y) && !(x && y)`.

### Comparison operators

Comparison operators are syntactic sugar for calling a method on one of the operands (Left or right depends on which operator is used as shown below). These methods must return a single `bool` value.

The comparison operators are:

- `x == y` Is equal to. `x.==(y)`
- `x != y` Is not equal to. `!x.==(y)`
- `x < y` Is less than. `x.<(y)`
- `x > y` Is greater than. `x.>(y)` or `y.<(x)`
- `x <= y` Is less than or equal to. `!x.>(y)` or `!y.<(x)`
- `x >= y` Is greater than or equal to. `!x.<(y)` or `!y.>(x)`

Boolean operators are evaluated before comparison operators. The order of evaluation can be changed with brackets.

### Mathematical operators

Mathematical operators perform mathematical operations on operands of the same type and return a result of that same type. These operators are syntactic sugar for calling a method on the left operand passing the right operand as the parameter.

The mathematical operators are:

- `x + y` Addition. `x.+(y)`
- `x - y` Subtraction. `x.-(y)`
- `x * y` Multiplication. `x.*(y)`
- `x / y` Division. `x./(y)`
- `x % y` Modulus. `x.%(y)`

Mathematical operators are evaluated before boolean or comparison operators and follow a conventional order as follows:

1. Multiplication, division and modulus from left to right.
1. Addition and subtraction from left to right.

### Assignment operators

Assignment operators will set the value of the left operand to the value of the right operand or to the result of an operation between the left and right operands.

The assignment operators are:

- `x = y` Assignment. Points `x` to the same value as `y`
- `x.a = y` Set a property. `x.a.=(y)`
- `x(a) = y` Set an internal value. `x.=(a, y)`
- `x(a)(b) = y` Set a nested internal value. `x(a).=(b, y)`

Assignment operations are performed last and cannot be nested further within other operations, they must stand as statements by themselves.

#### In-place assignments

The following operators can be used when the result should be assigned to the left most operator:

- `x += y` Add assign. `x = x + y`
- `x -= y` Subtract assign. `x = x - y`

#### Assignment increment and decrement shorthand

The following shorthands are available for assignments:

- `x++` Increment by one `x += 1`
- `x--` Decrement by one `x -= 1`


## Literals

Literals are sequences of characters that Lithium recognizes as a concrete value of a certain type.

### Boolean literals

The words `true` and `false` can be used as boolean literals.

Boolean literals can only be used for assignment of the type `bool`.

### Integer literals

Integer literals are a sequence of digits (0..9) that represent an integer constant. Integer literals can also be written in hexadecimal, octal or binary by using the prefix `0x`, `0o` or `0b` respectively.

Integer literals can be of any size and can be assigned to any type under `int` and can also be assigned to any type under `float` and `complex`. If the literal is too big to fit in the range of the the given type, a compile error should be raised.

For readability, spaces can be used in integer literals to denote groupings such as thousands.

Examples of integer literals:

    42
    0xFADE123
    0o600
    0b101010
    123 456 789 000
    0x AA BB CC DD

### Floating-point literals:

Floating-point literals are a sequence of digits, a decimal point, a fractional part and optionally and exponent part (prefixed by `e`). The decimal point can be omitted if the literal contains an exponent.

Floating-point literals can be of any size and can be assigned to any type under `float`. If the precision of literal is higher than what is possible in the given type, a lower precision rounded value will be used and the compiler should not raise any error.

As with integer literals, spaces can be used to denote groupings such as thousands. Unlike integer literals, floating-point literals must be decimal, they cannot be represented in hexadecimal, octal or binary.

Examples of floating-point literals:

    0.
    42.4242
    1.e+0
    42e+10
    0.003 e - 10

### Complex number literals:

Complex numbers are represented as two floating-point literals separated by a `+` where the second floating-point literal ends with the char `i`. The first floating-point literal is the real part and the second is the imaginary part.

The real part can be omitted if the number only has an imaginary part. The decimal point can also be omitted of the imaginary part if not needed.

Examples of complex number literals:

    1. + 2i
    4.2e+1 + 4.2e-2i
    3i

### String literals:

Strings are a sequence of unicode characters encapsulated by quotes `"`.

String literals cannot flow over a single line (Use embedded blocks for that). With string literals, the following escape sequences can be used to escape special chars. The following escape sequences are available:

- `\\` This represents a single backslash.
- `\"` This represents a double quote.
- `\a` Bell code.
- `\b` Backspace.
- `\e` Escape character.
- `\f` Page break.
- `\l` Linefeed.
- `\n` Newline.
- `\r` Carriage-return.
- `\t` Horizontal tab.
- `\v` Vertical tab.

Examples of string literals:

    "This is a string\n"
    "Strings are concatenated with a +, so this" + " is a single string literal"

#### String interpolation

String interpolation can be used to embed the result of an expression within a string. The syntax is as follows: `\(expression)`

Example:

    "1 + 2 = \(1 + 2)" # This represents "1 + 2 = 3"
    "The value of myVar is: \(myVar)"

#### Strings using embedded blocks 

To define an embbed string block use `string:` to begin the block and follow normal indentation rules within the block. This is using embedded blocks which can actually be custom created as long as the type referred to in the square brackets define a method with the signature `compile(string) (type.any, error)`. On compilation, this method will be called on the string literal and if the returned `error.ok` value is not `true`, a compile error will be raised.

Examples of string literals using embedded blocks:

    myString string:
        Use this if you need to create a string that contains lots of raw text
        such as for templates etc.
        
        The entire thing is one string literal up to the point where this is
        de-dented. All the line-breaks are preserved in the string literal, but
        the initials tabs / spaces used for indentation is stripped away. Any
        escape characters such as \\ or \n have no meaning within a string
        literal. Indicators for comments such as // and /* are also just treated
        as part of the raw text rather than as actual comments in the code.
    
    # This will run `web.css.embed(string) (string, error)` first to check if
    # this is a valid stylesheet. If it is not, the compiler will throw an error.
    myStyle string = web.css:
        body {
            font-size: 1.5em;
        }

### Tuple literals

Tuple literals are a collection of other literals enclosed in braces and separated by commas to indicate that they are part of a tuple.

Examples of tuple literals are:

    (1, 2, 3)
    
    # Nested tuples are also possible:
    (("a", "b", "c"), (1, 2, 3), (1.1, 2.2, 3.3))


## Constants

A constant is defined by the name of the constant followed by the keyword `const`, the operator `=` and then an assigned value. The assigned value can be any literal of any type or even a simplistic formula using literals, but the types cannot be mixed.

Examples of constants:

    debug const = true                   # Boolean constant
    pi const = 3.14                      # Floating-point constant
    hello const = "Hello world!"         # String constant
    day const = 24 * 60 * 60 * 1000      # Integer constant
    i const = 1i                         # Complex number constant
    html const = web.html:               # String constant from a block
        <div class="greeting">
            Hello world wide web!
            My name is %.
        </div>


## Types

Lithium uses static typing

### Tuples as types

It is possible to use tuples as types where a single variable may contain a more complex set of values.

Example of using a tuple as type:

    vector (float, float, fn(float, float) float)
        = 3, 4, func(x, y float) float:
            return math.sqrt(math.power(x, 2) + math.power(y, 2))
    vector.2(vector.0, vector.1) # This will return 5.

### Defining new types

New types can be defined with the `type` keyword.

    {type name} type({local object name}):
        {properties and methods...}

### Extends

Extends can be used for type inheritence.

    {type name} type({local object name}, {parent type name}) extends {parent type}:
        {properties and methods...}


## Tuples

A tuple is a finite ordered list of elements. The elements can be of any type including tuples. Tuples are defined as the set of elements encapsulated in brackets. A tuple with only one element is the same type as the element itself and can be used interchangeably. Tuples with a different number of elements or elements of a different type are considered different types.

Functions that return more than one result actually returns a tuple of the results. Assignment statements where multiple elements are assigned to are actually a destructuring assignment of tuple elements.

A single value in a tuple can be accessed with the name of the tuple followed by `.` and then the index number of the element starting at zero. The index number in this case must be an integer literal number.


## Scope

The following rules apply to determine the scope in which a variable is available

1. Variables declared at the top of a file may be preceded by `pub`. This will make the variable available to other packages that import this package.
1. Variables declared within an indented block is only available within that block. This includes functions and also any blocks created by `for`, `if` etc.
1. Variables declared within the block statement of a `for` or `if` is only available within that block, but also including any `else` statements.


## Conditionals

The `if` keyword is used to conditionally execute code. An `if` statement can take various forms:

    # Inline form
    if {init statement; }{condition}: {statement} {else:} {statement}
    
    # Block form
    if {init statement; }{condition}:
        {statement/(s)}
    {else if condition:}
        {statement/(s)}
    {else:}
        {statement/(s)}

The inline form may be used as an expression in a larger statement. The block form can only be used as a self-standing block.

The {initial statement ;} is optional and can contain a single statement that is always executed before the condition is evaluated. Any variables defined in this statement is available within the if statement/block as well as any else blocks, but is not available outside of it.

The {condition} expression must evaluate to a boolean. If the expression evaluates to `true`, the {statement/(s)} are executed.

The {else if condition:} and {else:} blocks are optional. There can be an arbitrary number of {else if condition:} blocks.

### if is

Another form of if statement can be used to check if a variable implements a certain interface. Inside the if statement, the functionality of that interface can then be used for the variable. Multiple variables can be used before the `is` keyword to check if they all implement the interface.

The `&&` operator can be used to check multiple variables against different kinds of interfaces or to check for additional conditions. Note that `||` cannot be used in the same way, as there would then be ambiguity within the if block whether a variable does implement the interface. `else` can be used just as with a normal `if`.

    if {init statement; }{variable}{, variable...} is {interface}{ && condition/(s)}:
        {statement/(s)}

### if in

Another form of if statement can be used to check if an element exists within a collection. Multiple elements can be used before the `in` keyword to check if they are all in the collection.

The `&&` and `||` operators can be used to check multiple elements in different collections or to check for additional conditions. `else` can be used just as with a normal `if`.

    if {init statement; }{element}{, element...} in {collection}{ && / || conditions/(s)}:
        {statement/(s)}


## Loops

The `for` keyword is used to execute code multiple times. `for` can take various forms:

    # Inline form
    for {init statement; }{condition}{; post statement}: {statement}
    
    # Shortened inline form
    for {condition}: {statement}
    
    # Block form
    for {init statement; }{condition}{; post statement}:
        {statement/(s)}

If neither an init nor a post statement is required, the for can be shortened.
In this case, for is the same as "while" in other languages.

### for else

The `else` keyword can be used to execute code if the for loop does not execute
even once.

    for {init statement; }{condition}{; post statement}:
        {statement/(s)}
    else:
        {statement/(s)}

In addition, it is possible that the code within the loop must be executed at
least once (Such as with a "do while" loop in other languages). In this case,
the `else` keyword can be used within the shortened form to indicate that the
code must run at least once:

    for else {condition}:
        {statement/(s)}

### for in

A special kind of for loop can be used to loop through all the items in an
iterable object.

    for {item tuple} in {iterable object}:
        {statement/(s)}

An object is iterable if it includes the following methods:

    first() (item tuple, index int, ok bool)
    # `index` is the index number of the first item
    # if `ok` is false, the object is empty
    
    next(prior int) (item tuple, index int, ok bool)
    # the `prior` is the index number of the previous item
    # if `ok` is false, there is no item after the `prior` index

An else can be used as with a normal `for` to run if the object was empty.


## Functions

Functions are first-class citizens in Lithium. Functions are defined using the `func` keyword. The function signature consists of the parameters it can take, including their type, and the return values it will send, including their type. The function signature is considered to be the function's type. Functions with the same signature are considered to be the same type and functions with signatures that are different are considered to be of different types.

There cannot be multiple functions with the same name as with programming languages like C or C++.

### Default parameter values

Parameters can have default values assigned and then be optionally omitted when called. If no default value is assigned, the parameter must be specified when called.

### Named parameters

Parameters can be specified by name when calling the function.

### Returning values

The `return` keyword is used to return from a function. If the function has no return values, `return` can be used as a single statment that will return back to the parent function. If the function should return values, these must be listed after the `return` keyword as a comma separated list.

### Function signatures

Functions with different signatures are considered different types


## Deferring execution

The `defer` keyword will defer execution of a function or block until the end of the current function. There are two forms a `defer` statement can take:

    # Inline form
    defer function({parameters...})
    
    # Block form
    defer:
        {statements/(s)}

In the inline form, the parameters of the function are evaluated immediately, but the execution is delayed until the current function returns. In the block form, the statements are only executed when the current function returns and no evaluation is done immediately


## Coroutines

The `co` keyword is used to create a new coroutine. Coroutines share the same address space, but can be executed in parallel.


## Interfaces

The `interface` keyword can be used to define a new interface. Interfaces are used in the place of concrete types when passing parameters or return values for functions so that more than one type can be used as long as the required capabilities are present.

    {name} interface:
        {signatures of methods required}

Any concrete type can also be used as an iterface where the signatures will then be the same as what the type provides. This can be used to provide a default implementation for an interface.

### any

`any` is an interface that contains no methods or properties at all and can be used to pass any type of value.


## Polymorphism

Parametric polymorphism allows for functions, types and interfaces to transparently support more than one concrete type as long as the required interface functionality is provided.
