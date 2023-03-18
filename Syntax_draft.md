# Syntax Draft

## Types

-- Based on Rust types - https://doc.rust-lang.org/book/ch03-02-data-types.html

i8 i16 i32 i64 i128 i256 // i integers
u8 u16 u32 u64 u128 u256 // u unsigned integers
c8 c16 // c chars -- TODO
s8 s16 // s strings -- TODO
b1 b8  // b boolean -- TODO

## Operators

+
-
/
*

## Functions

// Closures are defined with -> 
fun x y -> x + y

// Functions are defined with =>
fun sum x y =>
    x + y

fun sum x:i32 y:i32 : int =>
    x + y
   
// Defining new operators -- TODO
fun * x y =>
    100


