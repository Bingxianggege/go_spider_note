package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result   // strconv.Itoa 转字符串
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)  // panic 就是报错的意思
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader) // 一行一行的读

	for scanner.Scan() {  // 没有初始条件,没有递增条件,就相当于while
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {	// 什么都不加就是死循环
		fmt.Println("abc")
	}
}

// 可以省略初始条件,可以省略递增条件,什么都不写是死循环

func main() {
	fmt.Println("convertToBin results:")
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(72387885),
		convertToBin(0),
	)

	fmt.Println("abc.txt contents:")
	printFile("basic/branch/abc.txt")

	fmt.Println("printing a string:")
	s := `abc"d"
	kkkk
	123

	p`
	printFileContents(strings.NewReader(s))

	// Uncomment to see it runs forever
	// forever()
}
