package main

import (
	"fmt"
	//"constraints"
)

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func min2[T int32] (x, y T) T {
	fmt.Printf("%T %T\n", x, y)
	if x < y {
		return x
	}
	return y
}

type AnyString interface{
	~string
}
type MyString string


func main() {
	f := min2(1, 2)
	fmt.Println(f)
}