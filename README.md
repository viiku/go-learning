# go-learning
A series of go programming. 

#####
1. Keep in mind that := is a declaration, whereas = is an assignment.
2. In Go, := is used for declaring and initializing variables within a function but not for constants.


####
1. Signed Integers
Signed integers can represent both positive and negative numbers.
They use one bit (the most significant bit) to indicate the sign: 0 for positive and 1 for negative.
    In Go, signed integer types include:
    int8: Range from -128 to 127
    int16: Range from -32,768 to 32,767
    int32: Range from -2,147,483,648 to 2,147,483,647
    int64: Range from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
    int: Platform-dependent (either int32 or int64 depending on the system architecture)

2. Unsigned Integers
Unsigned integers can represent only non-negative values (0 and positive numbers).
All bits are used to store the number, allowing a larger positive range compared to signed integers of the same bit width.
In Go, unsigned integer types include:
    uint8 (or byte): Range from 0 to 255
    uint16: Range from 0 to 65,535
    uint32: Range from 0 to 4,294,967,295
    uint64: Range from 0 to 18,446,744,073,709,551,615
    uint: Platform-dependent (either uint32 or uint64 depending on the system architecture)

#### [] byte
Strings can be converted to byte slices and back again:

    s := "abc"
    b := []byte(s)
    s2 := string(b)

Conceptually, the []byte(s) conversion allocates a new byte array holding a copy of the bytes of s, and yields a slice that references the entirety of that array.

An optimizing compiler may be able to avoid the allocation and copying in some cases, but in general copying is required to ensure that the bytes of s remain unchanged even if those of b are subsequently modified.

The conversion from byte slice back to string with string(b) also makes a copy, to ensure immutability of the resulting string s2.