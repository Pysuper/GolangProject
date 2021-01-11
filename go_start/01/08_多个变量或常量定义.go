package main

import "fmt"

func main() {
	// 不同类型变量的声明(定义)
	//var a int
	//var b float64
	var (
		a int
		b float64
	)
	a, b = 10, 3.14
	fmt.Println(a, b)

	// 常量定义
	//const c int = 10
	//const d float64 = 20

	//const (
	//	c int     = 10
	//	d float64 = 20
	//)

	const (
		c = 10
		d = 20
	)
	fmt.Println(c, d)

}
