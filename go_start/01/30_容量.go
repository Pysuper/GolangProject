/*
@Time : 2021/1/15 21:39
@Author : Zhengxingtao
@File : 06_容量.go
@Software: GoLand
*/

package main

import (
	"fmt"
)

func main() {
	// 不使用切片容量
	s := make([]int, 0)

	fmt.Println(s == nil) // 这时候说明s是存在的，不是nil
	fmt.Printf("%p\n ", s)
	fmt.Println(len(s), cap(s)) // 查看切片的长度，容量

	// 指定切片的容量
	m := make([]int, 0, 3) // 设置切片的容量为3
	fmt.Println(m == nil)
	fmt.Printf("%p\n ", m)
	fmt.Println(len(m), cap(m)) // 没有长度，但是容量已经为3了 ==> 已经申请了一个长度为3的内存空间

	fmt.Println("===============================================")

	// append() 改变切片的长度
	/*
		可以向切片中添加一个或者多个值，添加后必须使用切片接收append()函数的返回值

		如果添加一次或者添加多个值，且添加后的长度大于扩容一次的大小，容量和长度相等
		等到下次添加内容时，如果不超出扩容大小，在现有的基础上进行翻倍

		也可以把一个切片的内容直接添加到另一个切片中，注意 ... 的使用
	*/
	l := make([]string, 0)
	fmt.Println(len(l), cap(l))

	l = append(l, "张三")
	fmt.Println(len(l), cap(l))

	l = append(l, "张三")
	fmt.Println(len(l), cap(l))

	fmt.Println(l)
	l = append(l, "张三")
	fmt.Println(len(l), cap(l))
	l = append(l, "张三")
	fmt.Println(len(l), cap(l))

	i := make([]int, 0)
	ii := []int{1, 2, 3}
	//iii := []string{"4", "5", "6"}
	i = append(i, ii...)
	fmt.Println(i)

	fmt.Println(len(i), cap(i))
	//i = append(i, iii...)
	//fmt.Println(i)
}
