# Parsing into intermediate code

The parser for Lithium will read a token file and produce a vastly simplified intermediate format that we call an assembly file. This contains a sort of assembly language in a similar vein that Java produces bytecode or C# produces CIL. The assembly file can then be used by compilers or interpreters for a specific target. The process to produce the assembly file is always the same regardless of the target, so this assembly file will be the final representation of Lithium code that is common to any target.

## Assembly files include dependencies

An assembly file is a single file that includes all the code from the current package, but not any code from any packages that have been imported. Those imported packages will produce their own assembly files.

## Assembly files are still high level

An assembly file does not contain true machine code or assembly language. The code is still at a higher level where there are concepts like variables and threads. The idea is that it represents a deeply simplified form, but still abstract enough that it can be efficiently compiled to any platform and apply things like garbage collection or parallel processing effectively.

## Commands

- `import {label} {path}` This indicates an external package that needs to be imported
- `func {label}` This indicates the start of a new function
- `funcp {label}` This indicates the start of a new function that is publicly accessible from other packages
- `label {label}` This indicates a new local label within a given function
- `jump {label}` This is a jump command to a given local label. This cannot be used to jump outside of the current function
- `var {type} {label}` This defines a variable as a given type. The label is used to identify the variable and must be unique within the current function
- `varg {type} {label}` This defines a variable as a given type in the global scope for the current package. The label is used to identify the variable and must be unique within the package
- `setl {label} {value}` Set the value of a variable to a literal value
- `setgl {label} {value}` Set the value of a global variable to a literal value 
- `push {label}` Push a value of a variable onto the stack
- `pushg {label}` Push a value of a global variable onto the stack
- `pushl {type} {value}` Push a literal value onto the stack
- `pop {label}` Pop a value off the stack into a variable
- `popg {label}` Pop a value off the stack into a global variable
- `call {label}` Call a function
- `calli {label}` Call a function from a different package
- `ret` Return to the previous function
- `add {type}` Add the last two items from the stack assuming that they are of the given type. Pop the numbers from the stack and push the result onto the stack
- `sub {type}` Subtract the last item on the stack from the second last from the stack assuming that they are of the given type. Pop the numbers from the stack and push the result onto the stack
- `mul {type}` Multiply the last two items from the stack assuming that they are of the given type. Pop the numbers from the stack and push the result onto the stack
- `div {type}` Divide the second last item on the stack by the last from the stack assuming that they are of the given type. Pop the numbers from the stack and push the result onto the stack

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

Source:

    Li 0
    
    # This is the main function
    main func:
        console.log("My favorite number is", int.rand(10))

Assembly:

    Li 0 core.assembly
    
    func main
        pushl str "My favorite number is"
        pushl int4s 10
        calli int.rand
        pushl int4u 2 # This is used for variadic functions to indicate how many parameters are passed
        calli console.log
        ret

### Example 2

Source:

    Li 0
    
    multiply func(x, y int) int:
        # This function is only available within the current package
        return x * y
    
    pub sqr func(x int) int:
        # This function can be used by other packages
        return multiply(x, x)

Assembly:

    Li 0 core.assembly
    
    func multiply
        mul int4s
        ret
    
    funcp sqr
        var int4s x
        pop x
        push x
        push x
        call multiply
        ret

### Example 3

Source:

    Li 0

    split func(sum int) (int, int):
        x int, y int = sum * 4 / 9, sum - x
        return x, y

    swap func(x, y string) (string, string): return y, x

    main func:
        console.log(split(17))

        a string, b string = swap("hello", "world")
        console.log(a, b)

Assembly:

    Li 0 core.assembly
    
    func split
        var int4s sum
        var int4s x
        var int4s y
        pop sum
        push sum
        pushl int4s 4
        mul int4s
        pushl int4s 9
        div int4s
        pop x
        push sum
        push x
        sub int4s
        pop y
        push x
        push y
        ret
    
    func swap
        var str x
        var str y
        pop x
        pop y
        push y
        push x
        ret
    
    func main
        var str a
        var str b
        pushl int4s 17
        call split
        pushl int4u 1
        calli console.log
        pushl str "hello"
        pushl str "world"
        call swap
        pop a
        pop b
        push a
        push b
        pushl int4u 2
        calli console.log
        ret

### Example 4

Source:

    Li 0
    
    i, j int, p float = 1, 2, math.pi
    
    main func:
        k, l bool, m, n string = true, false, "no!", "yes!"
        console.log(i, j, p, k, l, m, n)

Assembly:

    Li 0 core.assembly
    
    varg int4s i
    varg int4s j
    varg float4 p
    
    setgl i 1
    setgl j 2
    setgl p 3.141592653589793238462643383279502884197169399375105820974944
    
    func main
        var bit k
        var bit l
        var str m
        var str n
        set k 1
        set l 0
        set m "no!"
        set n "yes!"
        pushg i
        pushg j
        pushg p
        push k
        push l
        push m
        push n
        pushl int4u 7
        calli console.log
        ret
    
