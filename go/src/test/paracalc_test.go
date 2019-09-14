package test

import (
	. "Ophelia/go/src/internal/service"
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/9/14 22:51
 * @Description:
 */

func TestParaCalc(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	chanRes := make(chan int, 2)
	go Paracalc(values, chanRes)
	go Paracalc(values, chanRes)
	res1, res2 := <-chanRes, chanRes
	fmt.Println(res1, res2)
}
