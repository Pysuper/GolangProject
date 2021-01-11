package main

import "fmt"

func main() {
	a := 10

	// 一段一段输出，自动加换行
	fmt.Println(a)

	// 格式化输出，把变量的内容放到占位符那里，不会添加换行
	fmt.Printf("%d\n", a)

	//  用来处理占位
	b, c := 10, 20
	fmt.Printf("a=%d, b=%d,c=%d", a, b, c)
}
