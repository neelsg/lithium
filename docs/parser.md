# Parsing into intermediate code

The parser for Lithium will read a token file and produce a vastly simplified intermediate format that we call an assembly file. This contains a sort of assembly language in a similar vein that Java produces bytecode or C# produces CIL. The assembly file can then be used by compilers or interpreters for a specific target. The process to produce the assembly file is always the same regardless of the target, so this assembly file will be the final representation of Lithium code that is common to any target.

## Assembly files include dependencies

An assembly file is a single file that includes all the code from the current package, but not any code from any packages that have been imported. Those imported packages will produce their own assembly files.

## Assembly files are still high level

An assembly file does not contain true machine code or assembly language. The code is still at a higher level where there are concepts like variables and threads. The idea is that it represents a deeply simplified form, but still abstract enough that it can be efficiently compiled to any platform and apply things like garbage collection or parallel processing effectively.

## Commands

- `import {label} {path}` This indicates an external package that needs to be imported
- `func {label}` This indicates the start of a new function
- `funcpub {label}` This indicates the start of a new function that is publicly accessible from other packages
- `label {label}` This indicates a new local label within a given function
- `jump {label}` This is a jump command to a given local label. This cannot be used to jump outside of the current function
- `var {type} {number}` This defines a variable as a given type. The number is used to identify the variable and must be unique within the current function
- `set {type} {number} {value}` Set the value of a variable to a literal value
- `push {type} {number}` Push a value of a variable onto the stack
- `pushl {type} {value}` Push a literal value onto the stack
- `pop {type} {number}` Pop a value off the stack into a variable
- `call {label} {count}` Call a function with the given label and the given count of arguments
- `ret` Return to the previous function
- `add {type}` Add the last two items from the stack assuming that they are of the given type. Pop the numbers from the stack and push the result onto the stack
- `sub {type}` Subtract the last items on the stack from the second last from the stack assuming that they are of the given type. Pop the numbers from the stack and push the result onto the stack
- `mul {type}` Multiply the last two items from the stack assuming that they are of the given type. Pop the numbers from the stack and push the result onto the stack

## Types

- `bit` A single bit that can be 0 or 1
- `int1s` An 8-bit signed integer
- `int2s` A 16-bit signed integer
- `int4s` A 32-bit signed integer
- `int8s` A 64-bit signed integer
- `int1u` An 8-bit unsigned integer
- `int2u` A 16-bit unsigned integer
- `int4u` A 32-bit unsigned integer
- `int8u` A 64-bit unsigned integer
- `float4` A 32-bit floating point number
- `float8` A 64-bit floating point number
- `pointer` A pointer to a location on the heap
- `str` A string

## Examples

### Example 1

Source file:

    Li 0
    
    # This is the main function
    main func:
        console.log("My favorite number is", int.rand(10))

Assembly file:

    Li 0 core.assembly
    
    func main
        pushl [str] "My favorite number is"
        pushl [int4s] 10
        call int.rand 1
        call console.log 2
        ret

### Example 2

Source file:

    Li 0
    
    multiply func(x, y int) int:
        # This function is only available within the current package
        return x * y
    
    pub sqr func(x int) int:
        # This function can be used by other packages
        return multiply(x, x)

Assembly file:

    Li 0 core.assembly
    
    func multiply
        mul [int4s]
        ret
    
    funcpub sqr
        var [int4s] 0
        pop [int4s] 0
        push [int4s] 0
        push [int4s] 0
        call multiply 2
        ret
