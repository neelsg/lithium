# Go Programming Language translated code

This is a number of code snippets from [the Go tour](https://tour.golang.org/) translated into Lithium

## Snippet 0

Go:

    package main

    import "fmt"

    func main() {
      fmt.Println("Hello, 世界")
    }

Lithium:

    main func:
        console.log("Hello, 世界")

## Snippet 1

Go:

    package main

    import (
        "fmt"
        "time"
    )
    
    func main() {
        fmt.Println("Welcome to the playground!")
        
        fmt.Println("The time is", time.Now())
    }

Lithium:

    main func:
        console.log("Welcome to the playground!")
        
        console.log("The time is", time.now())

## Snippet 2

Go:

    package main
    
    import (
        "fmt"
        "math/rand"
    )
    
    func main() {
        fmt.Println("My favorite number is", rand.Intn(10))
    }

Lithium:

    main func:
        console.log("My favorite number is", int.random(10))

## Snippet 3

Go:

    package main
    
    import (
        "fmt"
        "math"
    )
    
    func main() {
        fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
    }

Lithium (Option A):

    main func:
        console.log("Now you have \% problems.\n".inset(math.sqrt(7)))

Lithium (Option B):

    main func:
        console.log("Now you have \(math.sqrt(7)) problems.\n")

## Snippet 4

Go:

    package main
    
    import (
        "fmt"
        "math"
    )
    
    func main() {
        fmt.Println(math.Pi)
    }

Lithium:

    main func:
        console.log(math.pi)
