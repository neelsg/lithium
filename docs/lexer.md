# Lexical analysis

The lexer converts raw source code into a tokenized format that can be used by
the parser. Each token is placed on a new line.

## Tokens

- `BOF` Beginning of file. Must be followed by file name.
- `EOF` End of file.
- `NEW` Newline. Must be followed by the line number.
- `IDR` Identifier.
- `LIT` Literal.
- `DEN` Indentation. Must indicate the new level.
- `SEP` Separator. `: ; , . ( ) [ ] { }`
- `OPE` Operator. `+ - * / ** % & | ^ << >> ~ = += -= *= /= **= %= &= |= ^= <<= >>= ! || && == < > <> != <= >= ++ --`
- `COM` Comment.

## Examples

### Example 1

File: myfile.li

    Li 0
    
    import:
        console
    
    main fn():
        console.log("My favorite number is", int.rand(10))

Tokens:

    BOF myfile.li
    NEW 1
    IDR Li
    LIT 0
    NEW 3
    IDR import
    SEP :
    NEW 4
    DEN 1
    IDR console
    NEW 6
    DEN 0
    IDR main
    IDR fn
    SEP (
    SEP )
    SEP :
    NEW 7
    DEN 1
    IDR console
    SEP .
    IDR log
    SEP (
    LIT "My favorite number is
    SEP ,
    IDR int
    SEP .
    IDR rand
    SEP (
    LIT 10
    SEP )
    SEP )
    EOF

### Example 2

File: myfile.li

    Li 0
    
    import console
    
    private:
        split fn(sum int) {int, int}:
            x int, y int = sum * 4 / 9, sum - x
            return x, y
    
    main fn(): console.log(split(17))

Tokens:

    BOF myfile.li
    NEW 1
    IDR Li
    LIT 0
    NEW 3
    IDR import
    IDR console
    NEW 5
    IDR private
    SEP :
    NEW 6
    DEN 1
    IDR split
    IDR fn
    SEP (
    IDR sum
    IDR int
    SEP )
    SEP {
    IDR int
    SEP ,
    IDR int
    SEP }
    SEP :
    NEW 7
    DEN 2
    IDR x
    IDR int
    SEP ,
    IDR y
    IDR int
    OPE =
    IDR sum
    OPE *
    LIT 4
    OPE /
    LIT 9
    SEP ,
    IDR sum
    OPE -
    IDR x
    NEW 8
    IDR return
    IDR x
    SEP ,
    IDR y
    NEW 10
    DEN 0
    IDR main
    IDR fn
    SEP (
    SEP )
    SEP :
    IDR console
    SEP .
    IDR log
    SEP (
    IDR split
    SEP (
    LIT 17
    SEP )
    SEP )
    EOF
