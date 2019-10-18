package base

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/18 14:57
 * @Description:
 */

type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct {
}

func (f *file) WriteData(data interface{}) error {
	fmt.Println("WriteData: ", data)
	return nil
}

func TestInterface(t *testing.T) {
	f := &file{}
	_ = f.WriteData("data test") // WriteData:  data test

}
