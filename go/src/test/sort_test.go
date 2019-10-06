package test

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/5 22:32
 * @Description:
 * TODO: sort
 */
type MyString []string
type i int

func (slice MyString) Len() int {
	return len(slice)
}

func (slice MyString) Less(i, j int) bool {
	return strings.ToLower(slice[i]) <= strings.ToLower(slice[j])
}

func (slice MyString) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func SortFoldedStrings(slice []string) {
	sort.Sort(MyString(slice))
}

func TestSort(t *testing.T) {
	str := []string{"a", "b", "c", "d", "E", "f", "G"}
	target := "E"
	sort.Strings(str)
	fmt.Println(str) // [E G a b c d f]
	SortFoldedStrings(str)
	fmt.Println(str) // [a b c d E f G]
	j := sort.Search(len(str), func(i int) bool {
		return strings.ToLower(str[i]) >= strings.ToLower(target)
	})
	fmt.Println(j, len(str), str[j])
	if j < len(str) && str[j] == target {
		fmt.Println(j, " ", str[j])
	}
	var s interface{} = 99
	fmt.Printf("%T\n", s)
	if res, ok := s.(i); ok {
		fmt.Printf("%T -- %d \n", res, res)
	} else {
		fmt.Println(ok)
	}
}
