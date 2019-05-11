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

    chan const defer else export extends false for fn if import in interface
    nil private property return true type


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
- `~x` Bitwise Ones Complement `x.bitFlip()`
- `x = y` Assignment
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
by the char `i`. The number before the `i` is the real part and the
number after is the imaginary part.

Examples of complex number literals:

    1i2
    4.2e+1i4.2e-1


## Constants


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
