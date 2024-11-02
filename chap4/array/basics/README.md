### Defining an array

var a[n] int --> n is some integer

fmt.Println(a[0]) --> First Element
fmt.Println(a[len(a)-1]) --> Last Element

### Array Literal
We can use an array literal to initialize an array with a list of values:
```
var q [3]int = [3]int{1, 2, 3}
var r [3]int = [3]int{1, 2}
fmt.Println(r[2]) // "0"
```
In an array literal, if an ellipsis ‘‘...’’ appears in place of the length, the array length is determined by the number of initializers. The definition of q can be simplified to
```
q := [...]int{1, 2, 3}
fmt.Printf("%T\n", q) // "[3]int"
```
#### Printing Array Elements
// Print the indices and elements.
```
for i,v := range a {
    fmt.Printf("%d %d\n", i, v)
}
```
// Print the elements only.
```
for _,v := range a {
    fmt.Printf("%d\n", v)
}
```
// By default, the elements of a new array variable are initially set to the zero value for the element type, which is 0 for numbers.