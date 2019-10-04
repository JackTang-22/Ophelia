package main

/**
 * @author: tangye
 * @Date: 2019/9/24 14:06
 * @Description:
 */
func outer() func() int {
	var x = 0
	return func() int {
		x++
		return x
	}
}
