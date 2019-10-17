package base

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

/**
 * @author: tangye
 * @Date: 2019/10/17 14:20
 * @Description:
 */

func TestString(t *testing.T) {
	//changeString()
	//stringByByte()
	//stringBySlice()
	stringOperating()
}

// 修改字符串错误的方式
func changeString() {
	var str1 = "hello "
	var str2 = " golang"
	fmt.Println(str1[0])        // 104
	fmt.Printf("%c\n", str1[0]) // h
	fmt.Println(len(str1))      // 6
	fmt.Println(str1 + str2)    // hello  golang
	// str1[0] = 'm'  cannot cast to str1[0]  编译错误
}

// 修改字符串的两种方式：
func stringByByte() {
	str := "golang"
	fmt.Println(&str) // 0xc00004e4d0
	strTmp := []byte(str)
	fmt.Println(strTmp) // [103 111 108 97 110 103]
	strTmp[0] = 'G'
	str = string(strTmp)
	fmt.Println(&str, str) // 0xc00004e4d0 Golang
	strRes := string(strTmp)
	fmt.Println(&strRes, strRes) // 0xc00004e4d0 Golang
}

func stringBySlice() {
	str := "hello"
	fmt.Println(&str, str) // 0xc0000f8020 hello
	str = "H" + str[1:]
	fmt.Println(&str, str) // 0xc0000f8020 Hello
}

// 字符串的常规操作
func stringOperating() {
	// golang默认以utf-8保存， 一个中文字符占3个字节
	var s1 = "hello"
	var s2 = "golang"
	var s3 = "我爱中国"
	fmt.Println(len(s1), len(s2), len(s3), utf8.RuneCountInString(s3)) // 5 6 12 4

	// 遍历string  unicode字符的遍历只能通过range来实现
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%d-%c ", i, s1[i]) // 0-h 1-e 2-l 3-l 4-o
	}
	fmt.Println()

	for i, v := range s1 {
		fmt.Printf("%d-%c ", i, v) // 0-h 1-e 2-l 3-l 4-o
	}

	a := 123
	fmt.Printf("%T-%v\n", string(a), string(a)) // string-{
}

//TODO strings strconv


