/* This is command line tool that can be used to parse lithium source files into go source files */

import "log"

func build(args []string) {
    // Compile into a target. Currently only supports compiling into Go source code
    
    if len(args) != 4 || args[2] != "-t" || args[3] != "\"go\"" {
        log.Println("ERROR: We do not currently support any parameters to the build command other than '-t \"go\"'")
        log.Println("Args provided:", args[2:])
        return
    }
    
    // TODO: Scan the current folder for .li files
    
    // TODO: Convert .li files into tokens
    
    // TODO: Create output file with compiled source
    
}

func main() {
    // Check the command line arguments for an appropriate action to take
    
    args := os.Args
    
    if len(args) < 2 {
        log.Println("ERROR: Please provide a command to execute")
        return
    }
    
    if args[1] == "build" {
        build(args)
        return
    }
    
    log.Println("ERROR: Command not recognized:")
    log.Println("Command provided:", args[1])
    return
}
