package base

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/18 14:09
 * @Description: 延迟执行
 */

func TestDefer(t *testing.T) {
	/* defer执行顺序 按照栈的方式先进后出
	main
	defer2
	defer1
	*/
	defer fmt.Println("defer1")
	// panic("111")
	defer fmt.Println("defer2")
	fmt.Println("main")

	// defer在入栈时，会将相关的值也copy到栈中
	num := 0
	defer fmt.Println("defer num = ", num) // defer num =  0
	num = 1
	fmt.Println("main num = ", num) // main num =  1
}
