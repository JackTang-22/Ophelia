package test

import (
	"fmt"
	"math"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/9/15 16:16
 * @Description:
 */

func Utf8Format(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxInt8 {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d is out of range uint8", x)
}

func TestString(t *testing.T) {
	fmt.Println('3' - '0')
	if u, err := Utf8Format(300); err != nil {
		t.Fatal(u)
	} else {
		t.Fatal(err)
	}
}
