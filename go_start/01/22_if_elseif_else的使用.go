/*
@Time : 2021/1/13 1:18
@Author : Zhengxingtao
@File : 22_if_elseif_else的使用.go
@Software: GoLand
*/

package main

import "fmt"

func main() {
	if a := 10; a == 100 {
		fmt.Println("True")
	} else if a <= 100 {
		fmt.Println("a <= 100")
	} else {
		fmt.Println("啥也不是")
	}
}
