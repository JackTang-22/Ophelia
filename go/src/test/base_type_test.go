package test

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/9/17 17:00
 * @Description:
 */

func TestBaseType(t *testing.T) {
	i := int(1234)
	j := 1234
	fmt.Println(i == j)
	var m map[string]string
	var m1 = make(map[string]string)
	fmt.Println(m == nil)  // true
	fmt.Println(m1 == nil) // false
	m = make(map[string]string)
	fmt.Println(m == nil) // false

}
