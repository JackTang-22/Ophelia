package base

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/17 13:49
 * @Description:
 */

func TestChar(t *testing.T) {
	var a = 'a'
	var b = 'b'
	fmt.Println(a, b)            // 97 98
	fmt.Printf("%c, %c\n", a, b) // a, b

	var c = '中'
	fmt.Println(c)  // 20013
	fmt.Printf("%T, %c\n", c, c)  // int32, 中
}
