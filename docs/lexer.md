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
- `SE` Separator. `: ; , . ( ) [ ]`
- `OP` Operator. `+ - * / ** % = += -= *= /= **= %= ! || && == < > != <= >= ++ --`
- `CO` Comment, single line
- `CM` Comment, multi-line
- `CC` Comment continued.

## Examples

### Example 1

File: myfile.li

    Li 0
    
    # This is the main function
    main func:
        console.log("My favorite number is", int.rand(10))

Tokens:

    BOF myfile.li
    ID 1 1 Li
    LI 1 4 0
    CO 3 1 " This is the main function"
    ID 4 1 main
    ID 4 6 func
    SE 4 10 :
    ID 5 5 console
    SE 5 12 .
    ID 5 13 log
    SE 5 16 (
    LI 5 17 "My favorite number is"
    SE 5 40 ,
    ID 5 42 int
    SE 5 45 .
    ID 5 46 rand
    SE 5 50 (
    LI 5 51 10
    SE 5 53 )
    SE 5 54 )
    EOF
