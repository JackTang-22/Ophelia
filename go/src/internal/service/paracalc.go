package service

/**
 * @author: tangye
 * @Date: 2019/9/14 22:43
 * @Description:
 */

func Paracalc(values []int, chanRes chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	chanRes <- sum
}
