package main

import (
	"fmt"
	"math"
	"math/cmplx"
)
// 变量定义
var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

//欧拉公式
func euler() {
	fmt.Printf("%.3f\n",  //f is format
		cmplx.Exp(1i*math.Pi)+1)  // 1i就代表是负数,1不是一个变量
		// import cmath cmath.exp(1j * cmath.pi) + 1   python j 代表负数
}

// 强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	// math.Sqrt(float64(a*a + b*b)) 返回的结果也是float64 c是一个int 要用显式类型转换
	fmt.Println(c)
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

//常量 const
func consts() {
	const (		// const 数值可作为各种类型使用
		filename = "abc.txt"   // 常量可以定义在函数外边 const filename = "abc.txt"
		a, b     = 3, 4		//一般常量不全部大写,大写另有含义
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))  //数值可作为各种类型使用,不用强制类型转换
	fmt.Println(filename, c)
}

//特殊的常量-枚举
func enums() {
	const (
		cpp = iota  //自增的
		_
		python
		golang
		javascript   // javascript value:4
	)

	const (
		b = 1 << (10 * iota) //iota 作为自增值的种子
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

// 变量类型写在变量名之后
// 没有char, 只有rune
// 原生支持负数类型

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
	consts()
	enums()
}
