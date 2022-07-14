package main

import "fmt"

/**
声明方式：
1、var 数组名 [数组长度]数组类型
2、var 数组名 = [数组长度] 数组类型{数值1,数值2....}
2、数组名 := [数组长度] 数组类型{数值1,数值2....}
*/
func main() {
	var array [5]int //申明一个int数组，长度为5。数组一旦声明，长度就不能变了
	fmt.Println(array)

	array[0] = 1
	fmt.Println(array) //指定第一个元素值为1

	var array2 = [2]int{1, 2}
	fmt.Println(array2)

	array3 := [2]int{1, 2}
	fmt.Println(array3)

	var array4 = [5]int{0, 0, 0, 0, 0} //在函数中传递的时候是传递的值，如果传递数组很大，这对内存是很大开销。
	modify(array4)
	fmt.Println(array4)

	array = array4 // 同样类型的数组（长度一样且每个元素类型也一样）才可以相互赋值，反之不可以
}

func modify(array [5]int) {
	array[2] = 1
}
