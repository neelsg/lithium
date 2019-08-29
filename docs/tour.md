# Lithium Programming Language Tour

## Welcome

This tour aims to cover the basics of the Lithium Programming Language.


## Start

    # This is the main function
    main func:
        console.log("My favorite number is", int.rand(10))

This code will print the text "My favorite number is" and a random number from 0 to 10.

Anything after a `#` on a line is a comment and is ignored by the compiler.

The `func` keyword is used to define functions (which are blocks of code that can be executed. When a program is executed, the first thing that will run is the "main" function.


## Exporting components

    multiply func(x, y int) int:
        # This function is only available within the current package
        return x * y
    
    sqr pub func(x int) int:
        # This function can be used by other packages
        return multiply(x, x)

Use the `pub` keyword to make your components available to other packages


## Importing packages

    import myMath "./myMathFolder"
    
    main func:
        console.log(myMath.sqr(3))

You can import your own packages or third-party libraries using the `import` keyword.


## Functions

    add func(x, y int) int:
        return x + y
    
    main func:
        console.log(add(42, 13))

A function can take zero or more arguments. In this example, `add` takes two arguments of type `int`. Note that the type comes after the variable name. If two or more consecutive arguments are of the same type, you only need to specify the type for the last argument.

A function that takes no arguments can optionally be declared without brackets:


## Multiple results

    split func(sum int) (int, int):
        x int = sum * 4 / 9
        y int = sum - x
        return x, y

    swap func(x, y string) (string, string): return y, x

    main func:
        console.log(split(17))
        
        a, b string = swap("hello", "world")
        console.log(a, b)

A function can return any number of results. If they do, these must be wrapped in brackets `( )`. Even if a function returns only one result, the result type can optionally be wrapped in brackets.


## Variables

    i, j, k bool
    
    l pub bool
    
    main func:
        m int
        console.log(i, j, k, l, m)

Variables are declared by specifying the name of the variable followed by the type.


## Initializers

    i, j int, p float = 1, 2, math.pi
    
    main func:
        k, l bool, m, n string = true, false, "no!", "yes!"
        console.log(i, j, p, k, l, m, n)

A variable declaration can include initializers. If an initializer is omitted, the variable will be assigned the default value of the type.


## Basic types

    toBe bool = false
    maxInt int.u64 = math.power[int.u64](2, 64) - 1
    z complex.p128 = math.sqrt[complex.p128](-5 + 12i)
    
    main func:
        console.log("Value: \(toBe)")
        console.log("Value: \(maxInt)")
        console.log("Value: \(z)")

Lithium basic types are:

    bool int string float complex

Each of these types are defined in a built-in library package.

There are also more specific data types such as:

    int.s8 int.s16 int.s32 int.s64
    int.u int.u8 int.u16 int.u32 int.u64
    
    float.p32 float.p64
    
    complex.p64 complex.p128


## Default values

    main func:
        i int
        f float
        b bool
        s string
        console.log("\(i) , \(f) , \(b) , \(s)")

Variables declared without an explicit initial value are given their default value.

The default values for built in types are:

    0 for numeric types
    false for the boolean type
    "" for strings.

Custom types can have their own defined default values.


## Type conversions

    main func:
        x, y int = 3, 4
        f float = math.sqrt((x * x + y * y).toFloat)
        z int.u = f.toUInt
        console.log(x, y, z)

The built in types have properties to convert to different types. Custom defined types should also include similar functionality.

Some numeric conversions:

    i int = 42
    f float.p64 = i.toFloat64
    u int.u = f.toUInt

Unlike in C, assignment in Lithium between items of different type requires an explicit conversion.


## Constants

    pi const = 3.14
    
    main func = func():
        world const = "Mars"
        console.log("Hello", world)
        console.log("Happy", pi, "Day")
        
        truth const = true
        console.log("Humans rock?", truth)

Constants are declared like variables, but with the `const` keyword as type. Constants can be string, boolean, or numeric values. Constants cannot be declared with an expression that requires runtime evaluation such as calling a function.


## For

    main func:
        sum int = 0
        for i int = 0; i < 10; i++:
            sum += i
        console.log(sum)

Lithium has only one looping construct, the `for` loop.

The basic for loop has three components separated by semicolons:

    The init statement: executed before the first iteration
    The condition expression: evaluated before every iteration
    The post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to `false`.


## For continued

    main func:
        sum int = 1
        for ; sum < 1000; :
            sum += sum
        console.log(sum)

The init and post statements are optional. At that point you can drop the semicolons:

    main func:
        sum int = 1
        for sum < 1000:
            sum += sum
        console.log(sum)


## Forever

    main func:
        for:
            # Whatever code goes here will run in an infinite loop

If you omit the loop condition it loops forever.


## If

    sqrt func(x float) string:
        if x < 0:
            return sqrt(-x).toString + "i"
        return math.sqrt(x).toString
    
    main func:
        console.log(sqrt(2), sqrt(-4))

Lithium's `if` statements are similar to its `for` loops


## If with an init statement

    pow func(x, n, lim float) float:
        if v float = math.power(x, n); v < lim:
            return v
        return lim
    
    main func:
        console.log(
            pow(3, 2, 10),
            pow(3, 3, 20),
            )

Like `for`, the `if` statement can start with a short statement to execute before the condition. Variables declared by the statement are only in scope until the end of the `if`.


## If and else

    pow func(x, n, lim float) float:
        if y = math.power(x, n); y < lim:
            return y
        else:
            console.log("\(y) >= \(lim)")
        return lim
    
    main func:
        console.log(pow(3, 2, 10), pow(3, 3, 20), )

Variables declared inside an `if` init statement are also available inside any of the else blocks.


## Function values

    compute func(f func(float, float) float) float:
        return f(3, 4)
    
    main func:
        hypot func(x, y float) float:
            return math.sqrt(x * x + y * y)
        ave func(x, y float) float:
            return (x + y) / 2
        
        console.log(hypot(5, 12))
        
        console.log(compute(hypot))
        console.log(compute(ave))

Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.


## Defer

    main func:
        defer console.log("world!")
        
        console.log("Hello")

The `defer` statement defers the execution of a function until the surrounding function returns.

The `defer` statement's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.


## Stacking defers

    main func:
        console.log("counting")
        
        for i int = 0; i < 10; i++:
            defer console.log(i)
        
        console.log("done")

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.


## Function closures

    adder func() (func(int) int):
        sum int = 0
        return func(x int) int:
            sum += x
            return sum
    
    main func:
        pos int, neg int = adder(), adder()
        for i int; i < 10; i++:
            console.log(
                pos(i),
                neg(-2 * i),
            )

Lithium functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables. In this sense the function is "bound" to the variables.

For example, the `adder` function returns a closure. Each closure is bound to its own `sum` variable.


## Defining new types

    vertex type:
        x pub int
        y pub int
    
    main func:
        v vertex
        v.x, v.y = 1, 2
        console.log(v)

Lithium allows simple object-oriented programming. A type is a collection of properties and methods and is often called a class in other languages. Note the `pub` keyword in front of the property names that make these properties available for use outside of the type itself.


## Methods

    vertex type:
        x, y pub float
        dist pub func() float:
                return math.sqrt(x * x + y * y)
    
    main func:
        v vertex
        v.x, v.y = 1, 2
        console.log(v.dist())

Methods are functions defined within a type.


## The setter function

    vertex type:
        
        x, y pub float
        
        .= func(xInit, yInit float):
            x, y = xInit, yInit
            return
        
        dist pub func() float:
            return math.sqrt(x * x + y * y)
    
    main func:
        v vertex = (1, 2)
        console.log(v.dist())

You can create a function with the label `.=` in a type which will act as the initializer for the type. The initializer is activated either by calling the type as a function in an expression such as `vertex(1, 2)` or by assigning to a variable of the specified type as in `v vertex = (1, 2)`.


## Private values

    vertex type:
        
        x, y float
        
        .= func(xInit, yInit float):
            x, y = xInit, yInit
            return
        
        dist pub func() float:
            return math.sqrt(x * x + y * y)
    
    main func:
        v vertex = (1, 2)
        console.log(v.dist())

If you omit the `pub` keyword for properties or methods, they are only
available to the methods in the object itself. If you try to use `v.x` in the
main function in this case, you will get a compile error.


## Getters and setters

    vertex type:
        .= func(xInit, yInit float):
            x, y = xInit, yInit # This will call the respective set functions
            return
        
        x float
        x get func() float:
            return x # This will get the underlying variable only because it is used inside the get method
        x set func(xVal float):
            if xVal >= 0:
                x = xVal # This will set the underlying variable only because it is used inside the set method
            else:
                x = -xVal
        
        y float
        y get float:
            return y
        y set func(yVal float):
            if yVal >= 0:
                y = yVal
            else:
                y = -yVal
        
        dist pub func() float:
            return math.sqrt(x * x + y * y)
    
    main func:
        v vertex = (1, 2)
        console.log(v.dist())

Properties can have a `get` and `set` method which allows for more control over what happens when you retrieve or assign a value to that property. These methods are public by definition.

If both a `get` method for a property as well as the actual property is defined, the `get` method will take preference whenever that method is used outside of the actual code run by the `get` method itself. Similarly, the `set` method also overrules an actual property in the type if such actual property is defined.


## Inheritance

    main func:
        myInt type(i int):
            double func() myInt:
                return i * 2
        
        i myInt = 3
        console.log(i.double().double()) # This will print 12

The `type` keyword can take parameters to specify which other type/(s) it should inherit from which allows you to create new types by extending existing ones.


## Interfaces

    hasX interface:
        x get float
    
    vertex type:
        x, y pub float
        .= func(xInit, yInit float):
            x, y = xInit, yInit
            return
    
    printX func(o hasX):
        console.log(o.x)
    
    main func:
        v vertex = (3, 4)
        printX(v) # This will print 3

Interfaces specify which properties and methods a type must have to satisfy it. You can use an interface as a parameter in a function and ensure that any object that is passed to that function must satisfy that interface. If you try to pass an object that does not satisfy the interface or try to use the parameter in the function in a way that is not allowed by the interface, you will get a compile error.

Note that a type does not need to specify that it implements the interface, all it needs to do is to actually conform to it. This allows for quite a bit of flexibility in using interfaces.


## Stringers

    person type:
        name pub string
        age pub int
        .= func(nameInit string, ageInit int):
            name, age = nameInit, ageInit
            return
        toString get string:
            return "\(name) (\(age) years)"
    
    main func:
        a person("Arthur Dent", 42)
        z person("Zaphod Beeblebrox", 9001)
        console.log(a, z)

One of the most ubiquitous interfaces is `any.string`.

    string pub interface:
        get toString string

Many packages look for this interface to print values.


## Errors

    myError type:
        .= func(when time, what string):
            at, msg = when, what
            return
        
        at time
        msg string
        
        ok pub bool = false
        class pub string = "MyError"
        
        toString get string:
            return "at \(at), \(msg)"
    
    run func() error:
        return myError(time.now(), "it didn't work")
    
    main func:
        if err error = run(); !err.ok:
            console.log(err)

Lithium programs express error state with `error` values.

The `error` type is a built-in interface:

    . interface:
        get toString string
        get ok bool

The name `.` as used for this interface is special in that it indicates that this is the default item in this package, so you can access it by simply using the package name which in this case is `error`.

Functions often return an `error` value, and calling code should handle errors by testing whether the `error.ok` is false.

    i int
    if i, err error = "42".toInt; !err.ok:
        console.log("couldn't convert number: \(err)")
        return
    console.log("Converted integer:", i)


## Parametric polymorphism

    plusable interface[p]:
        .+(p) p
    
    double func[p plusable](num p) p:
        return num + num
    
    main func:
        myInt int = 3
        myInt = double[int](myInt)
        console.log(myInt) # This will print 6
        
        myFloat float = 3.1
        myFloat = double[float](myFloat)
        console.log(myFloat) # This will print 6.2

This is often called generics in other languages. What is happening here is that we can pass any type of variable to `double()` as long as it implements the `plusable` interface.

The advantage over just using interfaces is that we can make sure that the variables retain their original type, so in the first case, `double()` actually returns an `int` and in the second, `double()` actually returns a `float`. If we used only interfaces, the return value would only allow properties and methods of the interface, so we would not be able to print it as `plusable` does not specify a `toString` property, but `float` does.


## Default types for polymorphism

    plusable interface[n]:
        .+(n) n
    
    double func[p plusable int](num p) p:
        return num + num
    
    main func:
        myInt int = 3
        myInt = double(myInt)
        console.log(myInt) # This will print 6
        
        myFloat float = 3.1
        myFloat = double[float](myFloat)
        console.log(myFloat) # This will print 6.2

A default type can be set for any polymorphic function. If the type is then omitted when calling the function, this default type is assumed.


## Parametric polymorphism for types

    vertex type[n any.string]:
        .= func(xInit, yInit n):
            x, y = xInit, yInit
            return
        pub x n
        pub y n
    
    main func:
        v vertex[int] = (1, 2)
        console.log(v.x)

Parametric polymorphism can also be used with type definitions to create more generic types. In the above example, `x` and `y` are of type `int`.


## Tuples

    main func:
        triplets (string, string, string) = "Anna", "Betty", "Carmen"
        console.log(triplets)

Tuples are ordered pairs of variables. You have actually already seen them used for multiple return values for functions as well as with assignments on multiple variables. Tuples defined as a set of types wrapped in brackets.

    main func:
        triplets (string, string, string) = "Anna", "Betty", "Carmen"
        console.log(triplets.0) # This prints Anna
        console.log(triplets.1) # This prints Betty
        console.log(triplets.2) # This prints Carmen

The individual values in a tuple can be accessed using the index number starting from zero. The index number must be a literal number, it cannot be another variable or calculated in some way.


## Lists

    main func:
        l list[string]
        l(0) = "Hello"
        l(1) = "World"
        console.log(l(0), l(1))
        console.log(l)
        
        primes list[int](2, 3, 5, 7, 11, 13)
        console.log(primes)

`list[t]` is a built-in type that stores a list or sometimes called an array of variables of type `t`.

The statement `l list[string]` declares a new list `l` of strings. The keys in a list are always of type `int`.


## List slices

    main func:
        abcd list[string] = ("a", "b", "c", "d")
        bc list[string] = abcd.slice(1, 3)
        console.log(bc)

A slice is formed by specifying two indexes, a low and high bound, separated by a comma: `list.slice(low, high int)`.

This selects a half-open range which includes the first element, but excludes the last one.


## List slice defaults

    main func:
        p list[int](2, 3, 5, 7, 11, 13)
        
        p = p.slice(1, 4)
        console.log(p)
        
        p = p.slice(, 2)
        console.log(p)
        
        p = p.slice(1, )
        console.log(p)

When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the list for the high bound.


## List length

    main func:
        p list[int](2, 3, 5, 7, 11, 13)
        console.log(p.length)

Lists have a length property that shows how many items are currently in the list.


## Lists of lists

    main func:
        # Create a tic-tac-toe board.
        board list[list[string]] = (
            ("_", "_", "_"),
            ("_", "_", "_"),
            ("_", "_", "_"),
        )
        
        # The players take turns.
        board(0)(0) = "X"
        board(2)(2) = "O"
        board(1)(2) = "X"
        board(1)(0) = "O"
        board(0)(2) = "X"
        
        for i int = 0; i < board.length; i++:
            console.log(board(i).join(" "))

Lists can contain any type, including other lists.


## Appending to a list

    main func:
        l list[int]
        console.log(l)
        
        l.push(3)
        console.log(l)
        
        l.push(4)
        console.log(l)
        
        # We can add more than one element at a time.
        l.push(5, 6, 7)
        console.log(l)

Lists have a `push` method to append items to it.


## Iterating over a list

    pow list[int] = (1, 2, 4, 8, 16, 32, 64, 128)
    
    main func:
        for i, v int in pow:
            console.log("2 ** \(i) = \(v)")

Lists are iterable, which means that they can be used in a for loop with the `in` keyword.


## Iterating continued

    main func:
        pow list[int] = (1, 2, 4, 8, 16, 32, 64, 128)
        
        for _, value int in pow:
            console.log(value)

You can skip the key or value by assigning to `_`.

    for key, _ in collection
    for _, value in collection

If only the value is required, this can be further shortened to:

    for value in collection


## Maps

    main func:
        m map[int, string]
        m(0) = "Hello"
        m(1) = "World"
        console.log(m(0), m(1))
        console.log(m)
        
        primes map[int, int] = ((0,2), (1,3), (2,5), (3,7), (4,11), (5,13))
        console.log(primes)

`map[k, v]` is a data type that stores a map with keys of type `k` that maps to values of type `v`.

The statement `m map[int, string]` declares a variable `m` as a map with `int` keys and `string` values.


## Map initializer with custom keys

    vertex type:
        lat, long float
    
    m map[string, vertex] = (
        ("Bell Labs", (40.68433, -74.39967)),
        ("Google", (37.42202, -122.08408)),
    )
    
    main func:
        console.log(m)

The map initializer accepts an arbitrary number of tuples that consist of the (key, value) pairs that need to be added.


## Mutating maps

    main func:
        m map[string, int]
        
        m("Answer") = 42
        console.log("The value: ", m("Answer"))
        
        m("Answer") = 48
        console.log("The value: ", m("Answer"))
        
        m.delete("Answer")
        console.log("The value: ", m("Answer"))
        
        ok bool = m.exists("Answer")
        console.log("Present? ", ok)
        
        v int, ok bool = m.get("Answer")
        console.log("The value: ", v, "Present? ", ok)

Insert or update an element in map `m`:

    map(key) = value

Retrieve an element:

    val = map(key)

If `key` is not in the map, then `val` is the zero value for the map's value
type.

Delete an element:

    map.delete(key)

Test that a key is present:

    ok = map.exists(key)

If `key` is in the map, `ok` is `true`. If not, `ok` is `false`.

You can use `val, ok = map.get(key)` as a shorthand for
`val, ok = map(key), map.exists(key)`.


## Variadic functions

    sum func(...nums int) int:
        s int
        for n int in nums:
            s += n
        return s
    
    main func:
        console.log(sum(1, 2, 3, 4, 5)) # This will print 15

Variadic functions are functions that can take an arbitrary number of parameters instead of a fixed set. The parameters are passed into the initializer method of a list.

    number interface[n]:
        plus(n) n
        power(n) n
    
    sumPower func[num number](...pairs (num, num)) num:
        sum num
        for (base, exp num) in pairs:
            sum += base ** exp
        return sum
    
    main func:
        console.log(sumPower[int]((2, 1), (2, 2), (2, 3), (2, 4))) # Output 30

Variadic functions can take any type of parameter that can be as a value in a list including tuples.


## Spread operator

    sum func(...nums int) int:
        s int
        for n int in nums:
            s += n
        return s
    
    main func:
        numbers array[int](1, 2, 3, 4, 5)
        console.log(sum(numbers...))

It is possible to use the spread operator (`...`) to pass the values in a collection as distinct parameters in a variadic function.


## Coroutines

    say func(s string):
        for i int = 0; i < 5; i++:
            time.sleep(100 * time.milliSecond)
            console.log(s)
    
    main func:
        co say("world")
        say("hello")

A coroutine is a lightweight thread managed by the Lithium runtime.

`co f(x, y, z)` starts a new coroutine running `f(x, y, z)`.

The evaluation of the passed parameters happen in the current coroutine, but the execution of the function happens in the new coroutine.

Coroutines run in the same address space, so access to shared memory must be synchronized. The `sync` package provides useful primitives, although you won't need them much in Lithium as there are other ways to handle concurrency as you will see shortly.

    say func(s string):
        for i int = 0; i < 5; i++:
            time.sleep(100 * time.milliSecond)
            console.log(s)
    
    main func:
        co:
            for i int in int.range(5):
                time.sleep(100 * time.milliSecond)
                console.log("world")
        say("hello")

It is also possible to use `co` with a block of code to be executed instead of calling a function.


## Channels

    sum func(s list[int], c chan[int]):
        sum int = 0
        for v int in s:
            sum += v
        c.push(sum) # Send sum to c
    
    main func:
        s array[int] = (7, 2, 8, -9, 4, 0)
        
        c chan[int]
        co sum(s.slice(, s.length / 2), c)
        co sum(s.slice(s.length / 2, ), c)
        x, y int = c.pop(), c.pop() # receive from c
        
        console.log(x, y, x + y)

Channels are a typed conduit through which you can send values with `push()` and receive values with `pop()`

    ch.push(v)     # Send v to channel ch.
    v = ch.pop()   # Receive from ch, and assign value to v.

By default, sends and receives block until the other side is ready. This allows coroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in an array, distributing the work between two coroutines. Once both coroutines have completed their computation, it calculates the final result.


## Buffered Channels

    main func:
        ch chan[int] = 2
        ch.push(1)
        ch.push(2)
        console.log(ch.pop())
        console.log(ch.pop())

Channels can be buffered. Provide the buffer length to the channel initializer to make a buffered channel:

    ch chan[int] = 100

Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.


## Closing a channel

    fibonacci func(n int, c chan[int]):
        x int, y int = 0, 1
        for i int = 0; i < n; i++:
            c.push(x)
            x, y = y, x + y
            c.close()
    
    main func:
        c chan[int] = 10
        co fibonacci(c.capacity, c)
        for i int in c:
            console.log(i)

A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by checking the `ok` property.

    if ch.ok: v int = ch.pop()

`ok` is `false` if there are no more values to receive and the channel is closed.

The loop `for i int in c` receives values from the channel repeatedly until it is closed.

Note: Only the sender should close a channel, never the receiver.

Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a loop.


## Checking multiple channels

    fibonacci func(c, quit chan[int]):
        x int, y int = 0, 1
        for:
            if c.capacity > 0: # This checks if you can push new values
                c.push(x)
                x, y = y, x+y
            else if quit.length > 0: # This checks if any values have been pushed
                console.log("quit")
                return
    
    main func:
        c, quit chan[int]
        co:
            for i int in int.range(10):
                console.log(c.pop())
            quit.push(0)
        fibonacci(c, quit)

The `for: if` pattern lets a coroutine wait on multiple communication operations. This will wait until one of its conditions are true, then it executes that block.


## More conditionals with channels

    main func:
        tick chan[bool] = time.tick(100 * time.millisecond)
        boom chan[bool] = time.after(500 * time.millisecond)
        for:
            if tick.length > 0:
                tick.pop()
                console.log("tick.")
            else if boom.length > 0:
                boom.pop()
                console.log("BOOM!")
                return
            else:
                console.log(".")
                time.sleep(25 * time.millisecond)

An `else` block can be used to run if no other condition is ready.


## Mutual exclusion

    # SafeCounter is safe to use concurrently.
    safeCounter type:
        counts map[string, int]
        mux sync.mutex
        
        # inc increments the counter for the given key.
        inc func(key string):
            mux.lock()
            # Lock so only one coroutine at a time can access the map.
            counts(key)++
            mux.unlock()
        
        # value returns the current value of the counter for the given key.
        value func(key string) int:
            mux.lock()
            # Lock so only one coroutine at a time can access the map.
            defer mux.unlock()
            return counts(key)
    
    main func:
        c safeCounter
        for i int = 0; i < 1000; i++:
            co c.inc("somekey")
        
        time.sleep(time.second)
        console.log(c.value("somekey"))

We've seen how channels are great for communication among coroutines.

But what if we don't need communication? What if we just want to make sure only one coroutine can access a variable at a time to avoid conflicts?

This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex.

The mutual exclusion type `sync.mutex` has two methods:

    lock()
    unlock()

We can define a block of code to be executed in mutual exclusion by surrounding it with a call to `lock` and `unlock` as shown on the `inc` method.

We can also use `defer` to ensure the mutex will be unlocked as in the `value` method.
