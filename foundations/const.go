package main

import "fmt"

func main() {
	const a = "a"
	const b = 1
	const (
		c = "c"
		d = 2
	)
	const e, f = "e", 3
	const g string = "g"

	fmt.Println(a, b, c, d, e, f, g)
}