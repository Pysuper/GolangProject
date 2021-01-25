/*
@Time : 2021/1/25 21:57
@Author : Zhengxingtao
@File : 12_循环链表.go
@Software: GoLand
*/

package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 代表整个循环列表，有代表循环链表中的第一个元素
	r := ring.New(5)

	r.Value = 0
	r.Next().Value = 1
	r.Next().Next().Value = 2
	//r.Next().Next().Next().Value=4
	r.Prev().Value = 4
	r.Prev().Prev().Value = 3

	// 输出: 循环链表有几个元素，func就执行几次，i代表当前执行元素的内容
	r.Do(func(i interface{}) {
		fmt.Println(i)
	})

	// 移动指针查找
	fmt.Println(r.Move(2).Value)

	// 增加元素
	r1 := ring.New(1)
	r1.Value = 5
	r.Next().Link(r1) // 添加到第一个元素的后面

	// 删除元素 ==> 取值：n%r.len(), 超出列表后面的元素，循环到开头
	r.Unlink(1) // 删除当前元素后面的 一 个元素
	r.Unlink(2) // 删除当前元素后面的 二 个元素

	r.Do(func(i interface{}) {
		fmt.Println(i)
	})

}
