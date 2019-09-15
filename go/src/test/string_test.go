package test

import (
	"fmt"
	"math"
	"strings"
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
	fmt.Printf("%c", 97)
	str := "golang"
	b := []byte(str) // 字符串转换成byte切片 O(1) 底层byte可以引用字符串的底层字节
	for i, v := range b {
		t.Logf("%d    %c\n", i, v)
	}
	s := string(b) // 同上
	t.Log(s)
	str = "i study golang"
	// 适用于ASCII
	i := strings.Index(str, " ")
	firstWord := str[:i]
	i = strings.LastIndex(str, " ")
	lastWord := str[i+1:]
	t.Log(firstWord, lastWord)
}
