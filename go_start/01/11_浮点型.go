package main

import "fmt"

func main() {
	// 声明变量
	var a float32
	var b float64
	a = 1.11
	b = 2.22
	fmt.Println(a)
	fmt.Println(b)

	// 自动推导类型
	c := 3.33
	fmt.Println(c)
	fmt.Printf("%T", c) // 默认的推导类型为float64

	// float64存储数据比float32更准确
}
