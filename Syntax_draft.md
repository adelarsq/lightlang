# Syntax Draft

## Types

i8 i16 i32 i64 i128 i256 // i integers
c8 c16 // c chars
s8 s16 // s strings
b1 b8  // b boolean

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


