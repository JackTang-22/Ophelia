package base

import (
	"bytes"
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/18 11:37
 * @Description:
 */

func init() {
	fmt.Println("init....")
}

type myFunc func(int, int) int // 定义函数类型

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func funcType(funcName myFunc, a, b int) {
	// fmt.Printf("%T\n", funcName)
	fmt.Println(funcName(a, b))
}

func TestFunc(t *testing.T) {
	fmt.Println(function6("a", "b", "c")) // abc
	// 匿名函数
	a := 1
	func1 := func(num int) {
		fmt.Println(num)
	}
	func1(a) // 1

	func() {
		fmt.Println("hello") // hello
	}()

	funcType(add, 1, 2) // 3
	funcType(sub, 2, 1) // 1

	f := outer(1)
	g := outer(2)
	fmt.Println(f(2)) // 0xc000018110 1  0xc000018120 2  1 2
	fmt.Println(g(2)) // 0xc000018118 2  0xc00009c1a0 2  2 2
	// 累加器
	o := funcOuter(0)
	fmt.Println(o()) // 1
	fmt.Println(o()) // 2
	fmt.Println(o()) // 3

	// TODO 常用函数
}

// 闭包函数
func outer(a int) func(b int) (int, int) {
	return func(b int) (int, int) {
		fmt.Println(&a, a)
		fmt.Println(&b, b)
		return a, b
	}
}

func funcOuter(a int) func() int {
	return func() int {
		a++
		return a
	}
}

// 函数的基本定义写法
func function() {}

func function1() int {
	return 0
}

func function2() (res int) {
	return res
}

func function3() (res int) {
	res = 1
	return
}

func function4() (a int, b int, c string) {
	return a, b, c
}

func function5(a, b int) int {
	return a + b
}

// 可变参数
func function6(s ...string) string {
	var buf bytes.Buffer
	for _, s := range s {
		buf.WriteString(s)
	}
	return buf.String()
}
