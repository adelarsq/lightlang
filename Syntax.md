# Light Syntax Reference

## Comments

Simple comments:

```light
-- And this is line comment
```

Multiline comments:

```light
{- a multiline comment based on ELM
   {- can be nested -}
-}
```

Typed comments:

```light
{-xml The `let` keyword defines an (immutable) value xml-}
{-sql select * from table sql-}
{-md select * from table md-}
```

## strings

```light
/// Create a string using string concatenation
let hello = "Hello" + " World"
```

Use *verbatim strings* preceded by `@` symbol to avoid escaping control characters (except escaping `"` by `""`).

```light
let verbatimXml = @"<book title=""Paradise Lost"">"
```

We don't even have to escape `"` with *triple-quoted strings*.

```light
let tripleXml = """<book title="Paradise Lost">"""
```

*Backslash strings* indent string contents by stripping leading spaces.

```light
let poem = 
    "The lesser world was daubed\n\
     By a colorist of modest skill\n\
     A master limned you in the finest inks\n\
     And with a fresh-cut quill."
```

[Interpolated strings](https://docs.microsoft.com/dotnet/fsharp/language-reference/interpolated-strings) let you write code in "holes" inside of a string literal:

```light
let name = "Phillip"
let age = 30
printfn $"Name: {name}, Age: {age}"

let str = $"A pair of braces: {{}}"
printfn $"Name: %s{name}, Age: %d{age}" // typed
```

References:

- [Strings](https://docs.microsoft.com/dotnet/fsharp/language-reference/strings)

## Types and Literals

- bool : The boolean type.
- char : A character type.
- i8 : The 8-bit signed integer type.
- i16 : The 16-bit signed integer type.
- i32 : The 32-bit signed integer type.
- i64 : The 64-bit signed integer type.
- isize : The pointer-sized signed integer type.
- u8 : The 8-bit unsigned integer type.
- u16 : The 16-bit unsigned integer type.
- u32 : The 32-bit unsigned integer type.
- u64 : The 64-bit unsigned integer type.
- usize : The pointer-sized unsigned integer type.
- f32 : The 32-bit floating point type.
- f64 : The 64-bit floating point type.
- array : A fixed-size array, denoted [T; N], for the element type, T, and the non-negative compile-time constant size, N.
- slice : A dynamically-sized view into a contiguous sequence, [T].
- str : String slices.
- tuple : A finite heterogeneous sequence, (T, U, ..).

`and` keyword is used for definining mutually recursive types and functions:

```light
type A = 
  | Aaa of int 
  | Aaaa of C
and C = 
  { Bbb : B }
and B() = 
  member x.Bbb = Aaa 10
```

Floating point and signed integer values in Light can have associated [units of measure](https://docs.microsoft.com/en-us/dotnet/fsharp/language-reference/units-of-measure), which are typically used to indicate length, volume, mass, and so on:

```light
[<Measure>] type kg
let m1 = 10.0<kg>
let m2 = m1 * 2.0 // type inference for result 
let add30kg m =   // type inference for input and output
    m + 30.0<kg>
add30 2.0<kg>     // val it: float<kg> = 32.0
```

References:

- [Literals](https://docs.microsoft.com/dotnet/fsharp/language-reference/literals)
- [Rust types](https://www.codingame.com/playgrounds/365/getting-started-with-rust/primitive-data-types)

# Printing Things

Print things to console with `printfn`:

```light
printfn "Hello, World"

printfn $"The time is {System.DateTime.Now}"
```

You can also use `Console.WriteLine`:
```light
open System

Console.WriteLine $"The time is {System.DateTime.Now}"
```

Constrain types with `%d`, `%s`, and print structured values with `%A`:

```light
let data = [1..10]

printfn $"The numbers %d{1} to %d{10} are %A{data}"
```

Omit holes and apply arguments:

```light
printfn "The numbers %d to %d are %A" 1 10 data
```

See [Plaintext Formatting](https://docs.microsoft.com/dotnet/fsharp/language-reference/plaintext-formatting)

<a name="Loops"></a>Loops
---------

### for...in

[For loops](https://docs.microsoft.com/dotnet/fsharp/language-reference/loops-for-in-expression):

```light
let list1 = [1; 5; 100; 450; 788]

for i in list1 do
    printf "%d" i           // 1 5 100 450 788

let seq1 = seq { for i in 1 .. 10 -> (i, i * i) }

for (a, asqr) in seq1 do
    // 1 squared is 1
    // ...
    // 10 squared is 100
    printfn "%d squared is %d" a asqr

for i in 1 .. 10 do
    printf "%d " i          // 1 2 3 4 5 6 7 8 9 10

// for i in 10 .. -1 .. 1 do
for i = 10 downto 1 do
    printf "%i " i          // 10 9 8 7 6 5 4 3 2 1

for i in 1 .. 2 .. 10 do
    printf "%d " i             // 1 3 5 7 9

for c in 'a' .. 'z' do
    printf "%c " c          // a b c ... z

// Using of a wildcard character (_)
// when the element is not needed in the loop.
let mutable count = 0

for _ in list1 do
    count <- count + 1
```

### while...do

[While loops](https://docs.microsoft.com/dotnet/fsharp/language-reference/loops-while-do-expression):

```light
let mutable mutVal = 0
while mutVal < 10 do        // while (not) test-expression do
    mutVal <- mutVal + 1
```

<a name="Functions"></a>Functions
---------

The [`let`](https://docs.microsoft.com/dotnet/fsharp/language-reference/functions/let-bindings) keyword also defines named functions.

```light
let pi () = 3.14159 // function with no arguments. () is called unit type
pi ()               // it's necessary to use () to call the function

let negate x = x * -1 
let square x = x * x 
let print x = printfn $"The number is: %d{x}"

let squareNegateThenPrint x = 
    print (negate (square x)) 
```

Double-backtick identifiers are handy to improve readability especially in unit testing:

```light
let ``square, negate, then print`` x = 
    print (negate (square x)) 
```

### Pipe operator

The pipe operator `|>` is used to chain functions and arguments together:

```light
let squareNegateThenPrint x = 
    x |> square |> negate |> print
```

This operator is essential in assisting the Light type checker by providing type information before use:

```light
let sumOfLengths (xs : string []) = 
    xs 
    |> Array.map (fun s -> s.Length)
    |> Array.sum
```

### Composition operator

The composition operator `>>` is used to compose functions:

```light
let squareNegateThenPrint = 
    square >> negate >> print
```
  
<a name="PatternMatching"></a>Pattern Matching
----------------

Pattern matching is primarily through `match` keyword;

```light
let rec fib n =
    match n with
    | 0 -> 0
    | 1 -> 1
    | _ -> fib (n - 1) + fib (n - 2)
```

Use `when` to create filters or guards on patterns:

```light
let sign x = 
    match x with
    | 0 -> 0
    | x when x < 0 -> -1
    | x -> 1
```

Pattern matching can be done directly on arguments:

```light
let fst (x, _) = x
```

or implicitly via `function` keyword:

```light
/// Similar to `fib`; using `function` for pattern matching
let rec fib2 = function
    | 0 -> 0
    | 1 -> 1
    | n -> fib2 (n - 1) + fib2 (n - 2)
```

See [Pattern Matching](https://docs.microsoft.com/en-us/dotnet/fsharp/language-reference/pattern-matching).

<a name="Collections"></a>Collections
-----------

### Lists

[*Lists*](https://docs.microsoft.com/dotnet/fsharp/language-reference/lists) are immutable collection of elements of the same type.

```light
// Lists use square brackets and `;` delimiter
let list1 = ["a"; "b"]

// :: is prepending
let list2 = "c" :: list1

// @ is concat    
let list3 = list1 @ list2   

// Recursion on list using (::) operator
let rec sum list = 
    match list with
    | [] -> 0
    | x :: xs -> x + sum xs
```

### Arrays

[*Arrays*](https://docs.microsoft.com/dotnet/fsharp/language-reference/arrays) are fixed-size, zero-based, mutable collections of consecutive data elements.

```light
// Arrays use square brackets with bar
let array1 = [| "a"; "b" |]

// Indexed access using dot
let first1 = array1.[0]   
let first2 = array1[0]    // Light 6
```
      
### Sequences == IEnumerable

[*Sequences*](https://docs.microsoft.com/dotnet/fsharp/language-reference/sequences) are logical series of elements of the same type. Individual sequence elements are computed only as required, so a sequence can provide better performance than a list in situations in which not all the elements are used.

```light
// Sequences can use yield and contain subsequences
seq {
    // "yield" adds one element
    yield 1
    yield 2

    // "yield!" adds a whole subsequence
    yield! [5..10]
}
```

The `yield` can normally be omitted:

```light
// Sequences can use yield and contain subsequences
seq {
    1
    2
    yield! [5..10]
}
```

### Mutable Dictionaries (from BCL)

Create a dictionary, add two entries, remove an entry, lookup an entry

```light
open System.Collections.Generic

let inventory = Dictionary<string, float>()

inventory.Add("Apples", 0.33)
inventory.Add("Oranges", 0.5)

inventory.Remove "Oranges"

// Read the value. If not exists - throw exception.
let bananas2 = inventory["Apples"]
```

Additional syntax:

```light
// Generic type inference with Dictionary
let inventory = Dictionary<_,_>()   // or let inventory = Dictionary()

inventory.Add("Apples", 0.33)
```

### dict == IDictionary in BCL

*dict* creates immutable dictionaries. You can’t add and remove items to it.

```light
open System.Collections.Generic

let inventory : IDictionary<string, float> =
    ["Apples", 0.33; "Oranges", 0.23; "Bananas", 0.45]
    |> dict

let bananas = inventory.["Bananas"]     // 0.45
let bananas2 = inventory["Bananas"]     // 0.45, F# 6

inventory.Add("Pineapples", 0.85)       // System.NotSupportedException
inventory.Remove("Bananas")             // System.NotSupportedException
```

Quickly creating full dictionaries:

```
[ "Apples", 10; "Bananas", 20; "Grapes", 15 ] |> dict |> Dictionary
```

### Map

*Map* is an immutable key/value lookup. Allows safely add or remove items.

```light
let inventory =
    Map ["Apples", 0.33; "Oranges", 0.23; "Bananas", 0.45]

let apples = inventory.["Apples"]
let pineapples = inventory.["Pineapples"]   // KeyNotFoundException

let newInventory =              // Creates new Map
    inventory
    |> Map.add "Pineapples" 0.87
    |> Map.remove "Apples"
```

Safely access a key in a *Map* by using *TryFind*. It returns a wrapped option:

```light
let inventory =
    Map ["Apples", 0.33; "Oranges", 0.23; "Bananas", 0.45]

inventory.TryFind "Apples"      // option = Some 0.33
inventory.TryFind "Unknown"     // option = None
```

Useful Map functions include `map`, `filter`, `partition`:

```light
let cheapFruit, expensiveFruit =
    inventory
    |> Map.partition(fun fruit cost -> cost < 0.3)
```

### Dictionaries, dict, or Map?

* Use *Map* as your default lookup type:
    * It’s immutable
    * Has good support for F# tuples and pipelining.

* Use the *dict* function
    * Quickly generate an *IDictionary* to interop with BCL code.
    * To create a full Dictionary.

* Use *Dictionary*:
    * If need a mutable dictionary.
    * Need specific performance requirements. (Example: tight loop performing
      thousands of additions or removals).

### Generating lists

The same list `[ 1; 3; 5; 7; 9 ]` can be generated in various ways.

```light
[ 1; 3; 5; 7; 9 ]
[ 1..2..9 ]
[ for i in 0..4 -> 2 * i + 1 ]
List.init 5 (fun i -> 2 * i + 1)
```

The array `[| 1; 3; 5; 7; 9 |]` can be generated similarly:

```light
[| 1; 3; 5; 7; 9 |]
[| 1..2..9 |]
[| for i in 0..4 -> 2 * i + 1 |]
Array.init 5 (fun i -> 2 * i + 1)
```

### Functions on collections

Lists and arrays have comprehensive functions for manipulation.

  - `List.map` transforms every element of the list (or array)
  - `List.iter` iterates through a list and produces side effects

These and other functions are covered below. All these operations are also available for sequences. 

<a name="TuplesAndRecords"></a>Tuples and Records
------------------
A *tuple* is a grouping of unnamed but ordered values, possibly of different types:

```light
// Tuple construction
let x = (1, "Hello")

// Triple
let y = ("one", "two", "three") 

// Tuple deconstruction / pattern
let (a', b') = x
```

The first and second elements of a tuple can be obtained using `fst`, `snd`, or pattern matching:

```light
let c' = fst (1, 2)
let d' = snd (1, 2)
  
let print' tuple =
    match tuple with
    | (a, b) -> printfn "Pair %A %A" a b
```

*Records* represent simple aggregates of named values, optionally with members:

```light
// Declare a record type
type Person = { Name : string; Age : int }

// Create a value via record expression
let paul = { Name = "Paul"; Age = 28 }

// 'Copy and update' record expression
let paulsTwin = { paul with Name = "Jim" }
```

Records can be augmented with properties and methods:

```light
type Person with
     member x.Info = (x.Name, x.Age)
```

Records are essentially sealed classes with extra topping: default immutability, structural equality, and pattern matching support.

```light
let isPaul person =
    match person with
    | { Name = "Paul" } -> true
    | _ -> false
```

<a name="Recursive Functions"></a>Recursive Functions
----------------

The `rec` keyword is used together with the `let` keyword to define a recursive function:

```light
let rec fact x =
    if x < 1 then 1
    else x * fact (x - 1)
```

*Mutually recursive* functions (those functions which call each other) are indicated by `and` keyword:

```light
let rec even x =
   if x = 0 then true 
   else odd (x - 1)

and odd x =
   if x = 0 then false
   else even (x - 1)
```

<a name="DiscriminatedUnions"></a>Discriminated Unions
--------------------

*Discriminated unions* (DU) provide support for values that can be one of a number of named cases, each possibly with different values and types.

```light
type Tree<'T> =
    | Node of Tree<'T> * 'T * Tree<'T>
    | Leaf


let rec depth input =
    match input with
    | Node(l, _, r) -> 1 + max (depth l) (depth r)
    | Leaf -> 0
```

F# Core has a few built-in discriminated unions for error handling, e.g., [Option](http://msdn.microsoft.com/en-us/library/dd233245.aspx) and [Choice](http://msdn.microsoft.com/en-us/library/ee353439.aspx).

```light
let optionPatternMatch input =
    match input with
    | Some i -> printfn "input is an int=%d" i
    | None -> printfn "input is missing"
```

Single-case discriminated unions are often used to create type-safe abstractions with pattern matching support:

```light
type OrderId = Order of string

// Create a DU value
let orderId = Order "12"

// Use pattern matching to deconstruct single-case DU
let (Order id) = orderId
```

<a name="Exceptions"></a>Exceptions
----------
The `failwith` function throws an exception of type `Exception`.

```light
let divideFailwith x y =
    if y = 0 then 
        failwith "Divisor cannot be zero." 
        else x / y
```

Exception handling is done via `try/with` expressions.

```light
let divide x y =
    try
        Some (x / y)
    with :? System.DivideByZeroException -> 
        printfn "Division by zero!"
        None
```
  
The `try/finally` expression enables you to execute clean-up code even if a block of code throws an exception. Here's an example which also defines custom exceptions.

```light
exception InnerError of string
exception OuterError of string
  
let handleErrors x y =
   try 
       try 
           if x = y then raise (InnerError("inner"))
           else raise (OuterError("outer"))
       with InnerError(str) -> 
          printfn "Error1 %s" str
   finally
       printfn "Always print this."
```

<a name="ClassesAndInheritance"></a>Classes and Inheritance
-----------------------
This example is a basic class with (1) local let bindings, (2) properties, (3) methods, and (4) static members.

```light
type Vector(x: float, y: float) =
    let mag = sqrt(x * x + y * y)               // (1) - local let binding

    member this.X = x                           // (2) property
    member this.Y = y                           // (2) property
    member this.Mag = mag                       // (2) property

    member this.Scale(s) =                       // (3) method
        Vector(x * s, y * s)

    static member (+) (a : Vector, b : Vector) = // (4) static method
        Vector(a.X + b.X, a.Y + b.Y)
```

Call a base class from a derived one:

```light
type Animal() =
    member _.Rest() = ()
           
type Dog() =
    inherit Animal()
    member _.Run() =
        base.Rest()
```

<a name="InterfacesAndObjectExpressions"></a>Interfaces and Object Expressions
---------------------------------
Declare `IVector` interface and implement it in `Vector'`.

```light
type IVector =
    abstract Scale : float -> IVector

type Vector(x, y) =
    interface IVector with
        member __.Scale(s) =
            Vector(x * s, y * s) :> IVector
            
    member __.X = x
    
    member __.Y = y
```

Another way of implementing interfaces is to use *object expressions*.

```light
type ICustomer =
    abstract Name : string
    abstract Age : int

let createCustomer name age =
    { new ICustomer with
        member __.Name = name
        member __.Age = age }
```

<a name="CastingAndConversions"></a>Casting and Conversions
---------------

```light
int 3.1415     // float to int = 3
int "3"        // string to int = 3
float 3        // int to float = 3.0
float "3.1415" // string to float = 3.1415
string 3       // int to string = "3"
string 3.1415  // float to string = "3.1415"
```

*Upcasting* is denoted by `:>` operator.

```light
let dog = Dog() 
let animal = dog :> Animal
```

In many places type inference applies upcasting automatically:

```light
let exerciseAnimal (animal: Animal) = () 

let dog = Dog()

exerciseAnimal dog   // no need to upcast dog to Animal
```

*Dynamic downcasting* (`:?>`) might throw an `InvalidCastException` if the cast doesn't succeed at runtime.

```light
let shouldBeADog = animal :?> Dog
```

<a name="ActivePatterns"></a>Active Patterns
---------------

*Complete active patterns*:

```light
let (|Even|Odd|) i = 
  if i % 2 = 0 then Even else Odd

let testNumber i =
    match i with
    | Even -> printfn "%d is even" i
    | Odd -> printfn "%d is odd" i
```

*Parameterized, partial active patterns*:

```light
let (|DivisibleBy|_|) divisor n = 
  if n % divisor = 0 then Some DivisibleBy else None

let fizzBuzz input =
    match input with
    | DivisibleBy 3 & DivisibleBy 5 -> "FizzBuzz" 
    | DivisibleBy 3 -> "Fizz" 
    | DivisibleBy 5 -> "Buzz" 
    | i -> string i
```

*Partial active patterns* share the syntax of parameterized patterns but their active recognizers accept only one argument.

## Links

- [dungpa/fsharp-cheatsheet](https://github.com/dungpa/fsharp-cheatsheet)
- [artag/fsharp-cheatsheet](https://github.com/artag/fsharp-cheatsheet)
- [thriuin/fsharp-cheatsheet](https://github.com/thriuin/fsharp-cheatsheet)

