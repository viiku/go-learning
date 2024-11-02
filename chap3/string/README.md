### Strings
A string is an immutable sequence of bytes.

Standard packages for manipulating strings:-
bytes
strings
strconv
unicode

```
s := "hello, world"
fmt.Println(len(s)) // "12"
fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
```

### Literals

A raw string literal is written `...`, using backquotes instead of double quotes. Within a raw
string literal, no escape sequences are processed; the contents are taken literally, including
backslashes and newlines, so a raw string literal may spread over several lines in the program
source.

### Strings and Byte Slices
s := "abc"
b := []byte(s)
s2 := string(b)
Conceptually, the []byte(s) conversion allocates a new byte array holding a copy of the bytes
of s, and yields a slice that references the entirety of that array.