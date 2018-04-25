package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)
// 四则运算
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(  // 返回报错信息
			"unsupported operation: %s", op)
	}
}

// 做带余除法
func div(a, b int) (q, r int) {  //(a, b int) 参数类型     (q, r int) 返回值类型
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer() //获得这个函数的指针
	opName := runtime.FuncForPC(p).Name() // 获得函数名
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {  // range 对数字渲染
		s += numbers[i]
	}
	return s
}

// 无法交换两个值,要使用指针类型,或是通过返回值来交换(如下函数)
func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println("Error handling")
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)  // 收两个返回值, q, _ := div(13, 3) 不想要的返回值用_(下划线)表示
	fmt.Printf("13 div 3 is %d mod %d\n", q, r)

	fmt.Println("pow(3, 4) is:", apply(
		func(a int, b int) int {  // 匿名函数
			return int(math.Pow(  // pow 3的4次方
				float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("1+2+...+5 =", sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println("a, b after swap is:", a, b)
}

// 函数是一等公民,函数式编程

//返回值类型写在最后面
//可返回多个值
//函数作为参数
//没有默认参数