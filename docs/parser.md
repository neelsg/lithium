# Parsing into intermediate code

The parser for Lithium will read a token file and produce a vastly simplified
intermediate format that we call an object file. This will almost represent
assembly language in a similar vein that Java produces bytecode or C# produces
CIL. The output file can then be used by compilers or interpreters for a
specific target. The process to produce the object file must always be the same
regardless of the target, so this object file will be the final representation
of Lithium code that is common to any target.

## Compile time checks

Compile time checks such as type and overflow checking is done by the parser.
This means that once we get to the object file, we assume that the file is
correct and no further checks are required. For this reason, object files will
never be a safe format to receive from untrusted sources and need to be treated
with caution.

## Object files include dependencies

An object file is a single file that includes all the code from the current
package as well as all the code from any packages that have been imported. It is
not possible to split an object file out so that it contains only the current
package and not any dependencies. It is also not possible to use object files
from dependencies instead of token files in the compilation of a given package.

## Object files are still high level

An object file does not contain true machine code or assembly language. The code
is still at a higher level where there are concepts like variables with scoping
and threads. The idea is that it represents a deeply simplified form, but still
abstract enough that it can be efficiently compiled to any platform and apply
things like garbage collection or parallel processing effectively.

## Commands

- `.{namespace}.{label}` This indicates a named address.
- `pushl {type} {value}` Push a literal value of a given type onto the stack
- `call {label} {count}` Call a function at the given label with the number of parameters
- `ret` Return to the previous function


## Examples

### Example 1

File: myfile.li

    Li 0
    
    import:
        console
    
    // This is the main function
    main fn():
        console.log("My favorite number is", int.rand(10))

Object file:

    {there would be more lines here for the "console" and "int" packages}
    
    ..main
        pushl str My favorite number is
        pushl i32 10
        call .int.rand 1
        call .console.log 2
        ret
