
/* This is command line tool that can be used to compile lithium source files into go source files */

import (
    "io/ioutil"
    "log"
    "os"
)

// File contains the contents of a file and may be extended later on to also contain meta about the file
type File struct {
    Contents string
}

// Token contains a single token generated by the lexer as well as meta to make troubleshooting easier
type Token struct {
    Token string
    File string
    Line int
    Column int
    Value string
}

// readSourceFiles will scan a given folder for .li files and put the contents into a map with filename as key and contents as value
func readSourceFiles(path string) (files map[string]File, err error) {
    // TODO
}

// tokenizeSourceFiles will convert source files into tokens that are simpler to compile
func tokenizeSourceFiles(files map[string]File) (tokens []Token, errors []error) {
    // TODO
}

// compileToGo will convert tokens into a Go Programming Language source code
func compileToGo(tokens []Token) (out string, errors []error) {
    // TODO
}

/* build is the command to compile source code into a target. Initially will only support compiling into Go source code
* Args:
*  -t {targetname}          Compile to the given target
*/
func build(args []string) {
    if len(args) != 4 || args[2] != "-t" || args[3] != "\"go\"" {
        log.Println("Error: We do not currently support any parameters to the build command other than '-t \"go\"'")
        log.Println("Args provided:", args[2:])
        return
    }
    
    var files map[string]File
    if files, err := readSourceFiles("./src"); err != nil {
        log.Println("Error reading source files:")
        log.Println(err.Error())
        return
    }
    
    var tokens []Token
    if tokens, errs := tokenizeSourceFiles(files); len(errs) != 0 {
        log.Println("Parse error/(s):")
        for _, err := range errs {
            log.Println(err.Error())
        }
        return
    }
    
    var out string
    if out, errs := compileToGo(tokens); len(errs) != 0 {
        log.Println("Compile error/(s):")
        for _, err := range errs {
            log.Println(err.Error())
        }
        return
    }
    
    if err := ioutil.WriteFile("./out/go/source.go", []byte(out), 0644); err != nil {
        log.Println("Error writing output to './out/go/source.go'")
        log.Println(err.Error())
        return
    }
    
    return
}

func main() {
    // Check the command line arguments for an appropriate action to take
    
    args := os.Args
    
    if len(args) < 2 {
        log.Println("Error: Please provide a command to execute")
        return
    }
    
    if args[1] == "build" {
        build(args)
        return
    }
    
    log.Println("Error: Command not recognized:")
    log.Println("Command provided:", args[1])
    return
}
