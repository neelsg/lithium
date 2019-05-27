# Lithium Programming Language Specification

This document defines the Lithium programming language.


## doctype indicator

The first line of a Lithium source file must start with the character sequence
`Li ` followed by the version of the language specification that was used.


## indents and dedents

Indentation is created by adding spaces or tabs to the front of a line. Where we
specify a line to be indented, this implies that there are more tabs or spaces
that precede the given line than there was for the previous non-blank line.
Blank lines are ignored completely, including for indentation purposes.

A line is dedented if it contains fewer tabs or spaces than the previous
non-blank line.

A block is a set of lines that have the same level of indentation. Blocks may be
nested within other blocks if they have a higher level of indentation.

For indentation purposes, either tabs or spaces may be used, but they may not
be mixed.


## statements

Statements are discrete steps that are executed by the program. Statements can
be terminated with a `;`, but will also end at the end of the current line
unless the next line is indented. If the next line is indented, this line and
any further lines at the deeper indentation level will still be considered to be
part of the same statement.


## expressions

An expression can be part of a statement or stand as a statement by itself.
Expressions evaluate to a value. For instance, in the statement `a = b + c`,
the part `b + c` is an expression. The part `b` is also an expression within the
`b + c` expression.


## blocks

A block is an indented set of statements or other text that is preceded by a
block statement followed by a `:`. The indentation of all the lines in a block
must be deeper than the block statement and must be the same for all the lines
within the block.

A block ends when the next line has the same or shallower indentation as the
block statement.


## comments

Single line comments start with `//` and end at the end of the line.

Block comments start with `/*` and end with `*/`


## files

All the files under a single folder is considered part of the same package.
Files must have the extension ".li".

The underscore ("_") is used in file names to indicate various special affixes
used by Lithium and should only be used for this purpose.

A file can have multiple affixes separated by underscores.

Affix for test files:

    test

Affixes for platforms, operating systems or architectures:

    amd64 android arduino arm jvm linux none riscv riscv64 webasm webjs windows


## keywords

The following keywords are reserved in Lithium:

    array bool chan co complex const default defer else error extends false
    float fn for get if implements import in init int interface map return set
    string true type


## import

An `import` statement is used to indicate packages that need to be imported into
the current package. This statement must be placed at the top level of each file
where the imported packages are used.

The package name is a variable name used within the current file to refer to the
imported package.

The package path is an optional string literal that specifies the path where the
importer package's source files are located. Package path can be omitted for any
packages within the standard library.

The `import` statement can take two forms:

### Inline form:

    import package name{ "package path"}{, package name...}

### Block form:

    import:
        package name{ "package path"}
        {...}


## built-in and standard library packages

There are both built-in as well as standard library packages. The availability
as well as specific format of any packages may vary depending on the
architecture and operating system which is targeted.

Built-in packages do not need to be imported with the `import` keyword and are
available implicitly. Standard packages are shipped with the language, but they
do need to be imported using an `import` statement to be used by any specific
file.

We do not list the standard packages here as their documentation is provided
separately.

The built-in packages are:

- `array` Provides a type and methods for working with generic arrays
- `bool` Provides a type and methods for working with booleans
- `chan` Provides a type and methods for working with channels
- `complex` Provides a type and methods for working with complex numbers
- `error` Provides a type and methods for working with errors
- `float` Provides a type and methods for working with floating-point numbers
- `int` Provides a type and methods for working with integer numbers
- `map` Provides a type and methods for working with generic maps
- `string` Provides a type and methods for working with strings
- `type` Provides methods for working with general types


## operators

The following character sequences are operators:

- `x + y` Addition `x.plus(y)`
- `x - y` Subtraction `x.minus(y)`
- `x * y` Multiplication `x.multiply(y)`
- `x / y` Division `x.divide(y)`
- `x ** y` Exponent `x.power(y)`
- `x % y` Modulus `x.modulo(y)`
- `x & y` Bitwise And `x.and(y)`
- `x | y` Bitwise Or `x.or(y)`
- `x ^ y` Bitwise Xor `x.xor(y)`
- `x << y` Left shift `x.left(y)`
- `x >> y` Right shift `x.right(y)`
- `~x` Bitwise ones complement `x.flip()`
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
- `x == y` Is equal to `x.equal(y)`
- `x < y` Is less than `x.less(y)`
- `x > y` Is greater than `y.less(x)`
- `x <> y` Is not equal to `!x.equal(y)`
- `x != y` Is not equal to `!x.equal(y)`
- `x <= y` Is less than or equal to `x.less(y) || x.equal(y)`
- `x >= y` Is greater than or equal to `y.less(x) || x.equal(y)`
- `x++` Increment by one `x = x + 1`
- `x--` Decrement by one `x = x - 1`


## literals

Literals are sequences of characters that Lithium recognizes as a concrete value
of a certain type.

### Boolean literals

The words `true` and `false` can be used as boolean literals.

Boolean literals can only be used for assignment of the type `bool`.

### Integer literals

Integer literals are a sequence of digits (0..9) that represent an integer
constant. Integer literals can also be written in hexadecimal, octal or
binary by using the prefix `0x`, `0o` or `0b` respectively.

Integer literals can be of any size and can be assigned to any type under `int`.
If the literal is too big to fit in the range of the the given type, a compile
error should be raised.

For readability, spaces can be used in integer literals to denote groupings such
as thousands.

Examples of integer literals:

    42
    0xFADE123
    0o600
    0b101010
    123 456 789 000
    0x AA BB CC DD

#### int.iota

The `int.iota` keyword can be used anywhere in the source code to signify that a
new integer number literal should be used starting with 0. Each time `int.iota`
is used in a single package, a unique integer is used, but this may overlap with
integers created with `iota` in imported packages. No guarantees are made that
the specific numbers will remain the same between compiles of the source.

Example:

    RED const = int.iota // This will set the RED const to 0
    GREEN, BLUE const = int.iota, int.iota // Will set GREEN to 1 and BLUE to 2
    myVal int = int.iota // The myVal variable will initially be set to 3

### Floating-point literals:

Floating-point literals are a sequence of digits, a decimal point, a fractional
part and optionally and exponent part (prefixed by `e`). The decimal point can
be omitted if the literal contains an exponent.

Floating-point literals can be of any size and can be assigned to any type under
`float`. If the precision of literal is higher than what is possible in the
given type, a lower precision rounded value will be used and the compiler should
not raise any error.

As with integer literals, spaces can be used to denote groupings such as
thousands. Unlike integer literals, floating-point literals must be decimal,
they cannot be represented in hexadecimal, octal or binary.

Examples of floating-point literals:

    0.
    42.4242
    1.e+0
    42e+10
    0.003 e - 10

### Complex number literals:

Complex numbers are represented as two floating-point literals separated
by a `+` where the second floating-point literal ends with the char `i`. The
first floating-point literal is the real part and the second is the imaginary
part.

The real part can be omitted if the number only has an imaginary part. The
decimal point can also be omitted of the imaginary part if not needed.

Examples of complex number literals:

    1. + 2i
    4.2e+1 + 4.2e-2i
    3i

### String literals:

Strings are a sequence of unicode characters that can be formatted in a few
different ways:

1. Characters encapsulated by single `'` or double `"` quotes.
1. Using the `string.embed` keyword.

Strings encapsulated in quotation marks cannot flow over a single line. With
strings encapsulated in quotation marks, the following escape sequences can be
used to escape special chars. The following escape sequences are available:

- `\\` This represents a single backslash.
- `\"` This represents a double quote (Only applies if the literal is
   encapsulated in double quotes).
- `\'` This represents a single quote (Only applies if the literal is
    encapsulated in single quotes).
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
    'This is also a string'
    'Strings are concatenated with a +, so this' + " is a single string literal"

#### string.embed

The `string.embed` keyword can be used in the following ways:

1. A docstring block indicated by `string.embed {validatable type}:`
1. Text imported from a file by `string.embed {validatable type} "filepath"`

The `{validatable type}` part is optional and can be any type where the type
has a method with the signature `validate(string) error`. On compilation, this
method will be called on the string literal and if the returned `error.ok` value
is not `true`, the value of `error.string` will be printed as a compile error.

Examples of string literals using `embed`:

    string.embed:
        Use this if you need to create a string that contains lots of raw text
        such as for templates etc.
        
        The entire thing is one string literal up to the point where this is
        de-dented. All the line-breaks are preserved in the string literal, but
        the initials tabs / spaces used for indentation is stripped away. Any
        escape characters such as \\ or \n have no meaning within a string
        literal. Indicators for comments such as // and /* are also just treated
        as part of the raw text rather than as actual comments in the code.
    
    //This will include the full text of the file in the literal at compile time
    string.embed "./MyTextFile.txt"
    
    /*
       This will run `web.css.validate(string) error` first to check if
       "./style.css" contains a valid stylesheet. If it does not, the
       compiler will throw an error.
    */
    string.embed web.css "./style.css"


## const

A constant is defined by the name of the constant followed by the keyword
`const`, the operator `=` and then an assigned value. The assigned value can be
any literal of any type or even a simplistic formula using literals, but the
types cannot be mixed.

Examples of constants:

    DEBUG const = true                   // Boolean constant
    PI const = 3.14                      // Floating-point constant
    HELLO const = "Hello world!"         // String constant
    DAY const = 24 * 60 * 60 * 1000      // Integer constant
    I const = 1i                         // Complex number constant
    HTML const = string.embed web.html:  // String constant from a block
        <div class="greeting">
            Hello world wide web!
            My name is %.
        </div>


## type

The following are built-in types in the language:

- `array[v bool.comparable]` An array with values of type `v`
- `bool` A boolean with a value of either `true` or `false`.
- `complex` A complex number. This can be 64 or 128-bit depending on the target
   architecture.
- `complex.p64` A 64-bit complex number.
- `complex.p128` A 128-bit complex number.
- `float` A floating-point number. This can be 32 or 64-bit depending on the
   target architecture.
- `float.p32` A 32-bit floating-point number.
- `float.p64` A 64-bit floating-point number.
- `fn(...) {...}` A function. Functions with different types of parameters or
   different types of return values are considered to be different types.
- `int` A signed integer number. This can be 16, 32 or 64-bit depending on the
   target architecture.
- `int.s8` A signed 8-bit integer.
- `int.s16` A signed 16-bit integer.
- `int.s32` A signed 32-bit integer.
- `int.s64` A signed 64-bit integer.
- `int.u` An unsigned integer. This can be 16, 32 or 64-bit depending on the
   target architecture.
- `int.u8` An unsigned 8-bit integer.
- `int.u16` An unsigned 16-bit integer.
- `int.u32` An unsigned 32-bit integer.
- `int.u64` An unsigned 64-bit integer.
- `map[k bool.comparable, v type.any]` A map with keys of type `k` and values of `v`
- `string` A string of text.

New types can be defined with the `type` keyword.


## scope

The following rules apply to determine the scope in which a variable is
available

1. Variables declared at the top of a file may be preceded by `private`. This
   will make the variable only available to the current package.
1. Variables declared within an indented block is only available within that
   block. This includes functions and also any blocks created by `for`, `if`
   etc.
1. Variables declared within the block statement of a `for` or `if` is only
   available within that block, but also including any `else` statements.


## if

The `if` keyword is used to conditionally execute code. An `if` statement can
take the following forms:

    // Inline form
    if {initial statement ;}{condition}: {expression} {else:} {expression}
    
    // Block form
    if {initial statement ;}{condition}:
        {statement/(s)}
    {else if condition:}
        {statement/(s)}
    {else:}
        {statement/(s)}

The inline form may be used as an expression in a larger statement. The block
form can only be used as a self-standing block.

The {initial statement ;} is optional and can contain a single statement that is
always executed before the condition is evaluated. Any variables defined in this
statement is available within the if statement/block as well as any else blocks,
but is not available outside of it.

The {condition} expression must evaluate to a boolean. If the expression
evaluates to `true`, the [statement/(s)} are executed.

The {else if condition:} and {else:} blocks are optional. There can be an
arbitrary number of {else if condition:} blocks.


## for

The `for` keyword is used to execute code multiple times.


## fn

Functions are first-class variables in Lithium. Functions are defined using the
`fn` keyword.


## return


## tuples


## defer


## co


## interface



