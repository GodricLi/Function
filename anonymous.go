package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func testDefer() {
	var i int = 1
	defer fmt.Println(i)
	i = 100
	fmt.Println(i)
	return
}

func testDeferAnonymous() {
	var i int = 1
	//defer定义的匿名函数，此时变量i在函数作用域外面，i被重新复制后会影响到defer里面的匿名函数
	defer func() {
		fmt.Println(i)
	}() //加括号调用
	i = 100
	fmt.Println(i)
	return
}

func reduce(a, b int) int {
	return a - b
}

func testFunc(a, b int, op func(int, int) int) int {
	//将函数当作参数传入另一个函数，被传入的函数的型式要与参数的型式一致：参数个数，类型，返回值
	return op(a, b)
}

func main() {
	f := add
	fmt.Println(f(2, 5))
	//匿名函数不需要函数名，直接写函数体
	anonym := func(a, b int) int {
		return a + b
	}
	fmt.Println(anonym(2, 5))
	testDefer()
	testDeferAnonymous()
	sum := testFunc(200, 300, add) //传入函数名
	fmt.Println(sum)
}
