<!--  -->
Go lets us declare a field with a type but no name; such fields are called anonymous fields. The
type of the field must be a named type or a pointer to a named type. Below, Circle and Wheel
have one anonymous field each. We say that a Point is embedded within Circle, and a
Circle is embedded within Wheel.

```
type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}
```

Example:- 
```
package main

import "fmt"

// type Circle struct {
// 	X, Y, Radius int
// }

// type Wheel struct {
// 	X, Y, Radius, Spokes int
// }

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var w Wheel
	// w.Circle.Center.X = 8
	// w.Circle.Center.Y = 8
	// w.Circle.Radius = 5
	// w.Spokes = 100

	// fmt.Println(w)

	// new equivalent
	w.X = 7      // equivalent to w.Circle.Point.X = 7
	w.Y = 7      // equivalent to w.Circle.Point.X = 7
	w.Radius = 4 // equivalent to w.Circle.Radius = 4
	w.Spokes = 99
	fmt.Println(w)
}
```