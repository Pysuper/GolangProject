/*
@Time : 2021/3/10 22:54
@Author : Zhengxingtao
@File : 08_方法.go
@Software: GoLand
*/

package main

import "fmt"

/*
方法==> 面向对象编程
go中的面向对象，属于一种特定类型的函数
func (变量名 结构体类型) 方法名(参数列表) 返回值列表{
	方法体
}
*/

type Person struct {
	Name string
	Age  int
}

func (person *Person) run() {
	// 如果需要从外面修改 结构体中的内容，需要把上面的结构体，修改为引用类型
	fmt.Println(person.Name, person.Age)
}

func main() {
	peo := Person{"zhang", 130}
	peo.run()
	peo.Age -= 1
	fmt.Println(peo.Age)
}
