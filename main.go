package main

import "fmt"

// 返回单个值
func fool(a string, b int) string {
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	return a
}

// 返回多个值
func foo2(a string, b int) (string, int) {
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	return a, b
}

// 返回多个带名称的
func foo3(a string, b int) (r1 string, r2 int) {
	r1 = "1000"
	r2 = 2000
	return
}

func foo4() (r1, r2 int) {
	r1 = 2
	r2 = 4
	return
}

// 二、函数
func main() {
	c := fool("hello", 666)
	fmt.Println("c: ", c)

	foo2("world", 888)
	foo3("122", 122)

	x, y := foo4()
	fmt.Printf("x: %v\ty: %v\n", x, y)
}
