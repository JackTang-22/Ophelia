package test

import (
	"fmt"
	"reflect"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/9/14 22:54
 * @Description:
 */

type Birds struct {
	Name string
	Life int
}

func (b *Birds) fly() {
	fmt.Println("i am flying")
}

func TestReflect(t *testing.T) {
	sparrow := &Birds{"name", 10}
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		t.Fatalf("%d:   %s   %s   =   %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
