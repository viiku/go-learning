### Defining an array

var a[n] int --> n is some integer

fmt.Println(a[0]) --> First Element
fmt.Println(a[len(a)-1]) --> Last Element

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