# Syntax Draft

## Comments

```
// Comments are markdown by default
```

## Types

-- Based on Rust types - https://doc.rust-lang.org/book/ch03-02-data-types.html

```
i8 i16 i32 i64 i128 i256 // i integers
u8 u16 u32 u64 u128 u256 // u unsigned integers
c8 c16 // c chars -- TODO
s8 s16 // s strings -- TODO
b1 b8  // b boolean -- TODO
```

## Operators

Aritmetic:

- +
- -
- /
- *
- %

Logical:

- &&
- ||
- !

Binaries:

- &
- |
- ^
- >>
- <<
- >>>

## Functions

```light
// Closures are defined with -> 
fun x y -> x + y

// Functions are defined with =>
fun sum x y =>
    x + y

fun sum x:i32 y:i32 : i32 =>
    x + y
   
// Defining new operators -- TODO
fun * x y =>
    100
```

## Generics

```light
// Closures are defined with -> 
fun x y -> x + y
fun x:a y:a :a-> x + y

// Functions are defined with =>
fun sum x y =>
    x + y
fun sum x:a y:a :a=>
    x + y

fun sum x:i32 y:i32 :i32=>
    x + y
   
// Defining new operators -- TODO
fun * x y =>
    100
fun * x:a y:a :a=>
    100
```

## Another Types

```light
 // Generic record, with the type parameter in angle brackets
type MyRecord:a =
     Field1:a
     Field2:a

 // Generic discriminated union
type MyUnion:a =
    | Choice1:a
    | Choice2:a*a
```

## Flux Control

### if

### switch

### for

### while

### do


