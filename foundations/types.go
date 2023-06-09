package main

import "fmt"

func main() {
	var _, _ bool = true, false
	var _ int8 = -128
	var _ int16 = -32768
	var _ int32 = -2147483648
	var _ int64 = -9223372036854775808
	var _ uint8 = 255
	var _ uint16 = 65535
	var _ uint32 = 4294967295
	var _ uint64 = 18446744073709551615
	var _ float32 = 3.40282346638528859811704183484516925440e+38
	var _ float64 = 1.797693134862315708145274237317043567981e+308
	var _ complex64 = 1 + 2i
	var _ complex128 = 1 + 2i
	var _ byte = 255
	var _ rune = 2147483647
	var _ string = "Hello, World!"
	var _ error = fmt.Errorf("error")

	type Date string
	var _ Date = "2019-01-01"
}