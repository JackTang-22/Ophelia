package base

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/17 18:19
 * @Description:
 */

func TestPointer(t *testing.T) {
	v := "a"
	value := &v
	ptr := *value
	fmt.Println(&v, value, &ptr, v, *value, ptr) // 0xc00004e4d0 0xc00004e4d0 0xc00004e4e0 a a a

	var a = 10
	var ptr1 *int = &a
	fmt.Println(ptr1, *ptr1) // 0xc000018100 10

	var p *int
	p = new(int)
	*p = 2
	fmt.Println(p, *p) // 0xc0000182c0 2

}

func swap(a, b *int) {
	*a, *b = *b, *a
}
