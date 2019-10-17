package base

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/17 16:51
 * @Description:
 */

func TestSlice(t *testing.T) {
	// 常用创建方式
	var s1 []int
	s2 := []int{1, 2, 3}
	// make
	s3 := make([]int, 5, 10)
	// 数组
	arr := [5]int{1, 2, 3, 4, 5}
	var a, b []int
	a = arr[1:4]
	b = arr[0:3]
	fmt.Println(len(s1), cap(s1), s1) // 0 0 []
	fmt.Println(len(s2), cap(s2), s2) // 3 3 [1 2 3]
	fmt.Println(len(s3), cap(s3), s3) // 5 10 [0 0 0 0 0]
	fmt.Println(len(a), cap(a), a)    // 3 4 [2 3 4]
	fmt.Println(len(b), cap(b), b)    // 3 5 [1 2 3]
	// 切片追加 容量超过cap 扩容
	a = append(a, 1)
	fmt.Println(len(a), cap(a), a) // 4 4 [2 3 4 1]
	a = append(a, 9)
	fmt.Println(len(a), cap(a), a) // 5 8 [2 3 4 1 9]
	a1 := make([]int, 10)
	fmt.Printf("%T - %T\n", a, a1) // []int - []int
	a11 := copy(a1, a)
	fmt.Println(a11, a1) // 5 [2 3 4 1 9 0 0 0 0 0]

	index := 2
	a1 = append(a1[:index], a1[index+1:]...) // 删除元素
	fmt.Println(a1)                          // [2 3 1 9 0 0 0 0 0]

	str := "hello,世界"
	strByte := []byte(str)        //字符串转换为[]byte类型切片
	strRune := []rune(str)        //字符串转换为[]rune类型切片
	fmt.Println(strByte, strRune) // [104 101 108 108 111 44 228 184 150 231 149 140] [104 101 108 108 111 44 19990 30028]

	/*
		切片存储结构
		type slice struct {
			arrary = unsafe.Pointer		//指向底层数组的指针
			len int						//切片元素数量
			cap int						//底层数组的容量
		}
	*/
}
