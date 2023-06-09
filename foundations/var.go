package main

import "fmt"

func main() {
	// var n1 
	var _ = 1
	var _ = "a"
	var _ int
	var _ int = 2
	var _ string = "b"
	var (
		_ int
		_ string
	)
	n := 3
	a, b := "c", 4

	fmt.Println(n)
	fmt.Println(a)
	fmt.Println(b)
}