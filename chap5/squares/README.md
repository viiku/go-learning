### Note

The function squares returns another function, of type func() int. A call to squares creates a local variable x and returns an anonymous function that, each time it is called, increments x and returns its square. A second call to squares would create a second variable x and return a new anonymous function which increments that variable.

### Explanation
The squares example demonstrates that function values are not just code but can have state. The anonymous inner function can access and update the local variables of the enclosing function squares. These hidden variable references are why we classify functions as reference types and why function values are not comparable. Function values like these are implemented using a technique called closures, and Go programmers often use this term for function values.