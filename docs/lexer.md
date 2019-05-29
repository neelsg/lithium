# Lexical analysis

The lexer converts raw source code into a tokenized format that can be used by
the parser. Each token is placed on a new line.

## Row and column values

Each token except `BOF` and `EOF` is followed by the row and column number where
the token starts. This is required to determine where new statements begin, the
indentation level of blocks and also providing more useful compile error
messages. For the purposes of the tokenized file, tabs should be counted as 4
spaces, but a check needs to be done before this to make sure that the source
does not mix tabs and spaces.

## Must be lossless

It is important not to lose any detail when converting from source to tokens as
these tokens are also used to generate formatted source files to ensure that the
source files all conform to a specific formatting convention.

## Tokens

- `BOF` Beginning of file. Must be followed by file name.
- `EOF` End of file.
- `ID` Identifier.
- `LI` Literal.
- `LC` Literal continued. Used for multi-line doc strings
- `SE` Separator. `: ; , . ( ) [ ] { }`
- `OP` Operator. `+ - * / ** % & | ^ << >> ~ = += -= *= /= **= %= &= |= ^= <<= >>= ! || && == < > <> != <= >= ++ --`
- `CO` Comment, single line
- `CM` Comment, multi-line
- `CC` Comment continued.

## Examples

### Example 1

File: myfile.li

    Li 0
    
    import:
        console
    
    // This is the main function
    main fn():
        console.log("My favorite number is", int.rand(10))

Tokens:

    BOF myfile.li
    ID 1 1 Li
    LI 1 4 0
    ID 3 1 import
    SE 3 7 :
    ID 4 5 console
    CO 6 1 This is the main function
    ID 7 1 main
    ID 7 6 fn
    SE 7 8 (
    SE 7 9 )
    SE 7 10 :
    ID 8 5 console
    SE 8 12 .
    ID 8 13 log
    SE 8 16 (
    LI 8 17 "My favorite number is"
    SE 8 40 ,
    ID 8 42 int
    SE 8 45 .
    ID 8 46 rand
    SE 8 50 (
    LI 8 51 10
    SE 8 53 )
    SE 8 54 )
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
    ID 1 1 Li
    LI 1 0
    ID 3 1 import
    ID 3 8 console
    ID 5 1 private
    SE 5 8 :
    ID 6 5 split
    ID 6 11 fn
    SE 6 13 (
    ID 6 14 sum
    ID 6 18 int
    SE 6 21 )
    SE 6 23 {
    ID 6 24 int
    SE 6 27 ,
    ID 6 29 int
    SE 6 32 }
    SE 6 33 :
    ID 7 9 x
    ID 7 11 int
    SE 7 14 ,
    ID 7 16 y
    ID 7 18 int
    OP 7 22 =
    ID 7 24 sum
    OP 7 28 *
    LI 7 30 4
    OP 7 32 /
    LI 7 34 9
    SE 7 35 ,
    ID 7 37 sum
    OP 7 41 -
    ID 7 43 x
    ID 8 9 return
    ID 8 16 x
    SE 8 17 ,
    ID 8 19 y
    ID 10 1 main
    ID 10 6 fn
    SE 10 8 (
    SE 10 9 )
    SE 10 10 :
    ID 10 12 console
    SE 10 19 .
    ID 10 20 log
    SE 10 23 (
    ID 10 24 split
    SE 10 29 (
    LI 10 30 17
    SE 10 32 )
    SE 10 33 )
    EOF
