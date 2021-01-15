/*
@Time : 2021/1/15 22:04
@Author : Zhengxingtao
@File : 07_数组产生切片.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	num := [5]int{1, 2, 3, 4, 5}
	fmt.Println(num)

	num2 := num[0:2]         // 数组取片段的时候，包左 不包右
	fmt.Printf("%T\n", num2) // []int ==> 没有数字 ==> 切片
	fmt.Printf("%T\n", num)  // [5]int ==> 有数字 ==> 数组

	fmt.Printf("%p, %p\n", num2, &num[0])

	// 切片发生改变的时候，数组也会跟着发生变化
	// 切片：指针，指向内存中数组某一片段的地址
	num2[0] = 9
	fmt.Println(num2, num)

	// 切片添加的数据，小于数组原有的内存空间
	// 当我们向切片中添加元素的时候，实际上是在内存中切片的 后面 添加一段
	// 所以给切片添加元素，会改变数组中的数据
	num2 = append(num2, 1, 2)
	fmt.Println(num2, num)

	// 切片中添加的元素，超过了数组自己已有的内存空间
	// 那这个时候，切片会新建一个内存空间存放数据，而原数组保持不变
	// 同时二者的内存空间也不再相同
	num2 = append(num2, 1, 4, 5, 1, 1, 3, 42)
	fmt.Println(num2, num)
	fmt.Printf("%p, %p\n", num2, &num[0])

	// 使用切片实现删除 --> 操作之后内存空间不变，所以我们操作的是同一个对象
	j := 1 // 需要删除的元素的索引
	newSlice := num[0:j]
	newSlice = append(newSlice, num[j+1:]...)
	fmt.Println(num, newSlice) // 导致原切片内容也发生变化 ==> 不要使用原切片
	fmt.Printf("%p, %p\n", &num, newSlice)

}
