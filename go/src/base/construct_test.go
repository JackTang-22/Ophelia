package base

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/18 14:31
 * @Description:
 */

type Student struct {
	Name string
	Age  int
}

type School struct {
	Student
	ClassRoom int
}

type Integer int

func (i Integer) Less(j Integer) bool {
	return i < j
}

func NewStudentByName(name string) *Student {
	return &Student{Name: name}
}

func NewStudentByAge(age int) *Student {
	return &Student{Age: age}
}

// 构造父类
func NewStudent(name string, age int) *Student {
	return &Student{Name: name, Age: age}
}

// 构造子类
func NewSchool(classRoom int) *School {
	return &School{
		ClassRoom: classRoom,
	}
}

/**
模拟构造函数
*/
func TestConstruct(t *testing.T) {
	s1 := NewStudentByAge(12)
	s2 := NewStudentByName("hello")
	fmt.Println(s1, s2, *s1, *s2) // &{ 12} &{hello 0} { 12} {hello 0}

	s := NewSchool(123)
	s.Student = *NewStudent("go", 12)
	fmt.Println(s, *s) // &{{go 12} 123} {{go 12} 123}

	var i Integer = 1
	var j Integer = 2
	fmt.Println(i.Less(j)) // true
}
