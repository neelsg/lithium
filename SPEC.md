# Lithium Programming Language Specification

**Version 0.0.1**

*Last updated 2019-05-07*

## Comments

Single line comments start with `//` and end at the end of the line.

Block comments start with `/*` and end with `*/`


## Statements and expressions

Statements can end with a `;`, but will also end at the end of the line unless the next line is indented.
If the next line is indented and the previous line does not end with a `:`, this line and any further
lines at the same indentation level will still be considered to be part of the previous line.

Examples:

    This is a statement; This is another statement on the same line
    
    This statement ends at the end of this line
    
    This statement carries over to the next lines
        as the next lines
        have been indented

An expression is part of a statement that can be evaluated without taking the rest of the statement into
account. For instance, in the statement `a = b + c`, the `b + c` is an expression.


## Keywords and operators

The following keywords are reserved:

    chan const defer else for func if import in interface type return yield

The following character sequences are operators:

    + - * / % ^ && || ! < > == <= >= != = += -= *= /= %= ^= ++ -- ...


## Literals and constants


## Variables and types


## Code blocks and indentation


## Conditionals


## Loops


## Functions


## Methods


## Packages and files


## Scope

