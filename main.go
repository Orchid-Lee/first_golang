package main

import (
	"fmt"
)

// 一、变量学习
func main() {
	var a = 10
	var c = "hello world"
	d := 100
	var e, f = 10, 12
	var (
		x int    = 10
		y string = "hello"
	)
	fmt.Printf("type of a = %T\n", a)
	fmt.Printf("type of c = %T\n", c)

	fmt.Printf("type of d = %T\n", d)
	fmt.Printf("type of e = %T, type of f = %T\n", e, f)
	fmt.Printf("type of x = %T, type of y = %T\n", x, y)
}
