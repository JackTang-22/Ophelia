package test

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
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

func a() {
	fmt.Printf("%c", 97)
	str := "golang"
	b := []byte(str) // 字符串转换成byte切片 O(1) 底层byte可以引用字符串的底层字节
	for i, v := range b {
		fmt.Printf("%d    %c\n", i, v)
	}
	s := string(b) // 同上
	fmt.Println(s)
	str = "i study golang"
	// 适用于ASCII
	i := strings.Index(str, " ")
	firstWord := str[:i]
	i = strings.LastIndex(str, " ")
	lastWord := str[i+1:]
	log.Fatal(firstWord, lastWord)
}

func TestString(t *testing.T) {
	fmt.Println(len("语言"))                    // 6  字节数
	fmt.Println(len([]rune("语言")))            // 2 字符个数
	fmt.Println(utf8.RuneCountInString("语言")) // 2 字符个数 速度快于len([]rune("语言"))
	for _, v := range []rune("语言") {
		// 35821 35328
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println(string([]rune{35821, 35328})) // 语言
	/*
	   比较操作符在内存中一个字节一个字节的比较字符串
	   字符串直接比较可能会出现的问题：
	   1. 有些Unicode编码的字符可以用多个不同的字节序列表示（字符串来源于外部时并且使用不同于go的码点映射时才可能引起）  exp/norm
	   2. 用户可能希望将不同的字符看成相同的  例如：5 ⑤
	   3. 某些字符的排序与语言有关
	*/
	fmt.Println("语言" == string([]rune{35821, 35328}))
	fmt.Println(strings.Index("语言", "语言"))
	s1 := "a,b,c,d,e"
	res1 := strings.Split(s1, ",")          //[a b c d e]
	res2 := strings.SplitN(s1, ",", 2)      //[a b,c,d,e]
	res3 := strings.SplitAfter(s1, ",")     //[a, b, c, d, e]
	res4 := strings.SplitAfterN(s1, ",", 2) //[a, b,c,d,e]
	fmt.Println(res1, res2, res3, res4)

	s2 := "a,b,c*d|e,f*g"
	res5 := strings.FieldsFunc(s2, func(r rune) bool {
		switch r {
		case ',', '|', '*':
			return true
		}
		return false
	})
	fmt.Println(res5) //[a b c d e f g]
	fmt.Println(strings.Replace(s1, ",", "*", -1))
	fmt.Println(strings.ReplaceAll(s1, ",", "|"))

	for _, truth := range []string{"1", "True", "t", "false", "F", "0", "5"} {
		if b, err := strconv.ParseBool(truth); err != nil {
			fmt.Printf("\n{%v}", err)
		} else {
			fmt.Print(b, " ")
		}
	}
	fmt.Println()

}
