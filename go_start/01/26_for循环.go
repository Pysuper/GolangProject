/*
@Time : 2021/1/14 20:21
@Author : Zhengxingtao
@File : 02_for循环.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	// TODO：经典for循环结构
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// 如果没有中间的判断语句，就是死循环
	// for ; ; {
	// 	fmt.Println("执行")
	// }

	// for循环的另一种写法
	//i := 0
	//for ; i < 5; {
	//	fmt.Println(i)
	//	i++
	//}
	//
	//fmt.Println("=========")
	//
	//j := 15
	//for j < 20 {
	//	fmt.Println(j)
	//	j++
	//}

	// 遍历数组(一)
	arr := [3]string{"张", "李", "郑"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 遍历数组，结合range
	// i 角标
	// n 元素内容 n = arr[i]
	//for i, n := range arr {
	for _, n := range arr {
		fmt.Println(n)
	}

	// TODO：continue：结束本次循环体，然后执行表达式3(i++)
	for i := 0; i < 5; i++ {
		if i <= 3 {
			continue
		}
		fmt.Println(i)
	}

	// 双重for循环中的continue
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i != j {
				continue
			}
			fmt.Println(i, j)
		}
	}

qwe:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i != j {
				continue qwe // 这个时候跳出循环就是跳出外层for循环了
			}
			fmt.Println(i, j)
		}
	}

	// break：终端for循环，不管还有几次需要执行，直接终止
	// 在双重for循环中，break也是默认影响 最近的 for循环
asd:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i != j {
				break asd // 这个时候break就是跳出外层for循环了
			}
			fmt.Println(i, j)
		}
	}
}
