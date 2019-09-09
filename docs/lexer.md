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
- `TAB` Indent or dedent
- `IDE` Identifier.
- `LIT` Literal.
- `LTC` Literal continued. Used for multi-line doc strings
- `SEP` Separator. `: ; , . ( ) [ ]`
- `OPE` Operator. `+ - * / ** % = += -= *= /= **= %= ! || && == < > != <= >= ++ --`
- `COM` Comment

## Examples

### Example 1

File: myfile.li

    # This is the main function
    main func:
        console.log("My favorite number is", int.rand(10))

Tokens:

    BOF myfile.li
    COM 1 1 " This is the main function"
    IDE 2 1 main
    IDE 2 6 func
    SEP 3 10 :
    TAB 3 5 +1
    IDE 3 5 console
    SEP 3 12 .
    IDE 3 13 log
    SEP 3 16 (
    LIT 3 17 "My favorite number is"
    SEP 3 40 ,
    IDE 3 42 int
    SEP 3 45 .
    IDE 3 46 rand
    SEP 3 50 (
    LIT 3 51 10
    SEP 3 53 )
    SEP 3 54 )
    EOF
