package depth

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/18 15:20
 * @Description:
 */

func TestString(t *testing.T) {
	s1 := "hello"
	s2 := "hello1"
	fmt.Println(&s1, s1, &s2, s2) // 0xc00004e330 hello 0xc00004e340 hello1
	s2 = s1
	fmt.Println(&s1, s1, &s2, s2) // 0xc00004e4a0 hello 0xc00004e4b0 hello
}
