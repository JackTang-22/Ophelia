package test

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/9/22 11:18
 * @Description:
 */

var v = "1 2 3"

func TestVar(t *testing.T) {
	var v = "1。 2。 3"
	if v != "" {
		fmt.Println(v)
		var v = "1， 2， 3"
		fmt.Println(v)
	}

	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Print(i)
		}()
		defer func(n int) {
			fmt.Print(n)
		}(i)
	}
}
