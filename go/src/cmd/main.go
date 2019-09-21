package main

import "fmt"

func main() {
	f := icur()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3

	f1 := icur()
	f2 := icur()
	fmt.Println(f1()) // 1
	fmt.Println(f2()) // 1
}

func icur() func() int {
	var x int
	return func() int {
		x++
		return x
	}

}
