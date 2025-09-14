package main

import "fmt"

// 定义枚举
const (
	BEIJING = iota
	SHANGHAI
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

// 二、常量
func main() {
	const len = 10
	fmt.Printf("type of len: %T\n", len)

	fmt.Printf("address: %d\n", SHANGHAI)

	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
}
