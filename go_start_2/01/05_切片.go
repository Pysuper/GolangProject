/*
@Time : 2021/1/14 23:50
@Author : Zhengxingtao
@File : 05_切片.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	// 声明时为空
	var sli []string  // 切片
	var arr [3]string // 数组

	fmt.Println(sli)
	fmt.Println(arr)
	fmt.Println(sli == nil) // == 的时候只能和nil比较
	fmt.Printf("%p\n", sli) // 没有内存地址

	// 切片赋值
	name := []string{"张三", "李四"}
	fmt.Println(name)
	fmt.Println(name[0])
	name[1] = "王二"
	fmt.Println(name)

	// 切片拷贝 ==> 指向的是同一个内存地址 ==> 切片是引用类型
	// 更改其中一个之后，另一个也会变化
	names := name
	names[0] = "麻子"
	fmt.Printf("%p\n", name)
	fmt.Printf("%p\n", names)
}
