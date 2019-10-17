package base

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/17 16:21
 * @Description:
 */

type Person struct {
	// 字段小写包外不可见
	name string
	age  int
}

type Name struct {
	name string
	Age
}

type Age struct {
	age int
}

func TestStruct(t *testing.T) {
	var p1 Person = Person{name: "person", age: 19}
	fmt.Println(p1) //{person 19}

	p2 := new(Person)      // 指针类型
	fmt.Printf("%T\n", p2) // *base.Person
	fmt.Println(p2, *p2)   // &{ 0} { 0}
	p2.name = "go"
	p2.age = 3
	fmt.Println(p2, *p2) // &{go 3} {go 3}

	var p3 Person
	fmt.Printf("%T\n", p3) // base.Person
	fmt.Println(p3)        // { 0}
	p3.name = "name"
	p3.age = 2
	fmt.Println(p3) // {name 2}

	// p4 := &Person{} == new(Person)
	p4 := &Person{name: "p3", age: 4}
	fmt.Printf("%T-%v\n", p4, *p4) // *base.Person-{p3 4}

	n1 := &Name{name: "n1", Age: Age{age: 10}}
	fmt.Println(*n1) // {n1 {10}}
	fmt.Println(n1.age) // 10 语法糖 实际上是(*n1).age

}
