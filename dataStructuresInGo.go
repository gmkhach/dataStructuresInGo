package main

import (
    "fmt"
)

func main() {
	// Here is how you can declare and use arrays in Go.
	// Keep in mind that the length of an array is part of its type!
    var x [5]int
	fmt.Printf("array: %v]\tlength: %v\tcapacity: %v\n", x, len(x), cap(x))
	x[3] = 42
	fmt.Printf("after assignment: %v\n", x)
	fmt.Printf("length: %v\n", len(x))

	// However, we don't really use arrays in Go. Instead we use slices.
	// Here is a definition of a composite literal, which create an instance of an expresion each time it is evaluated.
	// Composite literals are used to construct values for arrays, slices, maps, and structs.
	y := []int{4, 5, 6, 7, 8, 42}
	fmt.Printf("slice: %v\tlength: %v\tcapacity: %v\n", y, len(y), cap(y))
	fmt.Printf("length: %v\n", len(y))
	fmt.Println(y[0])
	fmt.Println(y[1])
	fmt.Println(y[2])
	fmt.Println(y[3])
	fmt.Println(y[4])
	fmt.Println(y[5])
	
	for i, v := range y {
		fmt.Printf("index: %v\tvalue: %v\n", i, v)
	}

	// We can slice slices using the : operator, like this: slice[a:b]
	// The way to read the : operator is "from a (included) up to b (not included)".
	fmt.Println(y[1:3])

	for i := 0; i < len(y); i++ {
		fmt.Printf("index: %v\tslice: %v\n", i, y[i:i+1])
	}

	// Now let's append a some values to our slice
	y = append(y, 55)
	
	// Notice how the capacity of the underlying array doubles when using append.
	fmt.Printf("slice: %v\tlength: %v\tcapacity: %v\n", y, len(y), cap(y))
	
	y = append(y, 77, 88, 99, 1014)

	z := []int{ 234, 456, 678 }
	fmt.Println(append(y, z...))

	// So, to remove an element from a slice we usethe : operator in combination with the append function.
	y = append(y[0:3], y[5:]...)
	fmt.Println(y)

	// If you know the ultimate size (capacity) of your slice you can use the make built in function to create the underlying array.
	a := make ([]int, 10, 100)
	fmt.Printf("slice: %v\tlength: %v\tcapacity: %v\n", a, len(a), cap(a))

	// Multidimensional slices
	b := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(b)
	c := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(c)
	d := [][]string{b, c}
	fmt.Println(d)

	// Maps are unordere, key-value lists
	m := map[string]int{
		"James": 32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Printf("James is %v years old\n", m["James"])

	// If you try to access a map through a key which doesn't exist the value returned will always be the zero value of the value type
	fmt.Println(m["Gabriel"])

	// to checkif there exists an entry for any key in a map use the "comma ok" idiom
	v, ok := m["Gabriel"]
	fmt.Printf("value: %v\tok: %v\n", v, ok)

	// This type of checks can be used in a conditional logic like this
	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Printf("THIS IS THE IF PRINT\nvalue: %v\tok: %v\n", v, ok)
	}

	// Adding another key-value pair to a map
	m["Gabriel"] = 32
	for k, v := range m{
		fmt.Printf("key: %v\tvalue: %v\n", k, v)
	}
}