# Fizz Buzz

Fizz if divisible by 3, Buzz if divisible by 5

    Li 0
    
    import:
        console
    
    LAST const = 30
    
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
        for i int = 1; i <= LAST; i++:
            console.log(fizzBuzz(i))
