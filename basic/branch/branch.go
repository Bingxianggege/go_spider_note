package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	// switch 后可以没有表达式
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	// If "abc.txt" is not found,
	// please check what current directory is,
	// and change filename accordingly.
	const filename = "branch/abc.txt"  // 相对于项目路径
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		// Uncomment to see it panics.
		// grade(-3),
	)
}
