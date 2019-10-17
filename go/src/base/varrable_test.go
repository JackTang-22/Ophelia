package base

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

/**
 * @author: tangye
 * @Date: 2019/10/16 23:56
 * @Description:
 */

// s := 1 expected declaration, found s

func TestVariable(t *testing.T) {
	var a int
	var b int = 11
	var c = 12
	d := 13
	a = 14
	t.Log(a, b, c, d) // 14 11 12 13
	e := 10.1
	f := "abc"
	// t.Log(s)
	// 1110--14-1.010000e+01-1.010000E+01-10.100000-16-0xc0000e8010-abc-int
	fmt.Printf("%b-%c-%d-%e-%E-%f-%o-%p-%s-%T", a, a, a, e, e, e, a, &a, f, a)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(a)) // 8

	g := 10.10
	h := 10.1
	fmt.Println(isEqual(g, h, 0.00000000001)) // true

	var t1  = complex(1, 2)
    fmt.Println(t1) // (1+2i)
}

func isEqual(a, b, p float64) bool {
	return math.Abs(a-b) < p
}
