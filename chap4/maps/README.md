# Maps in GO

### Imp. Points
1. In Go, a map is a reference to a hash table, and a map type is written map[K]V, where K and V are the types of its keys and values.
2. All of the keys in a given map are of the same type, and all of the values are of the same type, but the keys need not be of the same type as the values.

3. Empty maps
An alternative expression for a new empty map is...

```
map[string]int{}
```
#### Examples.
1. Creating a map from strings
```
ages := make(map[string]int) // mapping from strings to ints
```

