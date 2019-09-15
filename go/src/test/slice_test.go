package test

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/9/15 09:11
 * @Description:
 */

func TestSlice(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{7, 8, 9}
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
