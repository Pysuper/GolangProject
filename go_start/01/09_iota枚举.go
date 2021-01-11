package main

import "fmt"

func main() {
	// 1、iota常量自动生成器，每隔一行自动累加1
	// 2、给常量赋值使用
	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
	)
	fmt.Println(a, b, c)

	// 3、iota遇到const就会重置为0
	const d = iota
	fmt.Println(d) // 0

	// 4、可以只写一个iota
	const (
		e = iota
		f
		g
	)
	fmt.Println(e, f, g) // 0 1 2

	// 5、如果是同一行，值都一样
	const (
		h       = iota
		i, j, k = iota, iota, iota
		o       = iota
	)
	fmt.Println(h, i, j, k, o) // 0 1 1 1 2
}
