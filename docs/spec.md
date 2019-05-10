# Lithium Programming Language Specification

This spec was highly inspired by the [Go Programming Language Specification](https://golang.org/ref/spec)


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

    and chan class const default defer else export
    false for fn if import in inherit interface is
    Li local nil or property return true xor yield

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
- `x = y` Assignment `x.set(y)`
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
- `x == y` Is equal to `x.equalTo(y)`
- `x < y` Is less than `x.lessThan(y)`
- `x > y` Is greater than `x.greaterThan(y)`
- `x <> y` Is not equal to `x.notEqualTo(y)`
- `x != y` Is not equal to `x.notEqualTo(y)`
- `x <= y` Is less than or equal to `x.lessThanOrEqual(y)`
- `x >= y` Is greater than or equal to `x.greaterThanOrEqual(y)`
- `x++` Increment `x.increment()`
- `x--` Decrement `x.decrement()`


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


## Packages and files


## Scope

