package test

import (
	. "Ophelia/go/src/internal/service"
	"fmt"
	"os"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/9/14 22:48
 * @Description:
 */

func TestFly(t *testing.T) {
	fmt.Println("hello world")
	fmt.Println(os.Args)
	var bird Bird
	bird.Fly()
	var f IFly = new(Bird)
	f.Fly()
}
