package main

import (
		"fmt"
		"strings"
		"time"
		)
//闭包函数:
/*闭包的概念：是可以包含自由（未绑定到特定对象）变量的代码块，
	这些变量不在这个代码块内或者任何全局上下文中定义，而是在定义代码块的环境中定义。
	要执行的代码块（由于自由变量包含在代码块中，所以这些自由变量以及它们引用的对象没有被释放）
	为自由变量提供绑定的计算环境（作用域）。
	闭包的价值 : 闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，
	这意味着不仅要表示数据还要表示代码。支持闭包的多数语言都将函数作为第一级对象，
	就是说这些函数可以存储到变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回。
	Go语言中的闭包同样也会引用到函数外的变量。闭包的实现确保只要闭包还被使用，那么被闭包引用的变量会一直存在。*/

func adder()func(int)int{
	//调用外部变量
	var x int
	return func(d int) int{
		x += d
		return x
	}
}

func add2(a int) func (int)int  {
	//调用外部函数的参数
		return func (i int)int  {
			a+=i
			return a
		}
	} 

func makeStrings(suffix string) func (string) string {
	return func (name string) string {
		if !strings.HasSuffix(name,suffix){
			return name + suffix
		}
		return name
	}	
} 
	
func testclosure2(a int)(func (int)int,func(int)int)  {
	//多个返回值
	sum:=func (i int)int  {
		a +=i
		return a
	}
	sub:=func (i int)int  {
		a-=i
		return a
	}
	return sum,sub
}

func testclosure4()  {
	// 闭包的副作用
	for i:=0;i<5;i++{
		go func (index int)  {
			fmt.Println(index)
		}(i)
	}
	time.Sleep(time.Second)
}

func testclosure()  {
	f:=adder()
	ret:=f(1)
	fmt.Println(ret)
	ret=f(10)
	fmt.Println(ret)
	ret=f(100)
	fmt.Println(ret)
		
	f1:=add2(10)
	ret1:=f1(1)
	fmt.Println(ret1)
	ret1=f1(1)
	fmt.Println(ret1)
	//只要f还在被调用，x的值就不会被释放，会传入下次调用

	f2:=makeStrings(".txt")
	f3:=makeStrings(".jpg")
	ret2:=f2("test")
	ret3:=f3("test")
	fmt.Println(ret2)
	fmt.Println(ret3)
	ret4:=f2("test2")
	fmt.Println(ret4)

	c1,c2:=testclosure2(10)
	fmt.Println(c1(1),c2(2))
	fmt.Println(c1(3),c2(4))

}

func main(){

	testclosure()
	testclosure4()
}