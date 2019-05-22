# Fizz Buzz

Fizz if divisible by 3, Buzz if divisible by 5

    Li 0
    
    import:
        fmt
        num
    
    fizzBuzz fn(n int) string:
        if n % 3 == 0 && n % 5 == 0:
            return "FizzBuzz"
        else if n % 3 == 0:
            return "Fizz"
        else if n % 5 == 0:
            return "Buzz"
        else:
            return n.String
    
    main fn():
        for n int in num.range(30):
            fmt.printLn(fizzBuzz(n))
