package service

import "fmt"

/**
 * @author: tangye
 * @Date: 2019/9/14 22:38
 * @Description:
 */

type Bird struct {
}

type IFly interface {
	Fly()
}

func (b *Bird) Fly() {
	fmt.Printf("%v %T fly\n", b, b)
}
