package main

import "fmt"

func printArray(arr [5]int) {  // func printArray(arr *[5]int) {  使用指针可改变数组值
	arr[0] = 100
	for i, v := range arr {   // range 遍历数组并使i,v获得数组下标和值      for i:=0, i<=len(arr), i++{}
		fmt.Println(i, v)   // 可通过下划线来省略变量
	}
}

// 为什么要用range ?   1.简洁美观

// 数组是值类型
// [10]int  和  [20]int 是不同类型
// 在go语言中一般不直接使用数组


func main() {
	var arr1 [5]int  // 定义长度为5的一个数组 ,初始值为0
	arr2 := [3]int{1, 3, 5}  // 定义一个长度为3的数组 并赋值
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int  // 定义一个4行5列的二维数组

	fmt.Println("array definitions:")
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("printArray(arr3)")
	printArray(arr3)  //调用func f(arr [10]int) 会拷贝数组

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}

