# Lithium Programming Language Specification

This spec was inspired by the
[Go Programming Language Specification](https://golang.org/ref/spec)


## Doctype indicator

The first line of a Lithium source file must start with the character sequence
`li ` followed by the version of the language specification that was used.


## Comments

Single line comments start with `//` and end at the end of the line.

Block comments start with `/*` and end with `*/`


## Statements and expressions

Statements can end with a `;`, but will also end at the end of the line unless
the next line is indented. If the next line is indented and the previous line
does not end with a `:`, this line and any further lines at the same indentation
level will still be considered to be part of the previous line.

Examples:

    This is a statement; This is another statement on the same line
    
    This statement ends at the end of this line
    
    This statement carries over to the next lines
        as the next lines
        have been indented

An expression is part of a statement that can be evaluated without taking the
rest of the statement into account. For instance, in the statement `a = b + c`,
the `b + c` is an expression.


## Keywords

The following keywords are reserved:

    array bool byte chan complex const defer else embed export extends false
    float for fn if import int in interface iota map nil private property return
    string true type


## Operators

The following character sequences are operators:

- `x + y` Addition `x.plus(y)`
- `x - y` Subtraction `x.minus(y)`
- `x * y` Multiplication `x.multiply(y)`
- `x / y` Division `x.divide(y)`
- `x ** y` Exponent `x.power(y)`
- `x % y` Modulus `x.modulo(y)`
- `x & y` Bitwise And `x.bitAnd(y)`
- `x | y` Bitwise Or `x.bitOr(y)`
- `x ^ y` Bitwise Xor `x.bitXor(y)`
- `x << y` Left shift `x.shiftLeft(y)`
- `x >> y` Right shift `x.shiftRight(y)`
- `~x` Bitwise ones complement `x.bitFlip()`
- `x = y` Assignment. Points `x` to the same value as `y`
- `x.a = y` Set a property `x.a.set(y)`
- `x(a) = y` Set an internal value `x.set(a, y)`
- `x(a)(b) = y` Set a nested internal value `x(a).set(b, y)`
- `x += y` Add assign `x = x + y`
- `x -= y` Subtract assign `x = x - y`
- `x *= y` Multiply assign `x = x * y`
- `x /= y` Divide assign `x = x / y`
- `x **= y` Exponent assign `x = x ** y`
- `x %= y` Modulo assign `x = x % y`
- `x &= y` Bitwise And assign `x = x & y`
- `x |= y` Bitwise Or assign `x = x | y`
- `x ^= y` Bitwise Xor assign `x = x ^ y`
- `x <<= y` Left shift assign `x = x << y`
- `x >>= y` Right shift assign `x = x >> y`
- `!x` Boolean not
- `x || x` Boolean or
- `x && x` Boolean and
- `x == y` Is equal to `x.equals(y)`
- `x < y` Is less than `x.less(y)`
- `x > y` Is greater than `x.greater(y)`
- `x <> y` Is not equal to `!x.equals(y)`
- `x != y` Is not equal to `!x.equals(y)`
- `x <= y` Is less than or equal to `x.less(y) || x.equals(y)`
- `x >= y` Is greater than or equal to `x.greater(y) || x.equals(y)`
- `x++` Increment by one `x = x + 1`
- `x--` Decrement by one `x = x - 1`


## Literals

### Boolean literals

The words `true` and `false` can be used as boolean literals


### Integer literals

Integer literals are a sequence of digits (0..9) that represent an integer
constant. Integer literals can also be written in hexadecimal, octal or
binary by using the prefix `0x`, `0o` or `0b` respectively.

Examples of integer literals:

    42
    0xFADE123
    0o600
    0b101010

### Floating-point literals:

Floating-point literals are a sequence of digits, a decimal point, a
fractional part and optionally and exponent part (prefixed by `e`).
The decimal point can be omitted if the literal contains an exponent.

Examples of floating-point literals:

    0.
    42.4242
    1.e+0
    42e+10
    0.003e-10

### Complex number literals:

Complex numbers are represented as two floating-point literals separated
by a `+` where the second floating-point literal ends with the char `i`. The
first floating-point literal is the real part and the second is the imaginary
part.

Examples of complex number literals:

    1 + 2i
    4.2e+1 + 4.2e-1i

### String literals:

Strings are a sequence of characters that can be formatted in a few different
ways:

- Characters encapsulated by single `'` or double `"` quotes.
- A docstring block indicated by the statement `embed string:`

Strings encapsulated in quotation marks cannot flow over a single line.

With strings encapsulated in quotation marks, the following escape sequences can
be used to escape special chars. The following escape sequences are available:

- `\\` This represents a single backslash.
- `\"` This represents a double quote (Only applies if the literal is
    encapsulated in double quotes).
- `\'` This represents a single quote (Only applies if the literal is
    encapsulated in single quotes).
- `\n` A newline character.
- `\r` A carriage-return character.
- `\l` A linefeed character.

Examples of string literals:

    "This is a string\n"
    'This is also a string'
    'Strings are concatenated with a +, so this' + " is a single string literal"
    
    embed string:
        Use this if you need to create a string that contains lots of raw text
        such as for templates etc.
        
        The entire thing is one string literal up to the point where this is
        de-dented. All the line-breaks are preserved in the string literal, but
        the initials tabs / spaces used for indentation is stripped away. Any
        escape characters such as \\ or \n have no meaning within a string
        literal. Indicators for comments such as // and /* are also just treated
        as part of the raw text rather than as actual comments in the code.


## Constants

A constant is defined by the name of the constant followed by the keyword
`const`, the operator `=` and then an assigned value. The assigned value can be
any literal of any type or even a simplistic formula using literals.

Examples of constants:

    PI const = 3.14
    HELLO const = "Hello world!"
    DAY const = 24 * 60 * 60 * 1000


## Iota

The `iota` keyword can be used anywhere in the source code to signify that a new
integer number constant should be used starting with 0. Each time `iota` is used
in a single package, a new integer is used, but this may overlap with integers
created with `iota` in imported packages.

Example:

    RED const = iota // This will set the RED const to 0
    GREEN, BLUE const = iota, iota // This will set GREEN to 1 and BLUE to 2
    myVal int = iota // The myVal variable will initially be set to 3


## Variables and types


## Code blocks and indentation


## Conditionals


## Loops


## Functions


## Methods


## Files

All the files under a single folder is considered part of the same package.
Files must have the extension ".li".

The underscore ("_") is used in file names to indicate various special affixes
used by Lithium and should only be used for this purpose. Some of these affixes
are:

- "_test": This file is used for unit testing
- "_linux": This file only applies when compiling to the GNU/Linux OS
- "_riscv": This file only applies when compiling to the RISC-V architecture

File affixes can be combined such as "Filename_riscv_linux_test.li".


## Scope

The following rules apply to determine the scope in which a variable is
available

1. Variables declared at the top of a file must be either preceded by `private`
    or `export`. 
1. Variables preceded by `export` is available by any other package that imports
    the current package.
1. Top level variables preceded by `private` is available anywhere in the
    current package including any other files within this package
1. Variables declared within a function is only available within that function.
1. Variables declared within an indented block is only available within that
    block.
1. Variables declared within the initial statement of a `for` or `if` is only
    available within that statement, but including any `else` statements.
