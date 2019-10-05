package main

import (
	"fmt"
)

func init() {
	fmt.Println("init")
}

func main() {
	var x = 1
	f := icur(x)
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 4
	fmt.Println(x)

	f1 := icur(1)
	f2 := icur(1)
	fmt.Println(f1()) // 2
	fmt.Println(f2()) // 2
	/*inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(input)*/
	fmt.Println()
	f3 := outer()
	fmt.Println(f3())
	fmt.Println(f3())
}

func icur(x int) func() int {
	return func() int {
		x++
		return x
	}
}
