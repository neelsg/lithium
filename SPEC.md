# Lithium Programming Language Specification

**Version 0.0.1**

*Last updated 2019-05-07*

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


## Keywords and operators

The following keywords are reserved:

    and chan const default defer else export
    false for fn if import in interface is
    li local nil or return struct true type
    xor yield

The following character sequences are operators:

- `x + y` Addition `x.add(y)`
- `x - y` Subtraction `x.subtract(y)`
- `x * y` Multiplication `x.multiply(y)`
- `x / y` Division `x.divide(y)`
- `x % y` Modulus `x.modulo(y)`
- `x ^ y` Exponent `x.exponent(y)`
- `x & y` Binary And `x.binAnd(y)`
- `x | y` Binary Or `x.binOr(y)`
- `~x` Binary Ones Complement `x.binFlip()`
- `x = y` Assignment `x.set(y)`
- `x += y` Add assign `x = x + y`
- `x -= y` Subtract assign `x = x - y`
- `x *= y` Multiply assign `x = x * y`
- `x /= y` Divide assign `x = x / y`
- `x %= y` Modulo assign `x = x % y`
- `x ^= y` Power assign `x = x ^ y`
- `x &= y` Binary And assign `x = x & y`
- `x |= y` Binary Or assign `x = x | y`
- `x == y` Is equal to `x.equalTo(y)`
- `x < y` Is less than `x.lessThan(y)`
- `x > y` Is greater than `x.greaterThan(y)`
- `x <> y` Is not equal to `x.notEqualTo(y)`
- `x != y` Is not equal to `x.notEqualTo(y)`
- `x <= y` Is less than or equal to `x.lessThanOrEqual(y)`
- `x >= y` Is greater than or equal to `x.greaterThanOrEqual(y)`
- `x++` Increment `x.increment()`
- `x--` Decrement `x.decrement()`


## Literals and constants


## Variables and types


## Code blocks and indentation


## Conditionals


## Loops


## Functions


## Methods


## Packages and files


## Scope

