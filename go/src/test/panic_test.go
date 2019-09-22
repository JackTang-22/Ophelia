package test

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/9/22 15:08
 * @Description:
 */

func TestPanic(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	outerFunc()
}

func outerFunc() {
	innerFunc()
}

func innerFunc() {
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())
	panic(errors.New("test panic"))
}
