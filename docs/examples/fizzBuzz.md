# Fizz Buzz

Fizz if divisible by 3, Buzz if divisible by 5

    Li 0
    
    last const = 30
    
    fizzBuzz func(n int) string:
        if n % 3 == 0 && n % 5 == 0:
            return "FizzBuzz"
        else if n % 3 == 0:
            return "Fizz"
        else if n % 5 == 0:
            return "Buzz"
        else:
            return n.string
    
    main fn:
        for i in int.range(1, last):
            console.log(fizzBuzz(i))
