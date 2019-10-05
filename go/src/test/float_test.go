package test

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"testing"
)

/**
 * @author: tangye
 * @Date: 2019/10/5 15:14
 * @Description:
 * float32 float64 complex64 complex128
 */

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

func getStats(numbers []float64) (static statistics) {
	static.numbers = numbers
	sort.Float64s(static.numbers)
	static.mean = sum(numbers) / float64(len(numbers))
	static.median = median(numbers)
	return static
}

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2.0
	}
	return result
}

func homePage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(w, r)
	}
}

func TestFloat(t *testing.T) {
	fmt.Println(math.Modf(20.22))
	// 请求路径  路径请求时被执行的函数的引用， 这个函数的签名必须是func(http.ResponseWriter, *http.Request)
	http.HandleFunc("/", homePage)
	// 使用给定的tcp地址启动服务器  自定义服务器 nil表示默认类型
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("Failed to tart server", err)
	}
}
