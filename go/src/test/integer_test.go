package test

import (
	"fmt"
	"math/big"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/5 14:52
 * @Description:
 * byte=unit8 int(int32 or int 64) int16 int 32 int64
 * rune=unit32 uint(uint32 or uint64) uint8 uint16 uint32 uint64
 * uintptr 可以容纳指针值的无符号整数类型(uint32 or uint64)
 * big.Int:大整数  big.Rat:有理数
 */

func TestInteger(t *testing.T) {
	a := big.NewInt(10)
	b := big.NewInt(20)
	fmt.Println(a.Add(a, b))
}
