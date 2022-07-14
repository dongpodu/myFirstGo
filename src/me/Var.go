package main

import "fmt"

func main() {
	//testVar()
	testMultipleVar()
	//testConstant()
}

/**
变量申明方式有：1、var 变量名 变量类型=变量值，如果不赋值，使用的是该数据类型的默认值
2、var 变量名=变量值。根据变量值推断出类型，这种必须在声明时赋值。
3、变量名称:=变量值。省略了var和变量类型
*/
func testVar() {
	var str = "111111"
	fmt.Println(str)

	var bool = true
	fmt.Println(bool)

	var num int64 = 12
	fmt.Println(num)

	var floatNum float32 = 1.0
	fmt.Println(floatNum)

	test := "test"
	fmt.Println(test)
}

/**
第一种：var 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...

第二种：var 变量名称,变量名称 ... = 变量值,变量值 ...

第三种：变量名称,变量名称 ... := 变量值,变量值 ...
*/
func testMultipleVar() {
	var s, s1, s2 string = "s", "s2", "s3"
	fmt.Println(s, s1, s2)

	var s3, s4 = "s3", "s4"
	fmt.Println(s3, s4)

	s5, s6 := "s5", "s6"
	fmt.Println(s5, s6)
}

func testConstant() {
	const name string = "will" // constant 常量名 类型
	fmt.Println(name)

	const numConstant = 100 // constant 常量名
	fmt.Println(numConstant)

	const n, n1, n2 int32 = 1, 2, 3 // 多个常量声明
	fmt.Println(n, n1, n2)

	const m, m1, m2 = 1, 2, 3 // 多个常量声明
	fmt.Println(m, m1, m2)
}
