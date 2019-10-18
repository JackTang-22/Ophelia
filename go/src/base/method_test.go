package base

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/18 14:50
 * @Description:
 */

type Animal struct {
	Name string
	Age  int
}

func (a Animal) run(name string) {
	a.Name = name
	fmt.Println(a.Name, " is run")
}

func TestMethod(t *testing.T) {
	a := &Animal{}
	a.run("animal") // animal  is run
}
