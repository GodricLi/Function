package main

import (
	"fmt"
)

func calc(a, b int) (sum, sub int) {
	//使用sum，sub命名返回值，
	sum = a + b
	sub = a - b
	return
}

func calc_v2(a ...int) {
	//可变参数a...,可以传入一个或多个参数，也可以不传入参数,同类型calc(a,b ...int),calc(a,b,c ...int)
	fmt.Println(a)

}

func main() {
	// sum,sub:= calc(6, 2)
	// fmt.Println(sum,sub)
	//忽略某个返回值使用_,可以不调用
	sum, _ := calc(6, 2)
	fmt.Println(sum)
	calc_v2(3)
	calc_v2(6, 6, 6)
	calc_v2()
}
