package test

import (
	"fmt"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/9/15 13:32
 * @Description:
 */

type Rect struct {
	x, y float64
	w, h float64
}

// 相当于构造函数
func NewRect(x, y, w, h float64) *Rect {
	return &Rect{x, y, w, h}
}

func (rect *Rect) area() float64 {
	return rect.w * rect.h
}

func TestRect(t *testing.T) {
	rect := NewRect(1, 2, 3, 4)
	fmt.Println(rect.area())
}
