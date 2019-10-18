package base

import (
	. "fmt"
	"reflect"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/18 15:06
 * @Description:
 */

func TestReflect(t *testing.T) {
    reflectTest1()
    // TODO other reflect
}

func reflectTest1() {
	var a int
	Println(reflect.ValueOf(a))       // 0  变量值
	Println(reflect.TypeOf(a))        // int 变量类型对象名，其类型为 reflect.Type()
	Println(reflect.TypeOf(a).Name()) // int 变量类型对象的类型名
	Println(reflect.TypeOf(a).Kind()) // int 变量类型对象的种类名
}
