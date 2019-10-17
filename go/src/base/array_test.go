package base

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/17 15:10
 * @Description:
 */

func TestArray(t *testing.T) {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr) // [1 2 3]
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d-%d ", i, arr[i]) // 0-1 1-2 2-3
	}
	fmt.Println()
	for i, v := range arr {
		fmt.Printf("%d-%d ", i, v) // 0-1 1-2 2-3
	}
}
