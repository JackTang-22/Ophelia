package test

import (
	"fmt"
	"os/exec"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/9/22 17:39
 * @Description:
 */
func TestExec(t *testing.T) {
	cmd := exec.Command("echo", "-n", "hello")
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}
	stdout, err := cmd.StdoutPipe() // 返回一个输出管道
	if err != nil {
		fmt.Println(err)
		return
	}
	output := make([]byte, 30)
	n, err := stdout.Read(output)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output[:n])
}
